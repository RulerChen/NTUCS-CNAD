package db

import (
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type DB interface {
	CreateUser(username string) error
	GetUser(username string) (model.User, error)
	CreateListing(username string, title string, description string, price int, category string) (int, error)
	DeleteListing(username string, listingID int) error
	GetListing(username string, listingID int) (model.Listing, error)
	GetCategory(username string, category string) ([]model.Listing, error)
	GetTopCategory(username string) ([]string, error)
}
