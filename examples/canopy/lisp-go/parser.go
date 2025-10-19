// This file was generated from examples/canopy/lisp.peg
// See https://canopy.jcoglan.com/ for documentation

package lispgoparser

import (
    	"fmt"
	"regexp"
)

type NodeExtender func(TreeNode) TreeNode

type expectation struct {
    rule string
    expected string
}

type failureState struct {
    offset int
    expected []expectation
}

type cacheEntry struct {
    node TreeNode
    offset int
}

type ParseError struct {
    Input string
    Offset int
    Line int
    Column int
    Expected []expectation
    Message string
}

func (e *ParseError) Error() string {
    return e.Message
}

type LispGoParser struct {
    input []rune
    inputString string
    actions Actions
    types map[string]NodeExtender
    offset int
    cache map[string]map[int]cacheEntry
    failure failureState
    actionErr error
}


type Node1 struct {
    BaseNode
    Data TreeNode
}

var _ TreeNode = (*Node1)(nil)

func newNode1(text string, start int, elements []TreeNode) TreeNode {
    node := &Node1{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Data = elements[1]
    return node
}


type Node2 struct {
    BaseNode
    Cells TreeNode
}

var _ TreeNode = (*Node2)(nil)

func newNode2(text string, start int, elements []TreeNode) TreeNode {
    node := &Node2{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Cells = elements[1]
    return node
}


var REGEX_1 = regexp.MustCompile("^[1-9]")
var REGEX_2 = regexp.MustCompile("^[0-9]")
var REGEX_3 = regexp.MustCompile("^[^\"]")
var REGEX_4 = regexp.MustCompile("^[\\s]")

func (p *LispGoParser) _read_program() TreeNode {
    var address0 TreeNode
    address0 = nil
    var index0 int
    index0 = p.offset
    var cache0 map[int]cacheEntry
    cache0 = p.cache["program"]
    if cache0 == nil {
        cache0 = make(map[int]cacheEntry)
        p.cache["program"] = cache0
    }
    if entry, ok := cache0[index0]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index1 int
    index1 = p.offset
    var elements0 []TreeNode
    elements0 = nil
    var address1 TreeNode
    address1 = nil
    for {
        address1 = p._read_cell()
        if address1 != nil {
            elements0 = append(elements0, address1)
        } else {
            break
        }
    }
    if len(elements0) >= 1 {
        address0 = &BaseNode{text: p.slice(index1, p.offset), offset: index1, children: elements0}
        p.offset = p.offset
    } else {
        address0 = nil
    }
    cache0[index0] = cacheEntry{node: address0, offset: p.offset}
    return address0
}

func (p *LispGoParser) _read_cell() TreeNode {
    var address2 TreeNode
    address2 = nil
    var index2 int
    index2 = p.offset
    var cache1 map[int]cacheEntry
    cache1 = p.cache["cell"]
    if cache1 == nil {
        cache1 = make(map[int]cacheEntry)
        p.cache["cell"] = cache1
    }
    if entry, ok := cache1[index2]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index3 int
    index3 = p.offset
    var elements1 []TreeNode
    elements1 = make([]TreeNode, 3)
    var address3 TreeNode
    var index4 int
    index4 = p.offset
    var elements2 []TreeNode
    elements2 = nil
    var address4 TreeNode
    address4 = nil
    for {
        address4 = p._read_space()
        if address4 != nil {
            elements2 = append(elements2, address4)
        } else {
            break
        }
    }
    if len(elements2) >= 0 {
        address3 = &BaseNode{text: p.slice(index4, p.offset), offset: index4, children: elements2}
        p.offset = p.offset
    } else {
        address3 = nil
    }
    if address3 != nil {
        elements1[0] = address3
        var address5 TreeNode
        var index5 int
        index5 = p.offset
        address5 = p._read_list()
        if address5 == nil {
            p.offset = index5
            address5 = p._read_atom()
            if address5 == nil {
                p.offset = index5
            }
        }
        if address5 != nil {
            elements1[1] = address5
            var address6 TreeNode
            var index6 int
            index6 = p.offset
            var elements3 []TreeNode
            elements3 = nil
            var address7 TreeNode
            address7 = nil
            for {
                address7 = p._read_space()
                if address7 != nil {
                    elements3 = append(elements3, address7)
                } else {
                    break
                }
            }
            if len(elements3) >= 0 {
                address6 = &BaseNode{text: p.slice(index6, p.offset), offset: index6, children: elements3}
                p.offset = p.offset
            } else {
                address6 = nil
            }
            if address6 != nil {
                elements1[2] = address6
            } else {
                elements1 = nil
                p.offset = index3
            }
        } else {
            elements1 = nil
            p.offset = index3
        }
    } else {
        elements1 = nil
        p.offset = index3
    }
    if elements1 == nil {
        address2 = nil
    } else {
        address2 = newNode1(p.slice(index3, p.offset), index3, elements1)
        p.offset = p.offset
    }
    cache1[index2] = cacheEntry{node: address2, offset: p.offset}
    return address2
}

func (p *LispGoParser) _read_list() TreeNode {
    var address8 TreeNode
    address8 = nil
    var index7 int
    index7 = p.offset
    var cache2 map[int]cacheEntry
    cache2 = p.cache["list"]
    if cache2 == nil {
        cache2 = make(map[int]cacheEntry)
        p.cache["list"] = cache2
    }
    if entry, ok := cache2[index7]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index8 int
    index8 = p.offset
    var elements4 []TreeNode
    elements4 = make([]TreeNode, 3)
    var address9 TreeNode
    var chunk0 string
    chunk0 = ""
    var max0 int
    max0 = p.offset + 1
    if max0 <= len(p.input) {
        chunk0 = string(p.input[p.offset:max0])
    }
    if chunk0 == "(" {
        address9 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address9 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::list", expected: "\"(\""})
        }
    }
    if address9 != nil {
        elements4[0] = address9
        var address10 TreeNode
        var index9 int
        index9 = p.offset
        var elements5 []TreeNode
        elements5 = nil
        var address11 TreeNode
        address11 = nil
        for {
            address11 = p._read_cell()
            if address11 != nil {
                elements5 = append(elements5, address11)
            } else {
                break
            }
        }
        if len(elements5) >= 1 {
            address10 = &BaseNode{text: p.slice(index9, p.offset), offset: index9, children: elements5}
            p.offset = p.offset
        } else {
            address10 = nil
        }
        if address10 != nil {
            elements4[1] = address10
            var address12 TreeNode
            var chunk1 string
            chunk1 = ""
            var max1 int
            max1 = p.offset + 1
            if max1 <= len(p.input) {
                chunk1 = string(p.input[p.offset:max1])
            }
            if chunk1 == ")" {
                address12 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address12 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::list", expected: "\")\""})
                }
            }
            if address12 != nil {
                elements4[2] = address12
            } else {
                elements4 = nil
                p.offset = index8
            }
        } else {
            elements4 = nil
            p.offset = index8
        }
    } else {
        elements4 = nil
        p.offset = index8
    }
    if elements4 == nil {
        address8 = nil
    } else {
        address8 = newNode2(p.slice(index8, p.offset), index8, elements4)
        p.offset = p.offset
    }
    cache2[index7] = cacheEntry{node: address8, offset: p.offset}
    return address8
}

func (p *LispGoParser) _read_atom() TreeNode {
    var address13 TreeNode
    address13 = nil
    var index10 int
    index10 = p.offset
    var cache3 map[int]cacheEntry
    cache3 = p.cache["atom"]
    if cache3 == nil {
        cache3 = make(map[int]cacheEntry)
        p.cache["atom"] = cache3
    }
    if entry, ok := cache3[index10]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index11 int
    index11 = p.offset
    address13 = p._read_boolean_()
    if address13 == nil {
        p.offset = index11
        address13 = p._read_integer()
        if address13 == nil {
            p.offset = index11
            address13 = p._read_string()
            if address13 == nil {
                p.offset = index11
                address13 = p._read_symbol()
                if address13 == nil {
                    p.offset = index11
                }
            }
        }
    }
    cache3[index10] = cacheEntry{node: address13, offset: p.offset}
    return address13
}

func (p *LispGoParser) _read_boolean_() TreeNode {
    var address14 TreeNode
    address14 = nil
    var index12 int
    index12 = p.offset
    var cache4 map[int]cacheEntry
    cache4 = p.cache["boolean_"]
    if cache4 == nil {
        cache4 = make(map[int]cacheEntry)
        p.cache["boolean_"] = cache4
    }
    if entry, ok := cache4[index12]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index13 int
    index13 = p.offset
    var chunk2 string
    chunk2 = ""
    var max2 int
    max2 = p.offset + 2
    if max2 <= len(p.input) {
        chunk2 = string(p.input[p.offset:max2])
    }
    if chunk2 == "#t" {
        address14 = &BaseNode{text: p.slice(p.offset, p.offset + 2), offset: p.offset, children: nil}
        p.offset = p.offset + 2
    } else {
        address14 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::boolean_", expected: "\"#t\""})
        }
    }
    if address14 == nil {
        p.offset = index13
        var chunk3 string
        chunk3 = ""
        var max3 int
        max3 = p.offset + 2
        if max3 <= len(p.input) {
            chunk3 = string(p.input[p.offset:max3])
        }
        if chunk3 == "#f" {
            address14 = &BaseNode{text: p.slice(p.offset, p.offset + 2), offset: p.offset, children: nil}
            p.offset = p.offset + 2
        } else {
            address14 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::boolean_", expected: "\"#f\""})
            }
        }
        if address14 == nil {
            p.offset = index13
        }
    }
    cache4[index12] = cacheEntry{node: address14, offset: p.offset}
    return address14
}

func (p *LispGoParser) _read_integer() TreeNode {
    var address15 TreeNode
    address15 = nil
    var index14 int
    index14 = p.offset
    var cache5 map[int]cacheEntry
    cache5 = p.cache["integer"]
    if cache5 == nil {
        cache5 = make(map[int]cacheEntry)
        p.cache["integer"] = cache5
    }
    if entry, ok := cache5[index14]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index15 int
    index15 = p.offset
    var elements6 []TreeNode
    elements6 = make([]TreeNode, 2)
    var address16 TreeNode
    var chunk4 string
    chunk4 = ""
    var max4 int
    max4 = p.offset + 1
    if max4 <= len(p.input) {
        chunk4 = string(p.input[p.offset:max4])
    }
    if REGEX_1.MatchString(chunk4) {
        address16 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address16 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::integer", expected: "[1-9]"})
        }
    }
    if address16 != nil {
        elements6[0] = address16
        var address17 TreeNode
        var index16 int
        index16 = p.offset
        var elements7 []TreeNode
        elements7 = nil
        var address18 TreeNode
        address18 = nil
        for {
            var chunk5 string
            chunk5 = ""
            var max5 int
            max5 = p.offset + 1
            if max5 <= len(p.input) {
                chunk5 = string(p.input[p.offset:max5])
            }
            if REGEX_2.MatchString(chunk5) {
                address18 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address18 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::integer", expected: "[0-9]"})
                }
            }
            if address18 != nil {
                elements7 = append(elements7, address18)
            } else {
                break
            }
        }
        if len(elements7) >= 0 {
            address17 = &BaseNode{text: p.slice(index16, p.offset), offset: index16, children: elements7}
            p.offset = p.offset
        } else {
            address17 = nil
        }
        if address17 != nil {
            elements6[1] = address17
        } else {
            elements6 = nil
            p.offset = index15
        }
    } else {
        elements6 = nil
        p.offset = index15
    }
    if elements6 == nil {
        address15 = nil
    } else {
        address15 = &BaseNode{text: p.slice(index15, p.offset), offset: index15, children: elements6}
        p.offset = p.offset
    }
    cache5[index14] = cacheEntry{node: address15, offset: p.offset}
    return address15
}

func (p *LispGoParser) _read_string() TreeNode {
    var address19 TreeNode
    address19 = nil
    var index17 int
    index17 = p.offset
    var cache6 map[int]cacheEntry
    cache6 = p.cache["string"]
    if cache6 == nil {
        cache6 = make(map[int]cacheEntry)
        p.cache["string"] = cache6
    }
    if entry, ok := cache6[index17]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index18 int
    index18 = p.offset
    var elements8 []TreeNode
    elements8 = make([]TreeNode, 3)
    var address20 TreeNode
    var chunk6 string
    chunk6 = ""
    var max6 int
    max6 = p.offset + 1
    if max6 <= len(p.input) {
        chunk6 = string(p.input[p.offset:max6])
    }
    if chunk6 == "\"" {
        address20 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address20 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::string", expected: "\"\\\"\""})
        }
    }
    if address20 != nil {
        elements8[0] = address20
        var address21 TreeNode
        var index19 int
        index19 = p.offset
        var elements9 []TreeNode
        elements9 = nil
        var address22 TreeNode
        address22 = nil
        for {
            var index20 int
            index20 = p.offset
            var index21 int
            index21 = p.offset
            var elements10 []TreeNode
            elements10 = make([]TreeNode, 2)
            var address23 TreeNode
            var chunk7 string
            chunk7 = ""
            var max7 int
            max7 = p.offset + 1
            if max7 <= len(p.input) {
                chunk7 = string(p.input[p.offset:max7])
            }
            if chunk7 == "\\" {
                address23 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address23 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::string", expected: "\"\\\\\""})
                }
            }
            if address23 != nil {
                elements10[0] = address23
                var address24 TreeNode
                if p.offset < len(p.input) {
                    address24 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address24 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::string", expected: "<any char>"})
                    }
                }
                if address24 != nil {
                    elements10[1] = address24
                } else {
                    elements10 = nil
                    p.offset = index21
                }
            } else {
                elements10 = nil
                p.offset = index21
            }
            if elements10 == nil {
                address22 = nil
            } else {
                address22 = &BaseNode{text: p.slice(index21, p.offset), offset: index21, children: elements10}
                p.offset = p.offset
            }
            if address22 == nil {
                p.offset = index20
                var chunk8 string
                chunk8 = ""
                var max8 int
                max8 = p.offset + 1
                if max8 <= len(p.input) {
                    chunk8 = string(p.input[p.offset:max8])
                }
                if REGEX_3.MatchString(chunk8) {
                    address22 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address22 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::string", expected: "[^\"]"})
                    }
                }
                if address22 == nil {
                    p.offset = index20
                }
            }
            if address22 != nil {
                elements9 = append(elements9, address22)
            } else {
                break
            }
        }
        if len(elements9) >= 0 {
            address21 = &BaseNode{text: p.slice(index19, p.offset), offset: index19, children: elements9}
            p.offset = p.offset
        } else {
            address21 = nil
        }
        if address21 != nil {
            elements8[1] = address21
            var address25 TreeNode
            var chunk9 string
            chunk9 = ""
            var max9 int
            max9 = p.offset + 1
            if max9 <= len(p.input) {
                chunk9 = string(p.input[p.offset:max9])
            }
            if chunk9 == "\"" {
                address25 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address25 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::string", expected: "\"\\\"\""})
                }
            }
            if address25 != nil {
                elements8[2] = address25
            } else {
                elements8 = nil
                p.offset = index18
            }
        } else {
            elements8 = nil
            p.offset = index18
        }
    } else {
        elements8 = nil
        p.offset = index18
    }
    if elements8 == nil {
        address19 = nil
    } else {
        address19 = &BaseNode{text: p.slice(index18, p.offset), offset: index18, children: elements8}
        p.offset = p.offset
    }
    cache6[index17] = cacheEntry{node: address19, offset: p.offset}
    return address19
}

func (p *LispGoParser) _read_symbol() TreeNode {
    var address26 TreeNode
    address26 = nil
    var index22 int
    index22 = p.offset
    var cache7 map[int]cacheEntry
    cache7 = p.cache["symbol"]
    if cache7 == nil {
        cache7 = make(map[int]cacheEntry)
        p.cache["symbol"] = cache7
    }
    if entry, ok := cache7[index22]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index23 int
    index23 = p.offset
    var elements11 []TreeNode
    elements11 = nil
    var address27 TreeNode
    address27 = nil
    for {
        var index24 int
        index24 = p.offset
        var elements12 []TreeNode
        elements12 = make([]TreeNode, 2)
        var address28 TreeNode
        var index25 int
        index25 = p.offset
        address28 = p._read_delimiter()
        p.offset = index25
        if address28 == nil {
            address28 = &BaseNode{text: p.slice(p.offset, p.offset), offset: p.offset, children: nil}
            p.offset = p.offset
        } else {
            address28 = nil
        }
        if address28 != nil {
            elements12[0] = address28
            var address29 TreeNode
            if p.offset < len(p.input) {
                address29 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address29 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::symbol", expected: "<any char>"})
                }
            }
            if address29 != nil {
                elements12[1] = address29
            } else {
                elements12 = nil
                p.offset = index24
            }
        } else {
            elements12 = nil
            p.offset = index24
        }
        if elements12 == nil {
            address27 = nil
        } else {
            address27 = &BaseNode{text: p.slice(index24, p.offset), offset: index24, children: elements12}
            p.offset = p.offset
        }
        if address27 != nil {
            elements11 = append(elements11, address27)
        } else {
            break
        }
    }
    if len(elements11) >= 1 {
        address26 = &BaseNode{text: p.slice(index23, p.offset), offset: index23, children: elements11}
        p.offset = p.offset
    } else {
        address26 = nil
    }
    cache7[index22] = cacheEntry{node: address26, offset: p.offset}
    return address26
}

func (p *LispGoParser) _read_space() TreeNode {
    var address30 TreeNode
    address30 = nil
    var index26 int
    index26 = p.offset
    var cache8 map[int]cacheEntry
    cache8 = p.cache["space"]
    if cache8 == nil {
        cache8 = make(map[int]cacheEntry)
        p.cache["space"] = cache8
    }
    if entry, ok := cache8[index26]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var chunk10 string
    chunk10 = ""
    var max10 int
    max10 = p.offset + 1
    if max10 <= len(p.input) {
        chunk10 = string(p.input[p.offset:max10])
    }
    if REGEX_4.MatchString(chunk10) {
        address30 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address30 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::space", expected: "[\\s]"})
        }
    }
    cache8[index26] = cacheEntry{node: address30, offset: p.offset}
    return address30
}

func (p *LispGoParser) _read_paren() TreeNode {
    var address31 TreeNode
    address31 = nil
    var index27 int
    index27 = p.offset
    var cache9 map[int]cacheEntry
    cache9 = p.cache["paren"]
    if cache9 == nil {
        cache9 = make(map[int]cacheEntry)
        p.cache["paren"] = cache9
    }
    if entry, ok := cache9[index27]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index28 int
    index28 = p.offset
    var chunk11 string
    chunk11 = ""
    var max11 int
    max11 = p.offset + 1
    if max11 <= len(p.input) {
        chunk11 = string(p.input[p.offset:max11])
    }
    if chunk11 == "(" {
        address31 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address31 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::paren", expected: "\"(\""})
        }
    }
    if address31 == nil {
        p.offset = index28
        var chunk12 string
        chunk12 = ""
        var max12 int
        max12 = p.offset + 1
        if max12 <= len(p.input) {
            chunk12 = string(p.input[p.offset:max12])
        }
        if chunk12 == ")" {
            address31 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address31 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp::paren", expected: "\")\""})
            }
        }
        if address31 == nil {
            p.offset = index28
        }
    }
    cache9[index27] = cacheEntry{node: address31, offset: p.offset}
    return address31
}

func (p *LispGoParser) _read_delimiter() TreeNode {
    var address32 TreeNode
    address32 = nil
    var index29 int
    index29 = p.offset
    var cache10 map[int]cacheEntry
    cache10 = p.cache["delimiter"]
    if cache10 == nil {
        cache10 = make(map[int]cacheEntry)
        p.cache["delimiter"] = cache10
    }
    if entry, ok := cache10[index29]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index30 int
    index30 = p.offset
    address32 = p._read_paren()
    if address32 == nil {
        p.offset = index30
        address32 = p._read_space()
        if address32 == nil {
            p.offset = index30
        }
    }
    cache10[index29] = cacheEntry{node: address32, offset: p.offset}
    return address32
}

func New(input string, actions Actions) *LispGoParser {
    return &LispGoParser{
        input: []rune(input),
        inputString: input,
        actions: actions,
        cache: make(map[string]map[int]cacheEntry),
    }
}

func (p *LispGoParser) WithTypes(types map[string]NodeExtender) *LispGoParser {
    p.types = types
    return p
}

func Parse(input string, actions Actions, types map[string]NodeExtender) (TreeNode, error) {
    parser := New(input, actions)
    if types != nil {
        parser.types = types
    }
    return parser.Parse()
}

func (p *LispGoParser) Parse() (TreeNode, error) {
    node := p._read_program()
    if p.actionErr != nil {
        return nil, p.actionErr
    }
    if node != nil && p.offset == len(p.input) {
        return node, nil
    }
    if len(p.failure.expected) == 0 {
        p.failure.offset = p.offset
        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyLisp", expected: "<EOF>"})
    }
    return nil, p.newParseError()
}

func (p *LispGoParser) newParseError() error {
    line, column := 1, 1
    for i, r := range p.input {
        if i >= p.failure.offset {
            break
        }
        if r == '\n' {
            line++
            column = 1
        } else {
            column++
        }
    }
    message := fmt.Sprintf("parse error at line %d, column %d", line, column)
    if len(p.failure.expected) > 0 {
        message += ": expected "
        for i, exp := range p.failure.expected {
            if i > 0 {
                if i == len(p.failure.expected)-1 {
                    message += " or "
                } else {
                    message += ", "
                }
            }
            message += fmt.Sprintf("%s from %s", exp.expected, exp.rule)
        }
    }
    expected := make([]expectation, len(p.failure.expected))
    copy(expected, p.failure.expected)
    return &ParseError{
        Input: p.inputString,
        Offset: p.failure.offset,
        Line: line,
        Column: column,
        Expected: expected,
        Message: message,
    }
}

func (p *LispGoParser) slice(start, end int) string {
    if start < 0 { start = 0 }
    if end > len(p.input) { end = len(p.input) }
    if start > end { start = end }
    return string(p.input[start:end])
}

func (p *LispGoParser) extendNode(node TreeNode, name string) TreeNode {
    if node == nil {
        return nil
    }
    if p.types == nil {
        return node
    }
    if extender, ok := p.types[name]; ok && extender != nil {
        return extender(node)
    }
    return node
}

