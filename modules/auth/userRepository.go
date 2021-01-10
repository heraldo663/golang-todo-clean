package auth

import (
	"numtostr/gotodo/modules/todo"

	"gorm.io/gorm"
)

// User struct defines the user
type User struct {
	gorm.Model
	Name     string
	Email    string      `gorm:"uniqueIndex;not null"`
	Password string      `gorm:"not null"`
	Todos    []todo.Todo `gorm:"foreignKey:User"`
}

// IUserRepository -> user repository interface
type IUserRepository interface {
	CreateUser(user *User) (*User, error)
	FindUser(conds ...interface{}) (*User, error)
	FindUserByEmail(email string) (*User, error)
}

// NewUserRepository -> user repository factory
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *gorm.DB
}

// CreateUser create a user entry in the user's table
func (r *userRepository) CreateUser(user *User) (*User, error) {
	tx := r.db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

// FindUser searches the user's table with the condition given
func (r *userRepository) FindUser(conds ...interface{}) (*User, error) {
	var dest *User

	tx := r.db.Model(&User{}).Take(dest, conds...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return dest, nil
}

// FindUserByEmail searches the user's table with the email given
func (r *userRepository) FindUserByEmail(email string) (*User, error) {
	var dest *User

	dest, err := r.FindUser("email = ?", email)

	if err != nil {
		return nil, err
	}

	return dest, nil
}
