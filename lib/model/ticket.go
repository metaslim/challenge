package model

import (
	"github.com/zendesk/challenge/lib/loader"
)

var _ Model = (*Ticket)(nil)

//Tickets is array of Tickets
type Tickets struct {
	Items []Ticket
}

//Ticket is object to represent ticket
type Ticket struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	TicketType     string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

//Populate is
func (ticket *Ticket) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, ticket)
}

//Populate is
func (tickets *Tickets) Populate(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, &tickets.Items)
}
