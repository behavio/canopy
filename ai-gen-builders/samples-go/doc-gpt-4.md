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
package name is based on the path to the `.peg` file when you run `canopy`, for example if you run:

    $ canopy com/jcoglan/canopy/url.peg --lang go

then you will get a package named `com.jcoglan.canopy.url`. The `--output`
option can be used to override this:

    $ canopy com/jcoglan/canopy/url.peg --lang go --output some/dir/url

This will write the generated files into the directory `some/dir/url` with the
package name `some.dir.url`.

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
		fmt.Println(node.Offset, ", ", node.Text)
	}

	/*  prints:

	    0, http
	    4, ://
	    7, example.com
	    18, /search
	    25, ?q=hello
	    33, #page=1 */
}
```

This little example shows a few important things:

You invoke the parser by calling the package's `Parse()` function with a string.

The `Parse()` method returns a tree of *nodes*.

Each node has three properties:

* `Text`, the snippet of the input text that node represents
* `Offset`, the number of characters into the input text the node appears
* `Elements`, an array of nodes matching the sub-expressions

## Walking the parse tree

You can use `Elements` to walk into the structure of the tree:

```go
fmt.Println(tree.Elements[4].Elements[1].Text)
# -> 'q=hello'
```

Or, you can use the labels that Canopy generates, which can make your code clearer:

```go
fmt.Println(tree.Search.Query.Text)
# -> 'q=hello'
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

In Go, we handle actions using methods on the `Actions` struct. Each method corresponds to one action named in the grammar. A method must take one argument of type `[]Node` and return a `Node`.

```go
package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type Actions struct{}

func (a *Actions) MakeMap(nodes []maps.Node) maps.Node {
	n := make(map[string]int)
	n[nodes[1].Text] = nodes[3].Value.(int)
	return n
}

func (a *Actions) MakeString(nodes []maps.Node) maps.Node {
	return strings.Trim(nodes[0].Text, "'")
}

func (a *Actions) MakeList(nodes []maps.Node) maps.Node {
	list := make([]int, 0)
	list = append(list, nodes[1].Value.(int))

	for _, n := range nodes[2].Elements {
		list = append(list, n.Value.(int))
	}

	return list
}

func (a *Actions) MakeNumber(nodes []maps.Node) maps.Node {
	i, _ := strconv.Atoi(nodes[0].Text)
	return i
}

func main() {
	actions := &Actions{}
	result, err := maps.Parse("{'ints':[1,2,3]}", actions)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// -> map['ints']:[1, 2, 3]
}
```
Note that in Go, unlike in Python or Java, actions must return a `Node` that is suitable for inclusion in a parsing tree, in addition to performing any side-effects or computations.
