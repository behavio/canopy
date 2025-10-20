package test

import (
	"reflect"
	"testing"

	"nodeactionsgoparser"
)

var (
	nodeActionsAccessors = newNodeAccessors(
		func(n nodeactionsgoparser.TreeNode) string { return n.Text() },
		func(n nodeactionsgoparser.TreeNode) int { return n.Offset() },
		func(n nodeactionsgoparser.TreeNode) []nodeactionsgoparser.TreeNode { return n.Children() },
	)
	nodeActionsParse = func(input string, actions nodeactionsgoparser.Actions) (nodeactionsgoparser.TreeNode, error) {
		return nodeactionsgoparser.Parse(input, actions, nil)
	}
)

type actionNode struct {
	Tag      string
	Input    string
	Start    int
	End      int
	Elements []nodeactionsgoparser.TreeNode
	text     string
}

func newActionNode(tag, input string, start, end int, elements []nodeactionsgoparser.TreeNode) *actionNode {
	copied := append([]nodeactionsgoparser.TreeNode(nil), elements...)
	return &actionNode{
		Tag:      tag,
		Input:    input,
		Start:    start,
		End:      end,
		Elements: copied,
		text:     substring(input, start, end),
	}
}

func (n *actionNode) Text() string {
	return n.text
}

func (n *actionNode) Offset() int {
	return n.Start
}

func (n *actionNode) Children() []nodeactionsgoparser.TreeNode {
	return n.Elements
}

type valueNode struct {
	Value any
	Input string
	Start int
	End   int
	text  string
}

func newValueNode(value any, input string, start, end int) *valueNode {
	return &valueNode{
		Value: value,
		Input: input,
		Start: start,
		End:   end,
		text:  substring(input, start, end),
	}
}

func (n *valueNode) Text() string {
	return n.text
}

func (n *valueNode) Offset() int {
	return n.Start
}

func (n *valueNode) Children() []nodeactionsgoparser.TreeNode {
	return nil
}

type testActions struct{}

func (testActions) make(tag, input string, start, end int, elements []nodeactionsgoparser.TreeNode) nodeactionsgoparser.TreeNode {
	return newActionNode(tag, input, start, end, elements)
}

func (a testActions) Make0(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return newValueNode(0, input, start, end), nil
}

func (a testActions) MakeAny(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("any", input, start, end, elements), nil
}

func (a testActions) MakeChar(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("char", input, start, end, elements), nil
}

func (a testActions) MakeEmptyList(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return newValueNode([]any{}, input, start, end), nil
}

func (a testActions) MakeEmptyStr(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return newValueNode("", input, start, end), nil
}

func (a testActions) MakeFalse(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return newValueNode(false, input, start, end), nil
}

func (a testActions) MakeInt(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("int", input, start, end, elements), nil
}

func (a testActions) MakeMaybe(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("maybe", input, start, end, elements), nil
}

func (a testActions) MakeNull(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return newValueNode(nil, input, start, end), nil
}

func (a testActions) MakeParen(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("paren", input, start, end, elements), nil
}

func (a testActions) MakeRep(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("rep", input, start, end, elements), nil
}

func (a testActions) MakeRepParen(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("rep-paren", input, start, end, elements), nil
}

func (a testActions) MakeSeq(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("seq", input, start, end, elements), nil
}

func (a testActions) MakeStr(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("str", input, start, end, elements), nil
}

func (a testActions) MakeZero(input string, start, end int, elements []nodeactionsgoparser.TreeNode) (nodeactionsgoparser.TreeNode, error) {
	return a.make("zero", input, start, end, elements), nil
}

func parseNodeActionsRoot(t *testing.T, input string, actions nodeactionsgoparser.Actions) nodeactionsgoparser.TreeNode {
	t.Helper()

	tree, err := nodeActionsParse(input, actions)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}
	return tree
}

func parseNodeActionsResult(t *testing.T, input string, actions nodeactionsgoparser.Actions) nodeactionsgoparser.TreeNode {
	t.Helper()

	tree := parseNodeActionsRoot(t, input, actions)
	return secondChild(t, tree)
}

func secondChild(t *testing.T, node nodeactionsgoparser.TreeNode) nodeactionsgoparser.TreeNode {
	t.Helper()

	children := node.Children()
	if len(children) < 2 {
		t.Fatalf("expected at least two children, got %d", len(children))
	}
	return children[1]
}

func assertActionNode(t *testing.T, node nodeactionsgoparser.TreeNode, tag, input string, start, end int) *actionNode {
	t.Helper()

	result, ok := node.(*actionNode)
	if !ok {
		t.Fatalf("expected *actionNode, got %T", node)
	}
	if result.Tag != tag {
		t.Fatalf("expected tag %q, got %q", tag, result.Tag)
	}
	if result.Input != input {
		t.Fatalf("expected input %q, got %q", input, result.Input)
	}
	if result.Start != start || result.End != end {
		t.Fatalf("expected span (%d,%d), got (%d,%d)", start, end, result.Start, result.End)
	}
	return result
}

func valueFromNode(t *testing.T, node nodeactionsgoparser.TreeNode) any {
	t.Helper()

	value, ok := node.(*valueNode)
	if !ok {
		t.Fatalf("expected *valueNode, got %T", node)
	}
	return value.Value
}

func substring(text string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(text) {
		end = len(text)
	}
	if start >= end {
		return ""
	}
	return text[start:end]
}

func TestNodeActionsMakesNodesFromString(t *testing.T) {
	actions := &testActions{}
	input := "act-str: hello"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "str", input, 9, 14)

	if len(action.Elements) != 0 {
		t.Fatalf("expected no elements, got %d", len(action.Elements))
	}
}

func TestNodeActionsMakesNodesFromCharClass(t *testing.T) {
	actions := &testActions{}
	input := "act-class: k"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "char", input, 11, 12)

	if len(action.Elements) != 0 {
		t.Fatalf("expected no elements, got %d", len(action.Elements))
	}
}

func TestNodeActionsMakesNodesFromAnyChar(t *testing.T) {
	actions := &testActions{}
	input := "act-any: ?"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "any", input, 9, 10)

	if len(action.Elements) != 0 {
		t.Fatalf("expected no elements, got %d", len(action.Elements))
	}
}

func TestNodeActionsMakesNodesFromMaybeRule(t *testing.T) {
	actions := &testActions{}
	input := "act-maybe: hello"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "maybe", input, 11, 16)

	if len(action.Elements) != 0 {
		t.Fatalf("expected no elements, got %d", len(action.Elements))
	}
}

func TestNodeActionsDoesNotInvokeActionWhenMaybeRuleHasNoMatch(t *testing.T) {
	actions := &testActions{}
	input := "act-maybe: "

	root := parseNodeActionsRoot(t, input, actions)
	optional := secondChild(t, root)

	assertNodeMatches(t, nodeActionsAccessors, node("", 11), optional)
}

func TestNodeActionsMakesNodesFromRepetition(t *testing.T) {
	actions := &testActions{}
	input := "act-rep: abc"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "rep", input, 9, 12)

	if len(action.Elements) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(action.Elements))
	}

	assertNodeMatches(t, nodeActionsAccessors, node("a", 9), action.Elements[0])
	assertNodeMatches(t, nodeActionsAccessors, node("b", 10), action.Elements[1])
	assertNodeMatches(t, nodeActionsAccessors, node("c", 11), action.Elements[2])
}

func TestNodeActionsMakesNodesFromRepetitionInParentheses(t *testing.T) {
	actions := &testActions{}
	input := "act-rep-paren: abab"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "rep-paren", input, 15, 19)

	if len(action.Elements) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(action.Elements))
	}

	assertNodeMatches(t, nodeActionsAccessors, node("ab", 15, node("a", 15), node("b", 16)), action.Elements[0])
	assertNodeMatches(t, nodeActionsAccessors, node("ab", 17, node("a", 17), node("b", 18)), action.Elements[1])
}

func TestNodeActionsMakesNodesFromSequence(t *testing.T) {
	actions := &testActions{}
	input := "act-seq: xyz"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "seq", input, 9, 12)

	if len(action.Elements) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(action.Elements))
	}

	assertNodeMatches(t, nodeActionsAccessors, node("x", 9), action.Elements[0])
	assertNodeMatches(t, nodeActionsAccessors, node("y", 10), action.Elements[1])
	assertNodeMatches(t, nodeActionsAccessors, node("z", 11), action.Elements[2])
}

func TestNodeActionsMakesNodesFromSequenceWithMutedElements(t *testing.T) {
	actions := &testActions{}
	input := "act-seq-mute: xyz"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "seq", input, 14, 17)

	if len(action.Elements) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(action.Elements))
	}

	assertNodeMatches(t, nodeActionsAccessors, node("x", 14), action.Elements[0])
	assertNodeMatches(t, nodeActionsAccessors, node("z", 16), action.Elements[1])
}

func TestNodeActionsMakesNodesFromParenthesisedExpression(t *testing.T) {
	actions := &testActions{}
	input := "act-paren: !"

	result := parseNodeActionsResult(t, input, actions)
	action := assertActionNode(t, result, "paren", input, 11, 12)

	if len(action.Elements) != 0 {
		t.Fatalf("expected no elements, got %d", len(action.Elements))
	}
}

func TestNodeActionsBindsToChoiceOptions(t *testing.T) {
	actions := &testActions{}
	zeroInput := "act-choice: 0"

	zero := parseNodeActionsResult(t, zeroInput, actions)
	zeroAction := assertActionNode(t, zero, "zero", zeroInput, 12, 13)
	if len(zeroAction.Elements) != 0 {
		t.Fatalf("expected no elements, got %d", len(zeroAction.Elements))
	}

	intInput := "act-choice: 42"
	integer := parseNodeActionsResult(t, intInput, actions)
	intAction := assertActionNode(t, integer, "int", intInput, 12, 14)

	if len(intAction.Elements) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(intAction.Elements))
	}

	assertNodeMatches(t, nodeActionsAccessors, node("4", 12), intAction.Elements[0])
	assertNodeMatches(t, nodeActionsAccessors, node("2", 13, node("2", 13)), intAction.Elements[1])
}

func TestNodeActionsTreatsNullAsValidResult(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey: null", actions)

	if value := valueFromNode(t, result); value != nil {
		t.Fatalf("expected nil value, got %v", value)
	}
}

func TestNodeActionsTreatsFalseAsValidResult(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey: false", actions)

	if value := valueFromNode(t, result); value != false {
		t.Fatalf("expected false, got %v", value)
	}
}

func TestNodeActionsTreatsZeroAsValidResult(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey: 0", actions)

	if value := valueFromNode(t, result); value != 0 {
		t.Fatalf("expected 0, got %v", value)
	}
}

func TestNodeActionsTreatsEmptyStringAsValidResult(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey: ''", actions)

	if value := valueFromNode(t, result); value != "" {
		t.Fatalf("expected empty string, got %v", value)
	}
}

func TestNodeActionsTreatsEmptyListAsValidResult(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey: []", actions)

	if value := valueFromNode(t, result); !reflect.DeepEqual(value, []any{}) {
		t.Fatalf("expected empty list, got %v", value)
	}
}

func TestNodeActionsAcceptsFalseyLookaheadResults(t *testing.T) {
	actions := &testActions{}
	input := "act-falsey-pred: 0"

	root := parseNodeActionsRoot(t, input, actions)
	predicate := secondChild(t, root)

	children := predicate.Children()
	if len(children) < 2 {
		t.Fatalf("expected predicate to have at least two children, got %d", len(children))
	}

	assertActionNode(t, children[1], "zero", input, 17, 18)
}

func TestNodeActionsAcceptsFalseyRepetitionResults(t *testing.T) {
	actions := &testActions{}
	input := "act-falsey-rep: null0false''[]"

	root := parseNodeActionsRoot(t, input, actions)
	repetition := secondChild(t, root)

	values := make([]any, len(repetition.Children()))
	for i, child := range repetition.Children() {
		values[i] = valueFromNode(t, child)
	}

	expected := []any{nil, 0, false, "", []any{}}
	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %v, got %v", expected, values)
	}
}

func TestNodeActionsAcceptsFalseyMaybeResults(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey-opt: null", actions)

	if value := valueFromNode(t, result); value != nil {
		t.Fatalf("expected nil value, got %v", value)
	}
}

func TestNodeActionsAcceptsFalseySequenceResults(t *testing.T) {
	actions := &testActions{}
	input := "act-falsey-seq: (null)"

	sequence := parseNodeActionsResult(t, input, actions)
	children := sequence.Children()

	if len(children) != 3 {
		t.Fatalf("expected 3 children, got %d", len(children))
	}

	assertNodeMatches(t, nodeActionsAccessors, node("(", 16), children[0])
	if value := valueFromNode(t, children[1]); value != nil {
		t.Fatalf("expected nil middle value, got %v", value)
	}
	assertNodeMatches(t, nodeActionsAccessors, node(")", 21), children[2])
}

func TestNodeActionsAcceptsFalseyChoiceResults(t *testing.T) {
	actions := &testActions{}
	result := parseNodeActionsResult(t, "act-falsey-choice: null", actions)

	if value := valueFromNode(t, result); value != nil {
		t.Fatalf("expected nil value, got %v", value)
	}
}
