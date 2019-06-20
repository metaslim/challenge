package model

import (
	"testing"

	"github.com/metaslim/challenge/lib/loader"
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
			tickets:    GetMockTickets(),
			searchKey:  "type",
			searchTerm: "incident",
			expected:   2,
		},
		{
			desc:       "return 1 ticket with description=",
			tickets:    GetMockTickets(),
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
		expected           []string
	}{
		{
			desc:               "Ticket result is decorated succesfully, its Submitter, Assignee, and Organization is not nil",
			ticketSearchResult: GetMockTicketSearchResult(),
			expected:           []string{"Nick Fury", "Steve Rogers", "Shield", "Steve Rogers", "Nick Fury", "Avengers"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.ticketSearchResult.Decorate(GetMockDataSet())
			assert.Equal(t, testCase.expected[0], testCase.ticketSearchResult.Items[0].SubmitterName)
			assert.Equal(t, testCase.expected[1], testCase.ticketSearchResult.Items[0].AssigneeName)
			assert.Equal(t, testCase.expected[2], testCase.ticketSearchResult.Items[0].OrganizationName)
			assert.Equal(t, testCase.expected[3], testCase.ticketSearchResult.Items[1].SubmitterName)
			assert.Equal(t, testCase.expected[4], testCase.ticketSearchResult.Items[1].AssigneeName)
			assert.Equal(t, testCase.expected[5], testCase.ticketSearchResult.Items[1].OrganizationName)
		})
	}
}

func TestTicketPopulate(t *testing.T) {
	testCases := []struct {
		desc     string
		tickets  Tickets
		jsonFile string
		expected int
	}{
		{
			desc:     "Ticket is populated with 30 records from good json file",
			tickets:  Tickets{},
			jsonFile: `../../data/test-tickets.json`,
			expected: 30,
		},
		{
			desc:     "Ticket is not populated from bad json file",
			tickets:  Tickets{},
			jsonFile: `../../data/test-tickets-not-exists.json`,
			expected: 30,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			err := testCase.tickets.Populate(loader.JSONLoader{
				FileName: testCase.jsonFile,
			})

			if err != nil {
				assert.NotEqual(t, testCase.expected, testCase.tickets.GetSize())
			} else {
				assert.Equal(t, testCase.expected, testCase.tickets.GetSize())
			}
		})
	}
}
