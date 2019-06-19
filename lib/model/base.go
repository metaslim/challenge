package model

import (
	"io/ioutil"

	"github.com/metaslim/challenge/lib/loader"
)

//Records is an interface that require search to be implemented
type Records interface {
	Populate(loader.JSONLoader) error
	Search(string, string) SearchResult
}

type DecorateParams struct {
	Organizations Organizations
	Tickets       Tickets
	Users         Users
}

type SearchResult interface {
	Decorate(DecorateParams)
	GetSize() int
}

type BaseSearchResult struct {
	Size int `json:"number_of_result,omitempty"`
}

func (baseSearchResult BaseSearchResult) GetSize() int {
	return baseSearchResult.Size
}

//load
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
