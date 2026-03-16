package parser

import (
	"fmt"
	"strings"
)

// Cutting a cli command to be statement
func Parse(cli []string) (*Statement, error) {
	// Don't have a command
	if len(cli) < 2 {
		return nil, fmt.Errorf("not enough arguments: missing command")
	}

	cli = cli[1:]
	statement := NewStatement(cli[0])

	// Don't have any arguments
	if len(cli) == 1 {
		return statement, nil
	}

	cli = cli[1:]
	for i:=0; i<len(cli); i++ {
		switch {
		case strings.HasPrefix(cli[i], "--"):
			statement.SetBoolFlag(cli[i])
		case strings.HasPrefix(cli[i], "-"):
			if i+1 >= len(cli) || strings.HasPrefix(cli[i+1], "-") {
				return nil, fmt.Errorf("missing required value for flag: %s", cli[i])
			}

			statement.PushFlag(cli[i], cli[i+1])
			i++
		default:
			statement.Args = append(statement.Args, cli[i])
		}
	}
	
	return statement, nil
}