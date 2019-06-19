package schema

//User is object to represent an user
type User struct {
	ID              int           `json:"_id"`
	URL             string        `json:"url"`
	ExternalID      string        `json:"external_id"`
	Name            string        `json:"name"`
	Alias           string        `json:"alias"`
	CreatedAt       string        `json:"created_at"`
	Active          bool          `json:"active"`
	Verified        bool          `json:"verified"`
	Shared          bool          `json:"shared"`
	Locale          string        `json:"locale"`
	Timezone        string        `json:"timezone"`
	LastLoginAt     string        `json:"last_login_at"`
	Email           string        `json:"email"`
	Phone           string        `json:"phone"`
	Signature       string        `json:"signature"`
	OrganizationID  int           `json:"organization_id"`
	Tags            []string      `json:"tags"`
	Suspended       bool          `json:"suspended"`
	Role            string        `json:"role"`
	Organization    *Organization `json:"organization,omitempty"`
	SubmittedTicket []Ticket      `json:"submitted_by,omitempty"`
	AssignedTicket  []Ticket      `json:"assigned_to,omitempty"`
}
