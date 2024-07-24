package middleware

import (
	"crud-golang/exception"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	err := c.Next()
	if err == nil {
		return nil
	}

	return handleErrBase(c, err)
}

func handleErrBase(c *fiber.Ctx, err error) error {
	errBase := err.(*exception.ErrorBase)
	log.Print("ErrorBase", errBase)
	return c.Status(errBase.Status).
		JSON(map[string]any{
			"message": errBase.Message,
			"body":    errBase.Body,
		})
}
