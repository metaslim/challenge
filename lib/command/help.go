package command

import (
	"fmt"

	"github.com/metaslim/challenge/lib/model"
)

var _ Action = (*Help)(nil)

type Help struct {
	Base
	name string
}

func (action Help) Valid() bool {
	if action.Input == "help" {
		return true
	}

	return false
}

func (action Help) Run(params model.DecorateParams) {
	fmt.Println("Sample commands")
	fmt.Println("`search-organizations:tags=West` will return for organizations who has West in thier Tags")
	fmt.Println("`search-users:alias=Miss Coffey` will return for users whose alias is Miss Coffey")
	fmt.Println("`search-tickets:status=pending` will return for tickets are pending status")
}
