```markdown
---
layout: default
title: Go
---

## Go

To get an insight into utilizing Canopy with Go, let's explore an example with a simplified grammar for URLs:

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

You can compile this grammar into Go package using `canopy`:

    $ canopy url.peg --lang go

This will create a Go package under a directory `url`, comprising all necessary parser logic. Should you wish to alter the output location, you can do so with the `--output` option:

    $ canopy url.peg --lang go --output some/dir/url

This command writes the generated Go files into the specified directory `some/dir/url`.

### Testing Our Parser

Let's put our generated parser to the test:

```go
package main

import (
    "fmt"
    "url" // import the generated package
)

func main() {
    tree, err := url.Parse("http://example.com/search?q=hello#page=1")
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }

    for _, node := range tree.Elements {
        fmt.Println(node.Offset, ",", node.Text)
    }

    // Output:
    // 0 , http
    // 4 , ://
    // 7 , example.com
    // 18 , /search
    // 25 , ?q=hello
    // 33 , #page=1
}
```

This snippet elucidates several critical points:

- Parsing is initiated by invoking the `Parse()` function of the generated package, passing the target string.
- The `Parse()` function yields a tree of *nodes*.
- Each node encompasses several attributes:
  - `Text`, representing the corresponding snippet of the input text
  - `Offset`, indicating the position of the node within the input text
  - `Elements`, a slice of nodes matching the subexpressions

## Navigating the Parse Tree

To traverse the structure of the parse tree, you can either manipulate the `Elements` directly:

```go
fmt.Println(tree.Elements[4].Elements[1].Text)
// Output: 'q=hello'
```

Alternatively, use the generated labels for clearer and more manageable code:

```go
fmt.Println(tree.Search.Query.Text)
// Output: 'q=hello'
```

## Handling Parsing Errors

Feeding the parser with text that deviates from the defined grammar results in an error:

```go
_, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println("Parse error:", err)
}
// Parse error: Line 1: expected one of:
//
// - [a-z0-9-] from URL::segment
//
//   1 | https://example.com./
//                           ^
```

## Implementing Actions

Consider a grammar with action annotations:

###### maps.peg

    grammar Maps
      map     <-  "{" string ":" value "}" %make_map
      string  <-  "'" [^']* "'" %make_string
      value   <-  list / number
      list    <-  "[" value ("," value)* "]" %make_list
      number  <-  [0-9]+ %make_number

In Go, you integrate actions with the parser by defining functions that correspond to action annotations in the grammar. When compiling the grammar, the generator produces a `maps.go` package that contains an interface `Actions`. Implement this interface with functions that mirror your grammar annotations, each returning an instance of `Node` (or an implementation thereof) for flexibility.

Below is an example that processes the input string `{'ints':[1,2,3]}`:

```go
package main

import (
    "fmt"
    "maps"
)

type MapsActions struct{}

func (m *MapsActions) MakeMap(input string, start, end int, elements []*maps.Node) *maps.Node {
    // Implementation
}

func (m *MapsActions) MakeString(input string, start, end int, elements []*maps.Node) *maps.Node {
    // Implementation
}

func (m *MapsActions) MakeList(input string, start, end int, elements []*maps.Node) *maps.Node {
    // Implementation
}

func (m *MapsActions) MakeNumber(input string, start, end int, elements []*maps.Node) *maps.Node {
    // Implementation
}

func main() {
    actions := &MapsActions{}
    tree, err := maps.Parse("{'ints':[1,2,3]}", actions)
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }

    // Use tree...
}
```

## Extended Node Types

The Go version of Canopy does not support the `<Type>` grammar annotation directly. Instead, you define custom data structures or interfaces that your actions can return, tailored to the specific needs of your domain, and handle them accordingly when traversing the parse tree.

```

**Note:** This markdown document provides a foundational approach for working with Canopy in Go, emphasizing key operations like parsing input based on a grammar, traversing the parse tree, handling errors, and implementing custom actions. The examples show the general structure but omit detailed implementations for custom actions and error handling to keep them concise. You should adapt and expand these snippets based on your specific grammar and processing requirements.