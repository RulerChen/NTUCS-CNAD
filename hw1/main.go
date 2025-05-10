package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	db "github.com/RulerChen/NTUCS-CNAD/hw1/infra"
	cli "github.com/RulerChen/NTUCS-CNAD/hw1/interface"
	service "github.com/RulerChen/NTUCS-CNAD/hw1/service"
)

func main() {

	mockdb := db.NewMockDB()

	serviceManager := service.NewServiceManager(mockdb)

	handler := cli.NewCLIHandler(serviceManager)

	// scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		handler.ProcessCommand(line)
	}
}
