package presenter

import (
	"encoding/json"
	"fmt"

	"github.com/metaslim/challenge/lib/model"
	"github.com/tidwall/pretty"
)

var _ Output = (*Json)(nil)

//Json is a struct that will allow searh result to be outputted as JSON
type Json struct{}

//Flush will output the data
func (output Json) Flush(data model.SearchResult) {
	byteOutput, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(len(byteOutput))
	fmt.Println(len(pretty.Color(byteOutput, nil)))
	jsonString := string(pretty.Color(byteOutput, nil))

	fmt.Println(jsonString)
}
