package route

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/capstone-skincancer-2021/backend-skin-screener-app/src/controller"
)

func AuthenticateRoute(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v1.Post("/signin", controller.Signin)                 // /api/v1/signin
	v1.Post("/auth", controller.Auth)                     // /api/v1/auth
	v1.Get("/users/:username/reset", controller.GetReset) // /api/v1/users/:username/reset
	v1.Post("/password/reset", controller.ResetPassword)  // /api/v1/password/reset?resetToken={resetToken}
	v1.Get("/users/me", controller.Me)                    // /api/v1/users/me
	// v1.Get("/users")                                   // /api/v1/users
	// v1.Get("//users/{username}")                       // /api/v1/users/{username}
}
