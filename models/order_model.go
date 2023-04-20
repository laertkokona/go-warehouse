package models

import (
	"gorm.io/gorm"
	"time"
)

// Order model that has unique id as primary key, unique code, submitted date, deadline date, user id and order items
type Order struct {
	gorm.Model
	Code          string      `json:"code,omitempty" gorm:"uniqueIndex;not null"`
	SubmittedDate time.Time   `json:"submittedDate"`
	DeadlineDate  time.Time   `json:"deadlineDate"`
	UserID        int         `json:"user"`
	OrderItems    []OrderItem `json:"orderItems,omitempty"`
}

// ComparableOrder model that has unique id as primary key, unique code, submitted date, deadline date and user id
type ComparableOrder struct {
	ID            uint      `json:"id"`
	Code          string    `json:"code"`
	SubmittedDate time.Time `json:"submittedDate"`
	DeadlineDate  time.Time `json:"deadlineDate"`
	UserID        int       `json:"user"`
}
