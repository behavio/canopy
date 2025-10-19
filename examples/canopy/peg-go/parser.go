// This file was generated from examples/canopy/peg.peg
// See https://canopy.jcoglan.com/ for documentation

package peggoparser

import (
    	"fmt"
	"regexp"
	"strings"
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

type PegGoParser struct {
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
    GrammarName TreeNode
    Rules TreeNode
}

var _ TreeNode = (*Node1)(nil)

func newNode1(text string, start int, elements []TreeNode) TreeNode {
    node := &Node1{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.GrammarName = elements[1]
    node.Rules = elements[2]
    return node
}


type Node2 struct {
    BaseNode
    GrammarRule TreeNode
}

var _ TreeNode = (*Node2)(nil)

func newNode2(text string, start int, elements []TreeNode) TreeNode {
    node := &Node2{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.GrammarRule = elements[1]
    return node
}


type Node3 struct {
    BaseNode
    ObjectIdentifier TreeNode
}

var _ TreeNode = (*Node3)(nil)

func newNode3(text string, start int, elements []TreeNode) TreeNode {
    node := &Node3{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.ObjectIdentifier = elements[3]
    return node
}


type Node4 struct {
    BaseNode
    Identifier TreeNode
    Assignment TreeNode
    ParsingExpression TreeNode
}

var _ TreeNode = (*Node4)(nil)

func newNode4(text string, start int, elements []TreeNode) TreeNode {
    node := &Node4{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Identifier = elements[0]
    node.Assignment = elements[1]
    node.ParsingExpression = elements[2]
    return node
}


type Node5 struct {
    BaseNode
    ParsingExpression TreeNode
}

var _ TreeNode = (*Node5)(nil)

func newNode5(text string, start int, elements []TreeNode) TreeNode {
    node := &Node5{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.ParsingExpression = elements[2]
    return node
}


type Node6 struct {
    BaseNode
    FirstPart TreeNode
    ChoicePart TreeNode
    Rest TreeNode
}

var _ TreeNode = (*Node6)(nil)

func newNode6(text string, start int, elements []TreeNode) TreeNode {
    node := &Node6{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.FirstPart = elements[0]
    node.ChoicePart = elements[0]
    node.Rest = elements[1]
    return node
}


type Node7 struct {
    BaseNode
    Expression TreeNode
    ChoicePart TreeNode
}

var _ TreeNode = (*Node7)(nil)

func newNode7(text string, start int, elements []TreeNode) TreeNode {
    node := &Node7{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Expression = elements[3]
    node.ChoicePart = elements[3]
    return node
}


type Node8 struct {
    BaseNode
    TypeTag TreeNode
}

var _ TreeNode = (*Node8)(nil)

func newNode8(text string, start int, elements []TreeNode) TreeNode {
    node := &Node8{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.TypeTag = elements[1]
    return node
}


type Node9 struct {
    BaseNode
    ActionableExpression TreeNode
    ActionTag TreeNode
}

var _ TreeNode = (*Node9)(nil)

func newNode9(text string, start int, elements []TreeNode) TreeNode {
    node := &Node9{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.ActionableExpression = elements[0]
    node.ActionTag = elements[2]
    return node
}


type Node10 struct {
    BaseNode
    ActionableExpression TreeNode
}

var _ TreeNode = (*Node10)(nil)

func newNode10(text string, start int, elements []TreeNode) TreeNode {
    node := &Node10{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.ActionableExpression = elements[2]
    return node
}


type Node11 struct {
    BaseNode
    Identifier TreeNode
}

var _ TreeNode = (*Node11)(nil)

func newNode11(text string, start int, elements []TreeNode) TreeNode {
    node := &Node11{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Identifier = elements[1]
    return node
}


type Node12 struct {
    BaseNode
    ObjectIdentifier TreeNode
}

var _ TreeNode = (*Node12)(nil)

func newNode12(text string, start int, elements []TreeNode) TreeNode {
    node := &Node12{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.ObjectIdentifier = elements[1]
    return node
}


type Node13 struct {
    BaseNode
    FirstPart TreeNode
    SequencePart TreeNode
    Rest TreeNode
}

var _ TreeNode = (*Node13)(nil)

func newNode13(text string, start int, elements []TreeNode) TreeNode {
    node := &Node13{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.FirstPart = elements[0]
    node.SequencePart = elements[0]
    node.Rest = elements[1]
    return node
}


type Node14 struct {
    BaseNode
    Expression TreeNode
    SequencePart TreeNode
}

var _ TreeNode = (*Node14)(nil)

func newNode14(text string, start int, elements []TreeNode) TreeNode {
    node := &Node14{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Expression = elements[1]
    node.SequencePart = elements[1]
    return node
}


type Node15 struct {
    BaseNode
    Expression TreeNode
}

var _ TreeNode = (*Node15)(nil)

func newNode15(text string, start int, elements []TreeNode) TreeNode {
    node := &Node15{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Expression = elements[1]
    return node
}


type Node16 struct {
    BaseNode
    Atom TreeNode
}

var _ TreeNode = (*Node16)(nil)

func newNode16(text string, start int, elements []TreeNode) TreeNode {
    node := &Node16{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Atom = elements[0]
    return node
}


type Node17 struct {
    BaseNode
    Atom TreeNode
    Quantifier TreeNode
}

var _ TreeNode = (*Node17)(nil)

func newNode17(text string, start int, elements []TreeNode) TreeNode {
    node := &Node17{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Atom = elements[0]
    node.Quantifier = elements[1]
    return node
}


type Node18 struct {
    BaseNode
    Predicate TreeNode
    Atom TreeNode
}

var _ TreeNode = (*Node18)(nil)

func newNode18(text string, start int, elements []TreeNode) TreeNode {
    node := &Node18{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Predicate = elements[0]
    node.Atom = elements[1]
    return node
}


type Node19 struct {
    BaseNode
    Identifier TreeNode
}

var _ TreeNode = (*Node19)(nil)

func newNode19(text string, start int, elements []TreeNode) TreeNode {
    node := &Node19{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Identifier = elements[0]
    return node
}


type Node20 struct {
    BaseNode
    Identifier TreeNode
}

var _ TreeNode = (*Node20)(nil)

func newNode20(text string, start int, elements []TreeNode) TreeNode {
    node := &Node20{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Identifier = elements[0]
    return node
}


type Node21 struct {
    BaseNode
    Identifier TreeNode
}

var _ TreeNode = (*Node21)(nil)

func newNode21(text string, start int, elements []TreeNode) TreeNode {
    node := &Node21{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Identifier = elements[0]
    return node
}


type Node22 struct {
    BaseNode
    Identifier TreeNode
}

var _ TreeNode = (*Node22)(nil)

func newNode22(text string, start int, elements []TreeNode) TreeNode {
    node := &Node22{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.Identifier = elements[1]
    return node
}


var REGEX_1 = regexp.MustCompile("^[^\"]")
var REGEX_2 = regexp.MustCompile("^[^']")
var REGEX_3 = regexp.MustCompile("^[^`]")
var REGEX_4 = regexp.MustCompile("^[^\\]]")
var REGEX_5 = regexp.MustCompile("^[a-zA-Z_]")
var REGEX_6 = regexp.MustCompile("^[a-zA-Z0-9_]")
var REGEX_7 = regexp.MustCompile("^[\\s]")
var REGEX_8 = regexp.MustCompile("^[^\\n]")

func (p *PegGoParser) _read_grammar() TreeNode {
    var address0 TreeNode
    address0 = nil
    var index0 int
    index0 = p.offset
    var cache0 map[int]cacheEntry
    cache0 = p.cache["grammar"]
    if cache0 == nil {
        cache0 = make(map[int]cacheEntry)
        p.cache["grammar"] = cache0
    }
    if entry, ok := cache0[index0]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index1 int
    index1 = p.offset
    var elements0 []TreeNode
    elements0 = make([]TreeNode, 4)
    var address1 TreeNode
    var index2 int
    index2 = p.offset
    var elements1 []TreeNode
    elements1 = nil
    var address2 TreeNode
    address2 = nil
    for {
        address2 = p._read___()
        if address2 != nil {
            elements1 = append(elements1, address2)
        } else {
            break
        }
    }
    if len(elements1) >= 0 {
        address1 = &BaseNode{text: p.slice(index2, p.offset), offset: index2, children: elements1}
        p.offset = p.offset
    } else {
        address1 = nil
    }
    if address1 != nil {
        elements0[0] = address1
        var address3 TreeNode
        address3 = p._read_grammar_name()
        if address3 != nil {
            elements0[1] = address3
            var address4 TreeNode
            var index3 int
            index3 = p.offset
            var elements2 []TreeNode
            elements2 = nil
            var address5 TreeNode
            address5 = nil
            for {
                var index4 int
                index4 = p.offset
                var elements3 []TreeNode
                elements3 = make([]TreeNode, 2)
                var address6 TreeNode
                var index5 int
                index5 = p.offset
                var elements4 []TreeNode
                elements4 = nil
                var address7 TreeNode
                address7 = nil
                for {
                    address7 = p._read___()
                    if address7 != nil {
                        elements4 = append(elements4, address7)
                    } else {
                        break
                    }
                }
                if len(elements4) >= 0 {
                    address6 = &BaseNode{text: p.slice(index5, p.offset), offset: index5, children: elements4}
                    p.offset = p.offset
                } else {
                    address6 = nil
                }
                if address6 != nil {
                    elements3[0] = address6
                    var address8 TreeNode
                    address8 = p._read_grammar_rule()
                    if address8 != nil {
                        elements3[1] = address8
                    } else {
                        elements3 = nil
                        p.offset = index4
                    }
                } else {
                    elements3 = nil
                    p.offset = index4
                }
                if elements3 == nil {
                    address5 = nil
                } else {
                    address5 = newNode2(p.slice(index4, p.offset), index4, elements3)
                    p.offset = p.offset
                }
                if address5 != nil {
                    elements2 = append(elements2, address5)
                } else {
                    break
                }
            }
            if len(elements2) >= 1 {
                address4 = &BaseNode{text: p.slice(index3, p.offset), offset: index3, children: elements2}
                p.offset = p.offset
            } else {
                address4 = nil
            }
            if address4 != nil {
                elements0[2] = address4
                var address9 TreeNode
                var index6 int
                index6 = p.offset
                var elements5 []TreeNode
                elements5 = nil
                var address10 TreeNode
                address10 = nil
                for {
                    address10 = p._read___()
                    if address10 != nil {
                        elements5 = append(elements5, address10)
                    } else {
                        break
                    }
                }
                if len(elements5) >= 0 {
                    address9 = &BaseNode{text: p.slice(index6, p.offset), offset: index6, children: elements5}
                    p.offset = p.offset
                } else {
                    address9 = nil
                }
                if address9 != nil {
                    elements0[3] = address9
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

func (p *PegGoParser) _read_grammar_name() TreeNode {
    var address11 TreeNode
    address11 = nil
    var index7 int
    index7 = p.offset
    var cache1 map[int]cacheEntry
    cache1 = p.cache["grammar_name"]
    if cache1 == nil {
        cache1 = make(map[int]cacheEntry)
        p.cache["grammar_name"] = cache1
    }
    if entry, ok := cache1[index7]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index8 int
    index8 = p.offset
    var elements6 []TreeNode
    elements6 = make([]TreeNode, 4)
    var address12 TreeNode
    var chunk0 string
    chunk0 = ""
    var max0 int
    max0 = p.offset + 7
    if max0 <= len(p.input) {
        chunk0 = string(p.input[p.offset:max0])
    }
    if strings.EqualFold(chunk0, "grammar") {
        address12 = &BaseNode{text: p.slice(p.offset, p.offset + 7), offset: p.offset, children: nil}
        p.offset = p.offset + 7
    } else {
        address12 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::grammar_name", expected: "`grammar`"})
        }
    }
    if address12 != nil {
        elements6[0] = address12
        var address13 TreeNode
        var index9 int
        index9 = p.offset
        var chunk1 string
        chunk1 = ""
        var max1 int
        max1 = p.offset + 1
        if max1 <= len(p.input) {
            chunk1 = string(p.input[p.offset:max1])
        }
        if chunk1 == ":" {
            address13 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address13 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::grammar_name", expected: "\":\""})
            }
        }
        if address13 == nil {
            address13 = &BaseNode{text: p.slice(index9, index9), offset: index9, children: nil}
            p.offset = index9
        }
        if address13 != nil {
            elements6[1] = address13
            var address14 TreeNode
            var index10 int
            index10 = p.offset
            var elements7 []TreeNode
            elements7 = nil
            var address15 TreeNode
            address15 = nil
            for {
                address15 = p._read___()
                if address15 != nil {
                    elements7 = append(elements7, address15)
                } else {
                    break
                }
            }
            if len(elements7) >= 1 {
                address14 = &BaseNode{text: p.slice(index10, p.offset), offset: index10, children: elements7}
                p.offset = p.offset
            } else {
                address14 = nil
            }
            if address14 != nil {
                elements6[2] = address14
                var address16 TreeNode
                address16 = p._read_object_identifier()
                if address16 != nil {
                    elements6[3] = address16
                } else {
                    elements6 = nil
                    p.offset = index8
                }
            } else {
                elements6 = nil
                p.offset = index8
            }
        } else {
            elements6 = nil
            p.offset = index8
        }
    } else {
        elements6 = nil
        p.offset = index8
    }
    if elements6 == nil {
        address11 = nil
    } else {
        address11 = newNode3(p.slice(index8, p.offset), index8, elements6)
        p.offset = p.offset
    }
    cache1[index7] = cacheEntry{node: address11, offset: p.offset}
    return address11
}

func (p *PegGoParser) _read_grammar_rule() TreeNode {
    var address17 TreeNode
    address17 = nil
    var index11 int
    index11 = p.offset
    var cache2 map[int]cacheEntry
    cache2 = p.cache["grammar_rule"]
    if cache2 == nil {
        cache2 = make(map[int]cacheEntry)
        p.cache["grammar_rule"] = cache2
    }
    if entry, ok := cache2[index11]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index12 int
    index12 = p.offset
    var elements8 []TreeNode
    elements8 = make([]TreeNode, 3)
    var address18 TreeNode
    address18 = p._read_identifier()
    if address18 != nil {
        elements8[0] = address18
        var address19 TreeNode
        address19 = p._read_assignment()
        if address19 != nil {
            elements8[1] = address19
            var address20 TreeNode
            address20 = p._read_parsing_expression()
            if address20 != nil {
                elements8[2] = address20
            } else {
                elements8 = nil
                p.offset = index12
            }
        } else {
            elements8 = nil
            p.offset = index12
        }
    } else {
        elements8 = nil
        p.offset = index12
    }
    if elements8 == nil {
        address17 = nil
    } else {
        address17 = newNode4(p.slice(index12, p.offset), index12, elements8)
        p.offset = p.offset
    }
    cache2[index11] = cacheEntry{node: address17, offset: p.offset}
    return address17
}

func (p *PegGoParser) _read_assignment() TreeNode {
    var address21 TreeNode
    address21 = nil
    var index13 int
    index13 = p.offset
    var cache3 map[int]cacheEntry
    cache3 = p.cache["assignment"]
    if cache3 == nil {
        cache3 = make(map[int]cacheEntry)
        p.cache["assignment"] = cache3
    }
    if entry, ok := cache3[index13]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index14 int
    index14 = p.offset
    var elements9 []TreeNode
    elements9 = make([]TreeNode, 3)
    var address22 TreeNode
    var index15 int
    index15 = p.offset
    var elements10 []TreeNode
    elements10 = nil
    var address23 TreeNode
    address23 = nil
    for {
        address23 = p._read___()
        if address23 != nil {
            elements10 = append(elements10, address23)
        } else {
            break
        }
    }
    if len(elements10) >= 1 {
        address22 = &BaseNode{text: p.slice(index15, p.offset), offset: index15, children: elements10}
        p.offset = p.offset
    } else {
        address22 = nil
    }
    if address22 != nil {
        elements9[0] = address22
        var address24 TreeNode
        var chunk2 string
        chunk2 = ""
        var max2 int
        max2 = p.offset + 2
        if max2 <= len(p.input) {
            chunk2 = string(p.input[p.offset:max2])
        }
        if chunk2 == "<-" {
            address24 = &BaseNode{text: p.slice(p.offset, p.offset + 2), offset: p.offset, children: nil}
            p.offset = p.offset + 2
        } else {
            address24 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::assignment", expected: "\"<-\""})
            }
        }
        if address24 != nil {
            elements9[1] = address24
            var address25 TreeNode
            var index16 int
            index16 = p.offset
            var elements11 []TreeNode
            elements11 = nil
            var address26 TreeNode
            address26 = nil
            for {
                address26 = p._read___()
                if address26 != nil {
                    elements11 = append(elements11, address26)
                } else {
                    break
                }
            }
            if len(elements11) >= 1 {
                address25 = &BaseNode{text: p.slice(index16, p.offset), offset: index16, children: elements11}
                p.offset = p.offset
            } else {
                address25 = nil
            }
            if address25 != nil {
                elements9[2] = address25
            } else {
                elements9 = nil
                p.offset = index14
            }
        } else {
            elements9 = nil
            p.offset = index14
        }
    } else {
        elements9 = nil
        p.offset = index14
    }
    if elements9 == nil {
        address21 = nil
    } else {
        address21 = &BaseNode{text: p.slice(index14, p.offset), offset: index14, children: elements9}
        p.offset = p.offset
    }
    cache3[index13] = cacheEntry{node: address21, offset: p.offset}
    return address21
}

func (p *PegGoParser) _read_parsing_expression() TreeNode {
    var address27 TreeNode
    address27 = nil
    var index17 int
    index17 = p.offset
    var cache4 map[int]cacheEntry
    cache4 = p.cache["parsing_expression"]
    if cache4 == nil {
        cache4 = make(map[int]cacheEntry)
        p.cache["parsing_expression"] = cache4
    }
    if entry, ok := cache4[index17]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index18 int
    index18 = p.offset
    address27 = p._read_choice_expression()
    if address27 == nil {
        p.offset = index18
        address27 = p._read_choice_part()
        if address27 == nil {
            p.offset = index18
        }
    }
    cache4[index17] = cacheEntry{node: address27, offset: p.offset}
    return address27
}

func (p *PegGoParser) _read_parenthesised_expression() TreeNode {
    var address28 TreeNode
    address28 = nil
    var index19 int
    index19 = p.offset
    var cache5 map[int]cacheEntry
    cache5 = p.cache["parenthesised_expression"]
    if cache5 == nil {
        cache5 = make(map[int]cacheEntry)
        p.cache["parenthesised_expression"] = cache5
    }
    if entry, ok := cache5[index19]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index20 int
    index20 = p.offset
    var elements12 []TreeNode
    elements12 = make([]TreeNode, 5)
    var address29 TreeNode
    var chunk3 string
    chunk3 = ""
    var max3 int
    max3 = p.offset + 1
    if max3 <= len(p.input) {
        chunk3 = string(p.input[p.offset:max3])
    }
    if chunk3 == "(" {
        address29 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address29 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::parenthesised_expression", expected: "\"(\""})
        }
    }
    if address29 != nil {
        elements12[0] = address29
        var address30 TreeNode
        var index21 int
        index21 = p.offset
        var elements13 []TreeNode
        elements13 = nil
        var address31 TreeNode
        address31 = nil
        for {
            address31 = p._read___()
            if address31 != nil {
                elements13 = append(elements13, address31)
            } else {
                break
            }
        }
        if len(elements13) >= 0 {
            address30 = &BaseNode{text: p.slice(index21, p.offset), offset: index21, children: elements13}
            p.offset = p.offset
        } else {
            address30 = nil
        }
        if address30 != nil {
            elements12[1] = address30
            var address32 TreeNode
            address32 = p._read_parsing_expression()
            if address32 != nil {
                elements12[2] = address32
                var address33 TreeNode
                var index22 int
                index22 = p.offset
                var elements14 []TreeNode
                elements14 = nil
                var address34 TreeNode
                address34 = nil
                for {
                    address34 = p._read___()
                    if address34 != nil {
                        elements14 = append(elements14, address34)
                    } else {
                        break
                    }
                }
                if len(elements14) >= 0 {
                    address33 = &BaseNode{text: p.slice(index22, p.offset), offset: index22, children: elements14}
                    p.offset = p.offset
                } else {
                    address33 = nil
                }
                if address33 != nil {
                    elements12[3] = address33
                    var address35 TreeNode
                    var chunk4 string
                    chunk4 = ""
                    var max4 int
                    max4 = p.offset + 1
                    if max4 <= len(p.input) {
                        chunk4 = string(p.input[p.offset:max4])
                    }
                    if chunk4 == ")" {
                        address35 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address35 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::parenthesised_expression", expected: "\")\""})
                        }
                    }
                    if address35 != nil {
                        elements12[4] = address35
                    } else {
                        elements12 = nil
                        p.offset = index20
                    }
                } else {
                    elements12 = nil
                    p.offset = index20
                }
            } else {
                elements12 = nil
                p.offset = index20
            }
        } else {
            elements12 = nil
            p.offset = index20
        }
    } else {
        elements12 = nil
        p.offset = index20
    }
    if elements12 == nil {
        address28 = nil
    } else {
        address28 = newNode5(p.slice(index20, p.offset), index20, elements12)
        p.offset = p.offset
    }
    cache5[index19] = cacheEntry{node: address28, offset: p.offset}
    return address28
}

func (p *PegGoParser) _read_choice_expression() TreeNode {
    var address36 TreeNode
    address36 = nil
    var index23 int
    index23 = p.offset
    var cache6 map[int]cacheEntry
    cache6 = p.cache["choice_expression"]
    if cache6 == nil {
        cache6 = make(map[int]cacheEntry)
        p.cache["choice_expression"] = cache6
    }
    if entry, ok := cache6[index23]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index24 int
    index24 = p.offset
    var elements15 []TreeNode
    elements15 = make([]TreeNode, 2)
    var address37 TreeNode
    address37 = p._read_choice_part()
    if address37 != nil {
        elements15[0] = address37
        var address38 TreeNode
        var index25 int
        index25 = p.offset
        var elements16 []TreeNode
        elements16 = nil
        var address39 TreeNode
        address39 = nil
        for {
            var index26 int
            index26 = p.offset
            var elements17 []TreeNode
            elements17 = make([]TreeNode, 4)
            var address40 TreeNode
            var index27 int
            index27 = p.offset
            var elements18 []TreeNode
            elements18 = nil
            var address41 TreeNode
            address41 = nil
            for {
                address41 = p._read___()
                if address41 != nil {
                    elements18 = append(elements18, address41)
                } else {
                    break
                }
            }
            if len(elements18) >= 1 {
                address40 = &BaseNode{text: p.slice(index27, p.offset), offset: index27, children: elements18}
                p.offset = p.offset
            } else {
                address40 = nil
            }
            if address40 != nil {
                elements17[0] = address40
                var address42 TreeNode
                var chunk5 string
                chunk5 = ""
                var max5 int
                max5 = p.offset + 1
                if max5 <= len(p.input) {
                    chunk5 = string(p.input[p.offset:max5])
                }
                if chunk5 == "/" {
                    address42 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address42 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::choice_expression", expected: "\"/\""})
                    }
                }
                if address42 != nil {
                    elements17[1] = address42
                    var address43 TreeNode
                    var index28 int
                    index28 = p.offset
                    var elements19 []TreeNode
                    elements19 = nil
                    var address44 TreeNode
                    address44 = nil
                    for {
                        address44 = p._read___()
                        if address44 != nil {
                            elements19 = append(elements19, address44)
                        } else {
                            break
                        }
                    }
                    if len(elements19) >= 1 {
                        address43 = &BaseNode{text: p.slice(index28, p.offset), offset: index28, children: elements19}
                        p.offset = p.offset
                    } else {
                        address43 = nil
                    }
                    if address43 != nil {
                        elements17[2] = address43
                        var address45 TreeNode
                        address45 = p._read_choice_part()
                        if address45 != nil {
                            elements17[3] = address45
                        } else {
                            elements17 = nil
                            p.offset = index26
                        }
                    } else {
                        elements17 = nil
                        p.offset = index26
                    }
                } else {
                    elements17 = nil
                    p.offset = index26
                }
            } else {
                elements17 = nil
                p.offset = index26
            }
            if elements17 == nil {
                address39 = nil
            } else {
                address39 = newNode7(p.slice(index26, p.offset), index26, elements17)
                p.offset = p.offset
            }
            if address39 != nil {
                elements16 = append(elements16, address39)
            } else {
                break
            }
        }
        if len(elements16) >= 1 {
            address38 = &BaseNode{text: p.slice(index25, p.offset), offset: index25, children: elements16}
            p.offset = p.offset
        } else {
            address38 = nil
        }
        if address38 != nil {
            elements15[1] = address38
        } else {
            elements15 = nil
            p.offset = index24
        }
    } else {
        elements15 = nil
        p.offset = index24
    }
    if elements15 == nil {
        address36 = nil
    } else {
        address36 = newNode6(p.slice(index24, p.offset), index24, elements15)
        p.offset = p.offset
    }
    cache6[index23] = cacheEntry{node: address36, offset: p.offset}
    return address36
}

func (p *PegGoParser) _read_choice_part() TreeNode {
    var address46 TreeNode
    address46 = nil
    var index29 int
    index29 = p.offset
    var cache7 map[int]cacheEntry
    cache7 = p.cache["choice_part"]
    if cache7 == nil {
        cache7 = make(map[int]cacheEntry)
        p.cache["choice_part"] = cache7
    }
    if entry, ok := cache7[index29]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index30 int
    index30 = p.offset
    var elements20 []TreeNode
    elements20 = make([]TreeNode, 2)
    var address47 TreeNode
    var index31 int
    index31 = p.offset
    address47 = p._read_action_expression()
    if address47 == nil {
        p.offset = index31
        address47 = p._read_sequence_expression()
        if address47 == nil {
            p.offset = index31
            address47 = p._read_sequence_part()
            if address47 == nil {
                p.offset = index31
            }
        }
    }
    if address47 != nil {
        elements20[0] = address47
        var address48 TreeNode
        var index32 int
        index32 = p.offset
        var index33 int
        index33 = p.offset
        var elements21 []TreeNode
        elements21 = make([]TreeNode, 2)
        var address49 TreeNode
        var index34 int
        index34 = p.offset
        var elements22 []TreeNode
        elements22 = nil
        var address50 TreeNode
        address50 = nil
        for {
            address50 = p._read___()
            if address50 != nil {
                elements22 = append(elements22, address50)
            } else {
                break
            }
        }
        if len(elements22) >= 1 {
            address49 = &BaseNode{text: p.slice(index34, p.offset), offset: index34, children: elements22}
            p.offset = p.offset
        } else {
            address49 = nil
        }
        if address49 != nil {
            elements21[0] = address49
            var address51 TreeNode
            address51 = p._read_type_tag()
            if address51 != nil {
                elements21[1] = address51
            } else {
                elements21 = nil
                p.offset = index33
            }
        } else {
            elements21 = nil
            p.offset = index33
        }
        if elements21 == nil {
            address48 = nil
        } else {
            address48 = newNode8(p.slice(index33, p.offset), index33, elements21)
            p.offset = p.offset
        }
        if address48 == nil {
            address48 = &BaseNode{text: p.slice(index32, index32), offset: index32, children: nil}
            p.offset = index32
        }
        if address48 != nil {
            elements20[1] = address48
        } else {
            elements20 = nil
            p.offset = index30
        }
    } else {
        elements20 = nil
        p.offset = index30
    }
    if elements20 == nil {
        address46 = nil
    } else {
        address46 = &BaseNode{text: p.slice(index30, p.offset), offset: index30, children: elements20}
        p.offset = p.offset
    }
    cache7[index29] = cacheEntry{node: address46, offset: p.offset}
    return address46
}

func (p *PegGoParser) _read_action_expression() TreeNode {
    var address52 TreeNode
    address52 = nil
    var index35 int
    index35 = p.offset
    var cache8 map[int]cacheEntry
    cache8 = p.cache["action_expression"]
    if cache8 == nil {
        cache8 = make(map[int]cacheEntry)
        p.cache["action_expression"] = cache8
    }
    if entry, ok := cache8[index35]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index36 int
    index36 = p.offset
    var elements23 []TreeNode
    elements23 = make([]TreeNode, 3)
    var address53 TreeNode
    address53 = p._read_actionable_expression()
    if address53 != nil {
        elements23[0] = address53
        var address54 TreeNode
        var index37 int
        index37 = p.offset
        var elements24 []TreeNode
        elements24 = nil
        var address55 TreeNode
        address55 = nil
        for {
            address55 = p._read___()
            if address55 != nil {
                elements24 = append(elements24, address55)
            } else {
                break
            }
        }
        if len(elements24) >= 1 {
            address54 = &BaseNode{text: p.slice(index37, p.offset), offset: index37, children: elements24}
            p.offset = p.offset
        } else {
            address54 = nil
        }
        if address54 != nil {
            elements23[1] = address54
            var address56 TreeNode
            address56 = p._read_action_tag()
            if address56 != nil {
                elements23[2] = address56
            } else {
                elements23 = nil
                p.offset = index36
            }
        } else {
            elements23 = nil
            p.offset = index36
        }
    } else {
        elements23 = nil
        p.offset = index36
    }
    if elements23 == nil {
        address52 = nil
    } else {
        address52 = newNode9(p.slice(index36, p.offset), index36, elements23)
        p.offset = p.offset
    }
    cache8[index35] = cacheEntry{node: address52, offset: p.offset}
    return address52
}

func (p *PegGoParser) _read_actionable_expression() TreeNode {
    var address57 TreeNode
    address57 = nil
    var index38 int
    index38 = p.offset
    var cache9 map[int]cacheEntry
    cache9 = p.cache["actionable_expression"]
    if cache9 == nil {
        cache9 = make(map[int]cacheEntry)
        p.cache["actionable_expression"] = cache9
    }
    if entry, ok := cache9[index38]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index39 int
    index39 = p.offset
    var index40 int
    index40 = p.offset
    var elements25 []TreeNode
    elements25 = make([]TreeNode, 5)
    var address58 TreeNode
    var chunk6 string
    chunk6 = ""
    var max6 int
    max6 = p.offset + 1
    if max6 <= len(p.input) {
        chunk6 = string(p.input[p.offset:max6])
    }
    if chunk6 == "(" {
        address58 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address58 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::actionable_expression", expected: "\"(\""})
        }
    }
    if address58 != nil {
        elements25[0] = address58
        var address59 TreeNode
        var index41 int
        index41 = p.offset
        var elements26 []TreeNode
        elements26 = nil
        var address60 TreeNode
        address60 = nil
        for {
            address60 = p._read___()
            if address60 != nil {
                elements26 = append(elements26, address60)
            } else {
                break
            }
        }
        if len(elements26) >= 0 {
            address59 = &BaseNode{text: p.slice(index41, p.offset), offset: index41, children: elements26}
            p.offset = p.offset
        } else {
            address59 = nil
        }
        if address59 != nil {
            elements25[1] = address59
            var address61 TreeNode
            address61 = p._read_actionable_expression()
            if address61 != nil {
                elements25[2] = address61
                var address62 TreeNode
                var index42 int
                index42 = p.offset
                var elements27 []TreeNode
                elements27 = nil
                var address63 TreeNode
                address63 = nil
                for {
                    address63 = p._read___()
                    if address63 != nil {
                        elements27 = append(elements27, address63)
                    } else {
                        break
                    }
                }
                if len(elements27) >= 0 {
                    address62 = &BaseNode{text: p.slice(index42, p.offset), offset: index42, children: elements27}
                    p.offset = p.offset
                } else {
                    address62 = nil
                }
                if address62 != nil {
                    elements25[3] = address62
                    var address64 TreeNode
                    var chunk7 string
                    chunk7 = ""
                    var max7 int
                    max7 = p.offset + 1
                    if max7 <= len(p.input) {
                        chunk7 = string(p.input[p.offset:max7])
                    }
                    if chunk7 == ")" {
                        address64 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address64 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::actionable_expression", expected: "\")\""})
                        }
                    }
                    if address64 != nil {
                        elements25[4] = address64
                    } else {
                        elements25 = nil
                        p.offset = index40
                    }
                } else {
                    elements25 = nil
                    p.offset = index40
                }
            } else {
                elements25 = nil
                p.offset = index40
            }
        } else {
            elements25 = nil
            p.offset = index40
        }
    } else {
        elements25 = nil
        p.offset = index40
    }
    if elements25 == nil {
        address57 = nil
    } else {
        address57 = newNode10(p.slice(index40, p.offset), index40, elements25)
        p.offset = p.offset
    }
    if address57 == nil {
        p.offset = index39
        address57 = p._read_sequence_expression()
        if address57 == nil {
            p.offset = index39
            address57 = p._read_repeated_atom()
            if address57 == nil {
                p.offset = index39
                address57 = p._read_terminal_node()
                if address57 == nil {
                    p.offset = index39
                }
            }
        }
    }
    cache9[index38] = cacheEntry{node: address57, offset: p.offset}
    return address57
}

func (p *PegGoParser) _read_action_tag() TreeNode {
    var address65 TreeNode
    address65 = nil
    var index43 int
    index43 = p.offset
    var cache10 map[int]cacheEntry
    cache10 = p.cache["action_tag"]
    if cache10 == nil {
        cache10 = make(map[int]cacheEntry)
        p.cache["action_tag"] = cache10
    }
    if entry, ok := cache10[index43]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index44 int
    index44 = p.offset
    var elements28 []TreeNode
    elements28 = make([]TreeNode, 2)
    var address66 TreeNode
    var chunk8 string
    chunk8 = ""
    var max8 int
    max8 = p.offset + 1
    if max8 <= len(p.input) {
        chunk8 = string(p.input[p.offset:max8])
    }
    if chunk8 == "%" {
        address66 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address66 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::action_tag", expected: "\"%\""})
        }
    }
    if address66 != nil {
        elements28[0] = address66
        var address67 TreeNode
        address67 = p._read_identifier()
        if address67 != nil {
            elements28[1] = address67
        } else {
            elements28 = nil
            p.offset = index44
        }
    } else {
        elements28 = nil
        p.offset = index44
    }
    if elements28 == nil {
        address65 = nil
    } else {
        address65 = newNode11(p.slice(index44, p.offset), index44, elements28)
        p.offset = p.offset
    }
    cache10[index43] = cacheEntry{node: address65, offset: p.offset}
    return address65
}

func (p *PegGoParser) _read_type_tag() TreeNode {
    var address68 TreeNode
    address68 = nil
    var index45 int
    index45 = p.offset
    var cache11 map[int]cacheEntry
    cache11 = p.cache["type_tag"]
    if cache11 == nil {
        cache11 = make(map[int]cacheEntry)
        p.cache["type_tag"] = cache11
    }
    if entry, ok := cache11[index45]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index46 int
    index46 = p.offset
    var elements29 []TreeNode
    elements29 = make([]TreeNode, 3)
    var address69 TreeNode
    var chunk9 string
    chunk9 = ""
    var max9 int
    max9 = p.offset + 1
    if max9 <= len(p.input) {
        chunk9 = string(p.input[p.offset:max9])
    }
    if chunk9 == "<" {
        address69 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address69 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::type_tag", expected: "\"<\""})
        }
    }
    if address69 != nil {
        elements29[0] = address69
        var address70 TreeNode
        address70 = p._read_object_identifier()
        if address70 != nil {
            elements29[1] = address70
            var address71 TreeNode
            var chunk10 string
            chunk10 = ""
            var max10 int
            max10 = p.offset + 1
            if max10 <= len(p.input) {
                chunk10 = string(p.input[p.offset:max10])
            }
            if chunk10 == ">" {
                address71 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address71 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::type_tag", expected: "\">\""})
                }
            }
            if address71 != nil {
                elements29[2] = address71
            } else {
                elements29 = nil
                p.offset = index46
            }
        } else {
            elements29 = nil
            p.offset = index46
        }
    } else {
        elements29 = nil
        p.offset = index46
    }
    if elements29 == nil {
        address68 = nil
    } else {
        address68 = newNode12(p.slice(index46, p.offset), index46, elements29)
        p.offset = p.offset
    }
    cache11[index45] = cacheEntry{node: address68, offset: p.offset}
    return address68
}

func (p *PegGoParser) _read_sequence_expression() TreeNode {
    var address72 TreeNode
    address72 = nil
    var index47 int
    index47 = p.offset
    var cache12 map[int]cacheEntry
    cache12 = p.cache["sequence_expression"]
    if cache12 == nil {
        cache12 = make(map[int]cacheEntry)
        p.cache["sequence_expression"] = cache12
    }
    if entry, ok := cache12[index47]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index48 int
    index48 = p.offset
    var elements30 []TreeNode
    elements30 = make([]TreeNode, 2)
    var address73 TreeNode
    address73 = p._read_sequence_part()
    if address73 != nil {
        elements30[0] = address73
        var address74 TreeNode
        var index49 int
        index49 = p.offset
        var elements31 []TreeNode
        elements31 = nil
        var address75 TreeNode
        address75 = nil
        for {
            var index50 int
            index50 = p.offset
            var elements32 []TreeNode
            elements32 = make([]TreeNode, 2)
            var address76 TreeNode
            var index51 int
            index51 = p.offset
            var elements33 []TreeNode
            elements33 = nil
            var address77 TreeNode
            address77 = nil
            for {
                address77 = p._read___()
                if address77 != nil {
                    elements33 = append(elements33, address77)
                } else {
                    break
                }
            }
            if len(elements33) >= 1 {
                address76 = &BaseNode{text: p.slice(index51, p.offset), offset: index51, children: elements33}
                p.offset = p.offset
            } else {
                address76 = nil
            }
            if address76 != nil {
                elements32[0] = address76
                var address78 TreeNode
                address78 = p._read_sequence_part()
                if address78 != nil {
                    elements32[1] = address78
                } else {
                    elements32 = nil
                    p.offset = index50
                }
            } else {
                elements32 = nil
                p.offset = index50
            }
            if elements32 == nil {
                address75 = nil
            } else {
                address75 = newNode14(p.slice(index50, p.offset), index50, elements32)
                p.offset = p.offset
            }
            if address75 != nil {
                elements31 = append(elements31, address75)
            } else {
                break
            }
        }
        if len(elements31) >= 1 {
            address74 = &BaseNode{text: p.slice(index49, p.offset), offset: index49, children: elements31}
            p.offset = p.offset
        } else {
            address74 = nil
        }
        if address74 != nil {
            elements30[1] = address74
        } else {
            elements30 = nil
            p.offset = index48
        }
    } else {
        elements30 = nil
        p.offset = index48
    }
    if elements30 == nil {
        address72 = nil
    } else {
        address72 = newNode13(p.slice(index48, p.offset), index48, elements30)
        p.offset = p.offset
    }
    cache12[index47] = cacheEntry{node: address72, offset: p.offset}
    return address72
}

func (p *PegGoParser) _read_sequence_part() TreeNode {
    var address79 TreeNode
    address79 = nil
    var index52 int
    index52 = p.offset
    var cache13 map[int]cacheEntry
    cache13 = p.cache["sequence_part"]
    if cache13 == nil {
        cache13 = make(map[int]cacheEntry)
        p.cache["sequence_part"] = cache13
    }
    if entry, ok := cache13[index52]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index53 int
    index53 = p.offset
    var elements34 []TreeNode
    elements34 = make([]TreeNode, 2)
    var address80 TreeNode
    var index54 int
    index54 = p.offset
    address80 = p._read_label()
    if address80 == nil {
        address80 = &BaseNode{text: p.slice(index54, index54), offset: index54, children: nil}
        p.offset = index54
    }
    if address80 != nil {
        elements34[0] = address80
        var address81 TreeNode
        var index55 int
        index55 = p.offset
        address81 = p._read_maybe_atom()
        if address81 == nil {
            p.offset = index55
            address81 = p._read_repeated_atom()
            if address81 == nil {
                p.offset = index55
                address81 = p._read_atom()
                if address81 == nil {
                    p.offset = index55
                }
            }
        }
        if address81 != nil {
            elements34[1] = address81
        } else {
            elements34 = nil
            p.offset = index53
        }
    } else {
        elements34 = nil
        p.offset = index53
    }
    if elements34 == nil {
        address79 = nil
    } else {
        address79 = newNode15(p.slice(index53, p.offset), index53, elements34)
        p.offset = p.offset
    }
    cache13[index52] = cacheEntry{node: address79, offset: p.offset}
    return address79
}

func (p *PegGoParser) _read_maybe_atom() TreeNode {
    var address82 TreeNode
    address82 = nil
    var index56 int
    index56 = p.offset
    var cache14 map[int]cacheEntry
    cache14 = p.cache["maybe_atom"]
    if cache14 == nil {
        cache14 = make(map[int]cacheEntry)
        p.cache["maybe_atom"] = cache14
    }
    if entry, ok := cache14[index56]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index57 int
    index57 = p.offset
    var elements35 []TreeNode
    elements35 = make([]TreeNode, 2)
    var address83 TreeNode
    address83 = p._read_atom()
    if address83 != nil {
        elements35[0] = address83
        var address84 TreeNode
        var chunk11 string
        chunk11 = ""
        var max11 int
        max11 = p.offset + 1
        if max11 <= len(p.input) {
            chunk11 = string(p.input[p.offset:max11])
        }
        if chunk11 == "?" {
            address84 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address84 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::maybe_atom", expected: "\"?\""})
            }
        }
        if address84 != nil {
            elements35[1] = address84
        } else {
            elements35 = nil
            p.offset = index57
        }
    } else {
        elements35 = nil
        p.offset = index57
    }
    if elements35 == nil {
        address82 = nil
    } else {
        address82 = newNode16(p.slice(index57, p.offset), index57, elements35)
        p.offset = p.offset
    }
    cache14[index56] = cacheEntry{node: address82, offset: p.offset}
    return address82
}

func (p *PegGoParser) _read_repeated_atom() TreeNode {
    var address85 TreeNode
    address85 = nil
    var index58 int
    index58 = p.offset
    var cache15 map[int]cacheEntry
    cache15 = p.cache["repeated_atom"]
    if cache15 == nil {
        cache15 = make(map[int]cacheEntry)
        p.cache["repeated_atom"] = cache15
    }
    if entry, ok := cache15[index58]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index59 int
    index59 = p.offset
    var elements36 []TreeNode
    elements36 = make([]TreeNode, 2)
    var address86 TreeNode
    address86 = p._read_atom()
    if address86 != nil {
        elements36[0] = address86
        var address87 TreeNode
        var index60 int
        index60 = p.offset
        var chunk12 string
        chunk12 = ""
        var max12 int
        max12 = p.offset + 1
        if max12 <= len(p.input) {
            chunk12 = string(p.input[p.offset:max12])
        }
        if chunk12 == "*" {
            address87 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address87 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::repeated_atom", expected: "\"*\""})
            }
        }
        if address87 == nil {
            p.offset = index60
            var chunk13 string
            chunk13 = ""
            var max13 int
            max13 = p.offset + 1
            if max13 <= len(p.input) {
                chunk13 = string(p.input[p.offset:max13])
            }
            if chunk13 == "+" {
                address87 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address87 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::repeated_atom", expected: "\"+\""})
                }
            }
            if address87 == nil {
                p.offset = index60
            }
        }
        if address87 != nil {
            elements36[1] = address87
        } else {
            elements36 = nil
            p.offset = index59
        }
    } else {
        elements36 = nil
        p.offset = index59
    }
    if elements36 == nil {
        address85 = nil
    } else {
        address85 = newNode17(p.slice(index59, p.offset), index59, elements36)
        p.offset = p.offset
    }
    cache15[index58] = cacheEntry{node: address85, offset: p.offset}
    return address85
}

func (p *PegGoParser) _read_atom() TreeNode {
    var address88 TreeNode
    address88 = nil
    var index61 int
    index61 = p.offset
    var cache16 map[int]cacheEntry
    cache16 = p.cache["atom"]
    if cache16 == nil {
        cache16 = make(map[int]cacheEntry)
        p.cache["atom"] = cache16
    }
    if entry, ok := cache16[index61]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index62 int
    index62 = p.offset
    address88 = p._read_parenthesised_expression()
    if address88 == nil {
        p.offset = index62
        address88 = p._read_predicated_atom()
        if address88 == nil {
            p.offset = index62
            address88 = p._read_reference_expression()
            if address88 == nil {
                p.offset = index62
                address88 = p._read_terminal_node()
                if address88 == nil {
                    p.offset = index62
                }
            }
        }
    }
    cache16[index61] = cacheEntry{node: address88, offset: p.offset}
    return address88
}

func (p *PegGoParser) _read_terminal_node() TreeNode {
    var address89 TreeNode
    address89 = nil
    var index63 int
    index63 = p.offset
    var cache17 map[int]cacheEntry
    cache17 = p.cache["terminal_node"]
    if cache17 == nil {
        cache17 = make(map[int]cacheEntry)
        p.cache["terminal_node"] = cache17
    }
    if entry, ok := cache17[index63]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index64 int
    index64 = p.offset
    address89 = p._read_string_expression()
    if address89 == nil {
        p.offset = index64
        address89 = p._read_ci_string_expression()
        if address89 == nil {
            p.offset = index64
            address89 = p._read_char_class_expression()
            if address89 == nil {
                p.offset = index64
                address89 = p._read_any_char_expression()
                if address89 == nil {
                    p.offset = index64
                }
            }
        }
    }
    cache17[index63] = cacheEntry{node: address89, offset: p.offset}
    return address89
}

func (p *PegGoParser) _read_predicated_atom() TreeNode {
    var address90 TreeNode
    address90 = nil
    var index65 int
    index65 = p.offset
    var cache18 map[int]cacheEntry
    cache18 = p.cache["predicated_atom"]
    if cache18 == nil {
        cache18 = make(map[int]cacheEntry)
        p.cache["predicated_atom"] = cache18
    }
    if entry, ok := cache18[index65]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index66 int
    index66 = p.offset
    var elements37 []TreeNode
    elements37 = make([]TreeNode, 2)
    var address91 TreeNode
    var index67 int
    index67 = p.offset
    var chunk14 string
    chunk14 = ""
    var max14 int
    max14 = p.offset + 1
    if max14 <= len(p.input) {
        chunk14 = string(p.input[p.offset:max14])
    }
    if chunk14 == "&" {
        address91 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address91 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::predicated_atom", expected: "\"&\""})
        }
    }
    if address91 == nil {
        p.offset = index67
        var chunk15 string
        chunk15 = ""
        var max15 int
        max15 = p.offset + 1
        if max15 <= len(p.input) {
            chunk15 = string(p.input[p.offset:max15])
        }
        if chunk15 == "!" {
            address91 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address91 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::predicated_atom", expected: "\"!\""})
            }
        }
        if address91 == nil {
            p.offset = index67
        }
    }
    if address91 != nil {
        elements37[0] = address91
        var address92 TreeNode
        address92 = p._read_atom()
        if address92 != nil {
            elements37[1] = address92
        } else {
            elements37 = nil
            p.offset = index66
        }
    } else {
        elements37 = nil
        p.offset = index66
    }
    if elements37 == nil {
        address90 = nil
    } else {
        address90 = newNode18(p.slice(index66, p.offset), index66, elements37)
        p.offset = p.offset
    }
    cache18[index65] = cacheEntry{node: address90, offset: p.offset}
    return address90
}

func (p *PegGoParser) _read_reference_expression() TreeNode {
    var address93 TreeNode
    address93 = nil
    var index68 int
    index68 = p.offset
    var cache19 map[int]cacheEntry
    cache19 = p.cache["reference_expression"]
    if cache19 == nil {
        cache19 = make(map[int]cacheEntry)
        p.cache["reference_expression"] = cache19
    }
    if entry, ok := cache19[index68]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index69 int
    index69 = p.offset
    var elements38 []TreeNode
    elements38 = make([]TreeNode, 2)
    var address94 TreeNode
    address94 = p._read_identifier()
    if address94 != nil {
        elements38[0] = address94
        var address95 TreeNode
        var index70 int
        index70 = p.offset
        address95 = p._read_assignment()
        p.offset = index70
        if address95 == nil {
            address95 = &BaseNode{text: p.slice(p.offset, p.offset), offset: p.offset, children: nil}
            p.offset = p.offset
        } else {
            address95 = nil
        }
        if address95 != nil {
            elements38[1] = address95
        } else {
            elements38 = nil
            p.offset = index69
        }
    } else {
        elements38 = nil
        p.offset = index69
    }
    if elements38 == nil {
        address93 = nil
    } else {
        address93 = newNode19(p.slice(index69, p.offset), index69, elements38)
        p.offset = p.offset
    }
    cache19[index68] = cacheEntry{node: address93, offset: p.offset}
    return address93
}

func (p *PegGoParser) _read_string_expression() TreeNode {
    var address96 TreeNode
    address96 = nil
    var index71 int
    index71 = p.offset
    var cache20 map[int]cacheEntry
    cache20 = p.cache["string_expression"]
    if cache20 == nil {
        cache20 = make(map[int]cacheEntry)
        p.cache["string_expression"] = cache20
    }
    if entry, ok := cache20[index71]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index72 int
    index72 = p.offset
    var index73 int
    index73 = p.offset
    var elements39 []TreeNode
    elements39 = make([]TreeNode, 3)
    var address97 TreeNode
    var chunk16 string
    chunk16 = ""
    var max16 int
    max16 = p.offset + 1
    if max16 <= len(p.input) {
        chunk16 = string(p.input[p.offset:max16])
    }
    if chunk16 == "\"" {
        address97 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address97 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "'\"'"})
        }
    }
    if address97 != nil {
        elements39[0] = address97
        var address98 TreeNode
        var index74 int
        index74 = p.offset
        var elements40 []TreeNode
        elements40 = nil
        var address99 TreeNode
        address99 = nil
        for {
            var index75 int
            index75 = p.offset
            var index76 int
            index76 = p.offset
            var elements41 []TreeNode
            elements41 = make([]TreeNode, 2)
            var address100 TreeNode
            var chunk17 string
            chunk17 = ""
            var max17 int
            max17 = p.offset + 1
            if max17 <= len(p.input) {
                chunk17 = string(p.input[p.offset:max17])
            }
            if chunk17 == "\\" {
                address100 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address100 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "\"\\\\\""})
                }
            }
            if address100 != nil {
                elements41[0] = address100
                var address101 TreeNode
                if p.offset < len(p.input) {
                    address101 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address101 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "<any char>"})
                    }
                }
                if address101 != nil {
                    elements41[1] = address101
                } else {
                    elements41 = nil
                    p.offset = index76
                }
            } else {
                elements41 = nil
                p.offset = index76
            }
            if elements41 == nil {
                address99 = nil
            } else {
                address99 = &BaseNode{text: p.slice(index76, p.offset), offset: index76, children: elements41}
                p.offset = p.offset
            }
            if address99 == nil {
                p.offset = index75
                var chunk18 string
                chunk18 = ""
                var max18 int
                max18 = p.offset + 1
                if max18 <= len(p.input) {
                    chunk18 = string(p.input[p.offset:max18])
                }
                if REGEX_1.MatchString(chunk18) {
                    address99 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address99 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "[^\"]"})
                    }
                }
                if address99 == nil {
                    p.offset = index75
                }
            }
            if address99 != nil {
                elements40 = append(elements40, address99)
            } else {
                break
            }
        }
        if len(elements40) >= 0 {
            address98 = &BaseNode{text: p.slice(index74, p.offset), offset: index74, children: elements40}
            p.offset = p.offset
        } else {
            address98 = nil
        }
        if address98 != nil {
            elements39[1] = address98
            var address102 TreeNode
            var chunk19 string
            chunk19 = ""
            var max19 int
            max19 = p.offset + 1
            if max19 <= len(p.input) {
                chunk19 = string(p.input[p.offset:max19])
            }
            if chunk19 == "\"" {
                address102 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address102 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "'\"'"})
                }
            }
            if address102 != nil {
                elements39[2] = address102
            } else {
                elements39 = nil
                p.offset = index73
            }
        } else {
            elements39 = nil
            p.offset = index73
        }
    } else {
        elements39 = nil
        p.offset = index73
    }
    if elements39 == nil {
        address96 = nil
    } else {
        address96 = &BaseNode{text: p.slice(index73, p.offset), offset: index73, children: elements39}
        p.offset = p.offset
    }
    if address96 == nil {
        p.offset = index72
        var index77 int
        index77 = p.offset
        var elements42 []TreeNode
        elements42 = make([]TreeNode, 3)
        var address103 TreeNode
        var chunk20 string
        chunk20 = ""
        var max20 int
        max20 = p.offset + 1
        if max20 <= len(p.input) {
            chunk20 = string(p.input[p.offset:max20])
        }
        if chunk20 == "'" {
            address103 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address103 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "\"'\""})
            }
        }
        if address103 != nil {
            elements42[0] = address103
            var address104 TreeNode
            var index78 int
            index78 = p.offset
            var elements43 []TreeNode
            elements43 = nil
            var address105 TreeNode
            address105 = nil
            for {
                var index79 int
                index79 = p.offset
                var index80 int
                index80 = p.offset
                var elements44 []TreeNode
                elements44 = make([]TreeNode, 2)
                var address106 TreeNode
                var chunk21 string
                chunk21 = ""
                var max21 int
                max21 = p.offset + 1
                if max21 <= len(p.input) {
                    chunk21 = string(p.input[p.offset:max21])
                }
                if chunk21 == "\\" {
                    address106 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address106 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "\"\\\\\""})
                    }
                }
                if address106 != nil {
                    elements44[0] = address106
                    var address107 TreeNode
                    if p.offset < len(p.input) {
                        address107 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address107 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "<any char>"})
                        }
                    }
                    if address107 != nil {
                        elements44[1] = address107
                    } else {
                        elements44 = nil
                        p.offset = index80
                    }
                } else {
                    elements44 = nil
                    p.offset = index80
                }
                if elements44 == nil {
                    address105 = nil
                } else {
                    address105 = &BaseNode{text: p.slice(index80, p.offset), offset: index80, children: elements44}
                    p.offset = p.offset
                }
                if address105 == nil {
                    p.offset = index79
                    var chunk22 string
                    chunk22 = ""
                    var max22 int
                    max22 = p.offset + 1
                    if max22 <= len(p.input) {
                        chunk22 = string(p.input[p.offset:max22])
                    }
                    if REGEX_2.MatchString(chunk22) {
                        address105 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address105 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "[^']"})
                        }
                    }
                    if address105 == nil {
                        p.offset = index79
                    }
                }
                if address105 != nil {
                    elements43 = append(elements43, address105)
                } else {
                    break
                }
            }
            if len(elements43) >= 0 {
                address104 = &BaseNode{text: p.slice(index78, p.offset), offset: index78, children: elements43}
                p.offset = p.offset
            } else {
                address104 = nil
            }
            if address104 != nil {
                elements42[1] = address104
                var address108 TreeNode
                var chunk23 string
                chunk23 = ""
                var max23 int
                max23 = p.offset + 1
                if max23 <= len(p.input) {
                    chunk23 = string(p.input[p.offset:max23])
                }
                if chunk23 == "'" {
                    address108 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address108 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::string_expression", expected: "\"'\""})
                    }
                }
                if address108 != nil {
                    elements42[2] = address108
                } else {
                    elements42 = nil
                    p.offset = index77
                }
            } else {
                elements42 = nil
                p.offset = index77
            }
        } else {
            elements42 = nil
            p.offset = index77
        }
        if elements42 == nil {
            address96 = nil
        } else {
            address96 = &BaseNode{text: p.slice(index77, p.offset), offset: index77, children: elements42}
            p.offset = p.offset
        }
        if address96 == nil {
            p.offset = index72
        }
    }
    cache20[index71] = cacheEntry{node: address96, offset: p.offset}
    return address96
}

func (p *PegGoParser) _read_ci_string_expression() TreeNode {
    var address109 TreeNode
    address109 = nil
    var index81 int
    index81 = p.offset
    var cache21 map[int]cacheEntry
    cache21 = p.cache["ci_string_expression"]
    if cache21 == nil {
        cache21 = make(map[int]cacheEntry)
        p.cache["ci_string_expression"] = cache21
    }
    if entry, ok := cache21[index81]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index82 int
    index82 = p.offset
    var elements45 []TreeNode
    elements45 = make([]TreeNode, 3)
    var address110 TreeNode
    var chunk24 string
    chunk24 = ""
    var max24 int
    max24 = p.offset + 1
    if max24 <= len(p.input) {
        chunk24 = string(p.input[p.offset:max24])
    }
    if chunk24 == "`" {
        address110 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address110 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::ci_string_expression", expected: "\"`\""})
        }
    }
    if address110 != nil {
        elements45[0] = address110
        var address111 TreeNode
        var index83 int
        index83 = p.offset
        var elements46 []TreeNode
        elements46 = nil
        var address112 TreeNode
        address112 = nil
        for {
            var index84 int
            index84 = p.offset
            var index85 int
            index85 = p.offset
            var elements47 []TreeNode
            elements47 = make([]TreeNode, 2)
            var address113 TreeNode
            var chunk25 string
            chunk25 = ""
            var max25 int
            max25 = p.offset + 1
            if max25 <= len(p.input) {
                chunk25 = string(p.input[p.offset:max25])
            }
            if chunk25 == "\\" {
                address113 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address113 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::ci_string_expression", expected: "\"\\\\\""})
                }
            }
            if address113 != nil {
                elements47[0] = address113
                var address114 TreeNode
                if p.offset < len(p.input) {
                    address114 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address114 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::ci_string_expression", expected: "<any char>"})
                    }
                }
                if address114 != nil {
                    elements47[1] = address114
                } else {
                    elements47 = nil
                    p.offset = index85
                }
            } else {
                elements47 = nil
                p.offset = index85
            }
            if elements47 == nil {
                address112 = nil
            } else {
                address112 = &BaseNode{text: p.slice(index85, p.offset), offset: index85, children: elements47}
                p.offset = p.offset
            }
            if address112 == nil {
                p.offset = index84
                var chunk26 string
                chunk26 = ""
                var max26 int
                max26 = p.offset + 1
                if max26 <= len(p.input) {
                    chunk26 = string(p.input[p.offset:max26])
                }
                if REGEX_3.MatchString(chunk26) {
                    address112 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address112 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::ci_string_expression", expected: "[^`]"})
                    }
                }
                if address112 == nil {
                    p.offset = index84
                }
            }
            if address112 != nil {
                elements46 = append(elements46, address112)
            } else {
                break
            }
        }
        if len(elements46) >= 0 {
            address111 = &BaseNode{text: p.slice(index83, p.offset), offset: index83, children: elements46}
            p.offset = p.offset
        } else {
            address111 = nil
        }
        if address111 != nil {
            elements45[1] = address111
            var address115 TreeNode
            var chunk27 string
            chunk27 = ""
            var max27 int
            max27 = p.offset + 1
            if max27 <= len(p.input) {
                chunk27 = string(p.input[p.offset:max27])
            }
            if chunk27 == "`" {
                address115 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address115 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::ci_string_expression", expected: "\"`\""})
                }
            }
            if address115 != nil {
                elements45[2] = address115
            } else {
                elements45 = nil
                p.offset = index82
            }
        } else {
            elements45 = nil
            p.offset = index82
        }
    } else {
        elements45 = nil
        p.offset = index82
    }
    if elements45 == nil {
        address109 = nil
    } else {
        address109 = &BaseNode{text: p.slice(index82, p.offset), offset: index82, children: elements45}
        p.offset = p.offset
    }
    cache21[index81] = cacheEntry{node: address109, offset: p.offset}
    return address109
}

func (p *PegGoParser) _read_any_char_expression() TreeNode {
    var address116 TreeNode
    address116 = nil
    var index86 int
    index86 = p.offset
    var cache22 map[int]cacheEntry
    cache22 = p.cache["any_char_expression"]
    if cache22 == nil {
        cache22 = make(map[int]cacheEntry)
        p.cache["any_char_expression"] = cache22
    }
    if entry, ok := cache22[index86]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var chunk28 string
    chunk28 = ""
    var max28 int
    max28 = p.offset + 1
    if max28 <= len(p.input) {
        chunk28 = string(p.input[p.offset:max28])
    }
    if chunk28 == "." {
        address116 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address116 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::any_char_expression", expected: "\".\""})
        }
    }
    cache22[index86] = cacheEntry{node: address116, offset: p.offset}
    return address116
}

func (p *PegGoParser) _read_char_class_expression() TreeNode {
    var address117 TreeNode
    address117 = nil
    var index87 int
    index87 = p.offset
    var cache23 map[int]cacheEntry
    cache23 = p.cache["char_class_expression"]
    if cache23 == nil {
        cache23 = make(map[int]cacheEntry)
        p.cache["char_class_expression"] = cache23
    }
    if entry, ok := cache23[index87]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index88 int
    index88 = p.offset
    var elements48 []TreeNode
    elements48 = make([]TreeNode, 4)
    var address118 TreeNode
    var chunk29 string
    chunk29 = ""
    var max29 int
    max29 = p.offset + 1
    if max29 <= len(p.input) {
        chunk29 = string(p.input[p.offset:max29])
    }
    if chunk29 == "[" {
        address118 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address118 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::char_class_expression", expected: "\"[\""})
        }
    }
    if address118 != nil {
        elements48[0] = address118
        var address119 TreeNode
        var index89 int
        index89 = p.offset
        var chunk30 string
        chunk30 = ""
        var max30 int
        max30 = p.offset + 1
        if max30 <= len(p.input) {
            chunk30 = string(p.input[p.offset:max30])
        }
        if chunk30 == "^" {
            address119 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address119 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::char_class_expression", expected: "\"^\""})
            }
        }
        if address119 == nil {
            address119 = &BaseNode{text: p.slice(index89, index89), offset: index89, children: nil}
            p.offset = index89
        }
        if address119 != nil {
            elements48[1] = address119
            var address120 TreeNode
            var index90 int
            index90 = p.offset
            var elements49 []TreeNode
            elements49 = nil
            var address121 TreeNode
            address121 = nil
            for {
                var index91 int
                index91 = p.offset
                var index92 int
                index92 = p.offset
                var elements50 []TreeNode
                elements50 = make([]TreeNode, 2)
                var address122 TreeNode
                var chunk31 string
                chunk31 = ""
                var max31 int
                max31 = p.offset + 1
                if max31 <= len(p.input) {
                    chunk31 = string(p.input[p.offset:max31])
                }
                if chunk31 == "\\" {
                    address122 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address122 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::char_class_expression", expected: "\"\\\\\""})
                    }
                }
                if address122 != nil {
                    elements50[0] = address122
                    var address123 TreeNode
                    if p.offset < len(p.input) {
                        address123 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address123 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::char_class_expression", expected: "<any char>"})
                        }
                    }
                    if address123 != nil {
                        elements50[1] = address123
                    } else {
                        elements50 = nil
                        p.offset = index92
                    }
                } else {
                    elements50 = nil
                    p.offset = index92
                }
                if elements50 == nil {
                    address121 = nil
                } else {
                    address121 = &BaseNode{text: p.slice(index92, p.offset), offset: index92, children: elements50}
                    p.offset = p.offset
                }
                if address121 == nil {
                    p.offset = index91
                    var chunk32 string
                    chunk32 = ""
                    var max32 int
                    max32 = p.offset + 1
                    if max32 <= len(p.input) {
                        chunk32 = string(p.input[p.offset:max32])
                    }
                    if REGEX_4.MatchString(chunk32) {
                        address121 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                        p.offset = p.offset + 1
                    } else {
                        address121 = nil
                        if p.offset > p.failure.offset {
                            p.failure.offset = p.offset
                            p.failure.expected = nil
                        }
                        if p.offset == p.failure.offset {
                            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::char_class_expression", expected: "[^\\]]"})
                        }
                    }
                    if address121 == nil {
                        p.offset = index91
                    }
                }
                if address121 != nil {
                    elements49 = append(elements49, address121)
                } else {
                    break
                }
            }
            if len(elements49) >= 1 {
                address120 = &BaseNode{text: p.slice(index90, p.offset), offset: index90, children: elements49}
                p.offset = p.offset
            } else {
                address120 = nil
            }
            if address120 != nil {
                elements48[2] = address120
                var address124 TreeNode
                var chunk33 string
                chunk33 = ""
                var max33 int
                max33 = p.offset + 1
                if max33 <= len(p.input) {
                    chunk33 = string(p.input[p.offset:max33])
                }
                if chunk33 == "]" {
                    address124 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                    p.offset = p.offset + 1
                } else {
                    address124 = nil
                    if p.offset > p.failure.offset {
                        p.failure.offset = p.offset
                        p.failure.expected = nil
                    }
                    if p.offset == p.failure.offset {
                        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::char_class_expression", expected: "\"]\""})
                    }
                }
                if address124 != nil {
                    elements48[3] = address124
                } else {
                    elements48 = nil
                    p.offset = index88
                }
            } else {
                elements48 = nil
                p.offset = index88
            }
        } else {
            elements48 = nil
            p.offset = index88
        }
    } else {
        elements48 = nil
        p.offset = index88
    }
    if elements48 == nil {
        address117 = nil
    } else {
        address117 = &BaseNode{text: p.slice(index88, p.offset), offset: index88, children: elements48}
        p.offset = p.offset
    }
    cache23[index87] = cacheEntry{node: address117, offset: p.offset}
    return address117
}

func (p *PegGoParser) _read_label() TreeNode {
    var address125 TreeNode
    address125 = nil
    var index93 int
    index93 = p.offset
    var cache24 map[int]cacheEntry
    cache24 = p.cache["label"]
    if cache24 == nil {
        cache24 = make(map[int]cacheEntry)
        p.cache["label"] = cache24
    }
    if entry, ok := cache24[index93]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index94 int
    index94 = p.offset
    var elements51 []TreeNode
    elements51 = make([]TreeNode, 2)
    var address126 TreeNode
    address126 = p._read_identifier()
    if address126 != nil {
        elements51[0] = address126
        var address127 TreeNode
        var chunk34 string
        chunk34 = ""
        var max34 int
        max34 = p.offset + 1
        if max34 <= len(p.input) {
            chunk34 = string(p.input[p.offset:max34])
        }
        if chunk34 == ":" {
            address127 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
            p.offset = p.offset + 1
        } else {
            address127 = nil
            if p.offset > p.failure.offset {
                p.failure.offset = p.offset
                p.failure.expected = nil
            }
            if p.offset == p.failure.offset {
                p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::label", expected: "\":\""})
            }
        }
        if address127 != nil {
            elements51[1] = address127
        } else {
            elements51 = nil
            p.offset = index94
        }
    } else {
        elements51 = nil
        p.offset = index94
    }
    if elements51 == nil {
        address125 = nil
    } else {
        address125 = newNode20(p.slice(index94, p.offset), index94, elements51)
        p.offset = p.offset
    }
    cache24[index93] = cacheEntry{node: address125, offset: p.offset}
    return address125
}

func (p *PegGoParser) _read_object_identifier() TreeNode {
    var address128 TreeNode
    address128 = nil
    var index95 int
    index95 = p.offset
    var cache25 map[int]cacheEntry
    cache25 = p.cache["object_identifier"]
    if cache25 == nil {
        cache25 = make(map[int]cacheEntry)
        p.cache["object_identifier"] = cache25
    }
    if entry, ok := cache25[index95]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index96 int
    index96 = p.offset
    var elements52 []TreeNode
    elements52 = make([]TreeNode, 2)
    var address129 TreeNode
    address129 = p._read_identifier()
    if address129 != nil {
        elements52[0] = address129
        var address130 TreeNode
        var index97 int
        index97 = p.offset
        var elements53 []TreeNode
        elements53 = nil
        var address131 TreeNode
        address131 = nil
        for {
            var index98 int
            index98 = p.offset
            var elements54 []TreeNode
            elements54 = make([]TreeNode, 2)
            var address132 TreeNode
            var chunk35 string
            chunk35 = ""
            var max35 int
            max35 = p.offset + 1
            if max35 <= len(p.input) {
                chunk35 = string(p.input[p.offset:max35])
            }
            if chunk35 == "." {
                address132 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address132 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::object_identifier", expected: "\".\""})
                }
            }
            if address132 != nil {
                elements54[0] = address132
                var address133 TreeNode
                address133 = p._read_identifier()
                if address133 != nil {
                    elements54[1] = address133
                } else {
                    elements54 = nil
                    p.offset = index98
                }
            } else {
                elements54 = nil
                p.offset = index98
            }
            if elements54 == nil {
                address131 = nil
            } else {
                address131 = newNode22(p.slice(index98, p.offset), index98, elements54)
                p.offset = p.offset
            }
            if address131 != nil {
                elements53 = append(elements53, address131)
            } else {
                break
            }
        }
        if len(elements53) >= 0 {
            address130 = &BaseNode{text: p.slice(index97, p.offset), offset: index97, children: elements53}
            p.offset = p.offset
        } else {
            address130 = nil
        }
        if address130 != nil {
            elements52[1] = address130
        } else {
            elements52 = nil
            p.offset = index96
        }
    } else {
        elements52 = nil
        p.offset = index96
    }
    if elements52 == nil {
        address128 = nil
    } else {
        address128 = newNode21(p.slice(index96, p.offset), index96, elements52)
        p.offset = p.offset
    }
    cache25[index95] = cacheEntry{node: address128, offset: p.offset}
    return address128
}

func (p *PegGoParser) _read_identifier() TreeNode {
    var address134 TreeNode
    address134 = nil
    var index99 int
    index99 = p.offset
    var cache26 map[int]cacheEntry
    cache26 = p.cache["identifier"]
    if cache26 == nil {
        cache26 = make(map[int]cacheEntry)
        p.cache["identifier"] = cache26
    }
    if entry, ok := cache26[index99]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index100 int
    index100 = p.offset
    var elements55 []TreeNode
    elements55 = make([]TreeNode, 2)
    var address135 TreeNode
    var chunk36 string
    chunk36 = ""
    var max36 int
    max36 = p.offset + 1
    if max36 <= len(p.input) {
        chunk36 = string(p.input[p.offset:max36])
    }
    if REGEX_5.MatchString(chunk36) {
        address135 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address135 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::identifier", expected: "[a-zA-Z_]"})
        }
    }
    if address135 != nil {
        elements55[0] = address135
        var address136 TreeNode
        var index101 int
        index101 = p.offset
        var elements56 []TreeNode
        elements56 = nil
        var address137 TreeNode
        address137 = nil
        for {
            var chunk37 string
            chunk37 = ""
            var max37 int
            max37 = p.offset + 1
            if max37 <= len(p.input) {
                chunk37 = string(p.input[p.offset:max37])
            }
            if REGEX_6.MatchString(chunk37) {
                address137 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address137 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::identifier", expected: "[a-zA-Z0-9_]"})
                }
            }
            if address137 != nil {
                elements56 = append(elements56, address137)
            } else {
                break
            }
        }
        if len(elements56) >= 0 {
            address136 = &BaseNode{text: p.slice(index101, p.offset), offset: index101, children: elements56}
            p.offset = p.offset
        } else {
            address136 = nil
        }
        if address136 != nil {
            elements55[1] = address136
        } else {
            elements55 = nil
            p.offset = index100
        }
    } else {
        elements55 = nil
        p.offset = index100
    }
    if elements55 == nil {
        address134 = nil
    } else {
        address134 = &BaseNode{text: p.slice(index100, p.offset), offset: index100, children: elements55}
        p.offset = p.offset
    }
    cache26[index99] = cacheEntry{node: address134, offset: p.offset}
    return address134
}

func (p *PegGoParser) _read___() TreeNode {
    var address138 TreeNode
    address138 = nil
    var index102 int
    index102 = p.offset
    var cache27 map[int]cacheEntry
    cache27 = p.cache["__"]
    if cache27 == nil {
        cache27 = make(map[int]cacheEntry)
        p.cache["__"] = cache27
    }
    if entry, ok := cache27[index102]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index103 int
    index103 = p.offset
    var chunk38 string
    chunk38 = ""
    var max38 int
    max38 = p.offset + 1
    if max38 <= len(p.input) {
        chunk38 = string(p.input[p.offset:max38])
    }
    if REGEX_7.MatchString(chunk38) {
        address138 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address138 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::__", expected: "[\\s]"})
        }
    }
    if address138 == nil {
        p.offset = index103
        address138 = p._read_comment()
        if address138 == nil {
            p.offset = index103
        }
    }
    cache27[index102] = cacheEntry{node: address138, offset: p.offset}
    return address138
}

func (p *PegGoParser) _read_comment() TreeNode {
    var address139 TreeNode
    address139 = nil
    var index104 int
    index104 = p.offset
    var cache28 map[int]cacheEntry
    cache28 = p.cache["comment"]
    if cache28 == nil {
        cache28 = make(map[int]cacheEntry)
        p.cache["comment"] = cache28
    }
    if entry, ok := cache28[index104]; ok {
        p.offset = entry.offset
        return entry.node
    }
    var index105 int
    index105 = p.offset
    var elements57 []TreeNode
    elements57 = make([]TreeNode, 2)
    var address140 TreeNode
    var chunk39 string
    chunk39 = ""
    var max39 int
    max39 = p.offset + 1
    if max39 <= len(p.input) {
        chunk39 = string(p.input[p.offset:max39])
    }
    if chunk39 == "#" {
        address140 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
        p.offset = p.offset + 1
    } else {
        address140 = nil
        if p.offset > p.failure.offset {
            p.failure.offset = p.offset
            p.failure.expected = nil
        }
        if p.offset == p.failure.offset {
            p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::comment", expected: "\"#\""})
        }
    }
    if address140 != nil {
        elements57[0] = address140
        var address141 TreeNode
        var index106 int
        index106 = p.offset
        var elements58 []TreeNode
        elements58 = nil
        var address142 TreeNode
        address142 = nil
        for {
            var chunk40 string
            chunk40 = ""
            var max40 int
            max40 = p.offset + 1
            if max40 <= len(p.input) {
                chunk40 = string(p.input[p.offset:max40])
            }
            if REGEX_8.MatchString(chunk40) {
                address142 = &BaseNode{text: p.slice(p.offset, p.offset + 1), offset: p.offset, children: nil}
                p.offset = p.offset + 1
            } else {
                address142 = nil
                if p.offset > p.failure.offset {
                    p.failure.offset = p.offset
                    p.failure.expected = nil
                }
                if p.offset == p.failure.offset {
                    p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG::comment", expected: "[^\\n]"})
                }
            }
            if address142 != nil {
                elements58 = append(elements58, address142)
            } else {
                break
            }
        }
        if len(elements58) >= 0 {
            address141 = &BaseNode{text: p.slice(index106, p.offset), offset: index106, children: elements58}
            p.offset = p.offset
        } else {
            address141 = nil
        }
        if address141 != nil {
            elements57[1] = address141
        } else {
            elements57 = nil
            p.offset = index105
        }
    } else {
        elements57 = nil
        p.offset = index105
    }
    if elements57 == nil {
        address139 = nil
    } else {
        address139 = &BaseNode{text: p.slice(index105, p.offset), offset: index105, children: elements57}
        p.offset = p.offset
    }
    cache28[index104] = cacheEntry{node: address139, offset: p.offset}
    return address139
}

func New(input string, actions Actions) *PegGoParser {
    return &PegGoParser{
        input: []rune(input),
        inputString: input,
        actions: actions,
        cache: make(map[string]map[int]cacheEntry),
    }
}

func (p *PegGoParser) WithTypes(types map[string]NodeExtender) *PegGoParser {
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

func (p *PegGoParser) Parse() (TreeNode, error) {
    node := p._read_grammar()
    if p.actionErr != nil {
        return nil, p.actionErr
    }
    if node != nil && p.offset == len(p.input) {
        return node, nil
    }
    if len(p.failure.expected) == 0 {
        p.failure.offset = p.offset
        p.failure.expected = append(p.failure.expected, expectation{rule: "Canopy.PEG", expected: "<EOF>"})
    }
    return nil, p.newParseError()
}

func (p *PegGoParser) newParseError() error {
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

func (p *PegGoParser) slice(start, end int) string {
    if start < 0 { start = 0 }
    if end > len(p.input) { end = len(p.input) }
    if start > end { start = end }
    return string(p.input[start:end])
}

func (p *PegGoParser) extendNode(node TreeNode, name string) TreeNode {
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

