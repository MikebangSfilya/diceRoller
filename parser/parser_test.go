package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsonParse(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		contentType string
		expected    []PersonRequest
		expectError bool
	}{
		{
			name:        "Valid inputJSON",
			input:       `[{"name": "Джон","dext": 1, "wits" : 2}]`,
			contentType: "application/json",
			expected:    []PersonRequest{{Name: "Джон", Dext: 1, Wits: 2}},
			expectError: false,
		},
		{
			name:        "Bad JSON",
			input:       "",
			contentType: "application/json",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonParse := JSONParse{}
			pers, err := jsonParse.Parse([]byte(tc.input))
			if !tc.expectError {
				require.NotNil(t, pers)
				require.Equal(t, tc.expected, pers)
			} else {
				require.Error(t, err)
			}
		})
	}

}
