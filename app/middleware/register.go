package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/patrickcurl/gravel/app"
	"github.com/patrickcurl/gravel/app/auth"
	"github.com/patrickcurl/gravel/app/models"
)

func ValidateRegisterPost(c *fiber.Ctx) error {
	var register models.RegisterForm
	if err := c.BodyParser(&register); err != nil {
		return app.Http.Flash.WithError(c, fiber.Map{
			"message": err.Error(),
		}).Redirect("/register")
	}

	v := validate.Struct(register)
	if !v.Validate() {
		fmt.Println(v.Errors)
		return app.Http.Flash.WithError(c, fiber.Map{
			"message": v.Errors.One(),
		}).Redirect("/register")
	}
	c.Locals("register", register)
	return c.Next()
}

func ValidateConfirmToken(c *fiber.Ctx) error {
	t := app.Decrypt(c.Query("t"), app.Http.Server.Key)
	user, err := models.GetUserByEmail(t)
	if err != nil {
		return app.Http.Flash.WithError(c, fiber.Map{
			"message": err.Error(),
		}).Redirect("/login")
	}

	if user.EmailVerified {
		return app.Http.Flash.WithError(c, fiber.Map{
			"message": "Email was already validated",
		}).Redirect("/login")
	}
	user.EmailVerified = true
	app.Http.Database.DB.Save(&user)
	auth.Login(c, user.ID, app.Http.Server.Key) //nolint:wsl
	c.Locals("user", user)
	return c.Next()
}
