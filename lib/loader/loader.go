package loader

import (
	"encoding/json"
	"os"
)

var _ Loader = (*JSONLoader)(nil)

//JSONLoader is struct that will implement Loader interface
type JSONLoader struct {
	FileName string
}

//GetFileHandle return valid filehandle
func (jsonLoader JSONLoader) GetFileHandle() (*os.File, error) {
	fileHandle, err := os.Open(jsonLoader.FileName)

	if err != nil {
		return nil, err
	}

	return fileHandle, nil
}

//Parse return valid parser engine
func (jsonLoader JSONLoader) Parse() func([]byte, interface{}) error {
	return json.Unmarshal
}
