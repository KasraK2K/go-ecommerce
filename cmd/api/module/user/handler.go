package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"app/model"
	"app/pkg"
)

func (h *handler) List(c *fiber.Ctx) error {
	type JsonData struct {
		Filter model.UserFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	results, status, err := Logic.List(payload.Filter)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, results, status)
}

func (h *handler) Insert(c *fiber.Ctx) error {
	type jsonData struct {
		Data model.User `json:"data"`
	}
	var payload jsonData

	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	user := payload.Data
	result, status, err := Logic.Insert(user)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, result, status)
}

func (h *handler) Update(c *fiber.Ctx) error {
	type JsonData struct {
		Filter model.UserFilter `json:"filter"`
		Data   model.UserUpdate `json:"data"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	update := payload.Data
	result, status, err := Logic.Update(filter, update)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, result, status)
}

func (h *handler) Archive(c *fiber.Ctx) error {
	type JsonData struct {
		Filter model.UserFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	result, status, err := Logic.Archive(filter)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, result, status)
}

func (h *handler) Restore(c *fiber.Ctx) error {
	type JsonData struct {
		Filter model.UserFilter `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	filter := payload.Filter
	result, status, err := Logic.Restore(filter)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, result, status)
}

func (h *handler) Login(c *fiber.Ctx) error {
	type jsonData struct {
		Data model.UserLoginPayload `json:"data" validate:"required"`
	}
	var payload jsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	results, status, err := Logic.Login(payload.Data)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, results, status)
}

func (h *handler) ForgotPassword(c *fiber.Ctx) error {
	type forgetPass struct {
		Email string `json:"email" bson:"email" gorm:"type:string;unique;not null;" validate:"required,email,min=6,max=32"`
	}
	type JsonData struct {
		Filter forgetPass `json:"filter"`
	}
	var payload JsonData
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}

	// Validate
	validationError := pkg.Validator(payload)
	if validationError.Errors != nil {
		return pkg.JSON(c, validationError.Errors, http.StatusNotAcceptable)
	}

	results, status, err := Logic.ForgotPassword(payload.Filter.Email)
	if err != nil {
		return pkg.JSON(c, err.Error(), status)
	}

	return pkg.JSON(c, results, status)
}

func (h *handler) Upload(c *fiber.Ctx) error {
	type upload struct {
		M string `json:"m"`
		S string `json:"s"`
	}
	var payload upload
	err := c.BodyParser(&payload)
	if err != nil {
		return pkg.JSON(c, err.Error(), http.StatusBadRequest)
	}
	return pkg.JSON(c, payload, http.StatusOK)
}
