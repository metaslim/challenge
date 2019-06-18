package model

import (
	"reflect"

	"github.com/zendesk/challenge/lib/loader"
)

var _ Records = (*Organizations)(nil)
var _ SearchResult = (*OrganizationSearchResult)(nil)

//Organizations is array of Organization
type Organizations struct {
	Items []Organization
}

//Organization is object to represent an orgaization
type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
	Tickets       []Ticket `json:"tickets,omitempty"`
	Users         []User   `json:"users,omitempty"`
}

type OrganizationSearchResult struct {
	Items []Organization
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
