package db

import (
	"errors"
	"strings"
	"sync"
	"time"

	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type MockDB struct {
	users          map[string]model.User
	listings       map[int]model.Listing
	listingCounter int
	categoryCount  map[string]int
	mutex          sync.RWMutex
}

func NewMockDB() *MockDB {
	return &MockDB{
		users:          make(map[string]model.User),
		listings:       make(map[int]model.Listing),
		listingCounter: 100000,
		categoryCount:  make(map[string]int),
	}
}

func (db *MockDB) CreateUser(username string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	key := strings.ToLower(username)
	if _, exists := db.users[key]; exists {
		return errors.New("user already existing")
	}
	db.users[key] = model.User{Username: key}

	return nil
}

func (db *MockDB) GetUser(username string) (model.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	user, exists := db.users[strings.ToLower(username)]
	if !exists {
		return model.User{}, errors.New("unknown user")
	}

	return user, nil
}

func (db *MockDB) CreateListing(username string, title string, description string, price int, category string) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.listingCounter++
	listingID := db.listingCounter
	db.listings[listingID] = model.Listing{
		ID:          listingID,
		Username:    strings.ToLower(username),
		Title:       title,
		Description: description,
		Price:       price,
		Category:    category,
		CreatedAt:   time.Now(),
	}
	db.categoryCount[category]++

	return listingID, nil
}

func (db *MockDB) DeleteListing(username string, listingID int) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.listings[listingID]; !exists {
		return errors.New("listing not found")
	}
	if db.listings[listingID].Username != strings.ToLower(username) {
		return errors.New("listing owner mismatch")
	}
	db.categoryCount[db.listings[listingID].Category]--
	delete(db.listings, listingID)

	return nil
}

func (db *MockDB) GetListing(username string, listingID int) (model.Listing, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	listing, exists := db.listings[listingID]
	if !exists {
		return model.Listing{}, errors.New("not found")
	}

	return listing, nil
}

func (db *MockDB) GetCategory(username string, category string) ([]model.Listing, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	if _, exists := db.categoryCount[category]; !exists {
		return nil, errors.New("category not found")
	}

	var listings []model.Listing
	for _, listing := range db.listings {
		if listing.Category == category {
			listings = append(listings, listing)
		}
	}

	return listings, nil
}

func (db *MockDB) GetTopCategory(username string) (string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var topCategory string
	var topCount int
	for category, count := range db.categoryCount {
		if count >= topCount {
			topCategory = category
			topCount = count
		}
	}
	return topCategory, nil
}
