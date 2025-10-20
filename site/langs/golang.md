---
layout: default
title: Go
---

## Go

To get an overview of how to use Canopy with Go, consider this example of a
simplified grammar for URLs:

###### url.peg

    grammar URL
      url       <-  scheme "://" host pathname search hash?
      scheme    <-  "http" "s"?
      host      <-  hostname port?
      hostname  <-  segment ("." segment)*
      segment   <-  [a-z0-9-]+
      port      <-  ":" [0-9]+
      pathname  <-  "/" [^ ?]*
      search    <-  ("?" query:[^ #]*)?
      hash      <-  "#" [^ ]*

We can compile this grammar into a Go module using `canopy`:

    $ canopy url.peg --lang go

This creates a directory called `url-go` that contains all the parser logic as a
complete Go module. The `--output` option can be used to override the default
location:

    $ canopy url.peg --lang go --output some/dir/url-go

This will write the generated parser into the directory `some/dir/url-go`.

The generated module contains four files:

- `go.mod` - Go module definition
- `parser.go` - Main parser logic
- `treenode.go` - TreeNode interface and BaseNode struct
- `actions.go` - Actions interface (empty if no actions in grammar)

Let's try our parser out:

```go
package main

import (
    "fmt"
    "log"

    "urlgoparser"
)

func main() {
    tree, err := urlgoparser.Parse("http://example.com/search?q=hello#page=1", nil, nil)
    if err != nil {
        log.Fatal(err)
    }

    for _, node := range tree.Children() {
        fmt.Printf("%d %s\n", node.Offset(), node.Text())
    }

    /*  prints:

        0 http
        4 ://
        7 example.com
        18 /search
        25 ?q=hello
        33 #page=1      */
}
```

This little example shows a few important things:

You invoke the parser by calling the module's `Parse()` function with a string.
The second and third parameters are for actions and type extensions, which we'll
cover later.

The `Parse()` function returns a tree of _nodes_ and an error. If parsing
succeeds, the error is `nil`.

Each node implements the `TreeNode` interface with three methods:

- `Text()` returns the snippet of the input text that node represents
- `Offset()` returns the number of characters into the input text the node appears
- `Children()` returns a slice of nodes matching the sub-expressions

## Walking the parse tree

You can use `Children()` to walk into the structure of the tree:

```go
fmt.Println(tree.Children()[4].Children()[1].Text())
// -> "q=hello"
```

For sequences with labeled elements, Canopy generates typed structs with
exported fields for each label. You can use type assertions to access these
fields, which can make your code clearer:

```go
type Node struct {
    BaseNode
    Search TreeNode
}

if node, ok := tree.(*Node); ok {
    if search, ok := node.Search.(*SearchNode); ok {
        fmt.Println(search.Query.Text())
        // -> "q=hello"
    }
}
```

## Parsing errors

If you give the parser an input text that does not match the grammar, a
`ParseError` is returned. The error message will list any of the strings or
character classes the parser was expecting to find at the furthest position it
got to, along with the rule those expectations come from, and it will highlight
the line of the input where the syntax error occurs.

```go
import "errors"

tree, err := urlgoparser.Parse("https://example.com./", nil, nil)
if err != nil {
    var parseErr *urlgoparser.ParseError
    if errors.As(err, &parseErr) {
        fmt.Printf("Parse error at line %d, column %d\n",
                   parseErr.Line, parseErr.Column)
    }
}

// Parse error at line 1, column 20
// parse error at line 1, column 20: expected [a-z0-9-] from URL::segment
//
//      1 | https://example.com./
//                              ^
```

The `ParseError` struct contains the following fields:

- `Input` - the original input string
- `Offset` - the character offset where parsing failed
- `Line` - the line number where parsing failed (1-indexed)
- `Column` - the column number where parsing failed (1-indexed)
- `Expected` - a slice of expectations showing what was expected
- `Message` - a formatted error message

## Implementing actions

Say you have a grammar that uses action annotations, for example:

###### maps.peg

    grammar Maps
      map     <-  "{" string ":" value "}" %make_map
      string  <-  "'" [^']* "'" %make_string
      value   <-  list / number
      list    <-  "[" value ("," value)* "]" %make_list
      number  <-  [0-9]+ %make_number

In Go, compiling the above grammar produces an `Actions` interface in
`actions.go` with one method for each action annotation. Each method receives
the input string, start and end positions, and the matched elements, and must
return a `TreeNode` and an error.

You supply the action functions to the parser by implementing this interface:

```go
package main

import (
    "fmt"
    "strconv"
    "strings"

    "mapsgoparser"
)

type MapNode struct {
    mapsgoparser.BaseNode
    Value map[string][]int
}

type StringNode struct {
    mapsgoparser.BaseNode
    Value string
}

type ListNode struct {
    mapsgoparser.BaseNode
    Value []int
}

type NumberNode struct {
    mapsgoparser.BaseNode
    Value int
}

type MyActions struct{}

func (a *MyActions) MakeMap(input string, start, end int, elements []mapsgoparser.TreeNode) (mapsgoparser.TreeNode, error) {
    key := elements[1].(*StringNode).Value
    value := elements[3].(*ListNode).Value

    return &MapNode{
        BaseNode: mapsgoparser.BaseNode{},
        Value:    map[string][]int{key: value},
    }, nil
}

func (a *MyActions) MakeString(input string, start, end int, elements []mapsgoparser.TreeNode) (mapsgoparser.TreeNode, error) {
    return &StringNode{
        BaseNode: mapsgoparser.BaseNode{},
        Value:    elements[1].Text(),
    }, nil
}

func (a *MyActions) MakeList(input string, start, end int, elements []mapsgoparser.TreeNode) (mapsgoparser.TreeNode, error) {
    first := elements[1].(*NumberNode).Value
    list := []int{first}

    for _, node := range elements[2].Children() {
        if num, ok := node.Children()[1].(*NumberNode); ok {
            list = append(list, num.Value)
        }
    }

    return &ListNode{
        BaseNode: mapsgoparser.BaseNode{},
        Value:    list,
    }, nil
}

func (a *MyActions) MakeNumber(input string, start, end int, elements []mapsgoparser.TreeNode) (mapsgoparser.TreeNode, error) {
    text := input[start:end]
    value, err := strconv.Atoi(text)
    if err != nil {
        return nil, fmt.Errorf("invalid number: %w", err)
    }

    return &NumberNode{
        BaseNode: mapsgoparser.BaseNode{},
        Value:    value,
    }, nil
}

func main() {
    actions := &MyActions{}
    result, err := mapsgoparser.Parse("{'ints':[1,2,3]}", actions, nil)
    if err != nil {
        log.Fatal(err)
    }

    if mapNode, ok := result.(*MapNode); ok {
        fmt.Println(mapNode.Value)
        // -> map[ints:[1 2 3]]
    }
}
```

A few things to note about actions in Go:

- Each action method returns `(TreeNode, error)`. If an error is returned,
  parsing stops immediately.
- Custom node types should embed `BaseNode` and implement additional fields for
  semantic values.
- Action methods receive all matched elements, including literals and whitespace.
  You need to extract the relevant elements yourself.
- The `elements` slice should be copied if you need to store it, as the parser
  may reuse the underlying array.

## Extended node types

Say you have a grammar that contains type annotations:

###### words.peg

    grammar Words
      root  <-  first:"foo" second:"bar" <Extension>

To use this parser, you must pass in a map containing implementations of the
named types via the third parameter to `Parse()`. Each type is a function that
takes a `TreeNode` and returns a new `TreeNode` with extended functionality:

```go
package main

import (
    "fmt"
    "strings"

    "wordsgoparser"
)

type ExtensionNode struct {
    wordsgoparser.TreeNode
}

func (n *ExtensionNode) Convert() string {
    // Access the labeled fields through type assertion
    type NodeWithLabels interface {
        wordsgoparser.TreeNode
    }

    node := n.TreeNode.(NodeWithLabels)
    first := node.Children()[0].Text()
    second := node.Children()[1].Text()

    return first + strings.ToUpper(second)
}

func main() {
    extensions := map[string]wordsgoparser.NodeExtender{
        "Extension": func(node wordsgoparser.TreeNode) wordsgoparser.TreeNode {
            return &ExtensionNode{TreeNode: node}
        },
    }

    tree, err := wordsgoparser.Parse("foobar", nil, extensions)
    if err != nil {
        log.Fatal(err)
    }

    if ext, ok := tree.(*ExtensionNode); ok {
        fmt.Println(ext.Convert())
        // -> "fooBAR"
    }
}
```

Type extensions wrap existing nodes and add new methods or fields. They're
useful when you want to add behavior to nodes without using action annotations.

## Module structure and imports

The generated Go module is self-contained and depends only on Go's standard
library. To use it in your project:

1. Generate the parser:

   ```
   $ canopy url.peg --lang go
   ```

2. The generated module name is based on the grammar filename (e.g.,
   `url.peg` → `urlgoparser`). Update your `go.mod` to include it:

   ```
   require urlgoparser v0.0.0
   replace urlgoparser => ./url-go
   ```

3. Import and use it:

   ```go
   import "urlgoparser"

   tree, err := urlgoparser.Parse(input, nil, nil)
   ```

Alternatively, for simple use cases, you can call the `New()` constructor
directly and configure the parser with options:

```go
parser := urlgoparser.New(input, actions)
tree, err := parser.WithTypes(extensions).Parse()
```

## Performance considerations

The generated parsers use packrat parsing with full memoization, which
guarantees O(n) parsing time where n is the input length. However, this comes
with memory overhead:

- Each `(rule, position)` pair that is attempted is cached
- The input is stored as a `[]rune` slice for proper Unicode handling
- Parse tree nodes are allocated for each successful match

For best performance:

- Reuse parsers when parsing multiple inputs of the same grammar
- Consider streaming or chunking very large inputs
- Profile your grammar to identify rules that create excessive backtracking
- Use actions to transform nodes eagerly rather than walking large trees

The packrat approach is particularly beneficial for grammars with:

- Ambiguous rules that can match the same input in multiple ways
- Left-recursive rules
- Many lookahead predicates
- Complex backtracking scenarios

For simple, deterministic grammars, the memoization overhead may exceed the
benefits, but Canopy prioritizes correctness and ease of use over raw speed.
