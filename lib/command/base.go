package command

import (
	"github.com/metaslim/challenge/lib/model"
)

//Actionable is an interface to give ability to run command
type Actionable interface {
	Valid() bool
	SetInput(string)
	Run(model.DataSet)
}

//BaseCommand is an  base struct to save command from user
type BaseCommand struct {
	Input string
}

//SetInput  will save user command to be consumed later
func (action *BaseCommand) SetInput(command string) {
	(*action).Input = command
}
