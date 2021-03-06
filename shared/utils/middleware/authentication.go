package middleware

import (
	"heraldo663/todo/shared/utils/password"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	jwtImpl password.IJwt = password.NewJwt()
)

// Auth is the authentication middleware
func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	// Spliting the header
	chunks := strings.Split(h, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	user, err := jwtImpl.Verify(chunks[1])

	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("USER", user.ID)

	return c.Next()
}
