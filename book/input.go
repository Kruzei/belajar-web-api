package book

import "encoding/json"

type BookRequest struct { //Untuk menangkap data kita butuh model/struct
	Title       string      `json:"title" binding:"required"` //Ini kita memberi ketentun required = NOT NULL, number harus angka
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
}
