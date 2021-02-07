package todo

import "gorm.io/gorm"

// Todo struct defines the Todo Model
type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      *uint  `gorm:"not null" gorm:"index"`
	// this is a pointer because int == 0,
}
