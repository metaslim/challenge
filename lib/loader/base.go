package loader

import (
	"os"
)

//Loader is an interface that require Load to be implemented and return a file handle to be consumed
type Loader interface {
	GetFileHandle() (*os.File, error)
	Parse() func([]byte, interface{}) error
}
