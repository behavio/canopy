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

This creates a package called `url` that contains all the parser logic. The package name is based on the output
directory name when you run `canopy`, for example if you run:

    $ canopy url.peg --lang go --output myparser

then you will get a package named `myparser`.

Let's try out our parser:

```go
package main

import (
    "fmt"
    "myparser/url"
)

func main() {
    tree, err := url.Parse("http://example.com/search?q=hello#page=1")
    if err != nil {
        panic(err)
    }

    for _, node := range tree.Elements() {
        fmt.Printf("%d, %s\n", node.Offset(), node.Text())
    }

    // prints:
    //
    // 0, http
    // 4, ://
    // 7, example.com
    // 18, /search
    // 25, ?q=hello
    // 33, #page=1
}
```

This little example shows a few important things:

You invoke the parser by calling the package's `Parse` function with a string.

The `Parse` function returns an AST as a tree of `node` pointers, or an error if the input does not match the grammar.

Each node has three properties:

* `Text() string`, the snippet of input text the node represents
* `Offset() int`, the index into the input text where the node starts
* `Elements() []node`, an array of child nodes

## Walking the parse tree

You can use the `Elements()` method to walk into the structure of the AST:

```go
queryNode := tree.Elements()[4].Elements()[1]
fmt.Println(queryNode.Text())
// -> 'q=hello'
```

Or, you can use the labels that Canopy generates, which can make your code clearer:

```go
queryNode := tree.Elements()[3].(*URL_search).Query()
fmt.Println(queryNode.Text())
// -> 'q=hello'
```

## Parsing errors

If you give the parser an input that does not match the grammar, the `Parse` function will return an error.
The error message will tell you the furthest offset the parser reached successfully, and what it was expecting
to find at that point to continue.

```go
tree, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println(err)
    // url.ParseError: Line 1: expected [a-z0-9-]
}
```

## Implementing actions

Say you have a grammar that uses action annotations, for example:

###### maps.peg

    grammar Maps
      map    <- "{" string ":" value "}" %make_map
      string <- "'" [^']* "'" %make_string
      value  <- list / number
      list   <- "[" value ("," value)* "]" %make_list
      number <- [0-9]+ %make_number

In Go, compiling the above grammar produces a package that contains the following:

* Types `Maps`, `node`, and the various AST node types like `Maps_string`, `Maps_value`, etc
* An interface type `Actions` with a method for each action in the grammar
* The `Maps.Parse` function to parse an input string
* `ParseError` error type

You provide the action functions by passing an `Actions` implementation to `Parse`:

```go
package main

import (
    "fmt"
    "myparser/maps"
)

type MyActions struct {}

func (a MyActions) make_map(ast *maps.Maps_map) {
    pairs := make(map[string]int)
    pairs[ast.String().Text()] = ast.Value().Int()
    ast.SetPairs(pairs)
}

func (a MyActions) make_string(ast *maps.Maps_string) {
    ast.SetValue(ast.Elements()[1].Text())
}

func (a MyActions) make_list(ast *maps.Maps_list) {
    list := make([]int, len(ast.Elements()))
    for i, elem := range ast.Elements() {
        list[i] = elem.(*maps.Maps_value).Int()
    }
    ast.SetValue(list)
}

func (a MyActions) make_number(ast *maps.Maps_number) {
    num := 0
    for _, d := range ast.Text() {
        num = num * 10 + int(d - '0')
    }
    ast.SetValue(num)
}

func main() {
    input := "{'foo':42}"
    ast, err := maps.Parse(input, MyActions{})
    if err != nil {
        panic(err)
    }

    fmt.Println(ast.(*maps.Maps_map).Pairs())
    // prints: map[foo:42]
}
```

Each AST node type has:
- `Text() string` to get the matched text
- `Offset() int` to get the match position
- `Elements() []node` to get the child nodes
- `SetXXX(value)` methods to set properties populated by actions
- Typed getters for each labeled element in the grammar, like `String() *Maps_string`

## Extending the AST nodes

Using the `<Type>` annotation is not currently supported by the Go version of Canopy.
