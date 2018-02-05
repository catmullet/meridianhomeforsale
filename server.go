package main

import (
	"github.com/labstack/echo"
	"meridianhomeforsale/api/controllers/garage"
	"github.com/joho/godotenv"
)

// todo.go
func main() {


	err := godotenv.Load()
	if err != nil {
	}

	// Create a new instance of Echo
	server := echo.New()

	garage.Routes(server)

	server.Logger.Fatal(server.Start(":5000"))

}
