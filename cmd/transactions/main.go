package main

import (
	"fmt"
	"stori/api"
	"stori/internal/core/service"
	profilehdlr "stori/internal/handler/profile"
	registerRepository "stori/internal/storage"
	"stori/pkg/database"
)

func config() (*api.Stori, error) {
	db, err := database.ConnectInit("postgres://localhost:5432/stori?sslmode=disable", "postgres", "postgres", 3)
	if err != nil {
		panic(err)
	}

	profileRepo := registerRepository.NewProfileRepository(db)
	profileService := service.ProvideProfileService(profileRepo)
	profileHdlr := profilehdlr.ProvideProfileHandler(profileService)

	return &api.Stori{
		ProfileHandler: profileHdlr,
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
