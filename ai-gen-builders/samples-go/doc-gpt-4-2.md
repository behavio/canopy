golang.md
```markdown
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

We can compile this grammar into a Go package using `canopy`:

    $ canopy url.peg --lang go

This creates a file called `url.go` that contains all the parser logic. The `--output` option can be used to override the default location:

    $ canopy url.peg --lang go --output some/dir/url

This generates the parser code in the file `some/dir/url.go`.

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

    // prints:

    // 0, http
    // 4, ://
    // 7, example.com
    // 18, /search
    // 25, ?q=hello
    // 33, #page=1
}
```

This little example shows a few important things:

You invoke the parser by calling the module's `Parse()` function with a string.

The `Parse()` method returns a tree of *nodes*.

Each node has three properties:

* `Text`, the snippet of the input text that node represents
* `Offset`, the number of characters into the input text the node appears
* `Elements`, a slice of nodes matching the sub-expressions

## Walking the parse tree

You can use `Elements` to walk into the structure of the tree:

```go
fmt.Println(tree.Elements[4].Elements[1].Text)
// -> 'q=hello'
```

Or, you can use the labels that Canopy generates, which can make your code
clearer:

```go
fmt.Println(tree.Elements[url.Labels["search"]].Elements[url.Labels["query"]].text)
// -> 'q=hello'
```

## Parsing errors

If you give the parser an input text that does not match the grammar, a
`url.ParseError` is returned. The error message will list any of the strings or
character classes the parser was expecting to find at the furthest position it
got to, along with the rule those expectations come from, and it will highlight
the line of the input where the syntax error occurs.

```go
_, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println(err)
}

// url.ParseError: Line 1: expected one of:
//
//     - [a-z0-9-] from URL::segment
//
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

In Go, you give the action functions to the parser by using an instance of a
struct that implements these actions:

```go
type Actions struct{}

func (a *Actions) MakeMap(elements []*url.Node) url.Node {
    return url.Node{Tag: "map", Text: fmt.Sprintf("Key: %v, Value: %v", elements[1].Text, elements[3].Text)}
}

func (a *Actions) MakeString(elements []*url.Node) url.Node {
    return url.Node{Tag: "string", Text: elements[1].Text}
}

func (a *Actions) MakeList(elements []*url.Node) url.Node {
    values := make([]string, len(elements))
    for i, element := range elements {
        values[i] = element.Text
    }
    return url.Node{Tag: "list", Text: strings.Join(values, ", ")}
}

func (a *Actions) MakeNumber(elements []*url.Node) url.Node {
    return url.Node{Tag: "number", Text: elements[0].Text}
}

func main() {
    actions := &Actions{}
    tree, err := url.Parse("{'ints':[1,2,3]}", url.Actions(actions))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(tree.Text)
}
```

## Extended node types

As of the current version of Canopy, support for the `<Type>` grammar annotation is not available in the Go version.
```
