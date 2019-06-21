package model

import (
	"io/ioutil"

	"github.com/metaslim/challenge/lib/loader"
)

//load is a method that will be called by populate
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
