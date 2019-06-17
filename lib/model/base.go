package model

import (
	"io/ioutil"

	"github.com/zendesk/challenge/lib/loader"
)

//Model is an interface that require search to be implemented
type Model interface {
	Populate(loader.JSONLoader) error
}

//load return single item
func load(jsonLoader loader.JSONLoader, model interface{}) error {
	fileHandle, err := jsonLoader.GetFileHandle()
	defer fileHandle.Close()

	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(fileHandle)

	if err != nil {
		return err
	}

	parser := jsonLoader.Parse()
	err = parser(byteValue, model)

	if err != nil {
		return err
	}

	return nil
}
