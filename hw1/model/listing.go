package model

import "time"

type Listing struct {
	ID          int
	Title       string
	Description string
	Price       int
	CreatedAt   time.Time
	Category    string
	Username    string
}
