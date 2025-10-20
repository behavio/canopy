package {{name}}

// Actions defines the semantic callbacks for nodes with actions.
type Actions interface {
{{#each actions}}
	{{this}}(input string, start, end int, elements []TreeNode) (TreeNode, error)
{{/each}}
}
