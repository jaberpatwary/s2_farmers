package service

import (
	"app/src/model"
	"app/src/utils"
	"app/src/validation"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthService interface {
	CreateAuth(c *fiber.Ctx) (*model.AuthToken, error)
	GetAll(c *fiber.Ctx, params *validation.QueryAuth) ([]model.AuthToken, error)
	GetByAuthId(c *fiber.Ctx, id string) (*model.AuthToken, error)
	Update(c *fiber.Ctx, req *validation.UpdateAuth2, id string) (*model.AuthToken, error)

	//Update(c *fiber.Ctx)
	//Delete(c *fiber.Ctx)
}

// Define methods for user service

type authService struct {
	DB *gorm.DB
}

// DB servie init
func NewAuthService(db *gorm.DB) AuthService {
	return &authService{DB: db}
}

// Create
func (s *authService) CreateAuth(c *fiber.Ctx) (*model.AuthToken, error) {
	var auth model.AuthToken
	if err := c.BodyParser(&auth); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

// Get All
func (s *authService) GetAll(c *fiber.Ctx, params *validation.QueryAuth) ([]model.AuthToken, error) {

	var auth []model.AuthToken

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phoneNumber LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&auth).Offset(offset)

	if err := query.Find(&auth).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&auth)
	if result.Error != nil {

		return nil, result.Error
	}

	return auth, result.Error

}

// GetbyId

func (s *authService) GetByAuthId(c *fiber.Ctx, id string) (*model.AuthToken, error) {
	auth := new(model.AuthToken)

	result := s.DB.WithContext(c.Context()).First(&auth, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return auth, nil
}

// Update auth details
func (s *authService) Update(c *fiber.Ctx, req *validation.UpdateAuth2, id string) (*model.AuthToken, error) {
	if req.PhoneNumber == "" && req.FullName == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	if req.Password != "" { // Fixed: should be != not ==
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		req.Password = hashedPassword
	}

	// First get the AuthToken to find the UserID
	auth := new(model.AuthToken)
	result := s.DB.WithContext(c.Context()).First(&auth, "id = ?", id)
	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Auth token not found")
	}

	// Update the User model, not AuthToken
	updateData := map[string]interface{}{
		"phone_number": req.PhoneNumber,
		"full_name":    req.FullName,
	}

	if req.Password != "" {
		updateData["password"] = req.Password
	}

	result = s.DB.WithContext(c.Context()).
		Model(&model.User{}).
		Where("id = ?", auth.UserID).
		Updates(updateData)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return nil, fiber.NewError(fiber.StatusConflict, "Phone number already exists")
	}

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	// Get updated auth with user data
	updatedAuth, err := s.GetByAuthId(c, id)
	if err != nil {
		return nil, err
	}
	return updatedAuth, nil
}
