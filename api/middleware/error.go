package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err == nil {
		return nil
	}

	// switch err {
	// case reflect.TypeOf(err) == expection.ErrBadRequest:
	// 	return ctx.Status(err.(expection.ErrorBase).Status).
	// 		JSON(err.(expection.ErrorBase).Body)
	// }
	return nil
}
