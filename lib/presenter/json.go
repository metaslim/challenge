package presenter

import (
	"encoding/json"
	"fmt"

	"github.com/metaslim/challenge/lib/model"
)

var _ Output = (*Json)(nil)

type Json struct{}

func (output Json) Flush(data model.SearchResult) {
	byteOutput, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(byteOutput))
}
