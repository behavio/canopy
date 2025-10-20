package test

import (
	"errors"
	"testing"

	"terminalsgoparser"
)

var (
	terminalsNodeAccessors = newNodeAccessors(
		func(n terminalsgoparser.TreeNode) string { return n.Text() },
		func(n terminalsgoparser.TreeNode) int { return n.Offset() },
		func(n terminalsgoparser.TreeNode) []terminalsgoparser.TreeNode { return n.Children() },
	)
	terminalsParse = func(input string) (terminalsgoparser.TreeNode, error) {
		return terminalsgoparser.Parse(input, nil, nil)
	}
)

func parseTerminal(t *testing.T, input string) terminalsgoparser.TreeNode {
	t.Helper()

	tree, err := terminalsParse(input)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}

	children := tree.Children()
	if len(children) < 2 {
		t.Fatalf("parse(%q) expected at least two root children, got %d", input, len(children))
	}

	return children[1]
}

func assertTerminalMatches(t *testing.T, expected nodeMatcher, actual terminalsgoparser.TreeNode) {
	assertNodeMatches(t, terminalsNodeAccessors, expected, actual)
}

func expectTerminalParseError(t *testing.T, input string) {
	t.Helper()

	expectParseError(t, terminalsParse, func(err error) bool {
		var parseErr *terminalsgoparser.ParseError
		return errors.As(err, &parseErr)
	}, input)
}

func TestAnyCharParsesAnySingleCharacter(t *testing.T) {
	assertTerminalMatches(t, node("a", 5), parseTerminal(t, "any: a"))
	assertTerminalMatches(t, node("!", 5), parseTerminal(t, "any: !"))
}

func TestAnyCharRejectsTheEmptyString(t *testing.T) {
	expectTerminalParseError(t, "any: ")
}

func TestAnyCharRejectsTooManyCharacters(t *testing.T) {
	expectTerminalParseError(t, "any: ab")
}

func TestCharClassParsesCharactersWithinTheClass(t *testing.T) {
	assertTerminalMatches(t, node("x", 11), parseTerminal(t, "pos-class: x"))
}

func TestCharClassRejectsCharactersOutsideTheClass(t *testing.T) {
	expectTerminalParseError(t, "pos-class: 0")
}

func TestCharClassMatchesCaseSensitively(t *testing.T) {
	expectTerminalParseError(t, "pos-class: A")
}

func TestNegativeCharClassParsesCharactersOutsideTheClass(t *testing.T) {
	assertTerminalMatches(t, node("0", 11), parseTerminal(t, "neg-class: 0"))
}

func TestNegativeCharClassRejectsCharactersWithinTheClass(t *testing.T) {
	expectTerminalParseError(t, "neg-class: x")
}

func TestSingleQuotedStringParsesExactString(t *testing.T) {
	assertTerminalMatches(t, node("oat", 7), parseTerminal(t, "str-1: oat"))
}

func TestSingleQuotedStringMatchesCaseSensitively(t *testing.T) {
	expectTerminalParseError(t, "str-1: OAT")
}

func TestSingleQuotedStringRejectsAdditionalPrefixes(t *testing.T) {
	expectTerminalParseError(t, "str-1: boat")
}

func TestSingleQuotedStringRejectsAdditionalSuffixes(t *testing.T) {
	expectTerminalParseError(t, "str-1: oath")
}

func TestSingleQuotedStringRejectsTheEmptyString(t *testing.T) {
	expectTerminalParseError(t, "str-1: ")
}

func TestSingleQuotedStringRejectsPrefixes(t *testing.T) {
	expectTerminalParseError(t, "str-1: oa")
}

func TestDoubleQuotedStringParsesExactString(t *testing.T) {
	assertTerminalMatches(t, node("oat", 7), parseTerminal(t, "str-2: oat"))
}

func TestDoubleQuotedStringMatchesCaseSensitively(t *testing.T) {
	expectTerminalParseError(t, "str-2: OAT")
}

func TestDoubleQuotedStringRejectsAdditionalPrefixes(t *testing.T) {
	expectTerminalParseError(t, "str-2: boat")
}

func TestDoubleQuotedStringRejectsAdditionalSuffixes(t *testing.T) {
	expectTerminalParseError(t, "str-2: oath")
}

func TestDoubleQuotedStringRejectsTheEmptyString(t *testing.T) {
	expectTerminalParseError(t, "str-2: ")
}

func TestDoubleQuotedStringRejectsPrefixes(t *testing.T) {
	expectTerminalParseError(t, "str-2: oa")
}

func TestCaseInsensitiveStringParsesExactString(t *testing.T) {
	assertTerminalMatches(t, node("oat", 8), parseTerminal(t, "str-ci: oat"))
}

func TestCaseInsensitiveStringMatchesCaseInsensitively(t *testing.T) {
	assertTerminalMatches(t, node("OAT", 8), parseTerminal(t, "str-ci: OAT"))
}

func TestCaseInsensitiveStringRejectsAdditionalPrefixes(t *testing.T) {
	expectTerminalParseError(t, "str-ci: boat")
}

func TestCaseInsensitiveStringRejectsAdditionalSuffixes(t *testing.T) {
	expectTerminalParseError(t, "str-ci: oath")
}

func TestCaseInsensitiveStringRejectsTheEmptyString(t *testing.T) {
	expectTerminalParseError(t, "str-ci: ")
}

func TestCaseInsensitiveStringRejectsPrefixes(t *testing.T) {
	expectTerminalParseError(t, "str-ci: oa")
}
