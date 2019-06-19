package command

import (
	"github.com/metaslim/challenge/lib/model"
)

//Loader is an interface that require Load to be implemented and return a file handle to be consumed
type Action interface {
	valid(string) bool
	Run(string, model.DecorateParams)
}
