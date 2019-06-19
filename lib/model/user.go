package model

import (
	"reflect"

	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/schema"
)

var _ Records = (*Users)(nil)
var _ SearchResult = (*UserSearchResult)(nil)

//Users is array of User
type Users struct {
	Items []schema.User
}

type UserSearchResult struct {
	Items []schema.User
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
