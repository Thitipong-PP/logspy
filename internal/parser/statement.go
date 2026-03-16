package parser

import (
	"strings"
)

type Statement struct {
	Command   string
	Args      []string
	Flags     map[string][]string
	BoolFlags map[string]bool
}

// Create new statement
func NewStatement(command string) *Statement {
	return &Statement{
		Command:   command,
		Args:      make([]string, 0),
		Flags:     make(map[string][]string),
		BoolFlags: make(map[string]bool),
	}
}

// Set bool flag to true
func (s *Statement) SetBoolFlag(flag string) {
	flag = strings.TrimPrefix(flag, "--")
	s.BoolFlags[flag] = true
}

// Set flag to value
func (s *Statement) PushFlag(flag string, item string) {
	// Split item to item list
	itemList := strings.Split(item, ",")
	for i := range itemList {
		itemList[i] = strings.TrimSpace(itemList[i])
	}

	flag = strings.TrimPrefix(flag, "-")

	// Finding data in flag
	data, finding := s.Flags[flag]
	if !finding {
		s.Flags[flag] = make([]string, 0, len(itemList))
	}

	// Setting data to flag
	data = append(data, itemList...)
	s.Flags[flag] = data
}
