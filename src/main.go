package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/capstone-skincancer-2021/backend-skin-screener-app/src/config"
	_ "gitlab.com/capstone-skincancer-2021/backend-skin-screener-app/src/docs"
	"gitlab.com/capstone-skincancer-2021/backend-skin-screener-app/src/route"
)

// @title Skin Screener Backend Application
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config.Init()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	route.RegisterSwaggerRoute(app)
	route.RegisterHealthRoute(app)

	route.AuthenticateRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	config := config.GetServerConfig()
	err := app.Listen(fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		log.Printf("Application cannot listen at port %d", config.Server.Port)
		return
	}

}
