package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"

	"app/pkg"
)

func PullOutToken(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	if len(authorization) > 0 {
		token := authorization[7:]
		payload, err := pkg.Token.ParseToken(token)
		if err != nil {
			return pkg.JSON(c, err.Error(), http.StatusUnauthorized)
		}
		c.Locals("TokenPayload", payload)
	}

	return c.Next()
}

/* -------------------------------------------------------------------------- */
/*                              Route Middleware                              */
/* -------------------------------------------------------------------------- */
func OnlyLoggedUsers(c *fiber.Ctx, userType int) error {
	var payload pkg.Payload
	mapstructure.Decode(c.Locals("TokenPayload"), &payload)

	if payload.UserType == userType {
		return c.Next()
	}

	return pkg.JSON(c, "access denied", http.StatusForbidden)
}

func CheckPermission(c *fiber.Ctx, indexPermission int) error {
	var payload pkg.Payload
	mapstructure.Decode(c.Locals("TokenPayload"), &payload)

	if len(payload.Permission) <= 0 {
		return pkg.JSON(c, "invalid access", http.StatusForbidden)
	}

	if string(payload.Permission[indexPermission]) == "1" {
		return c.Next()
	}

	return pkg.JSON(c, "access denied", http.StatusForbidden)
}
