package auth

import (
	"numtostr/gotodo/shared/utils"

	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	Login(ctx *fiber.Ctx) error
	Signup(ctx *fiber.Ctx) error
}

type authController struct {
	useCase IAuthUseCase
}

// NewAuthController -> auth controller factory
func NewAuthController(useCase IAuthUseCase) IAuthController {
	return &authController{useCase: useCase}
}

// Login service logs in a user
func (c *authController) Login(ctx *fiber.Ctx) error {
	b := new(LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	res, err := c.useCase.Login(*b)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	return ctx.JSON(&res)
}

// Signup service creates a user
func (c *authController) Signup(ctx *fiber.Ctx) error {
	b := new(SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	res, err := c.useCase.Signup(*b)

	// If email already exists, return
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	return ctx.JSON(&res)
}
