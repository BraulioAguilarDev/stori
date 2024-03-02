package main

import (
	"fmt"
	"stori/api"
)

func config() (*api.Stori, error) {
	// db, err := postgres.ConnectInit("postgres://localhost:5432/sf_adtech_apptracking?sslmode=disable", "postgres", "root", 3)
	// if err != nil {
	// 	panic(err)
	// }

	// txns := transactionhdl.ProvideTransactionHandler()
	return &api.Stori{
		// TxnsHandler: txns,
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
