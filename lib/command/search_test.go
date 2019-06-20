package command

import (
	"testing"

	"github.com/metaslim/challenge/lib/mocks"
	"github.com/metaslim/challenge/lib/presenter"
	"github.com/stretchr/testify/assert"
)

func TestSearchValid(t *testing.T) {
	testCases := []struct {
		desc     string
		search   Search
		expected bool
	}{
		{
			desc: "return true for correct search command for organizations",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-organizations:tags=shield",
				},
			},
			expected: true,
		},
		{
			desc: "return true for correct search command for users",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-users:alias=Captain America",
				},
			},
			expected: true,
		},
		{
			desc: "return true for correct search command for tickets",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-tickets:status=pending",
				},
			},
			expected: true,
		},
		{
			desc: "return true for correct search command for tickets",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-unknow:_id=1",
				},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			valid := testCase.search.Valid()

			assert.Equal(t, testCase.expected, valid)
		})
	}
}

func TestSearcheRun(t *testing.T) {
	testCases := []struct {
		desc     string
		search   Search
		expected string
	}{
		{
			desc: "return organization tagged with shield",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-organizations:tags=shield",
				},
				Presenter: presenter.Json{},
			},
			expected: `"name": "Shield"`,
		},
		{
			desc: "return no organization tagged with wakanda",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-organizations:tags=wakanda",
				},
				Presenter: presenter.Json{},
			},
			expected: `No result found`,
		},
		{
			desc: "return user whose alias is Captain America",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-users:alias=Captain America",
				},
				Presenter: presenter.Json{},
			},
			expected: `"signature": "I can do this all day"`,
		},
		{
			desc: "return no user whose alias is Black Widow",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-users:alias=Black Widow",
				},
				Presenter: presenter.Json{},
			},
			expected: `No result found`,
		},
		{
			desc: "return ticket with pending status",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-tickets:status=solved",
				},
				Presenter: presenter.Json{},
			},
			expected: `"description": "Loki attack, Puny God",`,
		},
		{
			desc: "return no ticket with unknown status",
			search: Search{
				BaseCommand: BaseCommand{
					Input: "search-tickets:status=unknown",
				},
				Presenter: presenter.Json{},
			},
			expected: `No result found`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.search.Valid()

			output := captureOutput(func() {
				testCase.search.Run(mocks.MockDataSet)
			})

			assert.Contains(t, output, testCase.expected)
		})
	}
}
