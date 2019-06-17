package loader

import (
	"encoding/json"
	"os"
)

//Loader is an interface that require Load to be implemented and return a file handle to be consumed
type Loader interface {
	GetFileHandle() (*os.File, error)
	Parse() func([]byte, interface{}) error
}

var _ Loader = (*JSONLoader)(nil)

//JSONLoader is struct that will implement Loader interface
type JSONLoader struct {
	fileName string
}

//Load return valid filehandle
func (jsonLoader JSONLoader) GetFileHandle() (*os.File, error) {
	fileHandle, err := os.Open(jsonLoader.fileName)

	if err != nil {
		return nil, err
	}

	return fileHandle, nil
}

//Parse return valid parser engine
func (jsonLoader JSONLoader) Parse() func([]byte, interface{}) error {
	return json.Unmarshal
}
