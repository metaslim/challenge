package presenter

import "github.com/metaslim/challenge/lib/model"

//Flushable is an interface that require Flush to be implemented, Flush will be the method to display output
type Flushable interface {
	Flush(model.SearchResult)
}
