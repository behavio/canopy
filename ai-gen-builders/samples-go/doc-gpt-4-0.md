```markdown
---
layout: default
title: Go
---

## Go

To get an overview of how to use Canopy with Go, let's look at the example of a
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

This creates a Go package named `url` that contains all the parser logic. The
package name is based on the filename of the `.peg` file when you run `canopy`.
The `--output` option can be used to override the default location:

    $ canopy url.peg --lang go --output some/dir/url

This writes the generated parser into the package `some/dir/url`.

Let's try our parser:

```go
package main

import (
	"fmt"
	"url"
)

func main() {
	tree, _ := url.ParseUrl("http://example.com/search?q=hello#page=1")

	for _, node := range tree.Elements {
		fmt.Printf("%d, %s\n", node.Offset, node.Text)

		/*  prints:

		    0, http
		    4, ://
		    7, example.com
		    18, /search
		    25, ?q=hello
		    33, #page=1  */
	}
}
```

This example shows a few important things:

You invoke the parser by calling `ParseUrl()` function with a string.

The `ParseUrl()` function returns a tree of *nodes*.

Each node has three properties:

* `element.Text`, the snippet of the input text that node represents
* `element.Offset`, the number of characters into the input text the node appears
* `element.Elements`, an array of nodes matching the sub-expressions


## Walking the parse tree

You can use `Elements` to walk into the structure of the tree, or you can use
the labels that Canopy generates, which can make your code clearer:

```go
package main

import (
	"fmt"
	"url"
)

func main() {
	tree, _ := url.ParseUrl("http://example.com/search?q=hello#page=1")

	fmt.Println(tree.Elements[4].Elements[1].Text)
	// -> 'q=hello'

	fmt.Println(tree.Elements[url.Search].Get(url.Query).Text)
	// -> 'q=hello'
}
```

## Parsing errors

If you give the parser an input text that does not match the grammar, a
`url.ParseError` error is returned. The error message will list any of the strings or
character classes the parser was expecting to find at the furthest position it
got to, along with the rule those expectations come from, and it will highlight
the line of the input where the syntax error occurs.

```go
package main

import (
	"fmt"
	"url"
)

func main() {
	_, err := url.ParseUrl("https://example.com./")
	if err != nil {
		fmt.Println(err)
	}

	// url.ParseError: Line 1: expected one of:
	//
	//     - [a-z0-9-] from URL::segment
	//
	//      1 | https://example.com./
	//                              ^
}
```

Go version doesn't support action or type annotations features.

```