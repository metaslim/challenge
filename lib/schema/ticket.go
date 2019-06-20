package schema

//Ticket is object to represent ticket, json: will be the key to be interacted with such as search and display rather than using intenal struct field
type Ticket struct {
	ID             string        `json:"_id"`
	URL            string        `json:"url"`
	ExternalID     string        `json:"external_id"`
	CreatedAt      string        `json:"created_at"`
	TicketType     string        `json:"type"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	Priority       string        `json:"priority"`
	Status         string        `json:"status"`
	SubmitterID    int           `json:"submitter_id"`
	AssigneeID     int           `json:"assignee_id"`
	OrganizationID int           `json:"organization_id"`
	Tags           []string      `json:"tags"`
	HasIncidents   bool          `json:"has_incidents"`
	DueAt          string        `json:"due_at"`
	Via            string        `json:"via"`
	Submitter      *User         `json:"submitter,omitempty"`
	Assignee       *User         `json:"assignee,omitempty"`
	Organization   *Organization `json:"organization,omitempty"`
}
