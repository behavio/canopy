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

This creates a Go package that contains the parser logic. The resulting package will be in the path specified in `GOHOME` environment variable.

Let's try out our parser:

```go
package main

import (
	"fmt"
	"url"
)

func main() {
	tree, err := url.Parse("http://example.com/search?q=hello#page=1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, node := range tree.Elements {
		fmt.Println(node.Offset, node.Text)
	}

	/* prints:

	  0 http
	  4 ://
	  7 example.com
	  18 /search
	  25 ?q=hello
	  33 #page=1 */
}
```

This little example shows a few important things:

You invoke the parser by calling the module's `Parse()` function with a string.

The `Parse()` method returns a tree of *nodes*.

Each node has three properties:

- `Text`, the snippet of the input text that node represents
- `Offset`, the number of characters into the input text the node appears
- `Elements`, a slice of nodes matching the sub-expressions

## Walking the parse tree

You can use `Elements` to walk into the structure of the tree:

```go
fmt.Println(tree.Elements[4].Elements[1].Text)
// -> 'q=hello'
```

Or, you can use the labels that Canopy generates, which can make your code clearer:

```go
fmt.Println(tree.Search.Query.Text)
// -> 'q=hello'
```

## Parsing errors

If you give the parser an input text that does not match the grammar, a `ParseError` is thrown. The error message will list any of the strings or character classes the parser was expecting to find at the furthest position it got to, along with the rule those expectations come from, and it will highlight the line of the input where the syntax error occurs.

```go
_, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println(err)
}
// ParseError: Line 1: expected one of:

//     - [a-z0-9-] from URL::segment

//      1 | https://example.com./
//                              ^
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

In Go, you can implement the actions directly in your Go code:

```go
package main

import (
	"fmt"
	"maps"
)

type Actions struct{}

func (a *Actions) MakeMap(input string, start, end int, elements []*maps.TreeNode) *maps.TreeNode {
	result := make(map[string]interface{})
	result[elements[1].Text] = elements[3]

	return maps.NewTreeNode(result, start, end)
}

func (a *Actions) MakeString(input string, start, end int, elements []*maps.TreeNode) *maps.TreeNode {
	return maps.NewTreeNode(elements[1].Text, start, end)
}

func (a *Actions) MakeList(input string, start, end int, elements []*maps.TreeNode) *maps.TreeNode {
	list := []interface{}{elements[1].Value}
	for _, el := range elements[2].Elements {
		list = append(list, el.Value)
	}

	return maps.NewTreeNode(list, start, end)
}

func (a *Actions) MakeNumber(input string, start, end int, elements []*maps.TreeNode) *maps.TreeNode {
	num, _ := strconv.Atoi(input[start:end])
	return maps.NewTreeNode(num, start, end)
}

func main() {
	result, err := maps.Parse("{'ints':[1,2,3]}", &Actions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.Value)
	// -> map[ints:[1 2 3]]
}
```

## Extended node types

Using the `<Type>` grammar annotation is not supported in the Go version.
```