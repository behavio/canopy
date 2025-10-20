package test

import (
	"errors"
	"testing"

	"choicesgoparser"
)

var (
	choicesNodeAccessors = newNodeAccessors(
		func(n choicesgoparser.TreeNode) string { return n.Text() },
		func(n choicesgoparser.TreeNode) int { return n.Offset() },
		func(n choicesgoparser.TreeNode) []choicesgoparser.TreeNode { return n.Children() },
	)
	choicesParse = func(input string) (choicesgoparser.TreeNode, error) {
		return choicesgoparser.Parse(input, nil, nil)
	}
)

func parseChoice(t *testing.T, input string) choicesgoparser.TreeNode {
	t.Helper()

	tree, err := choicesParse(input)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}

	children := tree.Children()
	if len(children) < 2 {
		t.Fatalf("parse(%q) expected at least two root children, got %d", input, len(children))
	}

	return children[1]
}

func assertChoiceMatches(t *testing.T, expected nodeMatcher, actual choicesgoparser.TreeNode) {
	assertNodeMatches(t, choicesNodeAccessors, expected, actual)
}

func expectChoiceParseError(t *testing.T, input string) {
	t.Helper()

	expectParseError(t, choicesParse, func(err error) bool {
		var parseErr *choicesgoparser.ParseError
		return errors.As(err, &parseErr)
	}, input)
}

func TestChoiceStringsParsesAnyOfTheChoiceOptions(t *testing.T) {
	assertChoiceMatches(t, node("a", 12), parseChoice(t, "choice-abc: a"))
	assertChoiceMatches(t, node("b", 12), parseChoice(t, "choice-abc: b"))
	assertChoiceMatches(t, node("c", 12), parseChoice(t, "choice-abc: c"))
}

func TestChoiceStringsRejectsInputMatchingNoneOfTheOptions(t *testing.T) {
	expectChoiceParseError(t, "choice-abc: d")
}

func TestChoiceStringsRejectsSuperstringsOfTheOptions(t *testing.T) {
	expectChoiceParseError(t, "choice-abc: ab")
}

func TestChoiceStringsParsesAChoiceAsPartOfASequence(t *testing.T) {
	expected := node(
		"repeat",
		12,
		node("re", 12),
		node("peat", 14),
	)

	assertChoiceMatches(t, expected, parseChoice(t, "choice-seq: repeat"))
}

func TestChoiceStringsDoesNotBacktrackIfLaterRulesFail(t *testing.T) {
	expectChoiceParseError(t, "choice-seq: reppeat")
}

func TestChoiceRepetitionParsesDifferentOptionOnEachIteration(t *testing.T) {
	expected := node(
		"abcabba",
		12,
		node("a", 12),
		node("b", 13),
		node("c", 14),
		node("a", 15),
		node("b", 16),
		node("b", 17),
		node("a", 18),
	)

	assertChoiceMatches(t, expected, parseChoice(t, "choice-rep: abcabba"))
}

func TestChoiceRepetitionRejectsIfAnyIterationDoesNotMatchOptions(t *testing.T) {
	expectChoiceParseError(t, "choice-rep: abcadba")
}

func TestChoiceSequenceParsesOneBranchOfTheChoice(t *testing.T) {
	expected := node(
		"ab",
		13,
		node("a", 13),
		node("b", 14),
	)

	assertChoiceMatches(t, expected, parseChoice(t, "choice-bind: ab"))
}

func TestChoiceSequenceBindsSequencesTighterThanChoices(t *testing.T) {
	expectChoiceParseError(t, "choice-bind: abef")
}
