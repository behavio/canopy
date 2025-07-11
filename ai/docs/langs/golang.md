# Canopy Parser: Go Implementation Proposal

## 1. Overview

This document outlines a proposed design for a Canopy parser generator for the Go programming language. The design adheres to the core principles of Canopy while embracing idiomatic Go practices for clarity, performance, and maintainability.

The generated parser will be a self-contained Go package, allowing for easy integration into any Go project. The generation process will be driven by a new `src/builders/go.js` module, which will use Go-specific templates located in `templates/go/`.

**Reasoning**:
- **Self-contained Package**: Standard Go practice encourages modularity through packages. This makes the parser reusable and simple to import.
- **Builder and Templates**: Following the existing Canopy architecture (`src/builders/language.js` and `templates/language/`) provides consistency and leverages the current infrastructure for code generation.

## 2. Core Components

The generated Go package will consist of the following core components, implemented as structs and functions. The package should be named
as jsonparser for json.peg, the main struct being JsonParser. Thus we can use the jsonparser.New() syntax recommended in https://go.dev/blog/package-names.

### The `Parser`

This will be the main public-facing struct that users interact with.

```go
type Parser struct {
    text    []rune
    actions Actions
    cache   map[cacheKey]*cacheEntry
    err     *ParseError
}

func New(input string, actions Actions) *Parser {
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

Instead of a separate `Grammar` component, the parsing rule methods will be unexported methods on the `Parser` struct itself.

```go
func (p *Parser) readRoot() *TreeNode {
    // ...
}

func (p *Parser) readAnotherRule() *TreeNode {
    // ...
}
```

**Reasoning**:
- **Composition and Idiomatic Naming**: This approach favors composition over inheritance. The parsing logic is tightly coupled with the parser's state, so making them methods of `Parser` is a natural fit. The method names (e.g., `readRoot`) use `camelCase` and start with a lowercase letter, making them **unexported** according to Go's visibility rules. While other Canopy backends use a `_read_` prefix, this design opts for idiomatic Go naming to present a clean public API and improve integration within Go projects.

### The `TreeNode` Interface and Custom Nodes

To support a strongly-typed, semantic parse tree, the generated code will define a `TreeNode` interface and a `BaseNode` struct that other nodes can embed.

```go
// TreeNode is the interface that all nodes in the parse tree must implement.
type TreeNode interface {
    GetText() string
    GetOffset() int
    GetChildren() []TreeNode
}

// BaseNode provides the core fields and methods for a tree node.
// Custom node types will embed it to satisfy the TreeNode interface.
type BaseNode struct {
    Text     string
    Offset   int
    Children []TreeNode
}

func (n *BaseNode) GetText() string       { return n.Text }
func (n *BaseNode) GetOffset() int      { return n.Offset }
func (n *BaseNode) GetChildren() []TreeNode { return n.Children }
```

An action can then return a custom struct that embeds `BaseNode` and adds its own semantic value.

```go
// Example of a custom node returned by an action.
type IntegerNode struct {
    BaseNode
    Value int64
}
```

**Reasoning**:
- **Interface for Polymorphism**: Defining `TreeNode` as an interface allows the `Children` slice to hold different concrete node types (`[]TreeNode`), enabling a polymorphic tree structure. This is a common and idiomatic Go pattern for building flexible but type-safe data structures.
- **Struct Embedding for Reusability**: Embedding `BaseNode` allows custom node types (like `IntegerNode`) to automatically satisfy the `TreeNode` interface without boilerplate code. This keeps the design DRY (Don't Repeat Yourself) and makes it easy to add new node types.
- **Type Safety**: This approach is more type-safe than using `interface{}`. When you retrieve a child from the tree, you can use a type assertion to get back the specific, rich type (e.g., `*IntegerNode`), avoiding the need to check and cast a generic `Value` field.

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

Actions are handled via a Go interface. When a rule with an action is successfully parsed, the parser calls the corresponding method on the `Actions` interface. The method's return value replaces the generic node for that rule in the tree.

```go
// The Actions interface defines methods for transforming parsed nodes.
type Actions interface {
    // The method signature now returns a specific, custom node type
    // that must implement the TreeNode interface.
    MakeInteger(input string, start, end int, elements []TreeNode) (*IntegerNode, error)
}
```

The user provides a struct that implements this interface. The parser invokes the methods, and the returned custom nodes are integrated directly into the parse tree.

**Reasoning**:
- **Type-Safe Transformations**: By returning specific types (e.g., `*IntegerNode`) that implement the `TreeNode` interface, actions produce a semantic, strongly-typed tree from the very beginning. This eliminates the need for downstream type assertions on a generic `Value` field.
- **Decoupling via Interfaces**: The `Actions` interface decouples the generated parser from the user's semantic logic. The parser guarantees the structure and inputs, and the user provides the implementation that transforms the raw parse nodes into a meaningful data structure. The Go compiler enforces that the return values are valid `TreeNode` implementers, ensuring correctness at compile time.

## 5. Code Generation and Templates

The `src/builders/go.js` builder will generate the Go source files from templates.

- It will create a single directory for the new package (e.g., `my_grammar/`).
- It will generate `parser.go`, `tree_node.go`, and `actions.go` (for the interface).
- The builder will translate the Canopy AST into Go method calls, `if/else` blocks, and struct instantiations.
- The `_template` function will use Go's `text/template` or `html/template` conventions for placeholders.

**Reasoning**:
- **Standard Package Structure**: Generating code into a dedicated directory is the standard for Go packages. Separating structs and interfaces into their own files (`tree_node.go`, `actions.go`) is a common convention that improves code organization and readability.
- **Leveraging Go Tooling**: The generated code will be formatted with `gofmt` automatically as a final step of the build process, ensuring it conforms to Go's style guidelines.
