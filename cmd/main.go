package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())

	helloHandler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, fmt.Sprintf("Hello World!"))
	}

	e.GET("/hello", helloHandler)
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	port := os.Getenv("PORT")
	defaultPort := "8080"
	if port == "" {
		port = defaultPort
	}

	e.Logger.Fatal((e.Start(fmt.Sprintf(":%s", port))))
}
