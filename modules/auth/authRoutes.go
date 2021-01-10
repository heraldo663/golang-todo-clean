package auth

import (
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/shared/utils/password"

	"github.com/gofiber/fiber/v2"
)

var (
	jwt        password.IJwt    = password.NewJwt()
	bcrypt     password.IBcrypt = password.NewBcrypt()
	repository IUserRepository  = NewUserRepository(database.DB)
	useCase    IAuthUseCase     = NewAuthUseCase(repository, jwt, bcrypt)
	controller IAuthController  = NewAuthController(useCase)
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", controller.Signup)
	r.Post("/login", controller.Login)
}
