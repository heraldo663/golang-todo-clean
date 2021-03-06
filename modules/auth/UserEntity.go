// User struct defines the user
package auth

import (
	"heraldo663/todo/modules/todo"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string      `gorm:"uniqueIndex;not null"`
	Password string      `gorm:"not null"`
	Todos    []todo.Todo `gorm:"foreignKey:User"`
}
