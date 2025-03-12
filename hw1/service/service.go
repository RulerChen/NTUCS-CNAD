// service/service_manager.go
package service

import (
	db "github.com/RulerChen/NTUCS-CNAD/hw1/infra"
)

type ServiceManager struct {
	UserService    UserService
	ListingService ListingService
}

func NewServiceManager(database db.DB) *ServiceManager {
	userService := NewUserService(database)
	listingService := NewListingService(database, userService)
	return &ServiceManager{
		UserService:    userService,
		ListingService: listingService,
	}
}
