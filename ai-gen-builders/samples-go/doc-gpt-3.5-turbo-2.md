```---
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

We can compile this grammar into a Go package using `canopy`:

    $ canopy url.peg --lang go

This creates a directory structure in the current directory with the Go package,
following the Go conventions:

    .
    ├── url/
    │   ├── grammar.go
    │   ├── parser.go

Let's try our parser out:

```go
package main

import (
	"fmt"
	"url"
)

func main() {
	tree, err := url.Parse("http://example.com/search?q=hello#page=1")
	if err != nil {
		panic(err)
	}

	for _, node := range tree.Elements {
		fmt.Printf("%d, %s\n", node.Offset, node.Text)
	}
}
```

This little example shows a few important things:

You invoke the parser by calling the package's `Parse` function with a string.

The `Parse` method returns a tree of *nodes*.

Each node has three properties:

- `Text`, the snippet of the input text that the node represents
- `Offset`, the number of characters into the input text the node appears
- `Elements`, a slice of nodes matching the sub-expressions

## Walking the parse tree

You can use `Elements` to walk into the structure of the tree:

```go
fmt.Println(tree.Elements[4].Elements[1].Text)
// -> 'q=hello'
```

Or, you can use the labels that Canopy generates, which can make your code
clearer:

```go
fmt.Println(tree.Search.Query.Text)
// -> 'q=hello'
```

## Parsing errors

If you give the parser an input text that does not match the grammar, a
`ParseError` is returned. The error message will explain the issue.

```go
_, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println(err)
}

// Example output:
// url.ParseError: Line 1: expected one of:
//     - [a-z0-9-] from URL::segment
```

## Implementing actions

Say you have a grammar that uses action annotations, for example:

###### maps.peg

    grammar Maps
      map     <-  "{" string ":" value "}" %make_map
      string  <-  "'" [^']* "'" %make_string
      value   <-  list / number
      list    <-  "[" value ("," value)* "]" %make_list
      number  <-  [0-9]+ %make_number

In Go, you implement functions to handle actions directly within your code:

```go
package main

import (
	"fmt"
	"maps"
)

type Actions struct{}

func (a *Actions) MakeMap(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	return maps.Pair{Key: elements[1], Value: elements[3]}
}

func (a *Actions) MakeString(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	return maps.Text{Value: elements[1].Text}
}

func (a *Actions) MakeList(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	var list []int
	list = append(list, elements[1].Value)
	for _, el := range elements[2] {
		list = append(list, el.Value)
	}
	return maps.Array{List: list}
}

func (a *Actions) MakeNumber(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	return maps.Number(Number: elements[1].Value)
}

func main() {
	result, err := maps.Parse("{'ints':[1,2,3]}", Actions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// -> {'ints': [1, 2, 3]}
}
```

## Extended node types
Using the `<Type>` grammar annotation is not supported in the Go version.
```  