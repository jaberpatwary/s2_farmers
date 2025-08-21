package controller

import (
	"app/src/response"
	"app/src/service"
	"app/src/validation"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	_UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{

		_UserService: userService,
	}
}

// Create User
func (u *UserController) CreateUser(c *fiber.Ctx) error {

	user, err := u._UserService.CreateUser(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// Get All users
func (u *UserController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryUser{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	users, err := u._UserService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

// Get By UserId
func (u *UserController) GetByUserId(c *fiber.Ctx) error {

	UserId := c.Params("userId")

	if _, err := uuid.Parse(UserId); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	user, err := u._UserService.GetByUserId(c, UserId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

// Update User
func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	req := new(validation.UpdateUser2)
	userID := c.Params("userId")
	if _, err := uuid.Parse(userID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := u._UserService.Update(c, req, userID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// deleteUser
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("userId")

	if _, err := uuid.Parse(userID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}
	if err := u._UserService.DeleteUser(c, userID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(response.Common{
			Code:    fiber.StatusOK,
			Status:  "success",
			Message: "Delete user successfully",
		})

}
