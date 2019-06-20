package command

import (
	"testing"

	"github.com/metaslim/challenge/lib/model"
	"github.com/metaslim/challenge/lib/util"
	"github.com/stretchr/testify/assert"
)

func TestHelpValid(t *testing.T) {
	testCases := []struct {
		desc     string
		help     Help
		expected bool
	}{
		{
			desc: "return true for correct help command",
			help: Help{
				BaseCommand: BaseCommand{
					Input: "help",
				},
			},
			expected: true,
		},
		{
			desc: "return true for correct describe command for users",
			help: Help{
				BaseCommand: BaseCommand{
					Input: "unknown-help",
				},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			valid := testCase.help.Valid()

			assert.Equal(t, testCase.expected, valid)
		})
	}
}

func TestHelpRun(t *testing.T) {
	testCases := []struct {
		desc     string
		help     Help
		expected string
	}{
		{
			desc:     "return help message",
			help:     Help{},
			expected: "Sample commands",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			testCase.help.Valid()

			output := util.CaptureOutput(func() {
				testCase.help.Run(model.MockDataSet)
			})

			assert.Contains(t, output, testCase.expected)
		})
	}
}
