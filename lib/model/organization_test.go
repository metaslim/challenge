package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationsSearch(t *testing.T) {
	testCases := []struct {
		desc          string
		organizations Organizations
		searchKey     string
		searchTerm    string
		expected      int
	}{
		{
			desc:          "return 1 organization with _id=1",
			organizations: MockOrganizations,
			searchKey:     "_id",
			searchTerm:    "1",
			expected:      1,
		},
		{
			desc:          "return 2 organizations with shared_tickets=true",
			organizations: MockOrganizations,
			searchKey:     "shared_tickets",
			searchTerm:    "true",
			expected:      2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			result := testCase.organizations.Search(testCase.searchKey, testCase.searchTerm)
			assert.Equal(t, testCase.expected, result.GetSize())
		})
	}
}

func TestOrganizationDecorate(t *testing.T) {
	testCases := []struct {
		desc                     string
		organizationSearchResult OrganizationSearchResult
		expected                 []int
	}{
		{
			desc:                     "Organization result is decorated succesfully, its number of users and tickets is 1 not 0",
			organizationSearchResult: MockOrganizationSearchResult,
			expected:                 []int{1, 1, 1, 1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.organizationSearchResult.Decorate(MockDataSet)
			assert.Equal(t, testCase.expected[0], len(testCase.organizationSearchResult.Items[0].Tickets))
			assert.Equal(t, testCase.expected[1], len(testCase.organizationSearchResult.Items[1].Users))
			assert.Equal(t, testCase.expected[2], len(testCase.organizationSearchResult.Items[0].Tickets))
			assert.Equal(t, testCase.expected[3], len(testCase.organizationSearchResult.Items[1].Users))
		})
	}
}
