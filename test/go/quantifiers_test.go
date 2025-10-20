package test

import (
	"errors"
	"testing"

	"quantifiersgoparser"
)

var (
	quantifiersNodeAccessors = newNodeAccessors(
		func(n quantifiersgoparser.TreeNode) string { return n.Text() },
		func(n quantifiersgoparser.TreeNode) int { return n.Offset() },
		func(n quantifiersgoparser.TreeNode) []quantifiersgoparser.TreeNode { return n.Children() },
	)
	quantifiersParse = func(input string) (quantifiersgoparser.TreeNode, error) {
		return quantifiersgoparser.Parse(input, nil, nil)
	}
)

func parseQuantifier(t *testing.T, input string) quantifiersgoparser.TreeNode {
	t.Helper()

	tree, err := quantifiersParse(input)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}

	children := tree.Children()
	if len(children) < 2 {
		t.Fatalf("parse(%q) expected at least two root children, got %d", input, len(children))
	}

	return children[1]
}

func assertQuantifierMatches(t *testing.T, expected nodeMatcher, actual quantifiersgoparser.TreeNode) {
	assertNodeMatches(t, quantifiersNodeAccessors, expected, actual)
}

func expectQuantifierParseError(t *testing.T, input string) {
	t.Helper()

	expectParseError(t, quantifiersParse, func(err error) bool {
		var parseErr *quantifiersgoparser.ParseError
		return errors.As(err, &parseErr)
	}, input)
}

func TestMaybeParsesMatchingCharacter(t *testing.T) {
	assertQuantifierMatches(t, node("4", 7), parseQuantifier(t, "maybe: 4"))
}

func TestMaybeParsesTheEmptyString(t *testing.T) {
	assertQuantifierMatches(t, node("", 7), parseQuantifier(t, "maybe: "))
}

func TestMaybeRejectsNonMatchingString(t *testing.T) {
	expectQuantifierParseError(t, "maybe: a")
}

func TestZeroOrMoreParsesTheEmptyString(t *testing.T) {
	assertQuantifierMatches(t, node("", 7), parseQuantifier(t, "rep-0: "))
}

func TestZeroOrMoreParsesOneOccurrence(t *testing.T) {
	expected := node("z", 7, node("z", 7))
	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-0: z"))
}

func TestZeroOrMoreParsesManyOccurrencesSameInstance(t *testing.T) {
	expected := node(
		"zzzz",
		7,
		node("z", 7),
		node("z", 8),
		node("z", 9),
		node("z", 10),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-0: zzzz"))
}

func TestZeroOrMoreParsesManyOccurrencesDifferentInstances(t *testing.T) {
	expected := node(
		"wxyz",
		7,
		node("w", 7),
		node("x", 8),
		node("y", 9),
		node("z", 10),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-0: wxyz"))
}

func TestZeroOrMoreRejectsStringsWithNonMatchingPrefix(t *testing.T) {
	expectQuantifierParseError(t, "rep-0: 4x")
}

func TestZeroOrMoreRejectsStringsWithNonMatchingSuffix(t *testing.T) {
	expectQuantifierParseError(t, "rep-0: x4")
}

func TestZeroOrMoreParsesRepetitionsGreedily(t *testing.T) {
	expectQuantifierParseError(t, "greedy-0: xy")
}

func TestOneOrMoreRejectsTheEmptyString(t *testing.T) {
	expectQuantifierParseError(t, "rep-1: ")
}

func TestOneOrMoreParsesOneOccurrence(t *testing.T) {
	expected := node("z", 7, node("z", 7))
	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-1: z"))
}

func TestOneOrMoreParsesManyOccurrencesSameInstance(t *testing.T) {
	expected := node(
		"zzzz",
		7,
		node("z", 7),
		node("z", 8),
		node("z", 9),
		node("z", 10),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-1: zzzz"))
}

func TestOneOrMoreParsesManyOccurrencesDifferentInstances(t *testing.T) {
	expected := node(
		"wxyz",
		7,
		node("w", 7),
		node("x", 8),
		node("y", 9),
		node("z", 10),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-1: wxyz"))
}

func TestOneOrMoreRejectsStringsWithNonMatchingPrefix(t *testing.T) {
	expectQuantifierParseError(t, "rep-1: 4x")
}

func TestOneOrMoreRejectsStringsWithNonMatchingSuffix(t *testing.T) {
	expectQuantifierParseError(t, "rep-1: x4")
}

func TestOneOrMoreParsesRepetitionsGreedily(t *testing.T) {
	expectQuantifierParseError(t, "greedy-1: xy")
}

func TestOneOrMoreParsesRepeatedReference(t *testing.T) {
	expected := node(
		"#abc123",
		11,
		node("#", 11),
		node(
			"abc123",
			12,
			node("a", 12),
			node("b", 13),
			node("c", 14),
			node("1", 15),
			node("2", 16),
			node("3", 17),
		),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "color-ref: #abc123"))
}

func TestOneOrMoreParsesRepeatedChoice(t *testing.T) {
	expected := node(
		"#abc123",
		14,
		node("#", 14),
		node(
			"abc123",
			15,
			node("a", 15),
			node("b", 16),
			node("c", 17),
			node("1", 18),
			node("2", 19),
			node("3", 20),
		),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "color-choice: #abc123"))
}

func TestExactlyRejectsTheEmptyString(t *testing.T) {
	expectQuantifierParseError(t, "rep-exact: ")
}

func TestExactlyParsesRequiredNumberOfPattern(t *testing.T) {
	expected := node(
		"abc",
		11,
		node("a", 11),
		node("b", 12),
		node("c", 13),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-exact: abc"))
}

func TestExactlyRejectsTooFewCopies(t *testing.T) {
	expectQuantifierParseError(t, "rep-exact: ab")
}

func TestExactlyRejectsTooManyCopies(t *testing.T) {
	expectQuantifierParseError(t, "rep-exact: abcd")
}

func TestMinimumRejectsTheEmptyString(t *testing.T) {
	expectQuantifierParseError(t, "rep-min: ")
}

func TestMinimumParsesMinimumNumberOfPattern(t *testing.T) {
	expected := node(
		"abc",
		9,
		node("a", 9),
		node("b", 10),
		node("c", 11),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-min: abc"))
}

func TestMinimumParsesMoreCopiesOfPattern(t *testing.T) {
	expected := node(
		"abcdef",
		9,
		node("a", 9),
		node("b", 10),
		node("c", 11),
		node("d", 12),
		node("e", 13),
		node("f", 14),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-min: abcdef"))
}

func TestMinimumRejectsTooFewCopies(t *testing.T) {
	expectQuantifierParseError(t, "rep-min: ab")
}

func TestRangeRejectsTheEmptyString(t *testing.T) {
	expectQuantifierParseError(t, "rep-range: ")
}

func TestRangeParsesMinimumNumberOfPattern(t *testing.T) {
	expected := node(
		"abc",
		11,
		node("a", 11),
		node("b", 12),
		node("c", 13),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-range: abc"))
}

func TestRangeParsesMaximumNumberOfPattern(t *testing.T) {
	expected := node(
		"abcde",
		11,
		node("a", 11),
		node("b", 12),
		node("c", 13),
		node("d", 14),
		node("e", 15),
	)

	assertQuantifierMatches(t, expected, parseQuantifier(t, "rep-range: abcde"))
}

func TestRangeRejectsTooFewCopies(t *testing.T) {
	expectQuantifierParseError(t, "rep-range: ab")
}

func TestRangeRejectsTooManyCopies(t *testing.T) {
	expectQuantifierParseError(t, "rep-range: abcdef")
}
