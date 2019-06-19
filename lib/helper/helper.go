package helper

import (
	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/model"
)

func LoadOrganizations() (model.Organizations, error) {
	organizations := model.Organizations{}
	err := organizations.Populate(loader.JSONLoader{
		FileName: "data/organizations.json",
	})

	if err != nil {
		return organizations, err
	}

	return organizations, nil
}

func LoadUsers() (model.Users, error) {
	users := model.Users{}
	err := users.Populate(loader.JSONLoader{
		FileName: "data/users.json",
	})

	if err != nil {
		return users, err
	}

	return users, nil
}

func LoadTickets() (model.Tickets, error) {
	tickets := model.Tickets{}
	err := tickets.Populate(loader.JSONLoader{
		FileName: "data/tickets.json",
	})

	if err != nil {
		return tickets, err
	}

	return tickets, nil
}
