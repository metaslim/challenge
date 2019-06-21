package schema

//Organization is object to represent an organization, json: will be the key to be interacted with such as search and display rather than using intenal struct field
type Organization struct {
	ID            int      `json:"_id" header:"ID" search:"yes"`
	URL           string   `json:"url" search:"yes"`
	ExternalID    string   `json:"external_id" search:"yes"`
	Name          string   `json:"name" header:"Name" search:"yes"`
	DomainNames   []string `json:"domain_names" header:"Domain Names" search:"yes"`
	CreatedAt     string   `json:"created_at" search:"yes"`
	Details       string   `json:"details" header:"Details" search:"yes"`
	SharedTickets bool     `json:"shared_tickets" search:"yes"`
	Tags          []string `json:"tags" header:"Tags" search:"yes"`
	Tickets       []string `json:"tickets" header:"Tickets"`
	Users         []string `json:"users" header:"Users"`
}
