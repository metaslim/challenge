package model

import (
	"reflect"
	"strconv"

	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/schema"
)

var _ Records = (*Users)(nil)
var _ SearchResult = (*UserSearchResult)(nil)

//Users will contain records loaded from data source for users
type Users struct {
	BaseRecords
	Items []schema.User
}

//UserSearchResult will contain User search result
type UserSearchResult struct {
	BaseSearchResult
	Items []schema.User `json:"search_result"`
}

//Decorate will decorate the search result, in this case it will populate Organization, SubmittedTicket, and AssignedTicket properties
func (userSearchResult UserSearchResult) Decorate(dataSet DataSet) {
	for key, _ := range userSearchResult.Items {
		for _, organization := range dataSet.Organizations.Items {
			if userSearchResult.Items[key].OrganizationID == organization.ID {
				userSearchResult.Items[key].OrganizationName = organization.Name
			}
		}

		for _, ticket := range dataSet.Tickets.Items {
			if userSearchResult.Items[key].ID == ticket.SubmitterID {
				userSearchResult.Items[key].SubmittedTickets = append(userSearchResult.Items[key].SubmittedTickets, ticket.Subject)
			}

			if userSearchResult.Items[key].ID == ticket.AssigneeID {
				userSearchResult.Items[key].AssignedTickets = append(userSearchResult.Items[key].AssignedTickets, ticket.Subject)
			}
		}

	}
}

//Populate will load data from data source such as json
func (users *Users) Populate(jsonLoader loader.JSONLoader) error {
	err := load(jsonLoader, &users.Items)
	if err != nil {
		return err
	}
	users.Size = len(users.Items)
	return nil
}

//Search will be called to search data by search key and search field, business logic will be specific to Users
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
							results.Size++
							results.Items = append(results.Items, user)
							break
						}
					}
				} else {
					var castedSearchTerm interface{}
					var err error
					switch searchKey {
					case "_id", "organization_id":
						castedSearchTerm, err = strconv.Atoi(searchTerm)
						if err != nil {
							break
						}
					case "active", "verified", "shared", "suspended":
						castedSearchTerm, err = strconv.ParseBool(searchTerm)
						if err != nil {
							break
						}
					default:
						castedSearchTerm = searchTerm

					}
					if valueField.Interface() == castedSearchTerm {
						results.Size++
						results.Items = append(results.Items, user)
					}
				}
				break
			}
		}
	}

	return results
}
