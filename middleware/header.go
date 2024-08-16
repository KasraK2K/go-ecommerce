package middleware

import (
	"github.com/gofiber/fiber/v2"

	"app/config"
)

func AddHeaderMiddleware(c *fiber.Ctx) error {
	c.Response().Header.SetCanonical([]byte("Backend-Version"), []byte(config.AppConfig.AppVersion))
	c.Response().Header.SetCanonical([]byte("Frontend-Version"), []byte(config.AppConfig.AppVersion))
	c.Response().Header.SetCanonical([]byte("App-Version"), []byte(config.AppConfig.AppVersion))
	c.Response().Header.SetCanonical([]byte("Mode"), []byte(config.AppConfig.Mode))
	return c.Next()
}
