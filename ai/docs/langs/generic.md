# Canopy Parser: Generic Design and Structure

## 1. Overview

A Canopy-generated parser is fundamentally a top-down parsing-expression grammar (PEG) parser that uses the packrat parsing technique for memoization. This ensures linear-time parsing performance and provides support for left-recursive grammars.

The generation process is driven by a JavaScript-based builder that translates a Canopy grammar definition into source code for a target language.

The core design principles are consistent across all language implementations:

- **Builder Pattern**: A dedicated builder module in JavaScript (e.g., `src/builders/language.js`) is responsible for traversing the grammar's Abstract Syntax Tree (AST) and generating the source code.
- **Template-Based Code Generation**: The builder uses language-specific templates for the boilerplate code of the parser, syntax tree nodes, and other necessary components.
- **Memoization**: The parser uses a cache (typically a hash map or dictionary) to store the results of parsing rules at different positions in the input string. This avoids redundant parsing and is the key to linear-time performance.
- **Clear Separation of Concerns**: The generated code maintains a clear separation between the parsing logic, the data structures for the parse tree, and user-defined actions.

## 2. Core Components

The generated parser in any language consists of the following main components:

### The `Parser`

This is the main public-facing class that the user interacts with. It is responsible for:

- Initializing the parser state, including the input string, current offset, and memoization cache.
- Starting the parsing process by calling the method for the grammar's root rule.
- Verifying that the entire input was successfully consumed.
- Raising a `ParseError` if parsing fails, providing detailed information about the failure.

### The `Grammar`

This component contains the core parsing logic. For each rule in the Canopy grammar, a corresponding `_read_<rule_name>` method is generated. These methods implement the logic for matching sequences, choices, terminals, and other grammar expressions. The `Parser` uses or inherits from this component to access the parsing methods.

### The `TreeNode`

This is the base class for all nodes in the generated parse tree. It provides at least the following attributes:

- The portion of the input string that the node represents.
- The starting offset of the node's text in the input string.
- A collection of child `TreeNode` objects.

For grammar rules that have a specific name, a subclass of `TreeNode` is often created with that name to allow for more specific type-checking and interaction with the tree.

### The `ParseError`

A custom exception class that is raised when parsing fails. It contains information about the location of the error and what was expected at that point in the input.

## 3. Object Interaction and Parsing Flow

The interaction between the components follows a consistent pattern:

1.  A `Parser` object is instantiated with the input string and an optional object containing user-defined actions.
2.  The `parse()` method of the `Parser` is called.
3.  `parse()` invokes the method for the grammar's root rule.
4.  Each `_read_<rule_name>()` method attempts to match a part of the input string according to its rule.
5.  **Memoization**: Before attempting to parse a rule at a given offset, the parser checks if the result for that rule and offset is already in the cache. If so, the cached result is returned immediately.
6.  If a rule matches successfully:
    - A `TreeNode` object is created to represent the matched part of the input.
    - The result (the `TreeNode` and the new offset) is stored in the cache.
    - The `TreeNode` is returned.
7.  If a rule fails to match:
    - A special `FAILURE` object or `null` is returned.
    - The failure is recorded in the cache to avoid re-attempting the same failing parse.
    - The parser backtracks to try other alternatives in a choice, or fails the parent rule.
8.  If the root rule matches and the entire input string is consumed, the `parse()` method returns the root `TreeNode` of the complete parse tree.
9.  If the root rule does not match or the entire input is not consumed, a `ParseError` is raised.

## 4. Canopy Actions

Canopy allows you to embed actions in the grammar, which are code snippets that are executed when a rule is successfully matched. This feature allows for on-the-fly transformation of the parse tree.

- The user provides a class or object that contains methods corresponding to the action names in the grammar.
- When a rule with an action is matched, the generated code calls the corresponding user-provided method.
- This method is typically passed the input string, the start and end offsets of the match, and a list of the child elements.
- The value returned by the action method replaces the default `TreeNode` in the parse tree. The implementation details vary based on the target language's type system (e.g., a statically-typed interface in Java vs. a dynamic approach in Python).

## 5. Code Generation and Templates

The `src/builders/language.js` file is the heart of the code generator for a given language. It works as a "builder" that traverses the AST of the Canopy grammar and generates source code for each element.

Key methods in the builder typically include:

- `package_` or `grammarModule_`: Initializes the module or classes.
- `class_`, `method_`: Generate class and method definitions.
- `syntaxNode_`: Generates the code to create a `TreeNode` instance or, if an action is present, generates the call to the action method.
- `if_`, `loop_`: Generate conditional and loop statements.
- `cache_`: Generates the memoization logic for a rule.
- `_template`: Reads a template file and performs string substitution to insert generated code or grammar-specific names.

The final output is a self-contained set of files that can be used to parse input strings according to the defined grammar. If allowed by the host language, there is only one file produced.
