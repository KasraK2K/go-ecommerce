package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"app/cmd/api/routes"
	"app/config"
	"app/middleware"
	"app/pkg/storage/pg"
)

func (s *Server) setConfigs() {
	config.SetConfig()
}

func (s *Server) connectDatabases() {
	pg.Gorm.Connect()
}

func (s *Server) registerApplication() {
	s.App = fiber.New(fiber.Config{
		Prefork:               config.AppConfig.PREFORK,
		ServerHeader:          "Fiber",
		AppName:               "Go Blog v1.0.0",
		CaseSensitive:         true,
		StrictRouting:         false,
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})
}

func (s *Server) registerMiddleware() {
	s.App.Use(cache.New())
	s.App.Use(compress.New())
	s.App.Use(cors.New())
	s.App.Use(etag.New())
	s.App.Use(favicon.New())
	s.App.Use(limiter.New(limiter.Config{Max: 100, Expiration: 60 * time.Second}))
	s.App.Use(logger.New())
	s.App.Use(recover.New())
	s.App.Use(requestid.New())
	s.App.Use(middleware.HandleMultipart)
	s.App.Use(middleware.PullOutToken)
	s.App.Use(healthcheck.New())
}

func (s *Server) registerRoutes() {
	routes.Routes(s.App)
}

func (s *Server) start() {
	if !fiber.IsChild() {
		greenColor := color.New(color.FgGreen).SprintFunc()
		cyanColor := color.New(color.FgCyan).SprintFunc()
		domain := fmt.Sprintf("http://localhost:%s", config.AppConfig.PORT)
		message := fmt.Sprintf("%s %s", cyanColor("Server is starting on port"), greenColor(domain))
		fmt.Println(message)
	}

	log.Fatal(server.App.Listen(fmt.Sprintf("127.0.0.1:%s", config.AppConfig.PORT)))
}
