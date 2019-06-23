package loader

import (
	"encoding/json"
	"os"
)

var _ Loadable = (*JSONLoader)(nil)

//JSONLoader will implement Loadable interface and allow loading data from json file
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

//Parse return valid json parser engine
func (jsonLoader JSONLoader) Parse() func([]byte, interface{}) error {
	return json.Unmarshal
}
