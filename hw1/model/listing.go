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

type ListingInterface interface {
	CreateListing(username string, title string, description string, price int, category string) (int, error)
	DeleteListing(username string, listingID int) error
	GetListing(username string, listingID int) (Listing, error)
	GetCategory(username string, category string) ([]Listing, error)
	GetTopCategory(username string) (string, error)
}
