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

func TestTicketDecorate(t *testing.T) {
	testCases := []struct {
		desc               string
		ticketSearchResult TicketSearchResult
		expected           []int
	}{
		{
			desc:               "Ticket result is decorated succesfully, its number of users and tickets is 1 not 0",
			ticketSearchResult: MockTicketSearchResult,
			expected:           []int{1, 2, 1, 2, 1, 2},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.ticketSearchResult.Decorate(MockDataSet)
			assert.Equal(t, testCase.expected[0], testCase.ticketSearchResult.Items[0].Submitter.ID)
			assert.Equal(t, testCase.expected[1], testCase.ticketSearchResult.Items[0].Assignee.ID)
			assert.Equal(t, testCase.expected[2], testCase.ticketSearchResult.Items[0].Organization.ID)
			assert.Equal(t, testCase.expected[3], testCase.ticketSearchResult.Items[1].Submitter.ID)
			assert.Equal(t, testCase.expected[4], testCase.ticketSearchResult.Items[1].Assignee.ID)
			assert.Equal(t, testCase.expected[5], testCase.ticketSearchResult.Items[1].Organization.ID)
		})
	}
}
