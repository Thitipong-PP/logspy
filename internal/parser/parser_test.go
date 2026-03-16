package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		cli         []string
		want        *Statement
		expectError bool
	}{
		{
			name: "Basic command without flags",
			cli:  []string{"logspy", "help"},
			want: &Statement{
				Command:   "help",
				Args:      []string{},
				Flags:     map[string][]string{},
				BoolFlags: map[string]bool{},
			},
			expectError: false,
		},
		{
			name:        "Missing command",
			cli:         []string{"logspy"},
			want:        nil,
			expectError: true,
		},
		{
			name: "Command with flags",
			cli:  []string{"logspy", "find", "-file", "app.log", "-env", "development"},
			want: &Statement{
				Command: "find",
				Args:    []string{},
				Flags: map[string][]string{
					"file": {"app.log"},
					"env":  {"development"},
				},
				BoolFlags: map[string]bool{},
			},
			expectError: false,
		},
		{
			name: "Command with boolean flags",
			cli:  []string{"logspy", "find", "--test", "--dev"},
			want: &Statement{
				Command: "find",
				Args:    []string{},
				Flags:   map[string][]string{},
				BoolFlags: map[string]bool{
					"test": true,
					"dev":  true,
				},
			},
			expectError: false,
		},
		{
			name: "Command with arguments",
			cli:  []string{"logspy", "find", "query", "auto-fill"},
			want: &Statement{
				Command: "find",
				Args: []string{
					"query",
					"auto-fill",
				},
				Flags:     map[string][]string{},
				BoolFlags: map[string]bool{},
			},
			expectError: false,
		},
		{
			name: "Command with flags and arguments",
			cli:  []string{"logspy", "find", "-file", "app.log", "--ignore-case", "query"},
			want: &Statement{
				Command: "find",
				Args:    []string{"query"},
				Flags: map[string][]string{
					"file": {"app.log"},
				},
				BoolFlags: map[string]bool{
					"ignore-case": true,
				},
			},
			expectError: false,
		},
		{
			name:        "Missing flag value",
			cli:         []string{"logspy", "find", "-file", "--ignore-case", "query"},
			want:        nil,
			expectError: true,
		},
		{
			name: "Command with array flags",
			cli:  []string{"logspy", "find", "-find", "Panic,Error", "--ignore-case", "query"},
			want: &Statement{
				Command: "find",
				Args:    []string{"query"},
				Flags: map[string][]string{
					"find": {"Panic", "Error"},
				},
				BoolFlags: map[string]bool{
					"ignore-case": true,
				},
			},
			expectError: false,
		},
		{
			name: "Command with array flags with space",
			cli:  []string{"logspy", "find", "-find", "Panic, Error", "--ignore-case", "query"},
			want: &Statement{
				Command: "find",
				Args:    []string{"query"},
				Flags: map[string][]string{
					"find": {"Panic", "Error"},
				},
				BoolFlags: map[string]bool{
					"ignore-case": true,
				},
			},
			expectError: false,
		},
		// Add more edge cases here! (e.g., "-file" without a value)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.cli)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Parse() error = %v, expectError %v", err, tt.expectError)
				return
			}

			// Check if the parsed statement matches what we want
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
