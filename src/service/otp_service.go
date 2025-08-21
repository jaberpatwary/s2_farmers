package service

import (
	"app/src/model"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OtpService interface {
	CreateOtp(c *fiber.Ctx) (*model.OtpToken, error)
	GetAll(c *fiber.Ctx, params *validation.QueryOtp) ([]model.OtpToken, error)
	GetByOtpId(c *fiber.Ctx, id string) (*model.OtpToken, error)
	Update(c *fiber.Ctx, req *validation.UpdateOtp, id string) (*model.OtpToken, error)

	//Update(c *fiber.Ctx)
	DeleteOtp(c *fiber.Ctx, id string) error
}

// Define methods for user service

type otpService struct {
	DB *gorm.DB
}

// DB servie init
func NewOtpService(db *gorm.DB) OtpService {
	return &otpService{DB: db}
}

// Create
func (s *otpService) CreateOtp(c *fiber.Ctx) (*model.OtpToken, error) {
	var otp model.OtpToken
	if err := c.BodyParser(&otp); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&otp).Error; err != nil {
		return nil, err
	}

	return &otp, nil
}

// Get All
func (s *otpService) GetAll(c *fiber.Ctx, params *validation.QueryOtp) ([]model.OtpToken, error) {

	var otp []model.OtpToken

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phoneNumber LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&otp).Offset(offset)

	if err := query.Find(&otp).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&otp)
	if result.Error != nil {

		return nil, result.Error
	}

	return otp, result.Error

}

// GetbyId

func (s *otpService) GetByOtpId(c *fiber.Ctx, id string) (*model.OtpToken, error) {
	otp := new(model.OtpToken)

	result := s.DB.WithContext(c.Context()).First(&otp, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return otp, nil
}

// Update OTP token details
func (s *otpService) Update(c *fiber.Ctx, req *validation.UpdateOtp, id string) (*model.OtpToken, error) {
	// --- Validate: at least one field must be provided ---
	if req.OtpCode == "" && req.Purpose == "" && req.IsUsed == nil && req.ExpiresAt == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid Request: nothing to update")
	}

	// --- Prepare update body ---
	updateBody := &model.OtpToken{}

	if req.OtpCode != "" {
		updateBody.OtpCode = req.OtpCode
	}
	if req.Purpose != "" {
		updateBody.Purpose = req.Purpose
	}
	if req.IsUsed != nil {
		updateBody.IsUsed = *req.IsUsed
	}
	if req.ExpiresAt != nil {
		updateBody.ExpiresAt = *req.ExpiresAt
	}

	// --- Update query ---
	result := s.DB.WithContext(c.Context()).Where("id = ?", id).Updates(updateBody)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "OtpToken not found")
	}

	// --- Fetch updated OTP token ---
	otp := new(model.OtpToken)
	if err := s.DB.First(otp, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return otp, nil
}

// delete OTP token
func (s *otpService) DeleteOtp(c *fiber.Ctx, id string) error {
	otp := new(model.OtpToken)

	result := s.DB.WithContext(c.Context()).Delete(otp, "id = ?", id)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "otp not found")
	}

	if result.Error != nil {

		//s.Log.Errorf("Failed to delete user: %+v", result.Error)
	}

	return result.Error

}
