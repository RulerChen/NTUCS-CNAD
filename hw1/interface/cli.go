package cli

import (
	"fmt"
	"strconv"

	"github.com/RulerChen/NTUCS-CNAD/hw1/service"
)

type CLIHandler struct {
	userService    *service.UserService
	listingService *service.ListingService
}

func NewCLIHandler(userService *service.UserService, listingService *service.ListingService) *CLIHandler {
	return &CLIHandler{
		userService:    userService,
		listingService: listingService,
	}
}

func (h *CLIHandler) ProcessCommand(line string) {
	tokens := parseTokens(line)
	if len(tokens) == 0 {
		return
	}
	command := tokens[0]
	switch command {
	case "REGISTER":
		h.handleRegister(tokens)
	case "CREATE_LISTING":
		h.handleCreateListing(tokens)
	case "DELETE_LISTING":
		h.handleDeleteListing(tokens)
	case "GET_LISTING":
		h.handleGetListing(tokens)
	case "GET_CATEGORY":
		h.handleGetCategory(tokens)
	case "GET_TOP_CATEGORY":
		h.handleGetTopCategory(tokens)
	default:
		fmt.Println("Error - invalid command")
	}
}

func parseTokens(input string) []string {
	var tokens []string
	var current string
	inQuote := false
	var quoteChar byte

	for i := 0; i < len(input); i++ {
		c := input[i]
		if inQuote {
			if c == quoteChar {
				inQuote = false
				tokens = append(tokens, current)
				current = ""
			} else {
				current += string(c)
			}
		} else {
			if c == '\'' || c == '"' {
				inQuote = true
				quoteChar = c
			} else if c == ' ' {
				if len(current) > 0 {
					tokens = append(tokens, current)
					current = ""
				}
			} else {
				current += string(c)
			}
		}
	}
	if len(current) > 0 {
		tokens = append(tokens, current)
	}
	return tokens
}

func (h *CLIHandler) handleRegister(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Error - invalid parameters")
		return
	}
	username := tokens[1]

	err := h.userService.CreateUser(username)
	if err != nil {
		fmt.Println("Error - ", err.Error())
	} else {
		fmt.Println("Success")
	}
}

func (h *CLIHandler) handleCreateListing(tokens []string) {
	if len(tokens) != 6 {
		fmt.Println("Error - invalid parameters")
		return
	}
	username := tokens[1]
	title := tokens[2]
	description := tokens[3]
	price, err := strconv.Atoi(tokens[4])
	if err != nil {
		fmt.Println("Error - invalid price")
		return
	}
	category := tokens[5]

	id, err := h.listingService.CreateListing(username, title, description, price, category)
	if err != nil {
		fmt.Println("Error - ", err.Error())
	} else {
		fmt.Println(id)
	}
}

func (h *CLIHandler) handleDeleteListing(tokens []string) {
	if len(tokens) != 3 {
		fmt.Println("Error - invalid parameters")
		return
	}
	username := tokens[1]
	listingID, err := strconv.Atoi(tokens[2])
	if err != nil {
		fmt.Println("Error - invalid listing id")
		return
	}

	err = h.listingService.DeleteListing(username, listingID)
	if err != nil {
		fmt.Println("Error - ", err.Error())
	} else {
		fmt.Println("Success")
	}
}

func (h *CLIHandler) handleGetListing(tokens []string) {
	if len(tokens) != 3 {
		fmt.Println("Error - invalid parameters")
		return
	}
	username := tokens[1]
	listingID, err := strconv.Atoi(tokens[2])
	if err != nil {
		fmt.Println("Error - invalid listing id")
		return
	}
	listing, err := h.listingService.GetListing(username, listingID)
	if err != nil {
		fmt.Println("Error - ", err.Error())
	} else {
		timeStr := listing.CreatedAt.Format("2006-01-02 15:04:05")
		fmt.Printf("%s|%s|%d|%s|%s|%s\n", listing.Title, listing.Description, listing.Price, timeStr, listing.Category, listing.Username)
	}
}

func (h *CLIHandler) handleGetCategory(tokens []string) {
	if len(tokens) != 3 {
		fmt.Println("Error - invalid parameters")
		return
	}
	username := tokens[1]
	category := tokens[2]
	listings, err := h.listingService.GetCategory(username, category)
	if err != nil {
		fmt.Println("Error - ", err.Error())
	} else {
		for _, listing := range listings {
			timeStr := listing.CreatedAt.Format("2006-01-02 15:04:05")
			fmt.Printf("%s|%s|%d|%s|%s|%s\n", listing.Title, listing.Description, listing.Price, timeStr, listing.Category, listing.Username)
		}
	}
}

func (h *CLIHandler) handleGetTopCategory(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Error - invalid parameters")
		return
	}
	username := tokens[1]
	category, err := h.listingService.GetTopCategory(username)
	if err != nil {
		fmt.Println("Error - ", err.Error())
	} else {
		fmt.Println(category)
	}
}
