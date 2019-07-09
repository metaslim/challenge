package command

import (
	"fmt"
	"os"

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
	textcolor.Green(os.Stdout, "\ndescribe\n")
	textcolor.HiGreen(os.Stdout, "\t\tdescribe-organizations")
	fmt.Println("\t\twill return search fields for organizations")
	textcolor.HiGreen(os.Stdout, "\t\tdescribe-users")
	fmt.Println("\t\t\twill return search fields for users")
	textcolor.HiGreen(os.Stdout, "\t\tdescribe-tickets")
	fmt.Println("\t\twill return search fields for tickets")

	textcolor.Magenta(os.Stdout, "table\n")
	textcolor.HiMagenta(os.Stdout, "\t\ttable-organizations:tags=West")
	fmt.Println("\twill return any organizations who has West in their Tags in compact table")
	textcolor.HiMagenta(os.Stdout, "\t\ttable-users:alias=Miss Coffey")
	fmt.Println("\twill return any users whose alias is Miss Coffey in compact table")
	textcolor.HiMagenta(os.Stdout, "\t\ttable-tickets:status=pending")
	fmt.Println("\twill return any tickets with pending status in compact table")

	textcolor.Cyan(os.Stdout, "search\n")
	textcolor.HiCyan(os.Stdout, "\t\tsearch-organizations:tags=West")
	fmt.Println("\twill return any organizations who has West in their Tags in JSON")
	textcolor.HiCyan(os.Stdout, "\t\tsearch-users:alias=Miss Coffey")
	fmt.Println("\twill return any users whose alias is Miss Coffey in JSON")
	textcolor.HiCyan(os.Stdout, "\t\tsearch-tickets:status=pending")
	fmt.Println("\twill return any tickets with pending status in JSON")
}
