package presenter

import (
	"testing"

	"github.com/acarl005/stripansi"
	"github.com/metaslim/challenge/lib/model"
	"github.com/metaslim/challenge/lib/util"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationSearchResultFlush(t *testing.T) {
	testCases := []struct {
		desc                     string
		organizationSearchResult model.OrganizationSearchResult
		json                     Json
		expected                 string
	}{
		{
			desc:                     "Json Presenter is able to parse the organization search result",
			json:                     Json{},
			organizationSearchResult: model.GetMockOrganizationSearchResult(),
			expected:                 `"details": "Avengers"`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			output := util.CaptureOutput(func() {
				testCase.json.Flush(testCase.organizationSearchResult)
			})

			assert.Contains(t, stripansi.Strip(output), testCase.expected)

		})
	}
}

func TestUserSearchResultFlush(t *testing.T) {
	testCases := []struct {
		desc             string
		userSearchResult model.UserSearchResult
		json             Json
		expected         string
	}{
		{
			desc:             "Json Presenter is able to parse the user search result",
			json:             Json{},
			userSearchResult: model.GetMockUserSearchResult(),
			expected:         `"signature": "I still believe in heroes"`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			output := util.CaptureOutput(func() {
				testCase.json.Flush(testCase.userSearchResult)
			})

			assert.Contains(t, stripansi.Strip(output), testCase.expected)

		})
	}
}

func TestTicketSearchResultFlush(t *testing.T) {
	testCases := []struct {
		desc               string
		ticketSearchResult model.TicketSearchResult
		json               Json
		expected           string
	}{
		{
			desc:               "Json Presenter is able to parse the ticket search result",
			json:               Json{},
			ticketSearchResult: model.GetMockTicketSearchResult(),
			expected:           `"subject": "Thanos attack"`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			output := util.CaptureOutput(func() {
				testCase.json.Flush(testCase.ticketSearchResult)
			})

			assert.Contains(t, stripansi.Strip(output), testCase.expected)

		})
	}
}
