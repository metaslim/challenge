package presenter

import (
	"encoding/json"
	"fmt"

	"github.com/metaslim/challenge/lib/model"
	"github.com/tidwall/pretty"
)

var _ Flushable = (*Json)(nil)

//Json is a struct that will output searh result as Color JSON
type Json struct{}

//Flush will output the data
func (output Json) Flush(data model.SearchResult) {
	byteOutput, _ := json.MarshalIndent(data, "", "  ")
	jsonString := string(pretty.Color(byteOutput, nil))
	fmt.Println(jsonString)
}
