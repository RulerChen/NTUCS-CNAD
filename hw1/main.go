package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	db "github.com/RulerChen/NTUCS-CNAD/hw1/infra"
)

func main() {

	mockdb := db.NewMockDB()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		processCommand(line)
	}
}