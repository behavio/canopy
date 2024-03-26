
---
layout: default
title: Go
---

## Go

To get an overview of how to use Canopy with Go, consider this example of a simplified grammar for URLs:

###### url.peg

```
grammar URL
  url      <- scheme "://" host pathname search hash?
  scheme   <- "http" "s"?
  host     <- hostname port
To generate a Go parser with Canopy, run:

```bash
canopy url.peg --lang go
```

This will create a `url` directory containing the generated Go source files:

- `actions/actions.go` - Defines the interface for semantic actions
- `node.go` - Implementation of the syntax tree node struct
- `parser.go` - The parser itself, with a `Parse` function
- `grammar.go` - The implementation of the parsing rules
- `position.go` - Helper for tracking position and errors

To use the parser:

```go
import (
    "github.com/yourusername/url/actions"
    "github.com/yourusername/url/parser"
)

func main() {
    input := "http://example.com/search?q=hello#page=1"
    acts := &actions.Actions{
        // Provide action implementations here
    }
    
    tree, err := parser.Parse(input, acts)
    if err != nil {
        // Handle error
    }
    
    // Use the syntax tree
    println(tree.Text) // Prints the full input
}
```

## Implementing Actions

Let's add some semantic actions to our URL grammar:

```
grammar URL
  url      <- scheme "://" host pathname search hash? %make_url
  search   <- "?" query:[^ #]* %make_query
  
  scheme   <- "http" "s"?
  host     <- hostname port?  
  hostname <- segment ("." segment)*
  segment  <- [a-z0-9-]+
  port     <- ":" [0-9]+
  pathname <- "/" [^ ?]*
  hash     <- "#" [^ ]*
  
  query    <- [^ #]+ %make_string
```

We've added:

- A `%make_url` action to construct a URL value 
- A `%make_query` action to extract the query string
- A `%make_string` action for instantiating strings

To implement these actions, we'll create types in Go to represent URL and QueryString, and supply the action functions to the parser:

```go
package main

import (
    "github.com/yourusername/url/actions" 
    "github.com/yourusername/url/parser"
)

type URL struct {
    Scheme   string
    Host     string 
    Path     string
    Query    *QueryString
    Fragment string
}

type QueryString struct {
    value string
}

func makeURL(input string, start int, end int, elements []interface{}) interface{} {
    scheme := elements[0].(parser.Node).Text
    host := elements[1].(parser.Node).Text
    path := elements[2].(parser.Node).Text
    var query *QueryString
    if len(elements) > 4 {
        query = elements[4].(QueryString) 
    }
    fragment := elements[len(elements)-1].(parser.Node).Text
    
    return &URL{
        Scheme:   scheme,
        Host:     host,
        Path:     path, 
        Query:    query,
        Fragment: fragment,
    }
}

func makeQuery(input string, start int, end int, elements []interface{}) interface{} {
    value := input[start:end]
    return &QueryString{value: value}    
}

func makeString(input string, start int, end int, elements []interface{}) interface{} {
    return input[start:end]
}

func main() {
    input := "http://example.com/search?q=hello#page=1"
    acts := &actions.Actions{
        make_url:     makeURL,
        make_query:   makeQuery, 
        make_string:  makeString,
    }
    
    url, err := parser.Parse(input, acts)
    if err != nil {
        // Handle error
    }
    
    // Access properties of the URL value
    println(url.Query.value) // Prints "q=hello"
}
```

This shows how to supply action functions that instantiate your own data structures from the parser output.

## Extended Node Types 

To attach methods and fields to syntax nodes, use the `<Type>` syntax in your grammar:

```
expression <- ... <MyType>
```

This will make nodes matching `expression` have type `*MyType` instead of the default `*Node`. 

Define `MyType` by creating a named type and supplying it to the parser via the `types` argument:

```go
type MyType parser.Node 

func (n *MyType) MyMethod() string {
    // Implementation...
}

func main() {
    input := "..."
    acts := &actions.Actions{
        // Actions... 
    }
    types := struct{ MyType }{}
    
    node, err := parser.Parse(input, acts, types)
    // node will have type *MyType
    myNode := node.(*MyType) 
    println(myNode.MyMethod())
}
```