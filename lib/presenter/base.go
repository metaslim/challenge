package presenter

import "github.com/metaslim/challenge/lib/model"

//Output is an interface that require Flush to be implemented, FLush will be the method to display output
type Output interface {
	Flush(model.SearchResult)
}
