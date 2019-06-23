package command

import (
	"testing"

	"github.com/metaslim/challenge/lib/model"
	"github.com/metaslim/challenge/lib/util"
	"github.com/stretchr/testify/assert"
)

func TestDescribeValid(t *testing.T) {
	testCases := []struct {
		desc     string
		describe Describe
		expected bool
	}{
		{
			desc: "return true for correct describe command for organizations",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-organizations",
				},
			},
			expected: true,
		},
		{
			desc: "return true for correct describe command for users",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-users",
				},
			},
			expected: true,
		},
		{
			desc: "return true for correct describe command for tickets",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-tickets",
				},
			},
			expected: true,
		},
		{
			desc: "aaa",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-unknown",
				},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			valid := testCase.describe.Valid()

			assert.Equal(t, testCase.expected, valid)
		})
	}
}

func TestDescribeRun(t *testing.T) {
	testCases := []struct {
		desc     string
		describe Describe
		expected string
	}{
		{
			desc: "return correct search field for organizations",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-organizations",
				},
			},
			expected: "ORGANIZATIONS can be searched by any fields below",
		},
		{
			desc: "return correct search field for tickets",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-tickets",
				},
			},
			expected: "TICKETS can be searched by any fields below",
		},
		{
			desc: "return correct search field for users",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-users",
				},
			},
			expected: "USERS can be searched by any fields below",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.describe.Valid()

			output := util.CaptureOutput(func() {
				testCase.describe.Run(model.GetMockDataSet())
			})

			assert.Contains(t, output, testCase.expected)
		})
	}
}
