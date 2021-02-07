package todo

import (
	"errors"

	"gorm.io/gorm"
)

// ITodoRepository -> todo repository interface
type ITodoRepository interface {
	CreateTodo(todo *Todo) (*Todo, error)
	FindTodo(dest interface{}, conds ...interface{}) *gorm.DB
	FindTodoByUser(todoIden interface{}, userIden interface{}) (*Todo, error)
	FindTodosByUser(userIden interface{}) ([]*Todo, error)
	DeleteTodo(todoIden interface{}, userIden interface{}) error
	UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) error
}

type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository -> creates a todo repository using GORM
func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &todoRepository{db: db}
}

// CreateTodo create a todo entry in the todo's table
func (r *todoRepository) CreateTodo(todo *Todo) (*Todo, error) {
	tx := r.db.Create(todo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todo, nil
}

// FindTodo finds a todo with given condition
func (r *todoRepository) FindTodo(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.Model(&Todo{}).Take(dest, conds...)
}

// FindTodoByUser finds a todo with given todo and user identifier
func (r *todoRepository) FindTodoByUser(todoIden interface{}, userIden interface{}) (*Todo, error) {
	var dest *Todo

	tx := r.FindTodo(dest, "id = ? AND user = ?", todoIden, userIden)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return dest, nil
}

// FindTodosByUser finds the todos with user's identifier given
func (r *todoRepository) FindTodosByUser(userIden interface{}) ([]*Todo, error) {
	var dest []*Todo

	tx := r.db.Model(&Todo{}).Find(dest, "user = ?", userIden)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return dest, nil
}

// DeleteTodo deletes a todo from todos' table with the given todo and user identifier
func (r *todoRepository) DeleteTodo(todoIden interface{}, userIden interface{}) error {
	tx := r.db.Unscoped().Delete(&Todo{}, "id = ? AND user = ?", todoIden, userIden)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Unable to delete todo")
	}

	return nil
}

// UpdateTodo allows to update the todo with the given todoID and userID
func (r *todoRepository) UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) error {
	tx := r.db.Model(&Todo{}).Where("id = ? AND user = ?", todoIden, userIden).Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
