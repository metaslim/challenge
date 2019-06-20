package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketSearch(t *testing.T) {
	testCases := []struct {
		desc       string
		tickets    Tickets
		searchKey  string
		searchTerm string
		expected   int
	}{
		{
			desc:       "return 2 tickets with ticket_type=incident",
			tickets:    MockTickets,
			searchKey:  "type",
			searchTerm: "incident",
			expected:   2,
		},
		{
			desc:       "return 1 ticket with description=",
			tickets:    MockTickets,
			searchKey:  "description",
			searchTerm: "",
			expected:   1,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			result := testCase.tickets.Search(testCase.searchKey, testCase.searchTerm)
			assert.Equal(t, testCase.expected, result.GetSize())
		})
	}
}
