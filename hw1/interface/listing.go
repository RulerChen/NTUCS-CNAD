package handler

import (
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type ListingInterface interface {
	CreateListing(username string, title string, description string, price int, category string) (int, error)
	DeleteListing(listingID int) error
	GetListing(listingID int, username string) (model.Listing, bool)
	GetCategory(username string, category string) []model.Listing
	GetTopCategory(username string)
}
