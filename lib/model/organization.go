package model

import (
	"time"

	"github.com/zendesk/challenge/lib/loader"
)

var _ Model = (*Organization)(nil)

//Organization is object to represent an orgaization
type Organization struct {
	ID            int       `json:"_id"`
	URL           string    `json:"url"`
	ExternalID    string    `json:"external_id"`
	Name          string    `json:"name"`
	DomainNames   []string  `json:"domain_names"`
	CreatedAt     time.Time `json:"created_at"`
	Details       string    `json:"details"`
	SharedTickets bool      `json:"shared_tickets"`
	Tags          []string  `json:"tags"`
}

//New return model
func (organization *Organization) New(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, organization)
}
