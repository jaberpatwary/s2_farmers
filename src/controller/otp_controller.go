package controller

import (
	"app/src/response"
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OtpController struct {
	_OtpService service.OtpService
}

func NewOtpController(otpService service.OtpService) *OtpController {
	return &OtpController{

		_OtpService: otpService,
	}
}

// Create otp
func (u *OtpController) CreateOtp(c *fiber.Ctx) error {

	otp, err := u._OtpService.CreateOtp(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(otp)
}

// Get All
func (u *OtpController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryOtp{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	otp, err := u._OtpService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(otp)
}

// Get By OtpId

func (u *OtpController) GetByOtpId(c *fiber.Ctx) error {

	OtpId := c.Params("otpId")

	if _, err := uuid.Parse(OtpId); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	otp, err := u._OtpService.GetByOtpId(c, OtpId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(otp)

}

// Update OTP
func (u *OtpController) UpdateOtp(c *fiber.Ctx) error {
	req := new(validation.UpdateOtp)
	otpID := c.Params("otpId")
	if _, err := uuid.Parse(otpID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid otp ID")
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	otp, err := u._OtpService.Update(c, req, otpID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(otp)
}

// Delete OTP
func (u *OtpController) DeleteOtp(c *fiber.Ctx) error {
	otpID := c.Params("otpId")

	if _, err := uuid.Parse(otpID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid otp ID")
	}
	if err := u._OtpService.DeleteOtp(c, otpID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(response.Common{
			Code:    fiber.StatusOK,
			Status:  "success",
			Message: "Delete otp successfully",
		})

}
