package command

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/metaslim/challenge/lib/model"
	"github.com/metaslim/challenge/lib/textcolor"
)

var _ Actionable = (*Help)(nil)

//Help will implement Actioanble, the action is showing help message
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
	color.HiGreen("\ndescribe")
	textcolor.Green("\tdescribe-organizations")
	fmt.Println("\t\twill return search fields for organizations")
	textcolor.Green("\tdescribe-users")
	fmt.Println("\t\t\twill return search fields for users")
	textcolor.Green("\tdescribe-tickets")
	fmt.Println("\t\twill return search fields for tickets")

	color.HiMagenta("table")
	textcolor.Magenta("\ttable-organizations:tags=West")
	fmt.Println("\twill return any organizations who has West in their Tags in compact table")
	textcolor.Magenta("\ttable-users:alias=Miss Coffey")
	fmt.Println("\twill return any users whose alias is Miss Coffey in compact table")
	textcolor.Magenta("\ttable-tickets:status=pending")
	fmt.Println("\twill return any tickets with pending status in compact table")

	color.HiCyan("search")
	textcolor.Cyan("\tsearch-organizations:tags=West")
	fmt.Println("\twill return any organizations who has West in their Tags in JSON")
	textcolor.Cyan("\tsearch-users:alias=Miss Coffey")
	fmt.Println("\twill return any users whose alias is Miss Coffey in JSON")
	textcolor.Cyan("\tsearch-tickets:status=pending")
	fmt.Println("\twill return any tickets with pending status in JSON")
}
