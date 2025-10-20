package test

import (
	"slices"
	"testing"

	"extensionsgoparser"
)

var (
	extensionsTypes = map[string]extensionsgoparser.NodeExtender{
		"Ext": func(node extensionsgoparser.TreeNode) extensionsgoparser.TreeNode {
			if node == nil {
				return nil
			}
			return &extNode{TreeNode: node}
		},
		"NS.Ext": func(node extensionsgoparser.TreeNode) extensionsgoparser.TreeNode {
			if node == nil {
				return nil
			}
			return &nsExtNode{TreeNode: node}
		},
	}
	extensionsParse = func(input string) (extensionsgoparser.TreeNode, error) {
		return extensionsgoparser.Parse(input, nil, extensionsTypes)
	}
)

type extNode struct {
	extensionsgoparser.TreeNode
}

func (n *extNode) ExtFunc() (int, []string) {
	return len(n.Children()), splitChars(n.Text())
}

type nsExtNode struct {
	extensionsgoparser.TreeNode
}

func (n *nsExtNode) NsFunc() bool {
	return true
}

func splitChars(text string) []string {
	runes := []rune(text)
	chars := make([]string, len(runes))
	for i, r := range runes {
		chars[i] = string(r)
	}
	return chars
}

func parseExtensions(t *testing.T, input string) extensionsgoparser.TreeNode {
	t.Helper()

	tree, err := extensionsParse(input)
	if err != nil {
		t.Fatalf("parse(%q) returned unexpected error: %v", input, err)
	}

	children := tree.Children()
	if len(children) < 2 {
		t.Fatalf("parse(%q) expected at least two root children, got %d", input, len(children))
	}

	return children[1]
}

func assertExtFunc(t *testing.T, node extensionsgoparser.TreeNode, wantLen int, wantChars []string) {
	t.Helper()

	callable, ok := node.(interface {
		ExtFunc() (int, []string)
	})
	if !ok {
		t.Fatalf("node of type %T does not implement ExtFunc", node)
	}

	gotLen, gotChars := callable.ExtFunc()
	if gotLen != wantLen {
		t.Fatalf("ExtFunc length mismatch: want %d, got %d", wantLen, gotLen)
	}
	if !slices.Equal(gotChars, wantChars) {
		t.Fatalf("ExtFunc chars mismatch: want %v, got %v", wantChars, gotChars)
	}
}

func TestExtensionsAddsMethodsToString(t *testing.T) {
	result := parseExtensions(t, "ext-str: hello")
	assertExtFunc(t, result, 0, []string{"h", "e", "l", "l", "o"})
}

func TestExtensionsAddsMethodsToCharClass(t *testing.T) {
	result := parseExtensions(t, "ext-class: k")
	assertExtFunc(t, result, 0, []string{"k"})
}

func TestExtensionsAddsMethodsToAnyChar(t *testing.T) {
	result := parseExtensions(t, "ext-any: ?")
	assertExtFunc(t, result, 0, []string{"?"})
}

func TestExtensionsAddsMethodsToMaybeRule(t *testing.T) {
	result := parseExtensions(t, "ext-maybe: hello")
	assertExtFunc(t, result, 0, []string{"h", "e", "l", "l", "o"})
}

func TestExtensionsAddsMethodsToRepetition(t *testing.T) {
	result := parseExtensions(t, "ext-rep: abc")
	assertExtFunc(t, result, 3, []string{"a", "b", "c"})
}

func TestExtensionsAddsMethodsToSequence(t *testing.T) {
	result := parseExtensions(t, "ext-seq: xyz")
	assertExtFunc(t, result, 3, []string{"x", "y", "z"})
}

func TestExtensionsAddsMethodsToParenthesisedExpression(t *testing.T) {
	result := parseExtensions(t, "ext-paren: !")
	assertExtFunc(t, result, 0, []string{"!"})
}

func TestExtensionsAddsMethodsToChoiceOptions(t *testing.T) {
	first := parseExtensions(t, "ext-choice: 0")
	assertExtFunc(t, first, 0, []string{"0"})

	second := parseExtensions(t, "ext-choice: 42")
	assertExtFunc(t, second, 2, []string{"4", "2"})
}

func TestExtensionsAddsMethodsToReferenceResult(t *testing.T) {
	result := parseExtensions(t, "ext-ref: hello")
	assertExtFunc(t, result, 0, []string{"h", "e", "l", "l", "o"})
}

func TestExtensionsAddsFromNamespacedModule(t *testing.T) {
	result := parseExtensions(t, "ext-ns: hello")

	callable, ok := result.(interface {
		NsFunc() bool
	})
	if !ok {
		t.Fatalf("node of type %T does not implement NsFunc", result)
	}
	if !callable.NsFunc() {
		t.Fatalf("NsFunc() returned false")
	}
}
