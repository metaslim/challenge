package command

import (
	"fmt"

	"github.com/metaslim/challenge/lib/model"
)

var _ Action = (*Help)(nil)

//Help is action struct to return help message
type Help struct {
	BaseCommand
	name string
}

//Valid will control if the command will trigger Run
func (action Help) Valid() bool {
	if action.Input == "help" {
		return true
	}

	return false
}

//Run will be executed for help command
func (action Help) Run(dataSet model.DataSet) {
	fmt.Println("\n===================================================================================================================")
	fmt.Println("Sample commands")
	fmt.Println("describe")
	fmt.Println("\t`describe-organizations` will return search fields for organizations")
	fmt.Println("\t`describe-users` will return search fields for users")
	fmt.Println("\t`describe-tickets` will return search fields for tickets")
	fmt.Println("table")
	fmt.Println("\t`table-organizations:tags=West` will return any organizations who has West in their Tags in compact table")
	fmt.Println("\t`table-users:alias=Miss Coffey` will return any users whose alias is Miss Coffey in compact table")
	fmt.Println("\t`table-tickets:status=pending` will return any tickets with pending status in compact table")
	fmt.Println("search")
	fmt.Println("\t`search-organizations:tags=West` will return any organizations who has West in their Tags")
	fmt.Println("\t`search-users:alias=Miss Coffey` will return any users whose alias is Miss Coffey")
	fmt.Println("\t`search-tickets:status=pending` will return any tickets with pending status")
	fmt.Println("=================================================+++++++++++++++++++++++============================================")
}
