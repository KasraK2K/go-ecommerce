package user

import (
	"github.com/gofiber/fiber/v2"

	"app/model"
	"app/pkg/storage/pg"
)

func Routes(router fiber.Router) {
	pg.Gorm.Migrate(&model.User{})

	userGroup := router.Group("/user")

	userGroup.Post("/find", Handler.List)
	userGroup.Post("/create", Handler.Insert)
	userGroup.Patch("/update", Handler.Update)
	userGroup.Delete("/archive", Handler.Archive)
	userGroup.Patch("/restore", Handler.Restore)

	userGroup.Post("/login", Handler.Login)
	userGroup.Post("/forgot", Handler.ForgotPassword)

	userGroup.Post("/upload", Handler.Upload)
}
