package model

import (
	"reflect"
	"strconv"

	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/schema"
)

var _ Records = (*Organizations)(nil)
var _ SearchResult = (*OrganizationSearchResult)(nil)

//Organizations will contain Organization data source
type Organizations struct {
	Items []schema.Organization
}

//OrganizationSearchResult will contain Organization search result
type OrganizationSearchResult struct {
	Items []schema.Organization
	BaseSearchResult
}

//Decorate will decorate the search result, in this case it will populate Tickets and Users properties
func (organizationsSearchResult OrganizationSearchResult) Decorate(dataSet DataSet) {
	for key, _ := range organizationsSearchResult.Items {
		for _, ticket := range dataSet.Tickets.Items {
			if organizationsSearchResult.Items[key].ID == ticket.OrganizationID {
				organizationsSearchResult.Items[key].Tickets = append(organizationsSearchResult.Items[key].Tickets, ticket)
			}
		}

		for _, user := range dataSet.Users.Items {
			if organizationsSearchResult.Items[key].ID == user.OrganizationID {
				organizationsSearchResult.Items[key].Users = append(organizationsSearchResult.Items[key].Users, user)
			}
		}
	}
}

//Populate will load data from data source such as json
func (organizations *Organizations) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, &organizations.Items)
}

//Search will allow data source to be searched
func (organizations *Organizations) Search(searchKey string, searchTerm string) SearchResult {
	results := OrganizationSearchResult{}

	for _, organization := range organizations.Items {
		val := reflect.ValueOf(&organization).Elem()
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			tag := typeField.Tag

			if tag.Get("json") == searchKey {
				if searchKey == "domain_names" {
					for _, domainName := range organization.DomainNames {
						if domainName == searchTerm {
							results.Size++
							results.Items = append(results.Items, organization)
							break
						}
					}
				} else if searchKey == "tags" {
					for _, tag := range organization.Tags {
						if tag == searchTerm {
							results.Size++
							results.Items = append(results.Items, organization)
							break
						}
					}
				} else {
					var castedSearchTerm interface{}
					var err error
					switch searchKey {
					case "_id":
						castedSearchTerm, err = strconv.Atoi(searchTerm)
						if err != nil {
							break
						}
					case "shared_tickets":
						castedSearchTerm, err = strconv.ParseBool(searchTerm)
						if err != nil {
							break
						}
					default:
						castedSearchTerm = searchTerm

					}
					if valueField.Interface() == castedSearchTerm {
						results.Size++
						results.Items = append(results.Items, organization)
					}
				}
				break
			}
		}
	}

	return results
}
