package models

// Role model that has unique id as primary key, name and users
type Role struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignkey:role_id"`
}

// RoleDTO model that has unique id as primary key, name and users
type RoleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// TableName returns the name of the table
func (Role) TableName() string {
	return "go-warehouse.roles"
}
