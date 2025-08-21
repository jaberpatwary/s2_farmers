package service

import (
	"app/src/model"
	"app/src/utils"
	"app/src/validation"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(c *fiber.Ctx) (*model.User, error)
	GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]model.User, error)
	GetByUserId(c *fiber.Ctx, id string) (*model.User, error)
	Update(c *fiber.Ctx, req *validation.UpdateUser2, id string) (*model.User, error)
	DeleteUser(c *fiber.Ctx, id string) error
}

// Define methods for user service

type userService struct {
	DB *gorm.DB
}

// DB servie init
func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}

// Create
func (s *userService) CreateUser(c *fiber.Ctx) (*model.User, error) {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Get All
func (s *userService) GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]model.User, error) {

	var users []model.User

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phoneNumber LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&users).Offset(offset)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&users)
	if result.Error != nil {

		return nil, result.Error
	}

	return users, result.Error

}

// GetbyUserId

func (s *userService) GetByUserId(c *fiber.Ctx, id string) (*model.User, error) {
	user := new(model.User)

	result := s.DB.WithContext(c.Context()).First(&user, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update user details
func (s *userService) Update(c *fiber.Ctx, req *validation.UpdateUser2, id string) (*model.User, error) {
	if req.PhoneNumber == "" && req.FullName == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}
	if req.Password == "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		req.Password = hashedPassword
	}

	updateBody := &model.User{

		PhoneNumber: req.PhoneNumber,
		FullName:    req.FullName,
	}
	result := s.DB.WithContext(c.Context()).Where("id = ?", id).Updates(updateBody)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return nil, fiber.NewError(fiber.StatusConflict, "Phone number already exists")
	}

	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusConflict, "User Not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	user, err := s.GetByUserId(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

// DeleteUser
func (s *userService) DeleteUser(c *fiber.Ctx, id string) error {
	user := new(model.User)

	result := s.DB.WithContext(c.Context()).Delete(user, "id = ?", id)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if result.Error != nil {

		//s.Log.Errorf("Failed to delete user: %+v", result.Error)
	}

	return result.Error

}
