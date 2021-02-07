package todo

import (
	"heraldo663/todo/config/database"
	"heraldo663/todo/shared/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

var (
	repository ITodoRepository = NewTodoRepository(database.DB)
	useCase    ITodoUseCase    = NewTodoUseCase(repository)
	controller ITodoController = NewTodoController(useCase)
)

// TodoRoutes contains all routes relative to /todo
func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)

	r.Post("/create", controller.Create)
	r.Get("/list", controller.GetAll)
	r.Get("/:todoID", controller.Get)
	r.Patch("/:todoID", controller.Update)
	r.Patch("/:todoID/check", controller.Check)
	r.Delete("/:todoID", controller.Delete)
}
