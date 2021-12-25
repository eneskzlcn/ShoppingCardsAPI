package main

import "fmt"

func main(){
	InitializeModels()
	InitializeHandlers()
	app := InitializeShoppingCardServerWithGivenHandlers(shoppingCardsHandlers)
	app.Listen(fmt.Sprintf(":%d",5000))
}