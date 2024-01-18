package domain

import (
	"encoding/json"
	"time"
)

type Books struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Rating      int       `json:"rating"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
}
