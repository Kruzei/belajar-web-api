package domain

import (
	"encoding/json"
	"time"
)

type Users struct {
	ID        int       `json:"id" gorm:"primary key"`
	Name      string    `json:"name" gorm:"unique"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersRequests struct {
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Age      json.Number `json:"age"`
	Gender   string      `json:"gender"`
}
