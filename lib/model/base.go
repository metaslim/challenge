package model

import (
	"io/ioutil"

	"github.com/metaslim/challenge/lib/loader"
)

//Records is an interface that will give ability to be searched and populated
type Records interface {
	Populate(loader.JSONLoader) error
	Search(string, string) SearchResult
}

//SearchResult is an interface that will allow object to be decorated
type SearchResult interface {
	Decorate(DataSet)
	GetSize() int
}

//BaseSearchResult is SearchResult base struct that store number of records
type BaseSearchResult struct {
	Size int `json:"number_of_result,omitempty"`
}

//GetSize will return record size
func (baseSearchResult BaseSearchResult) GetSize() int {
	return baseSearchResult.Size
}

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
