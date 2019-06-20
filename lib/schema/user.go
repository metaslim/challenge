package schema

//User is object to represent an user, json: will be the key to be interacted with such as search and display rather than using intenal struct field
type User struct {
	ID               int      `json:"_id" header:"ID"`
	URL              string   `json:"url"`
	ExternalID       string   `json:"external_id"`
	Name             string   `json:"name" header:"Name"`
	Alias            string   `json:"alias" header:"Alias"`
	CreatedAt        string   `json:"created_at"`
	Active           bool     `json:"active" header:"Active"`
	Verified         bool     `json:"verified"`
	Shared           bool     `json:"shared"`
	Locale           string   `json:"locale"`
	Timezone         string   `json:"timezone" header:"Time Zone"`
	LastLoginAt      string   `json:"last_login_at"`
	Email            string   `json:"email" header:"Email"`
	Phone            string   `json:"phone" header:"Phone"`
	Signature        string   `json:"signature" header:"Signature"`
	OrganizationID   int      `json:"organization_id"`
	Tags             []string `json:"tags" header:"Tags"`
	Suspended        bool     `json:"suspended"`
	Role             string   `json:"role" header:"Role"`
	OrganizationName string   `json:"organization_name" header:"Organization"`
	SubmittedTickets []string `json:"submitted_tickets" header:"Submitted Tickets"`
	AssignedTickets  []string `json:"assigned_tickets" header:"Assigned Tickets"`
}
