package model

import (
	"reflect"
	"strconv"

	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/schema"
)

var _ Records = (*Tickets)(nil)
var _ SearchResult = (*TicketSearchResult)(nil)

//Tickets will contains Ticket data source
type Tickets struct {
	BaseRecords
	Items []schema.Ticket
}

//TicketSearchResult will contain Ticket search result
type TicketSearchResult struct {
	BaseSearchResult
	Items []schema.Ticket
}

//Decorate will decorate the search result, in this case it will populate Submitter, Assignee, and Organization properties
func (ticketSearchResult TicketSearchResult) Decorate(dataSet DataSet) {
	for key, _ := range ticketSearchResult.Items {
		for _, user := range dataSet.Users.Items {
			if ticketSearchResult.Items[key].SubmitterID == user.ID {
				matched_user := user
				ticketSearchResult.Items[key].Submitter = &matched_user
			}

			if ticketSearchResult.Items[key].AssigneeID == user.ID {
				matched_user := user
				ticketSearchResult.Items[key].Assignee = &matched_user
			}
		}

		for _, organization := range dataSet.Organizations.Items {
			if ticketSearchResult.Items[key].OrganizationID == organization.ID {
				matched_organization := organization
				ticketSearchResult.Items[key].Organization = &matched_organization
			}
		}
	}
}

//Populate will load data from data source such as json
func (tickets *Tickets) Populate(jsonLoader loader.JSONLoader) error {
	err := load(jsonLoader, &tickets.Items)
	if err != nil {
		return err
	}
	tickets.Size = len(tickets.Items)
	return nil
}

//Search will allow data source to be searched
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
