package model

import (
	"github.com/metaslim/challenge/lib/loader"
)

//Populatable is an interface that will give ability to be populated
type Populatable interface {
	Populate(loader.JSONLoader) error
}

//Searchable is an interface that will give ability to be searched
type Searchable interface {
	Search(string, string) SearchResult
}

//Records is an interface that will give ability to be searched and populated
type Records interface {
	Populatable
	Searchable
	Countable
}

//BaseRecords is Records base struct that store number of records
type BaseRecords struct {
	Size int
}

//GetSize will return record size
func (baseRecords BaseRecords) GetSize() int {
	return baseRecords.Size
}
