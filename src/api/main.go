package main

import (
	"anime-quote-api/src/api/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = 1331
)

func main() {

	serve(DEFAULT_HOST, DEFAULT_PORT)
}

// Load environment variables from the ".env" file
func loadEnvironmentToFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load environment variables")
	}
}

func serve(host string, port int) {
	// TODO: configure MONGO connection with docker container and api

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	apiGroup := router.Group("/api")

	routes := routes.New()
	routes.Load(apiGroup)

	address := getServerAddress(host, port)
	router.Logger.Fatal(router.Start(address))
}

func getServerAddress(host string, port int) string {
	return fmt.Sprintf("%v:%v", host, port)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
