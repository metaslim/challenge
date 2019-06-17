package model

import (
	"time"
)

//Ticket is object to represent ticket
type Ticket struct {
	id             int
	url            string
	externalId     string
	createdAt      time.Time
	ticketType     string
	subject        string
	description    string
	priority       string
	status         string
	submitterId    int
	assigneeId     int
	organizationId int
	tags           []string
	hasIncident    bool
	dueAt          time.Time
	via            string
}
