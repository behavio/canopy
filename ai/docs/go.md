# Canopy Parser: Go Implementation Proposal

## 1. Overview

This document outlines a proposed design for a Canopy parser generator for the Go programming language. The design adheres to the core principles of Canopy while embracing idiomatic Go practices for clarity, performance, and maintainability.

The generated parser will be a self-contained Go package, allowing for easy integration into any Go project. The generation process will be driven by a new `src/builders/go.js` module, which will use Go-specific templates located in `templates/go/`.

**Reasoning**:
- **Self-contained Package**: Standard Go practice encourages modularity through packages. This makes the parser reusable and simple to import.
- **Builder and Templates**: Following the existing Canopy architecture (`src/builders/language.js` and `templates/language/`) provides consistency and leverages the current infrastructure for code generation.

## 2. Core Components

The generated Go package will consist of the following core components, implemented as structs and functions.

### The `Parser`

This will be the main public-facing struct that users interact with.

```go
type Parser struct {
    text    []rune
    actions Actions
    cache   map[cacheKey]*cacheEntry
    err     *ParseError
}

func NewParser(input string, actions Actions) *Parser {
    // ...
}

func (p *Parser) Parse() (*TreeNode, error) {
    // ...
}
```

**Reasoning**:
- **Structs over Classes**: Go uses structs to represent data aggregates. The `Parser` struct will hold the state (input text, cache) required for parsing.
- **`[]rune` for Text**: Using a slice of runes correctly handles Unicode characters, which is crucial for a robust parser.
- **Explicit Initialization**: A `NewParser` constructor function is the idiomatic way to initialize a new struct instance, ensuring all fields are set correctly.

### The Grammar Logic

Instead of a separate `Grammar` component, the `_read_<rule_name>` methods will be unexported methods on the `Parser` struct itself.

```go
func (p *Parser) _read_root() *TreeNode {
    // ...
}

func (p *Parser) _read_anotherRule() *TreeNode {
    // ...
}
```

**Reasoning**:
- **Composition and Methods**: This approach favors composition over inheritance. The parsing logic is tightly coupled with the parser's state (cache, input text), so making them methods of `Parser` is a natural fit in Go. Unexported methods (starting with a lowercase letter) encapsulate the internal parsing logic, presenting a clean public API.

### The `TreeNode`

This struct represents a node in the generated parse tree.

```go
type TreeNode struct {
    Text     string
    Offset   int
    Children []*TreeNode
    // Rule-specific fields can be added here or handled via a more generic map
}
```

For named rules, specific node types can be generated using struct embedding to extend the base `TreeNode`.

```go
type ExpressionNode struct {
    TreeNode
    // Additional fields specific to 'Expression'
}
```

**Reasoning**:
- **Struct for Data**: `TreeNode` is a plain data structure, making a struct the perfect choice.
- **Slice for Children**: A slice (`[]*TreeNode`) is the natural and efficient way to handle a collection of child nodes in Go.
- **Struct Embedding for Specialization**: Go's struct embedding provides a form of composition that mimics inheritance, allowing for specialized node types that share the base `TreeNode`'s properties without the complexity of a class hierarchy.

### The `ParseError`

A custom error type that implements Go's built-in `error` interface.

```go
type ParseError struct {
    Line, Column int
    Message      string
}

func (e *ParseError) Error() string {
    return fmt.Sprintf("parse error at %d:%d: %s", e.Line, e.Column, e.Message)
}
```

**Reasoning**:
- **Idiomatic Error Handling**: Go's standard is to return `error` values, not to throw exceptions. Creating a custom `ParseError` struct that implements the `error` interface allows the parser to return rich error information that can be programmatically inspected by the caller, which is a hallmark of robust Go applications.

## 3. Object Interaction and Parsing Flow

The parsing flow will follow Go's idiomatic error handling and control flow patterns.

1.  A `Parser` instance is created using `NewParser(input, actions)`.
2.  The user calls the `Parse()` method, which returns `(*TreeNode, error)`.
3.  `Parse()` invokes the method for the grammar's root rule (e.g., `p._read_root()`).
4.  Each `_read_<rule_name>()` method attempts to match a part of the input.
5.  **Memoization**: The `cache` map within the `Parser` struct is checked before attempting a parse. If a result is found, it's returned immediately.
6.  If a rule matches successfully, a `*TreeNode` is created, the result is stored in the cache, and the node is returned.
7.  If a rule fails to match, it returns `nil`. The failure is also cached to prevent re-parsing. The calling function then handles the `nil` return, either by backtracking or propagating the failure.
8.  If the root rule returns a `*TreeNode` and the entire input is consumed, `Parse()` returns the node and a `nil` error.
9.  If parsing fails, `Parse()` returns `nil` for the node and a populated `*ParseError`.

**Reasoning**:
- **Explicit Error Returns**: The `(value, error)` return pattern is the most fundamental aspect of error handling in Go. It makes control flow explicit and forces the caller to handle errors, leading to more reliable code.
- **`nil` for Failure**: Using `nil` as the return value for a failed parse within the recursive descent is a simple and efficient way to signal failure without the overhead of creating error objects at every step. A formal `ParseError` is only constructed at the end if the entire parse fails.

## 4. Canopy Actions

Actions will be handled via a Go interface.

```go
type Actions interface {
    // Method for each action in the grammar
    MakeInteger(input string, start, end int, elements []*TreeNode) (interface{}, error)
}
```

The user provides a struct that implements this interface. The `Parser` calls the methods on the provided implementation.

**Reasoning**:
- **Interfaces for Behavior**: Interfaces are Go's way of specifying behavior. Defining an `Actions` interface decouples the parser from any specific implementation of the actions. This allows users to provide their own logic while ensuring a type-safe contract with the parser. The `(interface{}, error)` return signature allows actions to return custom types and signal failures gracefully.

## 5. Code Generation and Templates

The `src/builders/go.js` builder will generate the Go source files from templates.

- It will create a single directory for the new package (e.g., `my_grammar/`).
- It will generate `parser.go`, `tree_node.go`, and `actions.go` (for the interface).
- The builder will translate the Canopy AST into Go method calls, `if/else` blocks, and struct instantiations.
- The `_template` function will use Go's `text/template` or `html/template` conventions for placeholders.

**Reasoning**:
- **Standard Package Structure**: Generating code into a dedicated directory is the standard for Go packages. Separating structs and interfaces into their own files (`tree_node.go`, `actions.go`) is a common convention that improves code organization and readability.
- **Leveraging Go Tooling**: The generated code will be formatted with `gofmt` automatically as a final step of the build process, ensuring it conforms to Go's style guidelines.
