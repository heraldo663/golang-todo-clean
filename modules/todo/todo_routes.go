package todo

import (
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/shared/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

var (
	repository ITodoRepository = NewTodoRepository(database.DB)
	service    ITodoService    = NewTodoService(repository)
	controller ITodoController = NewTodoController(service)
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
