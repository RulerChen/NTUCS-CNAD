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
		listingCounter: 1000,
		categoryCount:  make(map[string]int),
	}
}

func (db *MockDB) CreateUser(username string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	key := strings.ToLower(username)
	if _, exists := db.users[key]; exists {
		return errors.New("user already exists")
	}
	db.users[key] = model.User{Username: username}
	return nil
}

func (db *MockDB) GetUser(username string) (model.User, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	key := strings.ToLower(username)
	user, exists := db.users[key]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (db *MockDB) CreateListing(username string, title string, description string, price int, category string) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	user, exists := db.users[strings.ToLower(username)]
	if !exists {
		return 0, errors.New("user not found")
	}
	db.listingCounter++
	listingID := db.listingCounter
	db.listings[listingID] = model.Listing{
		ID:       	 listingID,
		Username: 	 user.Username,
		Title:       title,
		Description: description,
		Price:       price,
		Category:    category,
		CreatedAt:   time.Now(),
	}

	db.categoryCount[category]++
	return listingID, nil
}

func (db *MockDB) DeleteListing(listingID int) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if _, exists := db.listings[listingID]; !exists {
		return errors.New("listing not found")
	}
	delete(db.listings, listingID)
	return nil
}

func (db *MockDB) GetListing(listingID int, username string) (model.Listing, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	listing, exists := db.listings[listingID]
	if !exists {
		return model.Listing{}, false
	}
	if listing.Username != username {
		return model.Listing{}, false
	}
	return listing, true
}

func (db *MockDB) GetCategory(username string, category string) []model.Listing {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var listings []model.Listing
	for _, listing := range db.listings {
		if listing.Username == username && listing.Category == category {
			listings = append(listings, listing)
		}
	}
	return listings
}

func (db *MockDB) GetTopCategory(username string) string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var topCategory string
	var topCount int
	for category, count := range db.categoryCount {
		if count > topCount {
			topCategory = category
			topCount = count
		}
	}
	return topCategory
}
