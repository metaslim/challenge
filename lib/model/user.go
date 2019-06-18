package model

import (
	"reflect"

	"github.com/zendesk/challenge/lib/loader"
)

var _ Records = (*Users)(nil)
var _ SearchResult = (*UserSearchResult)(nil)

//Users is array of User
type Users struct {
	Items []User
}

//User is object to represent an user
type User struct {
	ID              int           `json:"_id"`
	URL             string        `json:"url"`
	ExternalID      string        `json:"external_id"`
	Name            string        `json:"name"`
	Alias           string        `json:"alias"`
	CreatedAt       string        `json:"created_at"`
	Active          bool          `json:"active"`
	Verified        bool          `json:"verified"`
	Shared          bool          `json:"shared"`
	Locale          string        `json:"locale"`
	Timezone        string        `json:"timezone"`
	LastLoginAt     string        `json:"last_login_at"`
	Email           string        `json:"email"`
	Phone           string        `json:"phone"`
	Signature       string        `json:"signature"`
	OrganizationID  int           `json:"organization_id"`
	Tags            []string      `json:"tags"`
	Suspended       bool          `json:"suspended"`
	Role            string        `json:"role"`
	Organization    *Organization `json:"organization,omitempty"`
	SubmittedTicket []Ticket      `json:"submitted_by,omitempty"`
	AssignedTicket  []Ticket      `json:"assigned_to,omitempty"`
}

type UserSearchResult struct {
	Items []User
}

func (userSearchResult UserSearchResult) Decorate(decorateParams DecorateParams) {
	for key, _ := range userSearchResult.Items {
		for _, organization := range decorateParams.Organizations.Items {
			if userSearchResult.Items[key].OrganizationID == organization.ID {
				matched_organization := organization
				userSearchResult.Items[key].Organization = &matched_organization
			}
		}

		for _, ticket := range decorateParams.Tickets.Items {
			if userSearchResult.Items[key].ID == ticket.SubmitterID {
				userSearchResult.Items[key].SubmittedTicket = append(userSearchResult.Items[key].SubmittedTicket, ticket)
			}

			if userSearchResult.Items[key].ID == ticket.AssigneeID {
				userSearchResult.Items[key].AssignedTicket = append(userSearchResult.Items[key].AssignedTicket, ticket)
			}
		}

	}
}

//Populate is
func (users *Users) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, &users.Items)
}

func (users *Users) Search(searchKey string, searchTerm string) SearchResult {
	results := UserSearchResult{}

	for _, user := range users.Items {
		val := reflect.ValueOf(&user).Elem()
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			tag := typeField.Tag

			if tag.Get("json") == searchKey {
				if searchKey == "tags" {
					for _, tag := range user.Tags {
						if tag == searchTerm {
							results.Items = append(results.Items, user)
							break
						}
					}
				} else if valueField.Interface() == searchTerm {
					results.Items = append(results.Items, user)
				}

				break
			}
		}
	}

	return results
}
