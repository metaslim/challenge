package command

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/metaslim/challenge/lib/model"
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
			desc: "return true for correct describe command",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-organizations",
				},
			},
			expected: "organizations can be searched by any fields below",
		},
		{
			desc: "return false for incorrect describe command",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-users",
				},
			},
			expected: "users can be searched by any fields below",
		},
		{
			desc: "return false for incorrect describe command",
			describe: Describe{
				BaseCommand: BaseCommand{
					Input: "describe-users",
				},
			},
			expected: "users can be searched by any fields below",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.describe.Valid()

			output := captureOutput(func() {
				testCase.describe.Run(model.MockDataSet)
			})

			assert.Contains(t, output, testCase.expected)
		})
	}
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
