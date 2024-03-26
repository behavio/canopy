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
}
