package model

//Countable is an interface that will allow object to be counted
type Countable interface {
	GetSize() int
}

//Decoratable is an interface that will allow object to be decorated
type Decoratable interface {
	Decorate(DataSet)
}

//SearchResult is an interface that will allow object to be decorated
type SearchResult interface {
	Decoratable
	Countable
}

//BaseSearchResult is SearchResult base struct that store number of search result records
type BaseSearchResult struct {
	Size int `json:"number_of_result,omitempty"`
}

//GetSize will return search result record size
func (baseSearchResult BaseSearchResult) GetSize() int {
	return baseSearchResult.Size
}
