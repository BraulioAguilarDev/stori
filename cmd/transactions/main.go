package main

import (
	"fmt"
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
	api.Router.Run(":8080")
}
