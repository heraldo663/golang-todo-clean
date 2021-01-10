package todo

import (
	"errors"

	"gorm.io/gorm"
)

type ITodoUseCase interface {
	Create(c *CreateDTO, user *uint) (*Todo, error)
	FindByUser(user *uint) ([]*Todo, error)
	FindOne(user *uint, todoID string) (*Todo, error)
	Delete(user *uint, todoID string) error
	Update(user *uint, todoID string, data interface{}) error
	// Check()
}

type todoUseCase struct {
	repository ITodoRepository
}

// NewTodoUseCase -> Creates a new todo UseCase
func NewTodoUseCase(r ITodoRepository) ITodoUseCase {
	return &todoUseCase{repository: r}
}

func (s *todoUseCase) Create(c *CreateDTO, user *uint) (*Todo, error) {
	d := &Todo{
		Task: c.Task,
		User: user,
	}

	todo, err := s.repository.CreateTodo(d)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoUseCase) FindByUser(user *uint) ([]*Todo, error) {
	todos, err := s.repository.FindTodosByUser(user)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *todoUseCase) FindOne(user *uint, todoID string) (*Todo, error) {
	todo, err := s.repository.FindTodoByUser(user, todoID)
	if errors.Is(err, gorm.ErrRecordNotFound) {

	}

	return todo, nil

}

func (s *todoUseCase) Delete(user *uint, todoID string) error {
	return s.repository.DeleteTodo(user, todoID)
}

func (s *todoUseCase) Update(user *uint, todoID string, data interface{}) error {
	return s.repository.UpdateTodo(todoID, user, data)
}
