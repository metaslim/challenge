package command

import (
	"fmt"
	"regexp"

	"github.com/metaslim/challenge/lib/presenter"

	"github.com/metaslim/challenge/lib/model"
)

var _ Actionable = (*Table)(nil)

//Table will implement Actioanble, the action is searching and displaying result in Table
type Table struct {
	BaseCommand
	regex     *regexp.Regexp
	Presenter presenter.Flushable
}

//Valid will control if the command will trigger Run
func (action *Table) Valid() bool {
	if action.regex == nil {
		regex, _ := regexp.Compile(`^(?i)table-(organizations|users|tickets):(.*)=(.*)$`)
		action.regex = regex
	}

	return action.regex.MatchString(action.Input)
}

//Run will be executed for search command
func (action Table) Run(dataSet model.DataSet) {
	matches := action.regex.FindStringSubmatch(action.Input)

	if matches == nil {
		return
	}

	searchEngine := matches[1]
	field := matches[2]
	searchterm := matches[3]

	var search model.Records
	switch searchEngine {
	case "organizations":
		search = &dataSet.Organizations
	case "users":
		search = &dataSet.Users
	case "tickets":
		search = &dataSet.Tickets
	}

	result := search.Search(field, searchterm)

	if result.GetSize() < 1 {
		fmt.Println("No result found")
		return
	}

	result.Decorate(dataSet)

	fmt.Printf("\n")
	action.Presenter.Flush(result)
}
