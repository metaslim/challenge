package setup

import (
	"fmt"
	"sync"

	"github.com/metaslim/challenge/lib/command"
	"github.com/metaslim/challenge/lib/config"
	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/model"
	"github.com/metaslim/challenge/lib/presenter"
)

//loadAll is a helper method to load all the data sources, each data will be loaded concurrently
func loadAll(appConfig config.Config) (model.Organizations, model.Users, model.Tickets, error) {
	organizationsChannel := make(chan model.Organizations, 1)
	usersChannel := make(chan model.Users, 1)
	ticketsChannel := make(chan model.Tickets, 1)

	errOrganizationsChannel := make(chan error, 1)
	errUsersChannel := make(chan error, 1)
	errTicketsChannel := make(chan error, 1)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go loadOrganizations(appConfig.Data.Organization, organizationsChannel, errOrganizationsChannel, wg)
	organizations := <-organizationsChannel
	errOrganization := <-errOrganizationsChannel

	wg.Add(1)
	go loadUsers(appConfig.Data.User, usersChannel, errUsersChannel, wg)
	users := <-usersChannel
	errUsers := <-errUsersChannel

	wg.Add(1)
	go loadTickets(appConfig.Data.Ticket, ticketsChannel, errTicketsChannel, wg)
	tickets := <-ticketsChannel
	errTickets := <-errTicketsChannel

	wg.Wait()

	var err error
	if errOrganization != nil || errUsers != nil || errTickets != nil {
		err = fmt.Errorf("[%v, %v, %v]", errOrganization, errUsers, errTickets)
	}

	return organizations, users, tickets, err
}

//loadOrganizations is a helper method to load Organizations
func loadOrganizations(jsonFile string, outChannel chan model.Organizations, errChannel chan error, wg *sync.WaitGroup) {
	defer close(errChannel)
	defer close(outChannel)
	defer wg.Done()

	organizations := model.Organizations{}

	err := organizations.Populate(loader.JSONLoader{
		FileName: jsonFile,
	})

	if err != nil {
		errChannel <- err
		return
	}

	outChannel <- organizations
}

//loadUsers is a helper method to load Users
func loadUsers(jsonFile string, outChannel chan model.Users, errChannel chan error, wg *sync.WaitGroup) {
	defer close(errChannel)
	defer close(outChannel)
	defer wg.Done()

	users := model.Users{}

	err := users.Populate(loader.JSONLoader{
		FileName: jsonFile,
	})

	if err != nil {
		errChannel <- err
		return
	}

	outChannel <- users
}

//loadTickets is a helper method to load Tickets
func loadTickets(jsonFile string, outChannel chan model.Tickets, errChannel chan error, wg *sync.WaitGroup) {
	defer close(errChannel)
	defer close(outChannel)
	defer wg.Done()

	tickets := model.Tickets{}

	err := tickets.Populate(loader.JSONLoader{
		FileName: jsonFile,
	})

	if err != nil {
		errChannel <- err
		return
	}

	outChannel <- tickets
}

//LoadDataSet is a helper method to merge all the data into one set to be consumed
func LoadDataSet(appConfig config.Config) (model.DataSet, error) {
	organizations, users, tickets, err := loadAll(appConfig)

	if err != nil {
		return model.DataSet{}, err
	}

	return model.DataSet{
		Organizations: organizations,
		Users:         users,
		Tickets:       tickets,
	}, nil

}

//PrepareCommand is a helper method to set up all commands for user to be inputted
func PrepareCommand() []command.Actionable {
	commands := []command.Actionable{}

	commands = append(commands, &command.Help{})
	commands = append(commands, &command.Describe{})
	commands = append(commands, &command.Search{Presenter: presenter.Json{}})
	commands = append(commands, &command.Table{Presenter: presenter.Table{}})

	return commands
}
