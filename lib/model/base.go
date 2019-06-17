package model

import (
	"io/ioutil"

	"github.com/zendesk/challenge/lib/loader"
)

//Model is an interface that require search to be implemented
type Model interface {
	New(loader.JSONLoader) error
}

//New return model
func load(jsonLoader loader.JSONLoader, model Model) error {
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
	parser(byteValue, model)

	return nil
}
