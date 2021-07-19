package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patrickcurl/gravel/app/controllers"
	apiControllers "github.com/patrickcurl/gravel/app/controllers/api"
	"github.com/patrickcurl/gravel/app/middleware"
)

func WebAuthRoutes(App fiber.Router) {
	App.Get("/login",
		middleware.RedirectToHomePageOnLogin,
		controllers.LoginGet,
	)
	App.Post("/do/login",
		middleware.ValidateLoginPost,
		controllers.LoginPost,
	)
	App.Post("/do/logout",
		controllers.LogoutPost,
	)

	App.Get("/register", middleware.RedirectToHomePageOnLogin, controllers.RegisterGet)
	App.Post("/do/register",
		middleware.RedirectToHomePageOnLogin,
		middleware.ValidateRegisterPost,
		controllers.RegisterPost,
	)

	App.Get("/reset-password",
		middleware.ValidatePasswordReset,
		controllers.PasswordReset,
	)
	App.Post("/do/reset-password",
		controllers.RequestPasswordResetPost,
	)
	App.Get("/request-password-reset", middleware.RedirectToHomePageOnLogin, controllers.RequestPasswordReset)
	App.Post("/do/password-reset/:token",
		middleware.RedirectToHomePageOnLogin,
		middleware.ValidatePasswordResetPost,
		middleware.ValidateRegisterPost,
		controllers.PasswordResetPost)
	App.Get("/resend/confirm", controllers.ResendConfirmEmail)
	App.Get("/do/verify-email",
		middleware.ValidateConfirmToken,
		controllers.VerifyRegisteredEmail,
	)
}

func AuthRoutes(App fiber.Router) {
	App.Post("/do/login",
		middleware.ValidateApiLoginPost,
		apiControllers.ApiLoginPost,
	)
	App.Post("/me", controllers.Me)
	App.Post("/do/logout", controllers.LogoutPost)

	App.Get("/register", middleware.RedirectToHomePageOnLogin, controllers.RegisterGet)
	App.Post("/do/register",
		middleware.ValidateApiRegisterPost,
		apiControllers.ApiRegisterPost,
	)

	App.Get("/reset-password",
		middleware.ValidatePasswordReset,
		controllers.PasswordReset,
	)
	App.Post("/do/reset-password",
		controllers.RequestPasswordResetPost,
	)
	App.Get("/request-password-reset", middleware.RedirectToHomePageOnLogin, controllers.RequestPasswordReset)
	App.Post("/do/password-reset/:token",
		middleware.RedirectToHomePageOnLogin,
		middleware.ValidatePasswordResetPost,
		middleware.ValidateRegisterPost,
		controllers.PasswordResetPost)
	App.Get("/resend/confirm", controllers.ResendConfirmEmail)
	App.Get("/do/verify-email",
		middleware.ValidateConfirmToken,
		controllers.VerifyRegisteredEmail,
	)

}
