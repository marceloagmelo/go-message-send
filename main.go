package main

import (
	"github.com/marceloagmelo/go-message-send/logger"

	"github.com/marceloagmelo/go-message-send/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	logger.Info.Println("Listen 8080...")
	app.Run(":8080")
}
