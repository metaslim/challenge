package model

import (
	"time"
)

//Organization is object to represent an orgaization
type Organization struct {
	id            int
	url           string
	externalId    string
	name          string
	domainName    []string
	createdAt     time.Time
	details       string
	sharedTickets bool
	tags          []string
}
