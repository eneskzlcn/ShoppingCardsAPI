package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func InitializeShoppingCardServerWithGivenHandlers(handlers ShoppingCardsHandlers) * fiber.App{
	var app = fiber.New()
	app.Use(cors.New())
	for _ , handler := range handlers {
		app.Add(handler.Method,handler.Route,handler.Function)
	}
	return app
}
