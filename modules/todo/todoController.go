package todo

import (
	"errors"
	"heraldo663/todo/shared/types"
	"heraldo663/todo/shared/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ITodoController interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Check(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type todoController struct {
	useCase ITodoUseCase
}

// NewTodoController -> creates todo controller
func NewTodoController(useCase ITodoUseCase) ITodoController {
	return &todoController{useCase: useCase}
}

// CreateTodo is responsible for create todo
func (c *todoController) Create(ctx *fiber.Ctx) error {
	b := new(CreateDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	user := utils.GetUser(ctx)
	todo, err := c.useCase.Create(b, user)

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.JSON(&TodoCreateResponse{
		Todo: &TodoResponse{
			ID:        todo.ID,
			Task:      todo.Task,
			Completed: todo.Completed,
		},
	})
}

// GetTodos returns the todos list
func (c *todoController) GetAll(ctx *fiber.Ctx) error {
	res := []TodoResponse{}
	user := utils.GetUser(ctx)

	todos, err := c.useCase.FindByUser(user)

	for _, todo := range todos {
		todoRes := TodoResponse{
			ID:        todo.ID,
			Task:      todo.Task,
			Completed: todo.Completed,
		}

		res = append(res, todoRes)
	}

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.JSON(&TodosResponse{
		Todos: &res,
	})
}

// GetTodo return a single todo
func (c *todoController) Get(ctx *fiber.Ctx) error {
	todoID := ctx.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	user := utils.GetUser(ctx)

	todo, err := c.useCase.FindOne(user, todoID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(&TodoCreateResponse{})
	}

	d := &TodoResponse{
		ID:        todo.ID,
		Task:      todo.Task,
		Completed: todo.Completed,
	}

	return ctx.JSON(&TodoCreateResponse{
		Todo: d,
	})
}

// DeleteTodo deletes a single todo
func (c *todoController) Delete(ctx *fiber.Ctx) error {
	todoID := ctx.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	user := utils.GetUser(ctx)

	err := c.useCase.Delete(user, todoID)

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.JSON(&types.MsgResponse{
		Message: "Todo successfully deleted",
	})
}

// CheckTodo TODO
func (c *todoController) Check(ctx *fiber.Ctx) error {
	b := new(CheckTodoDTO)
	todoID := ctx.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	user := utils.GetUser(ctx)

	err := c.useCase.Update(user, todoID, map[string]interface{}{"completed": b.Completed})

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}

// UpdateTodoTitle TODO
func (c *todoController) Update(ctx *fiber.Ctx) error {
	b := new(CreateDTO)
	todoID := ctx.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	user := utils.GetUser(ctx)

	err := c.useCase.Update(user, todoID, &Todo{Task: b.Task})
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}
