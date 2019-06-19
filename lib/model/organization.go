package model

import (
	"reflect"

	"github.com/metaslim/challenge/lib/loader"
	"github.com/metaslim/challenge/lib/schema"
)

var _ Records = (*Organizations)(nil)
var _ SearchResult = (*OrganizationSearchResult)(nil)

//Organizations is array of Organization
type Organizations struct {
	Items []schema.Organization
}

type OrganizationSearchResult struct {
	Items []schema.Organization
}

func (organizationsSearchResult OrganizationSearchResult) Decorate(decorateParams DecorateParams) {
	for key, _ := range organizationsSearchResult.Items {
		for _, ticket := range decorateParams.Tickets.Items {
			if organizationsSearchResult.Items[key].ID == ticket.OrganizationID {
				organizationsSearchResult.Items[key].Tickets = append(organizationsSearchResult.Items[key].Tickets, ticket)
			}
		}

		for _, user := range decorateParams.Users.Items {
			if organizationsSearchResult.Items[key].ID == user.OrganizationID {
				organizationsSearchResult.Items[key].Users = append(organizationsSearchResult.Items[key].Users, user)
			}
		}
	}
}

//Populate is
func (organizations *Organizations) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, &organizations.Items)
}

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
							results.Items = append(results.Items, organization)
							break
						}
					}
				} else if searchKey == "tags" {
					for _, tag := range organization.Tags {
						if tag == searchTerm {
							results.Items = append(results.Items, organization)
							break
						}
					}
				} else if valueField.Interface() == searchTerm {
					results.Items = append(results.Items, organization)
				}
				break
			}
		}
	}

	return results
}
