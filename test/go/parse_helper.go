package test

import "testing"

type nodeMatcher struct {
	text     string
	offset   int
	children []nodeMatcher
}

func node(text string, offset int, children ...nodeMatcher) nodeMatcher {
	return nodeMatcher{text: text, offset: offset, children: children}
}

type nodeAccessors[T any] struct {
	text     func(T) string
	offset   func(T) int
	children func(T) []T
}

func newNodeAccessors[T any](text func(T) string, offset func(T) int, children func(T) []T) nodeAccessors[T] {
	return nodeAccessors[T]{
		text:     text,
		offset:   offset,
		children: children,
	}
}

func assertNodeMatches[T any](t *testing.T, accessors nodeAccessors[T], expected nodeMatcher, actual T) {
	t.Helper()

	if accessors.text(actual) != expected.text {
		t.Fatalf("expected node text %q, got %q", expected.text, accessors.text(actual))
	}

	if accessors.offset(actual) != expected.offset {
		t.Fatalf("expected node offset %d, got %d", expected.offset, accessors.offset(actual))
	}

	actualChildren := accessors.children(actual)
	if len(actualChildren) != len(expected.children) {
		t.Fatalf(
			"expected %d children for node %q, got %d",
			len(expected.children), expected.text, len(actualChildren),
		)
	}

	for i, child := range expected.children {
		assertNodeMatches(t, accessors, child, actualChildren[i])
	}
}

func expectParseError[T any](t *testing.T, parse func(string) (T, error), isParseError func(error) bool, input string) {
	t.Helper()

	if _, err := parse(input); err == nil {
		t.Fatalf("parse(%q) expected to fail", input)
	} else if !isParseError(err) {
		t.Fatalf("parse(%q) expected parse error, got %T", input, err)
	}
}
