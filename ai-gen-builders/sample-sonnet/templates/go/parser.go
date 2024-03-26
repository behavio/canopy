
package parser

import (
	"{{@root.importPath}}/actions"
	"fmt"
	"strings"
)

func Parse(input string, acts actions.Actions) (*Node, error) {
	g := &Grammar{Buffer: []rune(input), Actions: acts, cache: nil}
	result, err := g.{{root}}()

	if err == nil && g.Position.Offset != len(input) {
		expected := []Expected{}
		for _, exp := range g.Position.expected {
			expected = append(expected, Expected{Name: exp.Name, Value: exp.Value})
		}
		return nil, &parseError{Position: g.Position.Offset, Expected: expected}
	}

	return result, err
}

type parseError struct {
	Position int
	Expected []Expected
}

type Expected struct {
	Name  string
	Value string
}

func formatError(err *parseError, input string) string {
	lines := strings.Split(input, "\n")
	line, column := 0, 0
	position := 0
	for i, offset := range []int{0} {
		line, column = i, position
		if i < len(lines) {
			nextLine := len(lines[i]) + 1
			position += nextLine
			if position >= err.Position {
				break
			}
			offset = nextLine
		}
	}

	message := fmt.Sprintf("Line %d, column %d: Expected one of:\n", line+1, column)
	for _, exp := range err.Expected {
		message += fmt.Sprintf("  %s '%s'\n", exp.Name, exp.Value)
	}

	marker := "\x1b[0;31m\x1b[1m\u2192\x1b[0m" // Red bold arrow

	line = 0
	output := ""
	for _, rune := range input {
		if line != line {
			output += "\n"
		}
		if column == err.Position-1 {
			output += marker
		}
		output += string(rune)
		if column == err.Position {
			output += marker + "\x1b[0m"
		}
		column++
	}

	return message + output
}

func (err *parseError) Error() string {
	return formatError(err, "{{grammar}}")
}
