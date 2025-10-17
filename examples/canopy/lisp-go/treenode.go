// This file was generated from examples/canopy/lisp.peg
// See https://canopy.jcoglan.com/ for documentation

package lispgoparser

// TreeNode represents a node in the parse tree.
// Generated nodes satisfy this interface to allow custom action results.
type TreeNode interface {
	Text() string
	Offset() int
	Children() []TreeNode
}

// BaseNode is embedded by generated nodes to implement TreeNode.
type BaseNode struct {
	text     string
	offset   int
	children []TreeNode
}

// Text returns the source substring matched by the node.
func (n *BaseNode) Text() string {
	return n.text
}

// Offset returns the rune offset where the node starts.
func (n *BaseNode) Offset() int {
	return n.offset
}

// Children returns the node's child list.
func (n *BaseNode) Children() []TreeNode {
	return n.children
}
