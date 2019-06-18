package model

import (
	"reflect"

	"github.com/zendesk/challenge/lib/loader"
)

var _ Records = (*Tickets)(nil)
var _ SearchResult = (*TicketSearchResult)(nil)

//Tickets is array of Tickets
type Tickets struct {
	Items []Ticket
}

//Ticket is object to represent ticket
type Ticket struct {
	ID             string        `json:"_id"`
	URL            string        `json:"url"`
	ExternalID     string        `json:"external_id"`
	CreatedAt      string        `json:"created_at"`
	TicketType     string        `json:"type"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	Priority       string        `json:"priority"`
	Status         string        `json:"status"`
	SubmitterID    int           `json:"submitter_id"`
	AssigneeID     int           `json:"assignee_id"`
	OrganizationID int           `json:"organization_id"`
	Tags           []string      `json:"tags"`
	HasIncidents   bool          `json:"has_incidents"`
	DueAt          string        `json:"due_at"`
	Via            string        `json:"via"`
	Submitter      *User         `json:"submitter,omitempty"`
	Assignee       *User         `json:"assignee,omitempty"`
	Organization   *Organization `json:"organization,omitempty"`
}

type TicketSearchResult struct {
	Items []Ticket
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
							results.Items = append(results.Items, ticket)
							break
						}
					}
				} else if valueField.Interface() == searchTerm {
					results.Items = append(results.Items, ticket)
				}

				break
			}
		}
	}

	return results
}
