package routes

import (
	"goravel/app/http/controllers"
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
)

func AuthServiceRoutes(router route.Router) {
	authController := controllers.NewAuthController()

	// Semua route di bawah /auth-svc wajib lewat AuthMiddleware
	router.Prefix("/auth-svc").Middleware(middleware.AuthMiddleware()).Group(func(r route.Router) {
		r.Post("/register", authController.Register)
		r.Post("/login", authController.Login)
	})
}
