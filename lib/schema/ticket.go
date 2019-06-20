package schema

//Ticket is object to represent ticket, json: will be the key to be interacted with such as search and display rather than using intenal struct field
type Ticket struct {
	ID               string   `json:"_id" header:"ID"`
	URL              string   `json:"url"`
	ExternalID       string   `json:"external_id"`
	CreatedAt        string   `json:"created_at"`
	TicketType       string   `json:"type" header:"Type"`
	Subject          string   `json:"subject" header:"Subject"`
	Description      string   `json:"description" header:"Description"`
	Priority         string   `json:"priority" header:"Priority"`
	Status           string   `json:"status" header:"Status"`
	SubmitterID      int      `json:"submitter_id"`
	AssigneeID       int      `json:"assignee_id"`
	OrganizationID   int      `json:"organization_id"`
	Tags             []string `json:"tags" header:"Tags"`
	HasIncidents     bool     `json:"has_incidents"`
	DueAt            string   `json:"due_at"`
	Via              string   `json:"via" header:"Via"`
	SubmitterName    string   `json:"submitter_name" header:"Submitter"`
	AssigneeName     string   `json:"assignee_name" header:"Assignee"`
	OrganizationName string   `json:"organization_name" header:"Organization"`
}
