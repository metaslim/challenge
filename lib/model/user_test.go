package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSearch(t *testing.T) {
	testCases := []struct {
		desc       string
		users      Users
		searchKey  string
		searchTerm string
		expected   int
	}{
		{
			desc:       "return 1 user with signature=I can do this all day",
			users:      MockUsers,
			searchKey:  "signature",
			searchTerm: "I can do this all day",
			expected:   1,
		},
		{
			desc:       "return 2 users with tags=heroes",
			users:      MockUsers,
			searchKey:  "tags",
			searchTerm: "heroes",
			expected:   2,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			result := testCase.users.Search(testCase.searchKey, testCase.searchTerm)
			assert.Equal(t, testCase.expected, result.GetSize())
		})
	}
}
