package command

import (
	"github.com/metaslim/challenge/lib/model"
)

//Loader is an interface that require Load to be implemented and return a file handle to be consumed
type Action interface {
	Valid() bool
	SetInput(string)
	Run(model.DecorateParams)
}

type Base struct {
	Input string
}

func (action *Base) SetInput(command string) {
	(*action).Input = command
}
