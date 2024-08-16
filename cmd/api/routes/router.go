package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"app/cmd/api/module/users"
	"app/config"
	"app/pkg"
)

func Routes(app *fiber.App) {

	app.All("/_health", health)
	app.All("/_metrics", monitor.New(monitor.Config{Title: "Default Metrics Page"}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	users.Routes(v1)

	// Handle other routes
	app.Use("*", func(c *fiber.Ctx) error {
		return pkg.JSON(c, "This route is not exist", http.StatusNotFound)
	})
}

func health(c *fiber.Ctx) error {
	var domain string
	if config.AppConfig.Mode == "production" {
		domain = config.AppConfig.ServerDomain
	} else {
		domain = fmt.Sprintf("%s:%s", config.AppConfig.ServerDomain, config.AppConfig.Port)
	}

	result := struct {
		ServerLoadCheck  string `json:"server_load_check"`
		ServerReadyCheck string `json:"server_ready_check"`
	}{
		ServerLoadCheck:  fmt.Sprintf("%s/livez", domain),
		ServerReadyCheck: fmt.Sprintf("%s/readyz", domain),
	}

	return pkg.JSON(c, result, http.StatusOK)
}
