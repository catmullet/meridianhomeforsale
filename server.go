package main

import (
	"github.com/labstack/echo"
	"meridianhomeforsale/api/controllers/garage"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
)

// todo.go
func main() {


	err := godotenv.Load()
	if err != nil {
	}

	// Create a new instance of Echo
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	garage.Routes(server)


	server.Logger.Fatal(server.Start(":5000"))

}
