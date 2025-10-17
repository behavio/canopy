package test

import (
	"errors"
	"reflect"
	"testing"

	"sequencesgoparser"
)

var (
	sequencesNodeAccessors = newNodeAccessors(
		func(n sequencesgoparser.TreeNode) string { return n.Text() },
		func(n sequencesgoparser.TreeNode) int { return n.Offset() },
		func(n sequencesgoparser.TreeNode) []sequencesgoparser.TreeNode { return n.Children() },
	)
	sequencesParse = func(input string) (sequencesgoparser.TreeNode, error) {
		return sequencesgoparser.Parse(input, nil, nil)
	}
)

func parseSequence(t *testing.T, input string) sequencesgoparser.TreeNode {
	t.Helper()

	tree, err := sequencesParse(input)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}

	children := tree.Children()
	if len(children) < 2 {
		t.Fatalf("parse(%q) expected at least two root children, got %d", input, len(children))
	}

	return children[1]
}

func assertSequenceMatches(t *testing.T, expected nodeMatcher, actual sequencesgoparser.TreeNode) {
	assertNodeMatches(t, sequencesNodeAccessors, expected, actual)
}

func expectSequenceParseError(t *testing.T, input string) {
	t.Helper()

	expectParseError(t, sequencesParse, func(err error) bool {
		var parseErr *sequencesgoparser.ParseError
		return errors.As(err, &parseErr)
	}, input)
}

func TestSequenceStringsParsesMatchingSequence(t *testing.T) {
	expected := node(
		"abc",
		9,
		node("a", 9),
		node("b", 10),
		node("c", 11),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-str: abc"))
}

func TestSequenceStringsRejectsMissingPrefix(t *testing.T) {
	expectSequenceParseError(t, "seq-str: bc")
}

func TestSequenceStringsRejectsAdditionalPrefix(t *testing.T) {
	expectSequenceParseError(t, "seq-str: zabc")
}

func TestSequenceStringsRejectsMissingMiddle(t *testing.T) {
	expectSequenceParseError(t, "seq-str: ac")
}

func TestSequenceStringsRejectsAdditionalMiddle(t *testing.T) {
	expectSequenceParseError(t, "seq-str: azbzc")
}

func TestSequenceStringsRejectsMissingSuffix(t *testing.T) {
	expectSequenceParseError(t, "seq-str: ab")
}

func TestSequenceStringsRejectsAdditionalSuffix(t *testing.T) {
	expectSequenceParseError(t, "seq-str: abcz")
}

func TestSequenceMaybesParsesAtStart(t *testing.T) {
	expected := node(
		"bc",
		13,
		node("", 13),
		node("b", 13),
		node("c", 14),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-maybe-1: bc"))
}

func TestSequenceMaybesParsesInMiddle(t *testing.T) {
	expected := node(
		"ac",
		13,
		node("a", 13),
		node("", 14),
		node("c", 14),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-maybe-2: ac"))
}

func TestSequenceMaybesParsesAtEnd(t *testing.T) {
	expected := node(
		"ab",
		13,
		node("a", 13),
		node("b", 14),
		node("", 15),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-maybe-3: ab"))
}

func TestSequenceRepetitionAllowsEmptyMatches(t *testing.T) {
	expected := node(
		"0",
		11,
		node("", 11),
		node("0", 11),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-rep-1: 0"))
}

func TestSequenceRepetitionAllowsNonEmptyMatches(t *testing.T) {
	expected := node(
		"abc0",
		11,
		node(
			"abc",
			11,
			node("a", 11),
			node("b", 12),
			node("c", 13),
		),
		node("0", 14),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-rep-1: abc0"))
}

func TestSequenceRepetitionParsesRepetitionsGreedily(t *testing.T) {
	expectSequenceParseError(t, "seq-rep-2: aaa")
}

func TestSequenceRepeatedSubSequenceParsesNestedTree(t *testing.T) {
	expected := node(
		"ab1b2b3c",
		16,
		node("a", 16),
		node(
			"b1b2b3",
			17,
			node(
				"b1",
				17,
				node("b", 17),
				node("1", 18),
			),
			node(
				"b2",
				19,
				node("b", 19),
				node("2", 20),
			),
			node(
				"b3",
				21,
				node("b", 21),
				node("3", 22),
			),
		),
		node("c", 23),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-rep-subseq: ab1b2b3c"))
}

func TestSequenceRepeatedSubSequenceRejectsMismatchedInput(t *testing.T) {
	expectSequenceParseError(t, "seq-rep-subseq: ab1b2bc")
}

func TestSequenceLabellingCreatesNamedReferences(t *testing.T) {
	actual := parseSequence(t, "seq-label: v987")

	expected := node(
		"v987",
		11,
		node("v", 11),
		node(
			"987",
			12,
			node("9", 12),
			node("8", 13),
			node("7", 14),
		),
	)

	assertSequenceMatches(t, expected, actual)

	labelNode, ok := actual.(*sequencesgoparser.Node18)
	if !ok {
		t.Fatalf("seq-label node type %T does not expose labels", actual)
	}

	assertSequenceMatches(t, node(
		"987",
		12,
		node("9", 12),
		node("8", 13),
		node("7", 14),
	), labelNode.Num)
}

func TestSequenceLabellingCreatesReferencesInsideRepeatedSubSequences(t *testing.T) {
	actual := parseSequence(t, "seq-label-subseq: v.AB.CD.EF")

	expected := node(
		"v.AB.CD.EF",
		18,
		node("v", 18),
		node(
			".AB.CD.EF",
			19,
			node(
				".AB",
				19,
				node(".", 19),
				node(
					"AB",
					20,
					node("A", 20),
					node("B", 21),
				),
			),
			node(
				".CD",
				22,
				node(".", 22),
				node(
					"CD",
					23,
					node("C", 23),
					node("D", 24),
				),
			),
			node(
				".EF",
				25,
				node(".", 25),
				node(
					"EF",
					26,
					node("E", 26),
					node("F", 27),
				),
			),
		),
	)

	assertSequenceMatches(t, expected, actual)

	tailChildren := sequencesNodeAccessors.children(sequencesNodeAccessors.children(actual)[1])
	if len(tailChildren) != 3 {
		t.Fatalf("expected three labelled subsequences, got %d", len(tailChildren))
	}

	partMatchers := []nodeMatcher{
		node("AB", 20, node("A", 20), node("B", 21)),
		node("CD", 23, node("C", 23), node("D", 24)),
		node("EF", 26, node("E", 26), node("F", 27)),
	}

	for i, child := range tailChildren {
		labelled, ok := child.(*sequencesgoparser.Node19)
		if !ok {
			t.Fatalf("expected labelled subsequence node, got %T", child)
		}
		assertSequenceMatches(t, partMatchers[i], labelled.Part)
	}
}

func TestSequenceMutingRemovesChildNodes(t *testing.T) {
	expected := node(
		"key: 42",
		12,
		node(
			"key",
			12,
			node("k", 12),
			node("e", 13),
			node("y", 14),
		),
		node(
			"42",
			17,
			node("4", 17),
			node("2", 18),
		),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-mute-1: key: 42"))
}

func TestSequenceMutingRemovesChildSequences(t *testing.T) {
	expected := node(
		"key: 42",
		12,
		node(
			"key",
			12,
			node("k", 12),
			node("e", 13),
			node("y", 14),
		),
		node(
			"42",
			17,
			node("4", 17),
			node("2", 18),
		),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-mute-2: key: 42"))
}

func TestSequenceMutingRemovesNodesFromChildSequences(t *testing.T) {
	expected := node(
		"v.AB.CD.EF",
		12,
		node("v", 12),
		node(
			".AB.CD.EF",
			13,
			node(
				".AB",
				13,
				node("AB", 14, node("A", 14), node("B", 15)),
			),
			node(
				".CD",
				16,
				node("CD", 17, node("C", 17), node("D", 18)),
			),
			node(
				".EF",
				19,
				node("EF", 20, node("E", 20), node("F", 21)),
			),
		),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-mute-3: v.AB.CD.EF"))
}

func TestSequenceMutingHandlesNestedExpressions(t *testing.T) {
	expected := node(
		"abcde",
		12,
		node("a", 12),
		node("e", 16),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-mute-4: abcde"))
}

func TestSequenceMutingAllowsFirstElementMuted(t *testing.T) {
	expected := node(
		"abc",
		16,
		node("b", 17),
		node("c", 18),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-mute-first: abc"))
}

func TestSequenceMutingAllowsLastElementMuted(t *testing.T) {
	expected := node(
		"abc",
		15,
		node("a", 15),
		node("b", 16),
	)

	assertSequenceMatches(t, expected, parseSequence(t, "seq-mute-last: abc"))
}

func TestSequenceMutingRejectsMissingMutedExpressions(t *testing.T) {
	expectSequenceParseError(t, "seq-mute-4: ae")
	expectSequenceParseError(t, "seq-mute-4: abde")
}

func TestSequenceReferencesAssignLabels(t *testing.T) {
	actual := parseSequence(t, "seq-refs: ac")

	expected := node(
		"ac",
		10,
		node("a", 10),
		node("c", 11),
	)

	assertSequenceMatches(t, expected, actual)

	refsNode, ok := actual.(*sequencesgoparser.Node20)
	if !ok {
		t.Fatalf("seq-refs node type %T does not expose labels", actual)
	}

	assertSequenceMatches(t, node("a", 10), refsNode.A)
	assertSequenceMatches(t, node("c", 11), refsNode.B)
	assertSequenceMatches(t, node("c", 11), refsNode.C)
}

func TestSequenceReferencesMuteLabelsWhenRequested(t *testing.T) {
	actual := parseSequence(t, "seq-mute-refs: ac")

	expected := node(
		"ac",
		15,
		node("a", 15),
	)

	assertSequenceMatches(t, expected, actual)

	muteNode, ok := actual.(*sequencesgoparser.Node21)
	if !ok {
		t.Fatalf("seq-mute-refs node type %T does not match expected", actual)
	}

	assertSequenceMatches(t, node("a", 15), muteNode.A)

	if _, ok := reflect.TypeOf(muteNode).Elem().FieldByName("C"); ok {
		t.Fatalf("seq-mute-refs node should not expose label C")
	}
}
