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

func TestUserDecorate(t *testing.T) {
	testCases := []struct {
		desc             string
		userSearchResult UserSearchResult
		expected         []int
	}{
		{
			desc:             "User result is decorated succesfully, its number of SubmSubmittedTickets, AssignedTickets  is 1 and Organization is not nil",
			userSearchResult: MockUserSearchResult,
			expected:         []int{1, 1, 1, 2, 1, 1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.userSearchResult.Decorate(MockDataSet)
			assert.Equal(t, testCase.expected[0], testCase.userSearchResult.Items[0].Organization.ID)
			assert.Equal(t, testCase.expected[1], len(testCase.userSearchResult.Items[0].SubmittedTickets))
			assert.Equal(t, testCase.expected[2], len(testCase.userSearchResult.Items[0].AssignedTickets))
			assert.Equal(t, testCase.expected[0], testCase.userSearchResult.Items[0].Organization.ID)
			assert.Equal(t, testCase.expected[1], len(testCase.userSearchResult.Items[0].SubmittedTickets))
			assert.Equal(t, testCase.expected[2], len(testCase.userSearchResult.Items[0].AssignedTickets))
		})
	}
}
