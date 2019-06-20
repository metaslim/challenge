package command

import (
	"github.com/metaslim/challenge/lib/model"
)

//Action is interface to give ability to run command
type Action interface {
	Valid() bool
	SetInput(string)
	Run(model.DataSet)
}

//BaseCommand is action base struct that store command from user
type BaseCommand struct {
	Input string
}

//SetInput is will save user commnad to be consumed by object who embed base
func (action *BaseCommand) SetInput(command string) {
	(*action).Input = command
}
