package command

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/metaslim/challenge/lib/model"
)

var _ Action = (*Describe)(nil)

//Search is action struct to give ability to search data
type Describe struct {
	BaseCommand
	regex *regexp.Regexp
}

//Valid will control if the command will trigger Run
func (action *Describe) Valid() bool {
	if action.regex == nil {
		regex, _ := regexp.Compile(`^(?i)describe-(organizations|users|tickets)$`)
		action.regex = regex
	}

	return action.regex.MatchString(action.Input)
}

//Run will be executed for describe command
func (action Describe) Run(dataSet model.DataSet) {
	matches := action.regex.FindStringSubmatch(action.Input)

	if matches == nil {
		return
	}

	describeEngine := matches[1]

	var val reflect.Value
	switch describeEngine {
	case "organizations":
		val = reflect.ValueOf(&dataSet.Organizations.Items[0]).Elem()
	case "users":
		val = reflect.ValueOf(&dataSet.Users.Items[0]).Elem()
	case "tickets":
		val = reflect.ValueOf(&dataSet.Tickets.Items[0]).Elem()
	}

	regex, _ := regexp.Compile(`^(?i)json:"(\w+)".*?search:"yes"$`)

	fmt.Printf("\n%s can be searched by any fields below\n", describeEngine)
	fmt.Println("================================================")
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := string(typeField.Tag)
		matches := regex.FindStringSubmatch(tag)

		if matches != nil {
			fmt.Printf("\t%s\n", matches[1])
		}
	}
	fmt.Println("================================================")
}
