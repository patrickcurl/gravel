package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrickcurl/gravel/app/middleware"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("api").Use(middleware.AuthApi())
	web := app.Group("")
	ApiRoutes(api)
	WebRoutes(web)
}
