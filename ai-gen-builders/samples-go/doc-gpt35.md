---
layout: default
title: Go
---

## Go

To get an overview of how to use Canopy with Go, consider this example of a
simplified grammar for URLs:

###### url.peg

    grammar URL
      URL       <-  Scheme "://" Host Pathname Search Hash?
      Scheme    <-  "http" "s"?
      Host      <-  Hostname Port?
      Hostname  <-  Segment ("." Segment)*
      Segment   <-  [a-z0-9-]+
      Port      <-  ":" [0-9]+
      Pathname  <-  "/" [^ ?]*
      Search    <-  ("?" Query:[^ #]*)?
      Hash      <-  "#" [^ ]*

We can compile this grammar into a Go package using `canopy`:

    $ canopy url.peg --lang go

This creates a Go module that contains all the parser logic. The file name will be based on the grammar name (e.g., `url.peg` will generate `url.go`).

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
		panic(err)
	}

	for _, node := range tree.Elements {
		fmt.Println(node.Offset, node.Text)
	}

	/*  prints:

	0, http
	4, ://
	7, example.com
	18, /search
	25, ?q=hello
	33, #page=1       */
}
```

This little example shows a few important things:

You invoke the parser by calling the module's `Parse()` function with a string.

The `Parse()` method returns a tree of *nodes*.

Each node has three properties:

* `Text`, the snippet of the input text that node represents
* `Offset`, the number of characters into the input text the node appears
* `Elements`, an array of nodes matching the sub-expressions

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

If you give the parser an input text that does not match the grammar, a
`url.ParseError` is returned. The error message will indicate the expected strings or character classes the parser was looking for, along with the rule they are associated with, and it will highlight the line of the input where the syntax error occurs.

```go
_, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println(err)
}
// url.ParseError: Line 1: expected one of:
//
//     - [a-z0-9-] from URL::Segment
//
//      1 | https://example.com./
//                              ^
```

## Implementing actions

Say you have a grammar that uses action annotations, for example:

###### maps.peg

    grammar Maps
      Map     <-  "{" String ":" Value "}" %MakeMap
      String  <-  "'" [^']* "'" %MakeString
      Value   <-  List / Number
      List    <-  "[" Value ("," Value)* "]" %MakeList
      Number  <-  [0-9]+ %MakeNumber

In Go, you can implement actions by providing functions for each action named in the grammar:

```go
package main

import (
	"fmt"
	"maps"
)

type Actions struct{}

func (a *Actions) MakeMap(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	return map[maps.TreeNode]maps.TreeNode{elements[1]: elements[3]}
}

func (a *Actions) MakeString(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	return elements[1]
}

func (a *Actions) MakeList(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	list := [maps.TreeNode{elements[1]}]
	for _, el := range elements[2] {
		list = append(list, el.Value)
	}
}

func (a *Actions) MakeNumber(input string, start, end int, elements []maps.TreeNode) maps.TreeNode {
	number, _ := strconv.Atoi(input[start:end])
	return number
}

func main() {
	result, err := maps.Parse("{'ints':[1,2,3]}", &Actions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// -> {'ints': [1, 2, 3]}
}
```

## Extended node types

Using the `<Type>` grammar annotation is not supported in the Go version.
