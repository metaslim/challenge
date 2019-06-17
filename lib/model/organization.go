package model

import (
	"github.com/zendesk/challenge/lib/loader"
)

var _ Model = (*Organization)(nil)

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
}

//Populate is
func (organization *Organization) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, organization)
}

//Populate is
func (organizations *Organizations) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, &organizations.Items)
}
