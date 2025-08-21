package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func OtpRoutes(v1 fiber.Router, u service.OtpService) {
	// Initialize the UserController with the UserService
	otpController := controller.NewOtpController(u)
	// Define user-related routes
	otpGroup := v1.Group("/otp")
	otpGroup.Post("/", otpController.CreateOtp)
	otpGroup.Get("/", otpController.GetAll)
	otpGroup.Get("/:otpId", otpController.GetByOtpId)
	otpGroup.Put("/:otpId", otpController.UpdateOtp)
	otpGroup.Delete("/:optId", otpController.DeleteOtp)
}
