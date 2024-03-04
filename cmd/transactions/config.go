package main

import (
	"stori/api"
	"stori/internal/core/service"
	accounthdlr "stori/internal/handler/account"
	profilehdlr "stori/internal/handler/profile"
	s3hdlr "stori/internal/handler/s3"
	repository "stori/internal/storage"
	"stori/pkg/database"
)

func config() (*api.Stori, error) {
	db, err := database.ConnectInit("postgres://localhost:5432/stori?sslmode=disable", "postgres", "postgres", 3)
	if err != nil {
		panic(err)
	}

	// Profile setting
	profileRepo := repository.NewProfileRepository(db)
	profileService := service.ProvideProfileService(profileRepo)
	profileHdlr := profilehdlr.ProvideProfileHandler(profileService)

	//  S3 settings
	s3Repo := repository.NewAccountS3Repository(db)
	s3Service := service.ProvideAccountS3Service(s3Repo)

	// Account settings
	accountRepo := repository.NewAccountRepository(db)
	accountService := service.ProvideAccountService(accountRepo)

	accountHdlr := accounthdlr.ProvideAccountHandler(accountService, profileService, s3Service)
	s3Hdlr := s3hdlr.ProvideS3Handler(accountService, s3Service)

	return &api.Stori{
		ProfileHandler:   profileHdlr,
		AccountHandler:   accountHdlr,
		AccountS3Handler: s3Hdlr,
	}, nil
}
