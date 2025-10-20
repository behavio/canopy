package test

import (
	"errors"
	"testing"

	"predicatesgoparser"
)

var (
	predicatesNodeAccessors = newNodeAccessors(
		func(n predicatesgoparser.TreeNode) string { return n.Text() },
		func(n predicatesgoparser.TreeNode) int { return n.Offset() },
		func(n predicatesgoparser.TreeNode) []predicatesgoparser.TreeNode { return n.Children() },
	)
	predicatesParse = func(input string) (predicatesgoparser.TreeNode, error) {
		return predicatesgoparser.Parse(input, nil, nil)
	}
)

func parsePredicate(t *testing.T, input string) predicatesgoparser.TreeNode {
	t.Helper()

	tree, err := predicatesParse(input)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}

	children := tree.Children()
	if len(children) < 2 {
		t.Fatalf("parse(%q) expected at least two root children, got %d", input, len(children))
	}

	return children[1]
}

func assertPredicateMatches(t *testing.T, expected nodeMatcher, actual predicatesgoparser.TreeNode) {
	assertNodeMatches(t, predicatesNodeAccessors, expected, actual)
}

func expectPredicateParseError(t *testing.T, input string) {
	t.Helper()

	expectParseError(t, predicatesParse, func(err error) bool {
		var parseErr *predicatesgoparser.ParseError
		return errors.As(err, &parseErr)
	}, input)
}

func TestPositiveLookaheadChecksFirstCharacterOfWord(t *testing.T) {
	expected := node(
		"London",
		10,
		node("", 10),
		node(
			"London",
			10,
			node("L", 10),
			node("o", 11),
			node("n", 12),
			node("d", 13),
			node("o", 14),
			node("n", 15),
		),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "pos-name: London"))
}

func TestPositiveLookaheadRejectsWhenPredicateFails(t *testing.T) {
	expectPredicateParseError(t, "pos-name: london")
}

func TestPositiveLookaheadResetsCursorAfterMatching(t *testing.T) {
	expected := node(
		"<abc123>",
		9,
		node("", 9),
		node("<", 9),
		node(
			"abc123",
			10,
			node("a", 10),
			node("b", 11),
			node("c", 12),
			node("1", 13),
			node("2", 14),
			node("3", 15),
		),
		node(">", 16),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "pos-seq: <abc123>"))
}

func TestPositiveLookaheadUsesReferenceAsPredicate(t *testing.T) {
	expected := node(
		"c99",
		9,
		node("", 9),
		node(
			"c99",
			9,
			node("c", 9),
			node("9", 10),
			node("9", 11),
		),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "pos-ref: c99"))
}

func TestNegativeLookaheadChecksFirstCharacterOfWord(t *testing.T) {
	expected := node(
		"word",
		10,
		node("", 10),
		node(
			"word",
			10,
			node("w", 10),
			node("o", 11),
			node("r", 12),
			node("d", 13),
		),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "neg-name: word"))
}

func TestNegativeLookaheadRejectsWhenPredicateMatches(t *testing.T) {
	expectPredicateParseError(t, "neg-name: Word")
}

func TestNegativeLookaheadChecksTailString(t *testing.T) {
	expected := node(
		"word",
		14,
		node("word", 14),
		node("", 18),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "neg-tail-str: word"))
}

func TestNegativeLookaheadChecksTailClass(t *testing.T) {
	expected := node(
		"word",
		16,
		node("word", 16),
		node("", 20),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "neg-tail-class: word"))
}

func TestNegativeLookaheadChecksTailAnyChar(t *testing.T) {
	expected := node(
		"word",
		14,
		node("word", 14),
		node("", 18),
	)

	assertPredicateMatches(t, expected, parsePredicate(t, "neg-tail-any: word"))
}

func TestNegativeLookaheadRejectsMatchingNegativePattern(t *testing.T) {
	expectPredicateParseError(t, "neg-tail-str: wordmore text")
	expectPredicateParseError(t, "neg-tail-class: words")
	expectPredicateParseError(t, "neg-tail-any: word ")
}
