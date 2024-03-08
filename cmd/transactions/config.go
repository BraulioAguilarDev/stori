package main

import (
	"stori/api"
	"stori/internal/core/service"
	accounthdlr "stori/internal/handler/account"
	profilehdlr "stori/internal/handler/profile"
	s3hdlr "stori/internal/handler/s3"
	transactionhdlr "stori/internal/handler/transaction"
	repository "stori/internal/storage"
	"stori/pkg/cloud/aws"
	"stori/pkg/database"
	"time"
)

func config() (*api.Stori, error) {
	db, err := database.ConnectInit("postgres://localhost:5432/stori?sslmode=disable", "postgres", "postgres", 3)
	if err != nil {
		return nil, err
	}

	// db, err := database.ConnectInit("postgres://165.227.68.59:5432/stori?sslmode=disable", "postgres", "postgres", 3)
	// if err != nil {
	// 	return nil, err
	// }

	// AWS settings
	ses, err := aws.New(aws.Config{
		Region: "us-west-2",
		ID:     "AKIA3FLDYOQ4QT4SYBGS",
		Secret: "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",
	})
	if err != nil {
		return nil, err
	}

	s3 := aws.NewS3(ses, time.Second*15)

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

	// account s3
	accountS3Repo := repository.NewAccountS3Repository(db)
	accountS3Service := service.ProvideAccountS3Service(accountS3Repo)

	accountHdlr := accounthdlr.ProvideAccountHandler(accountService, profileService, s3Service)
	s3Hdlr := s3hdlr.ProvideS3Handler(accountService, s3Service, s3)

	// Transaction settings
	txnRepo := repository.NewTransactionRepository(db)
	txnService := service.ProvideTransactionService(txnRepo)
	txnsHdlr := transactionhdlr.ProvideTransactionHandler(txnService, accountService, accountS3Service, s3)

	return &api.Stori{
		TransactionHandler: txnsHdlr,
		ProfileHandler:     profileHdlr,
		AccountHandler:     accountHdlr,
		AccountS3Handler:   s3Hdlr,
	}, nil
}
