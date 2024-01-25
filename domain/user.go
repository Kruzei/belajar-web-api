package domain

import (
	"encoding/json"
	"time"
)

type Users struct {
	ID        int       `json:"id" gorm:"primary key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersRequests struct {
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Name     string      `json:"name"`
	Age      json.Number `json:"age"`
	Gender   string      `json:"gender"`
}
