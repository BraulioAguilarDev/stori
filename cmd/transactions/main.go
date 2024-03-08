package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

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
	lambda.Start(api.Handler)
}
