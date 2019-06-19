package main

import (
	"fmt"
	"os"

	"github.com/metaslim/challenge/lib/command"
	"github.com/metaslim/challenge/lib/helper"
	"github.com/metaslim/challenge/lib/model"
)

func main() {
	organizations, err := helper.LoadOrganizations()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	users, err := helper.LoadUsers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tickets, err := helper.LoadTickets()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dp := model.DecorateParams{
		Organizations: organizations,
		Users:         users,
		Tickets:       tickets,
	}

	commands := []command.Action{}
	help := command.Help{}

	commands = append(commands, help)

	var input string
	fmt.Println("Enter Command:")
	fmt.Println("==================================")
	fmt.Println("\t`quit` to exit")
	fmt.Println("\t`help` to get help")
	fmt.Println("==================================")
	fmt.Scanf("%s", &input)

	for input != "quit" {
		for _, command := range commands {
			command.Run(input, dp)
		}
		fmt.Scanf("%s", &input)
	}
}
