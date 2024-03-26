package test

import (
    "testing"
    "canopy/test/grammars/choices"
)

func TestChoiceStrings(t *testing.T) {
    // Tests parsing single-character choices
    expectParse(t, choices.Parse("choice-abc: a"), "a", 12)
    expectParse(t, choices.Parse("choice-abc: b"), "b", 12)
    expectParse(t, choices.Parse("choice-abc: c"), "c", 12)

    // Invalid choice
    expectError(t, choices.Parse("choice-abc: d"))

    // Superstring of a choice
    expectError(t, choices.Parse("choice-abc: ab"))
}

func TestChoiceSequence(t *testing.T) {
    expectParse(t, choices.Parse("choice-seq: repeat"),
        seq(
            node("re", 12),
            node("peat", 14),
        ),
    12)

    // Does not backtrack
    expectError(t, choices.Parse("choice-seq: reppeat"))
}

func TestChoiceRepetition(t *testing.T) {
    expectParse(t, choices.Parse("choice-rep: abcabba"),
        rep(
            node("a", 12),
            node("b", 13),
            node("c", 14),
            node("a", 15),
            node("b", 16),
            node("b", 17),
            node("a", 18),
        ),
    12)

    // Invalid character
    expectError(t, choices.Parse("choice-rep: abcadba"))
}

func TestChoiceSequence(t *testing.T) {
    expectParse(t, choices.Parse("choice-bind: ab"),
        seq(
            node("a", 13),
            node("b", 14),
        ),
    13)

    // Binds sequences tighter than choices
    expectError(t, choices.Parse("choice-bind: abef"))
}

func node(text string, offset int) *choices.Node {
    return &choices.Node{Text: text, Offset: offset}
}

func seq(nodes ...*choices.Node) *choices.Node {
    return &choices.Node{Elements: nodes}
}

func rep(nodes ...*choices.Node) *choices.Node {
    return &choices.Node{Elements: nodes}
}

func expectParse(t *testing.T, result *choices.Node, text string, offset int) {
    AssertEqual(t, text, result.Text)
    AssertEqual(t, offset, result.Offset)
}

func expectError(t *testing.T, _, err interface{}) {
    if err == nil {
        t.Errorf("Expected error")
    }
}
