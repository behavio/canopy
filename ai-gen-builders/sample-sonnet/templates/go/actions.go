
package actions

{{#each actions}}
type {{this}}Func func(input string, start int, end int, elements []interface{}) interface{}
{{/each}}

type Actions struct {
{{#each actions}}
	{{this}} {{this}}Func
{{/each}}
}
