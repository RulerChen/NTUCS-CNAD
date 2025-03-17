package service

import (
	db "github.com/RulerChen/NTUCS-CNAD/hw1/infra"
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type ListingService interface {
	CreateListing(username, title, description string, price int, category string) (int, error)
	DeleteListing(username string, listingID int) error
	GetListing(username string, listingID int) (model.Listing, error)
	GetCategory(username, category string) ([]model.Listing, error)
	GetTopCategory(username string) ([]string, error)
}

type ListingServiceImpl struct {
	DB          db.DB
	UserService UserService
}

func NewListingService(database db.DB, userService UserService) ListingService {
	return &ListingServiceImpl{
		DB:          database,
		UserService: userService,
	}
}

func (ls *ListingServiceImpl) CreateListing(username string, title string, description string, price int, category string) (int, error) {
	_, err := ls.UserService.GetUser(username)
	if err != nil {
		return 0, err
	}

	return ls.DB.CreateListing(username, title, description, price, category)
}

func (ls *ListingServiceImpl) DeleteListing(username string, listingID int) error {
	return ls.DB.DeleteListing(username, listingID)
}

func (ls *ListingServiceImpl) GetListing(username string, listingID int) (model.Listing, error) {
	_, err := ls.UserService.GetUser(username)
	if err != nil {
		return model.Listing{}, err
	}

	return ls.DB.GetListing(username, listingID)
}

func (ls *ListingServiceImpl) GetCategory(username string, category string) ([]model.Listing, error) {
	_, err := ls.UserService.GetUser(username)
	if err != nil {
		return nil, err
	}

	listings, err := ls.DB.GetCategory(username, category)
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

func (ls *ListingServiceImpl) GetTopCategory(username string) ([]string, error) {
	_, err := ls.UserService.GetUser(username)
	if err != nil {
		return nil, err
	}

	// sort by lexically desc (bubble sort)
	categories, err := ls.DB.GetTopCategory(username)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(categories); i++ {
		for j := i + 1; j < len(categories); j++ {
			if categories[i] > categories[j] {
				categories[i], categories[j] = categories[j], categories[i]
			}
		}
	}

	return categories, nil
}
