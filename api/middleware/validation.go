package middleware

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Error string `json:"error"`
}

var Validator = validator.New()

func ValidationBody[T any](c *fiber.Ctx) error {
	var errors []*ErrorMessage
	body := new(T)
	err := c.BodyParser(body)
	if err != nil {
		log.Println(err)
		message := ErrorMessage{
			Field: "body",
			Tag:   "parse",
			Error: "body parse error",
		}
		return c.Status(fiber.StatusBadRequest).
			JSON(message)
	}

	err = Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ErrorMessage
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Error = err.Error()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).
			JSON(errors)
	}
	return c.Next()
}
