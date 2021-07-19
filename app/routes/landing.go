package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrickcurl/gravel/app"
	"github.com/patrickcurl/gravel/app/controllers"
	"github.com/patrickcurl/gravel/app/middleware"
)

func LandingRoutes(web fiber.Router) {
	web.Get("/", controllers.Landing)
	web.Get("/ping", Pong)
	web.Get("/all-routes", AllRoutes)
	web.Get("/do/verify-email", middleware.ValidateConfirmToken, controllers.VerifyRegisteredEmail)
}

func Pong(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func AllRoutes(c *fiber.Ctx) error {
	return c.JSON(app.Http.Server.Stack())
}
