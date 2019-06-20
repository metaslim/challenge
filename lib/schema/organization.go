package schema

//Organization is object to represent an organization, json: will be the key to be interacted with such as search and display rather than using intenal struct field
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
