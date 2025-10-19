# Canopy Parser: Go Implementation

## 1. Overview

This document describes the implemented Go code generator for Canopy parsers. The design follows idiomatic Go practices while maintaining compatibility with Canopy's core principles of packrat parsing with memoization.

The generated parser is a self-contained Go module that can be imported into any Go project. Generation is driven by the `src/builders/go.js` module, which uses Go-specific templates in `templates/go/`.

**Key Features**:

- **Self-contained Go Module**: Each grammar generates a complete Go module with `go.mod`, making it easy to import and use.
- **Idiomatic Go**: Follows Go naming conventions, error handling patterns, and package structure.
- **Type-Safe Parse Trees**: Uses interfaces and struct embedding for strongly-typed, extensible parse trees.
- **Zero External Dependencies**: Generated parsers depend only on Go's standard library (`fmt`, `regexp`, `strings`).

## 2. Core Components

The generated Go module consists of three files and follows Go's package naming conventions.

### Package and File Structure

For a grammar file `json.peg`, Canopy generates a module named `jsongoparser` with the following structure:

```
json-go/
├── go.mod                    # Module definition
├── parser.go                 # Main parser struct and logic
├── treenode.go              # TreeNode interface and BaseNode
└── actions.go               # Actions interface definition
```

**Naming Conventions**:

- **Package Name**: Grammar filename + "goparser" (e.g., `jsongoparser`, `lispgoparser`)
- **Main Struct**: PascalCase filename + "Parser" (e.g., `JsonGoParser`, `LispGoParser`)
- **Constructor**: `New(input string, actions Actions) *JsonGoParser`
- **Parse Methods**: Private methods with `_read_` prefix (e.g., `_read_document()`)

### The Parser Struct

The main parser struct holds all parsing state:

```go
type JsonGoParser struct {
    input       []rune                       // Input as runes for Unicode support
    inputString string                       // Original input for action callbacks
    actions     Actions                      // User-provided semantic actions
    types       map[string]NodeExtender      // Type extensions (optional)
    offset      int                          // Current parsing position
    cache       map[string]map[int]cacheEntry // Memoization cache
    failure     failureState                 // Tracks parse failures for errors
    actionErr   error                        // Captures errors from action callbacks
}

func New(input string, actions Actions) *JsonGoParser {
    return &JsonGoParser{
        input:       []rune(input),
        inputString: input,
        actions:     actions,
        cache:       make(map[string]map[int]cacheEntry),
    }
}

func (p *JsonGoParser) Parse() (TreeNode, error) {
    node := p._read_document()  // Calls root rule
    if p.actionErr != nil {
        return nil, p.actionErr
    }
    if node != nil && p.offset == len(p.input) {
        return node, nil
    }
    return nil, p.newParseError()
}
```

**Design Decisions**:

- **`[]rune` for Unicode**: Go's `string` is a sequence of bytes; `[]rune` properly handles multi-byte Unicode characters.
- **Nested Cache Map**: Cache is `map[string]map[int]cacheEntry` where the outer key is the rule name and inner key is the offset.
- **Error Handling**: Actions return `(TreeNode, error)`, allowing them to fail gracefully. Parse errors are accumulated and formatted with line/column information.

### Grammar Rules as Methods

Each grammar rule becomes an unexported method on the parser struct:

```go
func (p *JsonGoParser) _read_document() TreeNode {
    var address0 TreeNode = nil
    var index0 int = p.offset
    var cache0 map[int]cacheEntry = p.cache["document"]

    // Check cache
    if cache0 == nil {
        cache0 = make(map[int]cacheEntry)
        p.cache["document"] = cache0
    }
    if entry, ok := cache0[index0]; ok {
        p.offset = entry.offset
        return entry.node
    }

    // Parsing logic...

    // Cache result
    cache0[index0] = cacheEntry{node: address0, offset: p.offset}
    return address0
}
```

**Design Decisions**:

- **`_read_` Prefix**: Maintains consistency with other Canopy backends (JavaScript, Python, Ruby, Java).
- **Unexported Methods**: The underscore prefix naturally makes these methods unexported in Go, hiding implementation details.
- **Explicit Memoization**: Each rule method includes cache lookup and storage logic for packrat parsing.

### The TreeNode Interface

The `TreeNode` interface defines the contract for all parse tree nodes:

```go
// TreeNode represents a node in the parse tree.
// Generated nodes satisfy this interface to allow custom action results.
type TreeNode interface {
    Text() string
    Offset() int
    Children() []TreeNode
}

// BaseNode is embedded by generated nodes to implement TreeNode.
type BaseNode struct {
    text     string
    offset   int
    children []TreeNode
}

func (n *BaseNode) Text() string       { return n.text }
func (n *BaseNode) Offset() int        { return n.offset }
func (n *BaseNode) Children() []TreeNode { return n.children }
```

**Generated Sequence Nodes**

For sequence expressions with labeled elements, the generator creates typed structs:

```go
// Generated for a sequence with labels
type Node5 struct {
    BaseNode
    String TreeNode   // Labeled element from sequence
    Value  TreeNode   // Another labeled element
}

func newNode5(text string, start int, elements []TreeNode) TreeNode {
    node := &Node5{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.String = elements[1]  // Extract labeled elements
    node.Value = elements[4]
    return node
}
```

**Custom Action Nodes**

Actions return custom types implementing `TreeNode`:

```go
// User-defined node type for semantic values
type IntegerNode struct {
    BaseNode
    Value int64
}

// Action implementation
func (a *MyActions) MakeInteger(input string, start, end int, elements []TreeNode) (TreeNode, error) {
    text := input[start:end]
    value, err := strconv.ParseInt(text, 10, 64)
    if err != nil {
        return nil, fmt.Errorf("invalid integer: %w", err)
    }
    return &IntegerNode{
        BaseNode: BaseNode{
            text:   text,
            offset: start,
        },
        Value: value,
    }, nil
}
```

**Design Benefits**:

- **Interface Polymorphism**: `TreeNode` allows mixing generated nodes and custom action nodes in the same tree.
- **Struct Embedding**: Embedding `BaseNode` eliminates boilerplate while maintaining type safety.
- **Compile-Time Safety**: The compiler enforces that action return values implement `TreeNode`.
- **Type Assertions**: Callers can use type assertions to access specific node types and their fields.

### Parse Errors

The `ParseError` struct provides detailed failure information:

```go
type ParseError struct {
    Input    string
    Offset   int
    Line     int
    Column   int
    Expected []expectation
    Message  string
}

type expectation struct {
    rule     string
    expected string
}

func (e *ParseError) Error() string {
    return e.Message
}
```

**Error Message Example**:

```
parse error at line 3, column 15: expected "}" from json::object or "," from json::object
```

**Error Handling Pattern**:

```go
tree, err := jsongoparser.Parse(input, nil, nil)
if err != nil {
    var parseErr *jsongoparser.ParseError
    if errors.As(err, &parseErr) {
        fmt.Printf("Parse failed at line %d, column %d\n", parseErr.Line, parseErr.Column)
        fmt.Printf("Expected: %v\n", parseErr.Expected)
    }
    return err
}
```

**Design Decisions**:

- **Idiomatic Error Interface**: Implements `error` interface for seamless integration with Go's error handling.
- **Rich Error Context**: Includes line, column, offset, and all expected tokens at the failure point.
- **Type Assertion Support**: Use `errors.As()` to access detailed `ParseError` fields programmatically.

## 3. Parsing Flow and Memoization

The parsing flow follows Go's idiomatic patterns for error handling and resource management.

### Typical Usage

```go
package main

import (
    "fmt"
    "log"

    "jsongoparser"
)

func main() {
    input := `{"name": "John", "age": 30}`

    // Parse without actions (structural parse tree only)
    tree, err := jsongoparser.Parse(input, nil, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Parsed: %s\n", tree.Text())
    fmt.Printf("Children: %d\n", len(tree.Children()))
}
```

### Parsing Steps

1. **Initialization**: `Parse()` creates a parser with `New(input, actions)`
2. **Root Rule**: Calls the grammar's root rule method (e.g., `_read_document()`)
3. **Memoization Check**: Each rule method first checks the cache
4. **Recursive Descent**: Rules call other rule methods, building the parse tree
5. **Success Path**: If root rule succeeds and all input consumed, return `(tree, nil)`
6. **Failure Path**: If parsing fails, construct and return `(*ParseError, nil)`

### Memoization (Packrat Parsing)

Every rule method implements packrat parsing with memoization:

```go
func (p *JsonGoParser) _read_value() TreeNode {
    var address0 TreeNode = nil
    var index0 int = p.offset
    var cache0 map[int]cacheEntry = p.cache["value"]

    // Initialize cache for this rule if needed
    if cache0 == nil {
        cache0 = make(map[int]cacheEntry)
        p.cache["value"] = cache0
    }

    // Check if we've already parsed this rule at this position
    if entry, ok := cache0[index0]; ok {
        p.offset = entry.offset
        return entry.node
    }

    // ... parsing logic ...

    // Cache the result (success or failure)
    cache0[index0] = cacheEntry{node: address0, offset: p.offset}
    return address0
}
```

**Memoization Benefits**:

- **O(n) Guarantee**: Each (rule, position) pair is parsed at most once
- **Handles Left Recursion**: Cache prevents infinite recursion in left-recursive grammars
- **Performance**: Dramatically improves parsing of ambiguous grammars

### Control Flow

- **Success**: Methods return a `TreeNode` (non-nil) and advance `p.offset`
- **Failure**: Methods return `nil` and record failure in `p.failure`
- **Backtracking**: On failure, parser resets offset and tries alternatives
- **Error Accumulation**: All expected tokens at the furthest offset are collected for error messages

## 4. Actions Interface

Actions provide semantic transformations during parsing. The generated `Actions` interface includes a method for each action defined in the grammar.

### Generated Actions Interface

For a grammar with actions:

```peg
grammar MyGrammar
  number <- [0-9]+ <makeNumber>
  string <- '"' chars:(!'"' .)* '"' <makeString>
end
```

Canopy generates:

```go
// Actions defines the semantic callbacks for nodes with actions.
type Actions interface {
    MakeNumber(input string, start, end int, elements []TreeNode) (TreeNode, error)
    MakeString(input string, start, end int, elements []TreeNode) (TreeNode, error)
}
```

### Implementing Actions

Users implement the interface with custom logic:

```go
type MyActions struct{}

func (a *MyActions) MakeNumber(input string, start, end int, elements []TreeNode) (TreeNode, error) {
    text := input[start:end]
    value, err := strconv.Atoi(text)
    if err != nil {
        return nil, fmt.Errorf("invalid number %q: %w", text, err)
    }

    return &NumberNode{
        BaseNode: BaseNode{text: text, offset: start},
        Value:    value,
    }, nil
}

func (a *MyActions) MakeString(input string, start, end int, elements []TreeNode) (TreeNode, error) {
    // Extract characters from labeled 'chars' element
    var chars strings.Builder
    for _, child := range elements {
        if child != nil {
            chars.WriteString(child.Text())
        }
    }

    return &StringNode{
        BaseNode: BaseNode{text: input[start:end], offset: start},
        Value:    chars.String(),
    }, nil
}
```

### Using Actions

```go
actions := &MyActions{}
tree, err := mygrammargoparser.Parse(input, actions, nil)
if err != nil {
    return err
}

// Type assert to access semantic values
if numNode, ok := tree.(*NumberNode); ok {
    fmt.Printf("Parsed number: %d\n", numNode.Value)
}
```

### Action Error Handling

Actions can return errors to abort parsing:

```go
func (a *MyActions) MakeDate(input string, start, end int, elements []TreeNode) (TreeNode, error) {
    dateStr := input[start:end]
    parsed, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return nil, fmt.Errorf("invalid date format: %w", err)
    }
    return &DateNode{Value: parsed}, nil
}
```

If an action returns an error, parsing stops immediately and `Parse()` returns that error.

**Design Principles**:

- **Callback Interface**: Actions are callbacks invoked during parsing, not post-processing
- **Error Propagation**: Action errors propagate immediately, stopping the parse
- **Type Safety**: Interface contract ensures actions return valid `TreeNode` implementations
- **Immutability**: `elements` slice should be copied if stored, as it may be reused internally

## 5. Code Generation

The `src/builders/go.js` module generates Go code by translating Canopy's AST into Go syntax.

### Builder Architecture

The builder extends `src/builders/base.js` and implements language-specific methods:

```javascript
class Builder extends Base {
  package_(name, actions, block) {
    this._grammarName = name;
    this._packageName = toPackageName(this._baseName);
    this._structName = toPascalCase(this._baseName) + 'Parser';
    // Generate parser.go, treenode.go, actions.go, go.mod
  }

  method_(name, args, block) {
    // Generate parser methods like _read_document()
  }

  syntaxNode_(address, start, end, elements, action, nodeClass) {
    // Handle node creation with optional actions
  }
}
```

### Template System

The builder uses Handlebars templates in `templates/go/`:

**`treenode.go.tpl`**:

```go
package {{name}}

type TreeNode interface {
    Text() string
    Offset() int
    Children() []TreeNode
}

type BaseNode struct {
    text     string
    offset   int
    children []TreeNode
}
// ... methods ...
```

**`actions.go.tpl`**:

```go
package {{name}}

type Actions interface {
{{#each actions}}
    {{this}}(input string, start, end int, elements []TreeNode) (TreeNode, error)
{{/each}}
}
```

### Go-Specific Translation

The builder maps Canopy constructs to Go:

| Canopy Construct | Go Output                                      |
| ---------------- | ---------------------------------------------- |
| String literal   | `if chunk0 == "{"` with failure handling       |
| Character class  | `REGEX_1.MatchString(chunk0)`                  |
| Sequence         | `elements0 := make([]TreeNode, 3)`             |
| Choice           | `if address1 == nil { ... }` with backtracking |
| Repetition       | `for { ... break }`                            |
| Optional         | `if address1 != nil { ... }`                   |
| Label            | Generated struct fields                        |
| Action           | `p.actions.MakeInteger(...)` call              |

### Generated Node Classes

For labeled sequences, the builder generates typed structs:

```go
type Node5 struct {
    BaseNode
    String TreeNode
    Value  TreeNode
}

func newNode5(text string, start int, elements []TreeNode) TreeNode {
    node := &Node5{
        BaseNode: BaseNode{text: text, offset: start, children: elements},
    }
    node.String = elements[1]
    node.Value = elements[4]
    return node
}
```

### Naming Transformations

The builder applies Go naming conventions:

- **Package Names**: `toPascalCase()` for filenames → lowercase + "parser"
- **Struct Names**: `toPascalCase()` for main parser struct
- **Action Methods**: `toPascalCase()` for action names (e.g., `make_number` → `MakeNumber`)
- **Regex Constants**: `REGEX_1`, `REGEX_2`, etc. for character classes

### Import Management

The builder tracks required imports:

```javascript
this._parserImports.add('fmt'); // Always included
this._parserImports.add('regexp'); // If character classes used
this._parserImports.add('strings'); // If case-insensitive matching used
```

Imports are injected into the generated `parser.go` file.

### Output Structure

For `json.peg`, the builder generates:

```
json-go/
├── go.mod                    # module jsongoparser; go 1.22.0
├── parser.go                 # ~1600 lines: structs, methods, helpers
├── treenode.go              # ~35 lines: TreeNode interface, BaseNode
└── actions.go               # ~8 lines: Actions interface (empty if no actions)
```

## 6. Advanced Features

### Type Extensions

The parser supports runtime type extensions via the `NodeExtender` function type:

```go
type NodeExtender func(TreeNode) TreeNode

// Define custom extensions
extensions := map[string]NodeExtender{
    "Integer": func(node TreeNode) TreeNode {
        text := node.Text()
        value, _ := strconv.Atoi(text)
        return &IntegerNode{
            BaseNode: BaseNode{text: text, offset: node.Offset()},
            Value:    value,
        }
    },
}

// Use with Parse()
tree, err := parser.Parse(input, nil, extensions)
```

Type extensions in the grammar use the `<TypeName>` syntax without the action decorator `<actionName>`.

### Convenience Parse Function

The module exports a top-level `Parse()` function for one-shot parsing:

```go
// Convenience function
func Parse(input string, actions Actions, types map[string]NodeExtender) (TreeNode, error) {
    parser := New(input, actions)
    if types != nil {
        parser.types = types
    }
    return parser.Parse()
}

// Usage
tree, err := jsongoparser.Parse(input, nil, nil)
```

### Fluent Parser API

The parser supports method chaining for configuration:

```go
parser := jsongoparser.New(input, actions)
tree, err := parser.WithTypes(extensions).Parse()
```

### Regex Pattern Handling

Character classes generate compiled regex patterns:

```go
// For character class [a-z0-9]
var REGEX_1 = regexp.MustCompile(`^[a-z0-9]`)

// Used in parsing
if REGEX_1.MatchString(chunk0) {
    address0 = &BaseNode{...}
}
```

Raw string literals (backticks) are used for patterns without backticks to avoid double-escaping.

## 7. Testing and Validation

### Test Structure

The `test/go/` directory contains comprehensive tests:

```
test/go/
├── go.mod                    # Test module
├── parse_helper.go          # Generic test helpers
├── choices_test.go          # Choice expression tests
├── sequences_test.go        # Sequence tests
├── terminals_test.go        # String/char class tests
├── quantifiers_test.go      # Repetition tests
├── predicates_test.go       # Lookahead tests
├── extensions_test.go       # Type extension tests
└── node_actions_test.go     # Action tests
```

### Test Helpers

Generic helpers use Go generics for type-safe assertions:

```go
type nodeAccessors[T any] struct {
    text     func(T) string
    offset   func(T) int
    children func(T) []T
}

func assertNodeMatches[T any](t *testing.T, accessors nodeAccessors[T],
                               expected nodeMatcher, actual T) {
    // Recursively validate node structure
}
```

### Example Test

```go
func TestChoiceStringsParsesAnyOfTheChoiceOptions(t *testing.T) {
    tree, err := choicesgoparser.Parse("foo", nil, nil)
    if err != nil {
        t.Fatalf("parse failed: %v", err)
    }

    assertNodeMatches(t, choicesNodeAccessors,
        node("foo", 0), tree.Children()[1])
}
```

### Benchmarking

The `examples/golang-example/bench_test.go` provides benchmarks:

```go
func BenchmarkCanopyParseJSON(b *testing.B) {
    benchmarkParse(b, jsonInput, func(input string) error {
        _, err := jsongoparser.Parse(input, nil, nil)
        return err
    })
}
```

**Sample Results**:

```
BenchmarkCanopyParseJSON-8    5000    250000 ns/op    125000 B/op    2500 allocs/op
BenchmarkCanopyParseLisp-8   50000     25000 ns/op     12000 B/op     250 allocs/op
BenchmarkCanopyParsePEG-8     2000    650000 ns/op    325000 B/op    6500 allocs/op
```

## 8. Comparison with Other Backends

### Naming Conventions

| Aspect         | Go              |      JavaScript |           Python |               Ruby |            Java |
| -------------- | --------------- | --------------: | ---------------: | -----------------: | --------------: |
| Package name   | `jsongoparser`  |               - |           `json` |             `JSON` |          `json` |
| Parser class   | `JsonGoParser`  |    `CanopyJson` |     `CanopyJson` | `CanopyJSONParser` |    `CanopyJson` |
| Rule methods   | `_read_value()` | `_read_value()` |  `_read_value()` |      `_read_value` | `_read_value()` |
| Action methods | `MakeInteger()` | `makeInteger()` | `make_integer()` |     `make_integer` | `makeInteger()` |

### Unique Go Features

- **No Class Inheritance**: Uses struct embedding and interfaces instead
- **Explicit Error Handling**: Methods return `(TreeNode, error)` tuples
- **Compile-Time Type Safety**: Action return types are enforced by the compiler
- **Module System**: Each grammar is a complete Go module with `go.mod`
- **Zero Dependencies**: Only uses standard library (`fmt`, `regexp`, `strings`)

## 9. Usage Examples

### Basic Parsing

```go
package main

import (
    "fmt"
    "log"
    "jsongoparser"
)

func main() {
    input := `{"name": "John", "age": 30}`
    tree, err := jsongoparser.Parse(input, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Parsed: %s\n", tree.Text())
}
```

### With Actions

```go
type MyActions struct{}

func (a *MyActions) MakeNumber(input string, start, end int,
                                elements []TreeNode) (TreeNode, error) {
    value, err := strconv.Atoi(input[start:end])
    if err != nil {
        return nil, err
    }
    return &NumberNode{Value: value}, nil
}

func main() {
    actions := &MyActions{}
    tree, err := myparser.Parse(input, actions, nil)
    // Process semantic tree...
}
```

### Error Handling

```go
tree, err := parser.Parse(input, nil, nil)
if err != nil {
    var parseErr *parser.ParseError
    if errors.As(err, &parseErr) {
        fmt.Fprintf(os.Stderr, "Parse error at line %d, column %d\n",
                    parseErr.Line, parseErr.Column)
        fmt.Fprintf(os.Stderr, "Expected: %v\n", parseErr.Expected)
        os.Exit(1)
    }
    return err
}
```

## 10. Performance Characteristics

### Memoization Overhead

- **Memory**: O(n × r) where n = input length, r = number of rules
- **Time**: Each (rule, position) pair computed at most once = O(n × r)
- **Benefit**: Essential for grammars with ambiguity or left recursion

### Allocation Patterns

- **[]rune Conversion**: One allocation for input conversion to runes
- **Node Creation**: One allocation per successful parse tree node
- **Cache Storage**: One allocation per (rule, position) pair attempted
- **Element Slices**: Pre-sized slices for sequence elements reduce allocations

### Optimization Opportunities

- **String Pooling**: Reuse common string literals
- **Object Pooling**: Pool and reuse `TreeNode` instances for hot paths
- **Lazy Cache Initialization**: Only create cache maps for rules that need them (already implemented)
- **Slice Capacity**: Pre-allocate element slices with capacity hints

## 11. Summary

The Go backend for Canopy generates production-ready, idiomatic Go parsers with:

✅ **Type Safety**: Compile-time enforcement of action return types
✅ **Zero Dependencies**: Only standard library imports
✅ **Idiomatic Code**: Follows Go conventions for naming, errors, and structure
✅ **Packrat Parsing**: O(n) parsing with complete memoization
✅ **Rich Errors**: Detailed parse errors with line/column information
✅ **Extensibility**: Actions and type extensions for semantic transformations
✅ **Testing**: Comprehensive test suite with benchmarks

The generated parsers are suitable for use in production Go applications, providing a clean API while maintaining the powerful parsing capabilities of the Canopy framework.
