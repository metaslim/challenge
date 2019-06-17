package main

import (
	"fmt"

	"github.com/zendesk/challenge/lib/loader"

	"github.com/zendesk/challenge/lib/model"
)

func main() {
	org := model.Organizations{}
	err := org.Populate(loader.JSONLoader{
		FileName: "data/organizations.json",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("org")
	fmt.Println(org)

	users := model.Users{}
	err = users.Populate(loader.JSONLoader{
		FileName: "data/users.json",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("users")
	fmt.Println(users)

	tickets := model.Tickets{}
	err = tickets.Populate(loader.JSONLoader{
		FileName: "data/tickets.json",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tickets")
	fmt.Println(tickets)
}
