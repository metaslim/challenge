package loader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoader(t *testing.T) {
	testCases := []struct {
		desc     string
		jsonFile string
		expected *os.File
	}{
		{
			desc:     "File handle is valid from bad json file",
			jsonFile: `../../data/test-organizations.json`,
			expected: nil,
		},
		{
			desc:     "File handle is nil from bad json file",
			jsonFile: `../../data/test-organizations-not-exists.json`,
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.desc, func(t *testing.T) {
			fileLoader := JSONLoader{
				FileName: testCase.jsonFile,
			}

			fileHandle, err := fileLoader.GetFileHandle()

			if err != nil {
				assert.Equal(t, testCase.expected, fileHandle)
			} else {
				assert.NotEqual(t, testCase.expected, fileHandle)
			}
		})
	}
}
