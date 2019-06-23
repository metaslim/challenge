package loader

import (
	"os"
)

//Parseable is an interface that will give ability to parse data from any I/O
type Parseable interface {
	Parse() func([]byte, interface{}) error
}

//Loadable is an interface that will give ability to load and parse data from any I/O
type Loadable interface {
	GetFileHandle() (*os.File, error)
	Parseable
}
