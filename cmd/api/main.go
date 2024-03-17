package main

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

var server Server

func main() {
	server.setConfigs()
	server.connectDatabases()
	server.registerApplication()
	server.registerMiddleware()
	server.registerRoutes()
	server.start()
}
