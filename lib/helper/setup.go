package helper

import (
	"bufio"
	"fmt"

	"github.com/metaslim/challenge/lib/command"
	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/model"
)

func loadAll() (model.Organizations, model.Users, model.Tickets, error) {
	organizations, errOrganization := loadOrganizations()
	users, errUsers := loadUsers()
	tickets, errTickets := loadTickets()

	var err error
	if errOrganization != nil || errUsers != nil || errTickets != nil {
		err = fmt.Errorf("[%v, %v, %v]", errOrganization, errUsers, errTickets)
	}

	return organizations, users, tickets, err
}

func loadOrganizations() (model.Organizations, error) {
	organizations := model.Organizations{}
	err := organizations.Populate(loader.JSONLoader{
		FileName: "data/organizations.json",
	})

	if err != nil {
		return organizations, err
	}

	return organizations, nil
}

func loadUsers() (model.Users, error) {
	users := model.Users{}
	err := users.Populate(loader.JSONLoader{
		FileName: "data/users.json",
	})

	if err != nil {
		return users, err
	}

	return users, nil
}

func loadTickets() (model.Tickets, error) {
	tickets := model.Tickets{}
	err := tickets.Populate(loader.JSONLoader{
		FileName: "data/tickets.json",
	})

	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

func LoadDecorateParams() (model.DecorateParams, error) {
	organizations, users, tickets, err := loadAll()

	if err != nil {
		return model.DecorateParams{}, err
	}

	return model.DecorateParams{
		Organizations: organizations,
		Users:         users,
		Tickets:       tickets,
	}, nil

}

func PrepareCommand() []command.Action {
	commands := []command.Action{}

	commands = append(commands, &command.Help{})
	commands = append(commands, &command.Search{})

	return commands
}

func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]

	return input
}
