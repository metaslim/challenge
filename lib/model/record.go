package model

import (
	"github.com/metaslim/challenge/lib/loader"
)

//Records is an interface that will give ability to be searched and populated
type Records interface {
	Populate(loader.JSONLoader) error
	Search(string, string) SearchResult
}

//BaseRecords is Records base struct that store number of records
type BaseRecords struct {
	Size int
}

//GetSize will return record size
func (baseRecords BaseRecords) GetSize() int {
	return baseRecords.Size
}
