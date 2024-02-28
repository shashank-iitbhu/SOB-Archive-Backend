package main

import (
	"fmt"
	"shashank-iitbhu/sob-archive/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Cannot load the environment variables")
		fmt.Println(err)
	}
	api.Start()
}
