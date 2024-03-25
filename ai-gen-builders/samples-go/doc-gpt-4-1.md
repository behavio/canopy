```markdown
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

This creates a package called `url` that contains all the parser logic. The
package name is based on the path to the `.peg` file when you run `canopy`. The `--output`
option can be used to override this:

    $ canopy url.peg --lang go --output some/dir/url

This will write the generated files into the directory `some/dir/url`.

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
        fmt.Println(err.Error())
        return
    }

	for _, node := range tree.Elements {
		fmt.Println(node.Text, ", ",node.Offset)
	}

    /*  prints:

        http ,  0
        :// ,  4
        example.com ,  7
        /search ,  18
        ?q=hello ,  25
        #page=1 ,  33
    */
}
```

This little example shows a few important things:

You invoke the parser by calling the package's `Parse()` function with a string.

The `Parse()` method returns a tree of *nodes*.

Each node has three properties:

* `Offset`, the number of characters into the input text the node appears
* `Text`, the snippet of the input text that node represents
* `Elements`, an array of nodes representing the child elements

## Walking the parse tree

You can use `Elements` to navigate deeper into the tree:

```go
fmt.Println(tree.Elements[4].Elements[1].Text)
// -> 'q=hello'
```

## Parsing errors

If you provide a text to the parser that doesn't match the grammar, an error is returned which identifies the strings or character classes the parser was expecting to find that it couldn't, as well as the rule those expectations come from, and it will highlight the line of the input where the syntax error occurs.

```go
_, err := url.Parse("https://example.com./")
if err != nil {
    fmt.Println(err.Error())
}

// Error: Line 1: expected one of:
//        
//     - [a-z0-9-] from URL::segment
//        
//      1 | https://example.com./
//                            ^
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

In Go, handling actions needs additional code which is defined as method receivers on the `Actions` type:

```go
package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type MapsActions struct{}

func (a MapsActions) Make_map(elements ...*maps.TreeNode) (*maps.TreeNode, error) {
	m := make(map[string]interface{})
	key := strings.Trim(elements[1].Text, "'")
	m[key] = elements[3].Value
	return &maps.TreeNode{Label: "map", Value: m}, nil
}

func (a MapsActions) Make_string(elements ...*maps.TreeNode) (*maps.TreeNode, error) {
	s := strings.Trim(elements[1].Text, "'")
	return &maps.TreeNode{Label: "string", Value: s}, nil
}

func (a MapsActions) Make_list(elements ...*maps.TreeNode) (*maps.TreeNode, error) {
	l := []int{elements[1].Value.(int)}
	for _, e := range elements[2].Elements {
		l = append(l, e.Value.(int))
	}
	return &maps.TreeNode{Label: "list", Value: l}, nil
}

func (a MapsActions) Make_number(elements ...*maps.TreeNode) (*maps.TreeNode, error) {
	i, _ := strconv.Atoi(elements[1].Text)
	return &maps.TreeNode{Label: "number", Value: i}, nil
}

func main() {
	tree, err := maps.Parse("{'ints':[1,2,3]}", MapsActions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tree.Value)
	// -> map[ints:[1 2 3]]
}
```

## Extended node types

Canopy's Go version does not support grammar that contains type annotations like `<Type>`.
```