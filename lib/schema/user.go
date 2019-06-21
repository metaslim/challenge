package schema

//User is object to represent an user, json: will be the key to be interacted with such as search and display rather than using intenal struct field
type User struct {
	ID               int      `json:"_id" header:"ID" search:"yes"`
	URL              string   `json:"url" search:"yes" search:"yes"`
	ExternalID       string   `json:"external_id" search:"yes"`
	Name             string   `json:"name" header:"Name" search:"yes"`
	Alias            string   `json:"alias" header:"Alias" search:"yes"`
	CreatedAt        string   `json:"created_at" search:"yes"`
	Active           bool     `json:"active" header:"Active" search:"yes"`
	Verified         bool     `json:"verified" search:"yes"`
	Shared           bool     `json:"shared" search:"yes"`
	Locale           string   `json:"locale" search:"yes"`
	Timezone         string   `json:"timezone" header:"Time Zone" search:"yes"`
	LastLoginAt      string   `json:"last_login_at" search:"yes"`
	Email            string   `json:"email" header:"Email" search:"yes"`
	Phone            string   `json:"phone" header:"Phone" search:"yes"`
	Signature        string   `json:"signature" header:"Signature" search:"yes"`
	OrganizationID   int      `json:"organization_id" search:"yes"`
	Tags             []string `json:"tags" header:"Tags" search:"yes"`
	Suspended        bool     `json:"suspended" search:"yes"`
	Role             string   `json:"role" header:"Role" search:"yes"`
	OrganizationName string   `json:"organization_name" header:"Organization"`
	SubmittedTickets []string `json:"submitted_tickets" header:"Submitted Tickets"`
	AssignedTickets  []string `json:"assigned_tickets" header:"Assigned Tickets"`
}
