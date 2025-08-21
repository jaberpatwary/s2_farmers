package router

import (
	"app/src/config"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {

	UserService := service.NewUserService(db)

	AuthService := service.NewAuthService(db)
	OtpService := service.NewOtpService(db)

	v1 := app.Group("/v1")

	//HealthCheckRoutes(v1, healthCheckService)

	UserRoutes(v1, UserService)
	AuthRoutes(v1, AuthService)
	OtpRoutes(v1, OtpService)
	// TODO: add another routes here...

	if !config.IsProd {
		DocsRoutes(v1)
	}
}
