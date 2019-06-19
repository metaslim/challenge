package command

import (
	"fmt"

	"github.com/metaslim/challenge/lib/model"
)

var _ Action = (*Help)(nil)

type Help struct {
	Description string
}

func (help Help) valid(command string) bool {
	if command == "help" {
		return true
	}
	return false
}

func (help Help) Run(command string, params model.DecorateParams) {
	if help.valid(command) {
		fmt.Println("This is a help")
	}
}
