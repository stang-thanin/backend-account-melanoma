package route

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterSwaggerRoute(app *fiber.App) {
	router := app.Group("/swagger")
	router.Get("*", swagger.Handler)
}
