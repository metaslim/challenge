package schema

//Ticket is object to represent ticket, json: will be the key to be interacted with such as search and display rather than using intenal struct field
type Ticket struct {
	ID               string   `json:"_id" header:"ID" search:"yes"`
	URL              string   `json:"url" search:"yes"`
	ExternalID       string   `json:"external_id" search:"yes"`
	CreatedAt        string   `json:"created_at" search:"yes"`
	TicketType       string   `json:"type" header:"Type" search:"yes"`
	Subject          string   `json:"subject" header:"Subject" search:"yes"`
	Description      string   `json:"description" header:"Description" search:"yes"`
	Priority         string   `json:"priority" header:"Priority" search:"yes"`
	Status           string   `json:"status" header:"Status" search:"yes"`
	SubmitterID      int      `json:"submitter_id" search:"yes"`
	AssigneeID       int      `json:"assignee_id" search:"yes"`
	OrganizationID   int      `json:"organization_id" search:"yes"`
	Tags             []string `json:"tags" header:"Tags" search:"yes"`
	HasIncidents     bool     `json:"has_incidents" search:"yes"`
	DueAt            string   `json:"due_at" search:"yes"`
	Via              string   `json:"via" header:"Via" search:"yes"`
	SubmitterName    string   `json:"submitter_name" header:"Submitter"`
	AssigneeName     string   `json:"assignee_name" header:"Assignee"`
	OrganizationName string   `json:"organization_name" header:"Organization"`
}
