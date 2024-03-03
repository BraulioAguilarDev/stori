package main

import (
	"fmt"
	"stori/api"
	"stori/internal/core/service"
	accounthdlr "stori/internal/handler/account"
	profilehdlr "stori/internal/handler/profile"
	repository "stori/internal/storage"
	"stori/pkg/database"
)

func config() (*api.Stori, error) {
	db, err := database.ConnectInit("postgres://localhost:5432/stori?sslmode=disable", "postgres", "postgres", 3)
	if err != nil {
		panic(err)
	}

	profileRepo := repository.NewProfileRepository(db)
	profileService := service.ProvideProfileService(profileRepo)
	profileHdlr := profilehdlr.ProvideProfileHandler(profileService)

	accountRepo := repository.NewAccountRepository(db)
	accountService := service.ProvideAccountService(accountRepo)
	accountHdlr := accounthdlr.ProvideAccountHandler(accountService, profileService)

	return &api.Stori{
		ProfileHandler: profileHdlr,
		AccountHandler: accountHdlr,
	}, nil
}

func main() {
	cleanup := func(err error) {
		fmt.Println(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Critical Error: %v\n", r)
		}
	}()

	api, err := config()
	if err != nil {
		cleanup(err)
	}

	api.SetupRouter()
	api.Router.Run(":8080")
	// lambda.Start(api.Handler)
}
