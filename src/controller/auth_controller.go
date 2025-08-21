package controller

import (
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthController struct {
	_AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{

		_AuthService: authService,
	}
}

// Create User
func (u *AuthController) CreateAuth(c *fiber.Ctx) error {

	auth, err := u._AuthService.CreateAuth(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(auth)
}

// Get All
func (u *AuthController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryAuth{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	auth, err := u._AuthService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(auth)
}

// Get By AuthId
func (u *AuthController) GetByAuthId(c *fiber.Ctx) error {

	AuthId := c.Params("authId")

	if _, err := uuid.Parse(AuthId); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	auth, err := u._AuthService.GetByAuthId(c, AuthId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(auth)

}

// Update auth details
func (u *AuthController) UpdateAuth(c *fiber.Ctx) error {
	req := new(validation.UpdateAuth2)
	authID := c.Params("authId")
	if _, err := uuid.Parse(authID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid auth ID")
	}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	auth, err := u._AuthService.Update(c, req, authID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(auth)
}
