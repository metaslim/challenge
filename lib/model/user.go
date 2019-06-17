package model

import (
	"time"

	"github.com/zendesk/challenge/lib/loader"
)

var _ Model = (*User)(nil)

//User is object to represent an user
type User struct {
	ID             int       `json:"_id"`
	URL            string    `json:"url"`
	ExternalID     string    `json:"external_id"`
	Name           string    `json:"name"`
	Alias          string    `json:"alias"`
	CreatedAt      time.Time `json:"created_at"`
	Active         bool      `json:"active"`
	Verified       bool      `json:"verified"`
	Shared         bool      `json:"shared"`
	Locale         string    `json:"locale"`
	Timezone       string    `json:"timezone"`
	LastLoginAt    time.Time `json:"last_login_at"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Signature      string    `json:"signature"`
	OrganizationID int       `json:"organization_id"`
	Tags           []string  `json:"tags"`
	Suspended      bool      `json:"suspended"`
	Role           string    `json:"role"`
}

//New return model
func (user *User) New(jsonLoader loader.JSONLoader) error {
	return load(jsonLoader, user)
}
