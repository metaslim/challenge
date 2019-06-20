package command

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/metaslim/challenge/lib/model"
)

var _ Action = (*Help)(nil)

type Search struct {
	Base
	regex *regexp.Regexp
}

func (action *Search) Valid() bool {
	if action.regex == nil {
		regex, _ := regexp.Compile(`^(?i)search-(organizations|users|tickets):(.*)=(.*)$`)
		action.regex = regex
	}

	return action.regex.MatchString(action.Input)
}

func (action Search) Run(params model.DecorateParams) {
	matches := action.regex.FindStringSubmatch(action.Input)

	searchEngine := matches[1]
	field := matches[2]
	searchterm := matches[3]

	var search model.Records
	switch searchEngine {
	case "organizations":
		search = &params.Organizations
	case "users":
		search = &params.Users
	case "tickets":
		search = &params.Tickets
	}

	result := search.Search(field, searchterm)

	if result.GetSize() < 1 {
		fmt.Println("No result found")
		return
	}

	result.Decorate(params)

	byteOutput, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteOutput))
}
