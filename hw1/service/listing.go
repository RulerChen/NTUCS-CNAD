package service

import (
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type ListingService struct {
	listingService model.ListingInterface
	userService    model.UserInterface
}

func NewListingService(listingService model.ListingInterface, userService model.UserInterface) *ListingService {
	return &ListingService{
		listingService: listingService,
		userService:    userService,
	}
}

func (ls *ListingService) CreateListing(username string, title string, description string, price int, category string) (int, error) {
	_, err := ls.userService.GetUser(username)
	if err != nil {
		return 0, err
	}

	return ls.listingService.CreateListing(username, title, description, price, category)
}

func (ls *ListingService) DeleteListing(username string, listingID int) error {
	return ls.listingService.DeleteListing(username, listingID)
}

func (ls *ListingService) GetListing(username string, listingID int) (model.Listing, error) {
	_, err := ls.userService.GetUser(username)
	if err != nil {
		return model.Listing{}, err
	}

	return ls.listingService.GetListing(username, listingID)
}

func (ls *ListingService) GetCategory(username string, category string) ([]model.Listing, error) {
	_, err := ls.userService.GetUser(username)
	if err != nil {
		return nil, err
	}

	listings, err := ls.listingService.GetCategory(username, category)
	if err != nil {
		return nil, err
	}

	// sort by create time desc (bubble sort)
	for i := 0; i < len(listings); i++ {
		for j := i + 1; j < len(listings); j++ {
			if listings[i].CreatedAt.Before(listings[j].CreatedAt) {
				listings[i], listings[j] = listings[j], listings[i]
			}
		}
	}

	return listings, nil
}

func (ls *ListingService) GetTopCategory(username string) (string, error) {
	_, err := ls.userService.GetUser(username)
	if err != nil {
		return "", err
	}

	return ls.listingService.GetTopCategory(username)
}
