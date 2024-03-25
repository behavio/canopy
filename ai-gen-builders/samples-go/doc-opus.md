---
layout: default
title: Go
---

## Go

To get an overview of how to use Canopy with Go, consider this example of a simplified grammar for URLs:

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

This creates a package called `url` that contains all the parser logic. The package name is based on the path to the `.peg` file when you run `canopy`, for example if you run:

    $ canopy github.com/jcoglan/canopy/url.peg --lang go

then you will get a package named `github.com/jcoglan/canopy/url`. The `--output` option can be used to override this:

    $ canopy github.com/jcoglan/canopy/url.peg --lang go --output some/pkg/url

This will write the generated files into the directory `some/pkg/url` with the package name `url`.

Let's try out our parser:

```go
package main

import (
    "fmt"
    "github.com/jcoglan/canopy/url"
)

func main() {
    tree, err := url.Parse("http://example.com/search?q=hello#page=1")
    if err != nil {
        panic(err)
    }

    for _, node := range tree.Elements {
        fmt.Printf("%d, %s\n", node.Offset, node.Text)
    }

    /* prints:

       0, http
       4, ://
       7, example.com
       18, /search
       25, ?q=hello
       33, #page=1
    */
}
```

This little example shows a few important things:

You invoke the parser by calling the module's `Parse()` function with a string.

The `Parse()` method returns a tree of *nodes*.

Each node has three properties:

* `Text string`, the snippet of the input text that node represents
* `Offset int`, the number of characters into the input text the node appears
* `Elements []Node`, a slice of nodes matching the sub-expressions

## Walking the parse tree

You can use `Elements` to walk into the structure of the tree, or, you can use the labels that Canopy generates, which can make your code clearer:

```go
package main

import (
    "fmt"
    "github.com/jcoglan/canopy/url"
)

func main() {
    tree, err := url.Parse("http://example.com/search?q=hello#page=1")
    if err != nil {
        panic(err)
    }

    fmt.Println(tree.Elements[4].Elements[1].Text)
    // -> "q=hello"

    fmt.Println(tree.Search.Query.Text)
    // -> "q=hello"
}
```

## Parsing errors

If you give the parser an input text that does not match the grammar, a `url.ParseError` is returned. The error message will list any of the strings or character classes the parser was expecting to find at the furthest position it got to, along with the rule those expectations come from, and it will highlight the line of the input where the syntax error occurs.

```go
package main

import (
    "fmt"
    "github.com/jcoglan/canopy/url"
)

func main() {
    _, err := url.Parse("https://example.com./")
    if err != nil {
        fmt.Println(err)
    }
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

In Go, compiling the above grammar produces a package called `maps` that contains an interface called `Actions`. You supply the action functions to the parser by implementing the `Actions` interface, which has one method for each action named in the grammar, each of which must return a `canopy.Node`.

The following example parses the input `{'ints':[1,2,3]}`. It implements the `Actions` interface to construct a `map[string][]int`.

```go
package main

import (
    "fmt"
    "strconv"

    "github.com/jcoglan/canopy"
    "github.com/jcoglan/canopy/maps"
)

type MapsActions struct{}

func (a MapsActions) MakeMap(input string, start, end int, elements []canopy.Node) canopy.Node {
    key := elements[1].(*canopy.TextNode).Text[1 : len(elements[1].(*canopy.TextNode).Text)-1]
    value := elements[3].(*maps.Array).List
    return maps.Map{key: value}
}

func (a MapsActions) MakeString(input string, start, end int, elements []canopy.Node) canopy.Node {
    return &canopy.TextNode{elements[1].Text()}
}

func (a MapsActions) MakeList(input string, start, end int, elements []canopy.Node) canopy.Node {
    list := []int{elements[1].(*maps.Number).Value}
    for _, el := range elements[2] {
        list = append(list, el.Elements[1].(*maps.Number).Value)
    }
    return &maps.Array{list}
}

func (a MapsActions) MakeNumber(input string, start, end int, elements []canopy.Node) canopy.Node {
    n, _ := strconv.Atoi(input[start:end])
    return &maps.Number{n}
}

func main() {
    result, _ := maps.Parse("{'ints':[1,2,3]}", MapsActions{})

    fmt.Println(result.(maps.Map))
    // -> map[ints:[1 2 3]]
}
```

## Extended node types

Go does not support the `<Type>` grammar annotation. However, you can achieve similar functionality by defining your own node types that embed `canopy.Node` and returning them from your action methods.

For example:

```go
type URL struct {
    canopy.Node
    Scheme string
    Host   string
}

func (a Actions) MakeURL(input string, start, end int, elements []canopy.Node) canopy.Node {
    return &URL{
        Node:   canopy.Node{Offset: start, Text: input[start:end], Elements: elements},
        Scheme: elements[0].Text(),
        Host:   elements[2].Text(),
    }
}
```

This allows you to add additional fields and methods to your node types while still being compatible with Canopy's parse tree.
