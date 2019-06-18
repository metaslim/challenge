package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zendesk/challenge/lib/helper"
	"github.com/zendesk/challenge/lib/model"
)

func main() {
	organizations, err := helper.LoadOrganizations()
	if err != nil {
		os.Exit(1)
	}
	sr := organizations.Search("tags", "West")

	users, err := helper.LoadUsers()
	if err != nil {
		os.Exit(1)
	}
	sr = users.Search("tags", "Southview")

	tickets, err := helper.LoadTickets()
	if err != nil {
		os.Exit(1)
	}
	sr = tickets.Search("tags", "Idaho")

	dp := model.DecorateParams{
		Organizations: organizations,
		Users:         users,
		Tickets:       tickets,
	}

	b, err := json.MarshalIndent(sr, "", "  ")
	fmt.Println(string(b))
	sr.Decorate(dp)
	b, err = json.MarshalIndent(sr, "", "  ")
	fmt.Println(string(b))

}
