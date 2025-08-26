package {{name}}

type Actions interface {
{{#each actions}}
    {{this}} (input string, start, end int, elements []TreeNode) (TreeNode, error)
{{/each}}
}
