package main

// @title Student API
// @version 1.0
// @description This is a sample server Student API
// @host localhost:8080
// @BasePath /
// @schemes http

import (
	"github.com/MayconVyctor/API-students/api"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

// hello godoc
// @Summary Hello world
// @Description Retorna Hello
// @Tags test
// @Produce plain
// @Success 200 {string} string "hello"
// @Router /hello [get]
func hello(c echo.Context) error {
	return c.String(200, "hello")
}
