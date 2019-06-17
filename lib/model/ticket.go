package model

import (
	"time"

	"github.com/zendesk/challenge/lib/loader"
)

var _ Model = (*Ticket)(nil)

//Ticket is object to represent ticket
type Ticket struct {
	ID             int       `json:"_id"`
	URL            string    `json:"url"`
	ExternalID     string    `json:"external_id"`
	CreatedAt      time.Time `json:"created_at"`
	TicketType     string    `json:"type"`
	Subject        string    `json:"subject"`
	Description    string    `json:"description"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	SubmitterID    int       `json:"submitter_id"`
	AssigneeID     int       `json:"assignee_id"`
	OrganizationID int       `json:"organization_id"`
	Tags           []string  `json:"tags"`
	HasIncidents   bool      `json:"has_incidents"`
	DueAt          time.Time `json:"due_at"`
	Via            string    `json:"via"`
}

//New return model
func (ticket *Ticket) New(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, ticket)
}
