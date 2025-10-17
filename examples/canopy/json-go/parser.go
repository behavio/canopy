// This file was generated from examples/canopy/json.peg
// See https://canopy.jcoglan.com/ for documentation

package jsongoparser

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

type JsonGoParser struct {
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
}

var _ TreeNode = (*Node1)(nil)

func newNode1(text string, start int, elements []TreeNode) TreeNode {
    node := &Node1{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    return node
}


type Node2 struct {
    BaseNode
    Pair TreeNode
}

var _ TreeNode = (*Node2)(nil)

func newNode2(text string, start int, elements []TreeNode) TreeNode {
    node := &Node2{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Pair = elements[1]
    return node
}


type Node3 struct {
    BaseNode
    Pair TreeNode
}

var _ TreeNode = (*Node3)(nil)

func newNode3(text string, start int, elements []TreeNode) TreeNode {
    node := &Node3{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Pair = elements[1]
    return node
}


type Node4 struct {
    BaseNode
}

var _ TreeNode = (*Node4)(nil)

func newNode4(text string, start int, elements []TreeNode) TreeNode {
    node := &Node4{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    return node
}


type Node5 struct {
    BaseNode
    String TreeNode
    Value TreeNode
}

var _ TreeNode = (*Node5)(nil)

func newNode5(text string, start int, elements []TreeNode) TreeNode {
    node := &Node5{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.String = elements[1]
    node.Value = elements[4]
    return node
}


type Node6 struct {
    BaseNode
    Value TreeNode
}

var _ TreeNode = (*Node6)(nil)

func newNode6(text string, start int, elements []TreeNode) TreeNode {
    node := &Node6{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Value = elements[1]
    return node
}


type Node7 struct {
    BaseNode
    Value TreeNode
}

var _ TreeNode = (*Node7)(nil)

func newNode7(text string, start int, elements []TreeNode) TreeNode {
    node := &Node7{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Value = elements[1]
    return node
}


type Node8 struct {
    BaseNode
}

var _ TreeNode = (*Node8)(nil)

func newNode8(text string, start int, elements []TreeNode) TreeNode {
    node := &Node8{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    return node
}


type Node9 struct {
    BaseNode
}

var _ TreeNode = (*Node9)(nil)

func newNode9(text string, start int, elements []TreeNode) TreeNode {
    node := &Node9{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    return node
}


var REGEX_1 = regexp.MustCompile("^[^\"]")
var REGEX_2 = regexp.MustCompile("^[1-9]")
var REGEX_3 = regexp.MustCompile("^[0-9]")
var REGEX_4 = regexp.MustCompile("^[0-9]")
var REGEX_5 = regexp.MustCompile("^[0-9]")
var REGEX_6 = regexp.MustCompile("^[\\s]")

func (p *JsonGoParser) _read_document() TreeNode {
    var address0 TreeNode
    address0 = nil
    var index0 int
    index0 = p.offset
    var cache0 map[int]cacheEntry
    cache0 = p.cache["document"]
    if cache0 == nil {
        cache0 = make(map[int]cacheEntry)
        p.cache["document"] = cache0
    }
    if entry, ok := cache0[index0]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index1 int
    index1 = p.offset
    var elements0 []TreeNode
    elements0 = make([]TreeNode, 3)
    var address1 TreeNode
    address1 = p._read___()
    if address1 != nil {
        elements0[0] = address1
        var address2 TreeNode
        var index2 int
        index2 = p.offset
        address2 = p._read_object()
        if address2 == nil {
            p.offset = index2
            address2 = p._read_array()
            if address2 == nil {
                p.offset = index2
            }
        }
        if address2 != nil {
            elements0[1] = address2
            var address3 TreeNode
            address3 = p._read___()
            if address3 != nil {
                elements0[2] = address3
            } else {
                elements0 = nil
                p.offset = index1
            }
        } else {
            elements0 = nil
            p.offset = index1
        }
    } else {
        elements0 = nil
        p.offset = index1
    }
    if elements0 == nil {
        address0 = nil
    } else {
        address0 = newNode1(p.slice(index1, p.offset), index1, elements0)
        p.offset = p.offset
    }
    cache0[index0] = cacheEntry{node: address0, offset: p.offset}
    return address0
}

func (p *JsonGoParser) _read_object() TreeNode {
    var address4 TreeNode
    address4 = nil
    var index3 int
    index3 = p.offset
    var cache1 map[int]cacheEntry
    cache1 = p.cache["object"]
    if cache1 == nil {
        cache1 = make(map[int]cacheEntry)
        p.cache["object"] = cache1
    }
    if entry, ok := cache1[index3]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index4 int
    index4 = p.offset
    var index5 int
    index5 = p.offset
    var elements1 []TreeNode
    elements1 = make([]TreeNode, 4)
    var address5 TreeNode
    var chunk0 string
    chunk0 = ""
    var max0 int
    max0 = p.offset + 1
    if max0 <= len(p.input) {
        chunk0 = string(p.input[p.offset:max0])
    }
    if chunk0 == "{" {
        address5 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address5 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::object", expected: "\"{\""})
        }
    }
    if address5 != nil {
        elements1[0] = address5
        var address6 TreeNode
        address6 = p._read_pair()
        if address6 != nil {
            elements1[1] = address6
            var address7 TreeNode
            var index6 int
            index6 = p.offset
            var elements2 []TreeNode
            elements2 = nil
            var address8 TreeNode
            address8 = nil
            for {
                var index7 int
                index7 = p.offset
                var elements3 []TreeNode
                elements3 = make([]TreeNode, 2)
                var address9 TreeNode
                var chunk1 string
                chunk1 = ""
                var max1 int
                max1 = p.offset + 1
                if max1 <= len(p.input) {
                    chunk1 = string(p.input[p.offset:max1])
                }
                if chunk1 == "," {
                    address9 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address9 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::object", expected: "\",\""})
                    }
                }
                if address9 != nil {
                    elements3[0] = address9
                    var address10 TreeNode
                    address10 = p._read_pair()
                    if address10 != nil {
                        elements3[1] = address10
                    } else {
                        elements3 = nil
                        p.offset = index7
                    }
                } else {
                    elements3 = nil
                    p.offset = index7
                }
                if elements3 == nil {
                    address8 = nil
                } else {
                    address8 = newNode3(p.slice(index7, p.offset), index7, elements3)
                    p.offset = p.offset
                }
                if address8 != nil {
                    elements2 = append(elements2, address8)
                } else {
                    break
                }
            }
            if len(elements2) >= 0 {
                address7 = &BaseNode{text: p.slice(index6, p.offset), offset: index6, children: elements2}
                p.offset = p.offset
            } else {
                address7 = nil
            }
            if address7 != nil {
                elements1[2] = address7
                var address11 TreeNode
                var chunk2 string
                chunk2 = ""
                var max2 int
                max2 = p.offset + 1
                if max2 <= len(p.input) {
                    chunk2 = string(p.input[p.offset:max2])
                }
                if chunk2 == "}" {
                    address11 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address11 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::object", expected: "\"}\""})
                    }
                }
                if address11 != nil {
                    elements1[3] = address11
                } else {
                    elements1 = nil
                    p.offset = index5
                }
            } else {
                elements1 = nil
                p.offset = index5
            }
        } else {
            elements1 = nil
            p.offset = index5
        }
    } else {
        elements1 = nil
        p.offset = index5
    }
    if elements1 == nil {
        address4 = nil
    } else {
        address4 = newNode2(p.slice(index5, p.offset), index5, elements1)
        p.offset = p.offset
    }
    if address4 == nil {
        p.offset = index4
        var index8 int
        index8 = p.offset
        var elements4 []TreeNode
        elements4 = make([]TreeNode, 3)
        var address12 TreeNode
        var chunk3 string
        chunk3 = ""
        var max3 int
        max3 = p.offset + 1
        if max3 <= len(p.input) {
            chunk3 = string(p.input[p.offset:max3])
        }
        if chunk3 == "{" {
            address12 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address12 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::object", expected: "\"{\""})
            }
        }
        if address12 != nil {
            elements4[0] = address12
            var address13 TreeNode
            address13 = p._read___()
            if address13 != nil {
                elements4[1] = address13
                var address14 TreeNode
                var chunk4 string
                chunk4 = ""
                var max4 int
                max4 = p.offset + 1
                if max4 <= len(p.input) {
                    chunk4 = string(p.input[p.offset:max4])
                }
                if chunk4 == "}" {
                    address14 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address14 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::object", expected: "\"}\""})
                    }
                }
                if address14 != nil {
                    elements4[2] = address14
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
            address4 = nil
        } else {
            address4 = newNode4(p.slice(index8, p.offset), index8, elements4)
            p.offset = p.offset
        }
        if address4 == nil {
            p.offset = index4
        }
    }
    cache1[index3] = cacheEntry{node: address4, offset: p.offset}
    return address4
}

func (p *JsonGoParser) _read_pair() TreeNode {
    var address15 TreeNode
    address15 = nil
    var index9 int
    index9 = p.offset
    var cache2 map[int]cacheEntry
    cache2 = p.cache["pair"]
    if cache2 == nil {
        cache2 = make(map[int]cacheEntry)
        p.cache["pair"] = cache2
    }
    if entry, ok := cache2[index9]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index10 int
    index10 = p.offset
    var elements5 []TreeNode
    elements5 = make([]TreeNode, 5)
    var address16 TreeNode
    address16 = p._read___()
    if address16 != nil {
        elements5[0] = address16
        var address17 TreeNode
        address17 = p._read_string()
        if address17 != nil {
            elements5[1] = address17
            var address18 TreeNode
            address18 = p._read___()
            if address18 != nil {
                elements5[2] = address18
                var address19 TreeNode
                var chunk5 string
                chunk5 = ""
                var max5 int
                max5 = p.offset + 1
                if max5 <= len(p.input) {
                    chunk5 = string(p.input[p.offset:max5])
                }
                if chunk5 == ":" {
                    address19 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address19 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::pair", expected: "\":\""})
                    }
                }
                if address19 != nil {
                    elements5[3] = address19
                    var address20 TreeNode
                    address20 = p._read_value()
                    if address20 != nil {
                        elements5[4] = address20
                    } else {
                        elements5 = nil
                        p.offset = index10
                    }
                } else {
                    elements5 = nil
                    p.offset = index10
                }
            } else {
                elements5 = nil
                p.offset = index10
            }
        } else {
            elements5 = nil
            p.offset = index10
        }
    } else {
        elements5 = nil
        p.offset = index10
    }
    if elements5 == nil {
        address15 = nil
    } else {
        address15 = newNode5(p.slice(index10, p.offset), index10, elements5)
        p.offset = p.offset
    }
    cache2[index9] = cacheEntry{node: address15, offset: p.offset}
    return address15
}

func (p *JsonGoParser) _read_array() TreeNode {
    var address21 TreeNode
    address21 = nil
    var index11 int
    index11 = p.offset
    var cache3 map[int]cacheEntry
    cache3 = p.cache["array"]
    if cache3 == nil {
        cache3 = make(map[int]cacheEntry)
        p.cache["array"] = cache3
    }
    if entry, ok := cache3[index11]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index12 int
    index12 = p.offset
    var index13 int
    index13 = p.offset
    var elements6 []TreeNode
    elements6 = make([]TreeNode, 4)
    var address22 TreeNode
    var chunk6 string
    chunk6 = ""
    var max6 int
    max6 = p.offset + 1
    if max6 <= len(p.input) {
        chunk6 = string(p.input[p.offset:max6])
    }
    if chunk6 == "[" {
        address22 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address22 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::array", expected: "\"[\""})
        }
    }
    if address22 != nil {
        elements6[0] = address22
        var address23 TreeNode
        address23 = p._read_value()
        if address23 != nil {
            elements6[1] = address23
            var address24 TreeNode
            var index14 int
            index14 = p.offset
            var elements7 []TreeNode
            elements7 = nil
            var address25 TreeNode
            address25 = nil
            for {
                var index15 int
                index15 = p.offset
                var elements8 []TreeNode
                elements8 = make([]TreeNode, 2)
                var address26 TreeNode
                var chunk7 string
                chunk7 = ""
                var max7 int
                max7 = p.offset + 1
                if max7 <= len(p.input) {
                    chunk7 = string(p.input[p.offset:max7])
                }
                if chunk7 == "," {
                    address26 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address26 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::array", expected: "\",\""})
                    }
                }
                if address26 != nil {
                    elements8[0] = address26
                    var address27 TreeNode
                    address27 = p._read_value()
                    if address27 != nil {
                        elements8[1] = address27
                    } else {
                        elements8 = nil
                        p.offset = index15
                    }
                } else {
                    elements8 = nil
                    p.offset = index15
                }
                if elements8 == nil {
                    address25 = nil
                } else {
                    address25 = newNode7(p.slice(index15, p.offset), index15, elements8)
                    p.offset = p.offset
                }
                if address25 != nil {
                    elements7 = append(elements7, address25)
                } else {
                    break
                }
            }
            if len(elements7) >= 0 {
                address24 = &BaseNode{text: p.slice(index14, p.offset), offset: index14, children: elements7}
                p.offset = p.offset
            } else {
                address24 = nil
            }
            if address24 != nil {
                elements6[2] = address24
                var address28 TreeNode
                var chunk8 string
                chunk8 = ""
                var max8 int
                max8 = p.offset + 1
                if max8 <= len(p.input) {
                    chunk8 = string(p.input[p.offset:max8])
                }
                if chunk8 == "]" {
                    address28 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address28 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::array", expected: "\"]\""})
                    }
                }
                if address28 != nil {
                    elements6[3] = address28
                } else {
                    elements6 = nil
                    p.offset = index13
                }
            } else {
                elements6 = nil
                p.offset = index13
            }
        } else {
            elements6 = nil
            p.offset = index13
        }
    } else {
        elements6 = nil
        p.offset = index13
    }
    if elements6 == nil {
        address21 = nil
    } else {
        address21 = newNode6(p.slice(index13, p.offset), index13, elements6)
        p.offset = p.offset
    }
    if address21 == nil {
        p.offset = index12
        var index16 int
        index16 = p.offset
        var elements9 []TreeNode
        elements9 = make([]TreeNode, 3)
        var address29 TreeNode
        var chunk9 string
        chunk9 = ""
        var max9 int
        max9 = p.offset + 1
        if max9 <= len(p.input) {
            chunk9 = string(p.input[p.offset:max9])
        }
        if chunk9 == "[" {
            address29 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address29 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::array", expected: "\"[\""})
            }
        }
        if address29 != nil {
            elements9[0] = address29
            var address30 TreeNode
            address30 = p._read___()
            if address30 != nil {
                elements9[1] = address30
                var address31 TreeNode
                var chunk10 string
                chunk10 = ""
                var max10 int
                max10 = p.offset + 1
                if max10 <= len(p.input) {
                    chunk10 = string(p.input[p.offset:max10])
                }
                if chunk10 == "]" {
                    address31 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address31 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::array", expected: "\"]\""})
                    }
                }
                if address31 != nil {
                    elements9[2] = address31
                } else {
                    elements9 = nil
                    p.offset = index16
                }
            } else {
                elements9 = nil
                p.offset = index16
            }
        } else {
            elements9 = nil
            p.offset = index16
        }
        if elements9 == nil {
            address21 = nil
        } else {
            address21 = newNode8(p.slice(index16, p.offset), index16, elements9)
            p.offset = p.offset
        }
        if address21 == nil {
            p.offset = index12
        }
    }
    cache3[index11] = cacheEntry{node: address21, offset: p.offset}
    return address21
}

func (p *JsonGoParser) _read_value() TreeNode {
    var address32 TreeNode
    address32 = nil
    var index17 int
    index17 = p.offset
    var cache4 map[int]cacheEntry
    cache4 = p.cache["value"]
    if cache4 == nil {
        cache4 = make(map[int]cacheEntry)
        p.cache["value"] = cache4
    }
    if entry, ok := cache4[index17]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index18 int
    index18 = p.offset
    var elements10 []TreeNode
    elements10 = make([]TreeNode, 3)
    var address33 TreeNode
    address33 = p._read___()
    if address33 != nil {
        elements10[0] = address33
        var address34 TreeNode
        var index19 int
        index19 = p.offset
        address34 = p._read_object()
        if address34 == nil {
            p.offset = index19
            address34 = p._read_array()
            if address34 == nil {
                p.offset = index19
                address34 = p._read_string()
                if address34 == nil {
                    p.offset = index19
                    address34 = p._read_number()
                    if address34 == nil {
                        p.offset = index19
                        address34 = p._read_boolean_()
                        if address34 == nil {
                            p.offset = index19
                            address34 = p._read_null_()
                            if address34 == nil {
                                p.offset = index19
                            }
                        }
                    }
                }
            }
        }
        if address34 != nil {
            elements10[1] = address34
            var address35 TreeNode
            address35 = p._read___()
            if address35 != nil {
                elements10[2] = address35
            } else {
                elements10 = nil
                p.offset = index18
            }
        } else {
            elements10 = nil
            p.offset = index18
        }
    } else {
        elements10 = nil
        p.offset = index18
    }
    if elements10 == nil {
        address32 = nil
    } else {
        address32 = newNode9(p.slice(index18, p.offset), index18, elements10)
        p.offset = p.offset
    }
    cache4[index17] = cacheEntry{node: address32, offset: p.offset}
    return address32
}

func (p *JsonGoParser) _read_string() TreeNode {
    var address36 TreeNode
    address36 = nil
    var index20 int
    index20 = p.offset
    var cache5 map[int]cacheEntry
    cache5 = p.cache["string"]
    if cache5 == nil {
        cache5 = make(map[int]cacheEntry)
        p.cache["string"] = cache5
    }
    if entry, ok := cache5[index20]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index21 int
    index21 = p.offset
    var elements11 []TreeNode
    elements11 = make([]TreeNode, 3)
    var address37 TreeNode
    var chunk11 string
    chunk11 = ""
    var max11 int
    max11 = p.offset + 1
    if max11 <= len(p.input) {
        chunk11 = string(p.input[p.offset:max11])
    }
    if chunk11 == "\"" {
        address37 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address37 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::string", expected: "'\"'"})
        }
    }
    if address37 != nil {
        elements11[0] = address37
        var address38 TreeNode
        var index22 int
        index22 = p.offset
        var elements12 []TreeNode
        elements12 = nil
        var address39 TreeNode
        address39 = nil
        for {
            var index23 int
            index23 = p.offset
            var index24 int
            index24 = p.offset
            var elements13 []TreeNode
            elements13 = make([]TreeNode, 2)
            var address40 TreeNode
            var chunk12 string
            chunk12 = ""
            var max12 int
            max12 = p.offset + 1
            if max12 <= len(p.input) {
                chunk12 = string(p.input[p.offset:max12])
            }
            if chunk12 == "\\" {
                address40 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address40 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::string", expected: "\"\\\\\""})
                }
            }
            if address40 != nil {
                elements13[0] = address40
                var address41 TreeNode
                if p.offset < len(p.input) {
                    address41 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address41 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::string", expected: "<any char>"})
                    }
                }
                if address41 != nil {
                    elements13[1] = address41
                } else {
                    elements13 = nil
                    p.offset = index24
                }
            } else {
                elements13 = nil
                p.offset = index24
            }
            if elements13 == nil {
                address39 = nil
            } else {
                address39 = &BaseNode{text: p.slice(index24, p.offset), offset: index24, children: elements13}
                p.offset = p.offset
            }
            if address39 == nil {
                p.offset = index23
                var chunk13 string
                chunk13 = ""
                var max13 int
                max13 = p.offset + 1
                if max13 <= len(p.input) {
                    chunk13 = string(p.input[p.offset:max13])
                }
                if REGEX_1.MatchString(chunk13) {
                    address39 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address39 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::string", expected: "[^\"]"})
                    }
                }
                if address39 == nil {
                    p.offset = index23
                }
            }
            if address39 != nil {
                elements12 = append(elements12, address39)
            } else {
                break
            }
        }
        if len(elements12) >= 0 {
            address38 = &BaseNode{text: p.slice(index22, p.offset), offset: index22, children: elements12}
            p.offset = p.offset
        } else {
            address38 = nil
        }
        if address38 != nil {
            elements11[1] = address38
            var address42 TreeNode
            var chunk14 string
            chunk14 = ""
            var max14 int
            max14 = p.offset + 1
            if max14 <= len(p.input) {
                chunk14 = string(p.input[p.offset:max14])
            }
            if chunk14 == "\"" {
                address42 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address42 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::string", expected: "'\"'"})
                }
            }
            if address42 != nil {
                elements11[2] = address42
            } else {
                elements11 = nil
                p.offset = index21
            }
        } else {
            elements11 = nil
            p.offset = index21
        }
    } else {
        elements11 = nil
        p.offset = index21
    }
    if elements11 == nil {
        address36 = nil
    } else {
        address36 = &BaseNode{text: p.slice(index21, p.offset), offset: index21, children: elements11}
        p.offset = p.offset
    }
    cache5[index20] = cacheEntry{node: address36, offset: p.offset}
    return address36
}

func (p *JsonGoParser) _read_number() TreeNode {
    var address43 TreeNode
    address43 = nil
    var index25 int
    index25 = p.offset
    var cache6 map[int]cacheEntry
    cache6 = p.cache["number"]
    if cache6 == nil {
        cache6 = make(map[int]cacheEntry)
        p.cache["number"] = cache6
    }
    if entry, ok := cache6[index25]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index26 int
    index26 = p.offset
    var elements14 []TreeNode
    elements14 = make([]TreeNode, 4)
    var address44 TreeNode
    var index27 int
    index27 = p.offset
    var chunk15 string
    chunk15 = ""
    var max15 int
    max15 = p.offset + 1
    if max15 <= len(p.input) {
        chunk15 = string(p.input[p.offset:max15])
    }
    if chunk15 == "-" {
        address44 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address44 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"-\""})
        }
    }
    if address44 == nil {
        address44 = &BaseNode{text: p.slice(index27, index27), offset: index27, children: nil}
        p.offset = index27
    }
    if address44 != nil {
        elements14[0] = address44
        var address45 TreeNode
        var index28 int
        index28 = p.offset
        var chunk16 string
        chunk16 = ""
        var max16 int
        max16 = p.offset + 1
        if max16 <= len(p.input) {
            chunk16 = string(p.input[p.offset:max16])
        }
        if chunk16 == "0" {
            address45 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address45 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"0\""})
            }
        }
        if address45 == nil {
            p.offset = index28
            var index29 int
            index29 = p.offset
            var elements15 []TreeNode
            elements15 = make([]TreeNode, 2)
            var address46 TreeNode
            var chunk17 string
            chunk17 = ""
            var max17 int
            max17 = p.offset + 1
            if max17 <= len(p.input) {
                chunk17 = string(p.input[p.offset:max17])
            }
            if REGEX_2.MatchString(chunk17) {
                address46 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address46 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "[1-9]"})
                }
            }
            if address46 != nil {
                elements15[0] = address46
                var address47 TreeNode
                var index30 int
                index30 = p.offset
                var elements16 []TreeNode
                elements16 = nil
                var address48 TreeNode
                address48 = nil
                for {
                    var chunk18 string
                    chunk18 = ""
                    var max18 int
                    max18 = p.offset + 1
                    if max18 <= len(p.input) {
                        chunk18 = string(p.input[p.offset:max18])
                    }
                    if REGEX_3.MatchString(chunk18) {
                        address48 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address48 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "[0-9]"})
                        }
                    }
                    if address48 != nil {
                        elements16 = append(elements16, address48)
                    } else {
                        break
                    }
                }
                if len(elements16) >= 0 {
                    address47 = &BaseNode{text: p.slice(index30, p.offset), offset: index30, children: elements16}
                    p.offset = p.offset
                } else {
                    address47 = nil
                }
                if address47 != nil {
                    elements15[1] = address47
                } else {
                    elements15 = nil
                    p.offset = index29
                }
            } else {
                elements15 = nil
                p.offset = index29
            }
            if elements15 == nil {
                address45 = nil
            } else {
                address45 = &BaseNode{text: p.slice(index29, p.offset), offset: index29, children: elements15}
                p.offset = p.offset
            }
            if address45 == nil {
                p.offset = index28
            }
        }
        if address45 != nil {
            elements14[1] = address45
            var address49 TreeNode
            var index31 int
            index31 = p.offset
            var index32 int
            index32 = p.offset
            var elements17 []TreeNode
            elements17 = make([]TreeNode, 2)
            var address50 TreeNode
            var chunk19 string
            chunk19 = ""
            var max19 int
            max19 = p.offset + 1
            if max19 <= len(p.input) {
                chunk19 = string(p.input[p.offset:max19])
            }
            if chunk19 == "." {
                address50 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address50 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\".\""})
                }
            }
            if address50 != nil {
                elements17[0] = address50
                var address51 TreeNode
                var index33 int
                index33 = p.offset
                var elements18 []TreeNode
                elements18 = nil
                var address52 TreeNode
                address52 = nil
                for {
                    var chunk20 string
                    chunk20 = ""
                    var max20 int
                    max20 = p.offset + 1
                    if max20 <= len(p.input) {
                        chunk20 = string(p.input[p.offset:max20])
                    }
                    if REGEX_4.MatchString(chunk20) {
                        address52 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address52 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "[0-9]"})
                        }
                    }
                    if address52 != nil {
                        elements18 = append(elements18, address52)
                    } else {
                        break
                    }
                }
                if len(elements18) >= 1 {
                    address51 = &BaseNode{text: p.slice(index33, p.offset), offset: index33, children: elements18}
                    p.offset = p.offset
                } else {
                    address51 = nil
                }
                if address51 != nil {
                    elements17[1] = address51
                } else {
                    elements17 = nil
                    p.offset = index32
                }
            } else {
                elements17 = nil
                p.offset = index32
            }
            if elements17 == nil {
                address49 = nil
            } else {
                address49 = &BaseNode{text: p.slice(index32, p.offset), offset: index32, children: elements17}
                p.offset = p.offset
            }
            if address49 == nil {
                address49 = &BaseNode{text: p.slice(index31, index31), offset: index31, children: nil}
                p.offset = index31
            }
            if address49 != nil {
                elements14[2] = address49
                var address53 TreeNode
                var index34 int
                index34 = p.offset
                var index35 int
                index35 = p.offset
                var elements19 []TreeNode
                elements19 = make([]TreeNode, 3)
                var address54 TreeNode
                var index36 int
                index36 = p.offset
                var chunk21 string
                chunk21 = ""
                var max21 int
                max21 = p.offset + 1
                if max21 <= len(p.input) {
                    chunk21 = string(p.input[p.offset:max21])
                }
                if chunk21 == "e" {
                    address54 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address54 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"e\""})
                    }
                }
                if address54 == nil {
                    p.offset = index36
                    var chunk22 string
                    chunk22 = ""
                    var max22 int
                    max22 = p.offset + 1
                    if max22 <= len(p.input) {
                        chunk22 = string(p.input[p.offset:max22])
                    }
                    if chunk22 == "E" {
                        address54 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address54 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"E\""})
                        }
                    }
                    if address54 == nil {
                        p.offset = index36
                    }
                }
                if address54 != nil {
                    elements19[0] = address54
                    var address55 TreeNode
                    var index37 int
                    index37 = p.offset
                    var chunk23 string
                    chunk23 = ""
                    var max23 int
                    max23 = p.offset + 1
                    if max23 <= len(p.input) {
                        chunk23 = string(p.input[p.offset:max23])
                    }
                    if chunk23 == "+" {
                        address55 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address55 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"+\""})
                        }
                    }
                    if address55 == nil {
                        p.offset = index37
                        var chunk24 string
                        chunk24 = ""
                        var max24 int
                        max24 = p.offset + 1
                        if max24 <= len(p.input) {
                            chunk24 = string(p.input[p.offset:max24])
                        }
                        if chunk24 == "-" {
                            address55 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                            p.offset = p.offset + 1
                        } else {
                            address55 = nil
                            if p.offset > p.failure.offset {
                                p.failure.offset = p.offset
                                p.failure.expected = nil
                            }
                            if p.offset == p.failure.offset {
                                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"-\""})
                            }
                        }
                        if address55 == nil {
                            p.offset = index37
                            var chunk25 string
                            chunk25 = ""
                            var max25 int
                            max25 = p.offset + 0
                            if max25 <= len(p.input) {
                                chunk25 = string(p.input[p.offset:max25])
                            }
                            if chunk25 == "" {
                                address55 = &BaseNode{text: p.slice(p.offset, p.offset + 0), offset: p.offset, children: nil}
                                p.offset = p.offset + 0
                            } else {
                                address55 = nil
                                if p.offset > p.failure.offset {
                                    p.failure.offset = p.offset
                                    p.failure.expected = nil
                                }
                                if p.offset == p.failure.offset {
                                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "\"\""})
                                }
                            }
                            if address55 == nil {
                                p.offset = index37
                            }
                        }
                    }
                    if address55 != nil {
                        elements19[1] = address55
                        var address56 TreeNode
                        var index38 int
                        index38 = p.offset
                        var elements20 []TreeNode
                        elements20 = nil
                        var address57 TreeNode
                        address57 = nil
                        for {
                            var chunk26 string
                            chunk26 = ""
                            var max26 int
                            max26 = p.offset + 1
                            if max26 <= len(p.input) {
                                chunk26 = string(p.input[p.offset:max26])
                            }
                            if REGEX_5.MatchString(chunk26) {
                                address57 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                                p.offset = p.offset + 1
                            } else {
                                address57 = nil
                                if p.offset > p.failure.offset {
                                    p.failure.offset = p.offset
                                    p.failure.expected = nil
                                }
                                if p.offset == p.failure.offset {
                                    p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::number", expected: "[0-9]"})
                                }
                            }
                            if address57 != nil {
                                elements20 = append(elements20, address57)
                            } else {
                                break
                            }
                        }
                        if len(elements20) >= 1 {
                            address56 = &BaseNode{text: p.slice(index38, p.offset), offset: index38, children: elements20}
                            p.offset = p.offset
                        } else {
                            address56 = nil
                        }
                        if address56 != nil {
                            elements19[2] = address56
                        } else {
                            elements19 = nil
                            p.offset = index35
                        }
                    } else {
                        elements19 = nil
                        p.offset = index35
                    }
                } else {
                    elements19 = nil
                    p.offset = index35
                }
                if elements19 == nil {
                    address53 = nil
                } else {
                    address53 = &BaseNode{text: p.slice(index35, p.offset), offset: index35, children: elements19}
                    p.offset = p.offset
                }
                if address53 == nil {
                    address53 = &BaseNode{text: p.slice(index34, index34), offset: index34, children: nil}
                    p.offset = index34
                }
                if address53 != nil {
                    elements14[3] = address53
                } else {
                    elements14 = nil
                    p.offset = index26
                }
            } else {
                elements14 = nil
                p.offset = index26
            }
        } else {
            elements14 = nil
            p.offset = index26
        }
    } else {
        elements14 = nil
        p.offset = index26
    }
    if elements14 == nil {
        address43 = nil
    } else {
        address43 = &BaseNode{text: p.slice(index26, p.offset), offset: index26, children: elements14}
        p.offset = p.offset
    }
    cache6[index25] = cacheEntry{node: address43, offset: p.offset}
    return address43
}

func (p *JsonGoParser) _read_boolean_() TreeNode {
    var address58 TreeNode
    address58 = nil
    var index39 int
    index39 = p.offset
    var cache7 map[int]cacheEntry
    cache7 = p.cache["boolean_"]
    if cache7 == nil {
        cache7 = make(map[int]cacheEntry)
        p.cache["boolean_"] = cache7
    }
    if entry, ok := cache7[index39]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index40 int
    index40 = p.offset
    var chunk27 string
    chunk27 = ""
    var max27 int
    max27 = p.offset + 4
    if max27 <= len(p.input) {
        chunk27 = string(p.input[p.offset:max27])
    }
    if chunk27 == "true" {
        address58 = &BaseNode{text: p.slice(p.offset, p.offset + 4), offset: p.offset, children: nil}
        p.offset = p.offset + 4
    } else {
        address58 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::boolean_", expected: "\"true\""})
        }
    }
    if address58 == nil {
        p.offset = index40
        var chunk28 string
        chunk28 = ""
        var max28 int
        max28 = p.offset + 5
        if max28 <= len(p.input) {
            chunk28 = string(p.input[p.offset:max28])
        }
        if chunk28 == "false" {
            address58 = &BaseNode{text: p.slice(p.offset, p.offset + 5), offset: p.offset, children: nil}
            p.offset = p.offset + 5
        } else {
            address58 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::boolean_", expected: "\"false\""})
            }
        }
        if address58 == nil {
            p.offset = index40
        }
    }
    cache7[index39] = cacheEntry{node: address58, offset: p.offset}
    return address58
}

func (p *JsonGoParser) _read_null_() TreeNode {
    var address59 TreeNode
    address59 = nil
    var index41 int
    index41 = p.offset
    var cache8 map[int]cacheEntry
    cache8 = p.cache["null_"]
    if cache8 == nil {
        cache8 = make(map[int]cacheEntry)
        p.cache["null_"] = cache8
    }
    if entry, ok := cache8[index41]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var chunk29 string
    chunk29 = ""
    var max29 int
    max29 = p.offset + 4
    if max29 <= len(p.input) {
        chunk29 = string(p.input[p.offset:max29])
    }
    if chunk29 == "null" {
        address59 = &BaseNode{text: p.slice(p.offset, p.offset + 4), offset: p.offset, children: nil}
        p.offset = p.offset + 4
    } else {
        address59 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::null_", expected: "\"null\""})
        }
    }
    cache8[index41] = cacheEntry{node: address59, offset: p.offset}
    return address59
}

func (p *JsonGoParser) _read___() TreeNode {
    var address60 TreeNode
    address60 = nil
    var index42 int
    index42 = p.offset
    var cache9 map[int]cacheEntry
    cache9 = p.cache["__"]
    if cache9 == nil {
        cache9 = make(map[int]cacheEntry)
        p.cache["__"] = cache9
    }
    if entry, ok := cache9[index42]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index43 int
    index43 = p.offset
    var elements21 []TreeNode
    elements21 = nil
    var address61 TreeNode
    address61 = nil
    for {
        var chunk30 string
        chunk30 = ""
        var max30 int
        max30 = p.offset + 1
        if max30 <= len(p.input) {
            chunk30 = string(p.input[p.offset:max30])
        }
        if REGEX_6.MatchString(chunk30) {
            address61 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address61 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson::__", expected: "[\\s]"})
            }
        }
        if address61 != nil {
            elements21 = append(elements21, address61)
        } else {
            break
        }
    }
    if len(elements21) >= 0 {
        address60 = &BaseNode{text: p.slice(index43, p.offset), offset: index43, children: elements21}
        p.offset = p.offset
    } else {
        address60 = nil
    }
    cache9[index42] = cacheEntry{node: address60, offset: p.offset}
    return address60
}

func New(input string, actions Actions) *JsonGoParser {
    return &JsonGoParser{
        input: []rune(input),
        inputString: input,
        actions: actions,
        cache: make(map[string]map[int]cacheEntry),
    }
}

func (p *JsonGoParser) WithTypes(types map[string]NodeExtender) *JsonGoParser {
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

func (p *JsonGoParser) Parse() (TreeNode, error) {
    node := p._read_document()
    if p.actionErr != nil {
        return nil, p.actionErr
    }
    if node != nil && p.offset == len(p.input) {
        return node, nil
    }
    if len(p.failure.expected) == 0 {
        p.failure.offset = p.offset
        p.failure.expected = append(p.failure.expected, expectation{rule: "CanopyJson", expected: "<EOF>"})
    }
    return nil, p.newParseError()
}

func (p *JsonGoParser) newParseError() error {
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

func (p *JsonGoParser) slice(start, end int) string {
    if start < 0 { start = 0 }
    if end > len(p.input) { end = len(p.input) }
    if start > end { start = end }
    return string(p.input[start:end])
}

func (p *JsonGoParser) extendNode(node TreeNode, name string) TreeNode {
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

