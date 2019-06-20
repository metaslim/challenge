package setup

import (
	"fmt"
	"sync"

	"github.com/metaslim/challenge/lib/command"
	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/model"
)

func loadAll() (model.Organizations, model.Users, model.Tickets, error) {
	organizationsChannel := make(chan model.Organizations, 1)
	usersChannel := make(chan model.Users, 1)
	ticketsChannel := make(chan model.Tickets, 1)

	errOrganizationsChannel := make(chan error, 1)
	errUsersChannel := make(chan error, 1)
	errTicketsChannel := make(chan error, 1)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go loadOrganizations(organizationsChannel, errOrganizationsChannel, wg)
	organizations := <-organizationsChannel
	errOrganization := <-errOrganizationsChannel

	wg.Add(1)
	go loadUsers(usersChannel, errUsersChannel, wg)
	users := <-usersChannel
	errUsers := <-errUsersChannel

	wg.Add(1)
	go loadTickets(ticketsChannel, errTicketsChannel, wg)
	tickets := <-ticketsChannel
	errTickets := <-errTicketsChannel

	wg.Wait()

	var err error
	if errOrganization != nil || errUsers != nil || errTickets != nil {
		err = fmt.Errorf("[%v, %v, %v]", errOrganization, errUsers, errTickets)
	}

	fmt.Println("abc")
	return organizations, users, tickets, err
}

func loadOrganizations(outChannel chan model.Organizations, errChannel chan error, wg *sync.WaitGroup) {
	defer close(errChannel)
	defer close(outChannel)
	defer wg.Done()

	organizations := model.Organizations{}

	err := organizations.Populate(loader.JSONLoader{
		FileName: "data/organizations.json",
	})

	if err != nil {
		errChannel <- err
		return
	}

	outChannel <- organizations
}

func loadUsers(outChannel chan model.Users, errChannel chan error, wg *sync.WaitGroup) {
	defer close(errChannel)
	defer close(outChannel)
	defer wg.Done()

	users := model.Users{}

	err := users.Populate(loader.JSONLoader{
		FileName: "data/users.json",
	})

	if err != nil {
		errChannel <- err
		return
	}

	outChannel <- users
}

func loadTickets(outChannel chan model.Tickets, errChannel chan error, wg *sync.WaitGroup) {
	defer close(errChannel)
	defer close(outChannel)
	defer wg.Done()

	tickets := model.Tickets{}

	err := tickets.Populate(loader.JSONLoader{
		FileName: "data/tickets.json",
	})

	if err != nil {
		errChannel <- err
		return
	}

	outChannel <- tickets
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
