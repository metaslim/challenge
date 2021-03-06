package model

import (
	"testing"

	"github.com/metaslim/challenge/lib/loader"
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
			organizations: GetMockOrganizations(),
			searchKey:     "_id",
			searchTerm:    "1",
			expected:      1,
		},
		{
			desc:          "return 2 organizations with shared_tickets=true",
			organizations: GetMockOrganizations(),
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
			organizationSearchResult: GetMockOrganizationSearchResult(),
			expected:                 []int{1, 1, 1, 1},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.organizationSearchResult.Decorate(GetMockDataSet())
			assert.Equal(t, testCase.expected[0], len(testCase.organizationSearchResult.Items[0].Tickets))
			assert.Equal(t, testCase.expected[1], len(testCase.organizationSearchResult.Items[1].Users))
			assert.Equal(t, testCase.expected[2], len(testCase.organizationSearchResult.Items[0].Tickets))
			assert.Equal(t, testCase.expected[3], len(testCase.organizationSearchResult.Items[1].Users))
		})
	}
}

func TestOrganizationPopulate(t *testing.T) {
	testCases := []struct {
		desc          string
		organizations Organizations
		jsonFile      string
		expected      int
	}{
		{
			desc:          "Organization is populated with 10 records from good json file",
			organizations: Organizations{},
			jsonFile:      `../../data/test-organizations.json`,
			expected:      10,
		},
		{
			desc:          "Organization is not populated from bad json file",
			organizations: Organizations{},
			jsonFile:      `../../data/test-organizations-not-exists.json`,
			expected:      10,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			err := testCase.organizations.Populate(loader.JSONLoader{
				FileName: testCase.jsonFile,
			})

			if err != nil {
				assert.NotEqual(t, testCase.expected, testCase.organizations.GetSize())
			} else {
				assert.Equal(t, testCase.expected, testCase.organizations.GetSize())
			}
		})
	}
}
