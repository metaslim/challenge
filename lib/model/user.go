package model

import (
	"time"
)

//User is object to represent an user
type User struct {
	id             int
	url            string
	externalId     string
	name           string
	alias          string
	createdAt      time.Time
	active         bool
	verified       bool
	shared         bool
	locale         string
	timezone       string
	lastLoginAt    time.Time
	email          string
	phone          string
	signature      string
	organizationId int
	tags           []string
	suspended      bool
	role           string
}
