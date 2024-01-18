package book

import "time"

type Books struct {
	ID          int			
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt	time.Time
}