package domain

import (
	"encoding/json"
	"time"
)

type Users struct {
	Name      string    `json:"name" gorm:"primary key"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

type UsersRequests struct {
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Age      json.Number `json:"age"`
	Gender   string      `json:"gender"`
}
