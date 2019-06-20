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
