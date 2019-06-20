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
	fmt.Println("Sample commands")
	fmt.Println("`search-organizations:tags=West` will return for organizations who has West in thier Tags")
	fmt.Println("`search-users:alias=Miss Coffey` will return for users whose alias is Miss Coffey")
	fmt.Println("`search-tickets:status=pending` will return for tickets are pending status")
}
