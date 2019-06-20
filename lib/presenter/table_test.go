package presenter

import (
	"testing"

	"github.com/acarl005/stripansi"
	"github.com/metaslim/challenge/lib/model"
	"github.com/metaslim/challenge/lib/util"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationTableSearchResultFlush(t *testing.T) {
	testCases := []struct {
		desc                     string
		organizationSearchResult model.OrganizationSearchResult
		table                    Table
		expected                 string
	}{
		{
			desc:                     "Table Presenter is able to parse the organization search result",
			table:                    Table{},
			organizationSearchResult: model.GetMockOrganizationSearchResult(),
			expected:                 `shield.gov.us`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			output := util.CaptureOutput(func() {
				testCase.table.Flush(testCase.organizationSearchResult)
			})

			assert.Contains(t, stripansi.Strip(output), testCase.expected)

		})
	}
}

func TestUserTableSearchResultFlush(t *testing.T) {
	testCases := []struct {
		desc             string
		userSearchResult model.UserSearchResult
		table            Table
		expected         string
	}{
		{
			desc:             "Table Presenter is able to parse the user search result",
			table:            Table{},
			userSearchResult: model.GetMockUserSearchResult(),
			expected:         `I can do this all day`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			output := util.CaptureOutput(func() {
				testCase.table.Flush(testCase.userSearchResult)
			})

			assert.Contains(t, stripansi.Strip(output), testCase.expected)

		})
	}
}

func TestTicketTableSearchResultFlush(t *testing.T) {
	testCases := []struct {
		desc               string
		ticketSearchResult model.TicketSearchResult
		table              Table
		expected           string
	}{
		{
			desc:               "Table Presenter is able to parse the ticket search result",
			table:              Table{},
			ticketSearchResult: model.GetMockTicketSearchResult(),
			expected:           `Loki attack, Puny God`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			output := util.CaptureOutput(func() {
				testCase.table.Flush(testCase.ticketSearchResult)
			})

			assert.Contains(t, stripansi.Strip(output), testCase.expected)

		})
	}
}
