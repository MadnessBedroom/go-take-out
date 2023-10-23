// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/go-take-out/internal/server"
)

// Injectors from wire.go:

func NewApp() (*fiber.App, func(), error) {
	app := server.NewHTTPServer()
	return app, func() {
	}, nil
}
