package cli

import (
	"fmt"
)

type CLIHandler struct {
	userUsecase    *service.UserUsecase
	listingUsecase *service.ListingUsecase
}

func NewCLIHandler(userUsecase *usecase.UserUsecase, listingUsecase *usecase.ListingUsecase) *CLIHandler {
	return &CLIHandler{
		userUsecase:    userUsecase,
		listingUsecase: listingUsecase,
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

