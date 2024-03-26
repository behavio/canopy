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
}
