package main

import (
	"github.com/joho/godotenv" // swagger embed files
	// gin-swagger middleware

	"github.com/uzixCode/gocode/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")

	}

	cmd.Execute()

}
