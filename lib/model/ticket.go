package model

import (
	"reflect"
	"strconv"

	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/schema"
)

var _ Records = (*Tickets)(nil)
var _ SearchResult = (*TicketSearchResult)(nil)

//Tickets is array of Tickets
type Tickets struct {
	Items []schema.Ticket
}

type TicketSearchResult struct {
	Items []schema.Ticket
	BaseSearchResult
}

func (ticketSearchResult TicketSearchResult) Decorate(decorateParams DecorateParams) {
	for key, _ := range ticketSearchResult.Items {
		for _, user := range decorateParams.Users.Items {
			if ticketSearchResult.Items[key].SubmitterID == user.ID {
				matched_user := user
				ticketSearchResult.Items[key].Submitter = &matched_user
			}

			if ticketSearchResult.Items[key].AssigneeID == user.ID {
				matched_user := user
				ticketSearchResult.Items[key].Assignee = &matched_user
			}
		}

		for _, organization := range decorateParams.Organizations.Items {
			if ticketSearchResult.Items[key].OrganizationID == organization.ID {
				matched_organization := organization
				ticketSearchResult.Items[key].Organization = &matched_organization
			}
		}
	}
}

//Populate is
func (tickets *Tickets) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, &tickets.Items)
}

func (tickets *Tickets) Search(searchKey string, searchTerm string) SearchResult {
	results := TicketSearchResult{}

	for _, ticket := range tickets.Items {
		val := reflect.ValueOf(&ticket).Elem()
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			tag := typeField.Tag

			if tag.Get("json") == searchKey {
				if searchKey == "tags" {
					for _, tag := range ticket.Tags {
						if tag == searchTerm {
							results.Size++
							results.Items = append(results.Items, ticket)
							break
						}
					}
				} else {
					var castedSearchTerm interface{}
					var err error
					switch searchKey {
					case "assignee_id", "submitter_id", "organization_id":
						castedSearchTerm, err = strconv.Atoi(searchTerm)
						if err != nil {
							break
						}
					case "has_incidents":
						castedSearchTerm, err = strconv.ParseBool(searchTerm)
						if err != nil {
							break
						}
					default:
						castedSearchTerm = searchTerm

					}
					if valueField.Interface() == castedSearchTerm {
						results.Size++
						results.Items = append(results.Items, ticket)
					}
				}
				break
			}
		}
	}

	return results
}
