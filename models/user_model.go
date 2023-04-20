package models

import "gorm.io/gorm"

// User model that has unique id as primary key, first name, last name, unique username, password and role id
type User struct {
	gorm.Model
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Username  string `json:"username" gorm:"uniqueIndex" example:"johndoe"`
	Password  string `json:"password,omitempty" example:"Password123!"`
	RoleID    int    `json:"role" example:"1"`
}

type UserDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Username  string `json:"username" example:"johndoe"`
	RoleID    int    `json:"role" example:"1"`
}

type Login struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"Password123!"`
}

// TableName overrides the table name used by User to `go-warehouse.users`
func (User) TableName() string {
	return "go-warehouse.users"
}
