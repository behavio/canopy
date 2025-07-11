# Go Implementation Modification Plan

This document outlines the plan to add support for Go (version 1.24) as a target language for the Canopy parser generator, based on the existing `Makefile` structure.

## Affected Files

The following files will be created or modified during this implementation:

**New Files:**
- `src/builders/go.js`
- `templates/go/actions.go.tpl`
- `templates/go/parser.go.tpl`
- `templates/go/treenode.go.tpl`
- `test/go/choices_test.go`
- `test/go/extensions_test.go`
- `test/go/go.mod`
- `test/go/node_actions_test.go`
- `test/go/parse_helper.go`
- `test/go/predicates_test.go`
- `test/go/quantifiers_test.go`
- `test/go/sequences_test.go`
- `test/go/terminals_test.go`
- `.github/workflows/go.yml`
- `site/langs/go.md`

**Modified Files:**
- `bin/canopy`
- `Makefile`
- `README.md`
- `site/_config.yml`

---

## Phase 1: Build Integration and Scaffolding

The goal of this phase is to integrate Go into the project's build system. This phase will **not** generate any functional Go code; it will only produce the empty directories and module files required for a valid Go module structure, triggered by new `Makefile` rules.

- **Update the Makefile:**
    - Define `example_grammars_go` as `$(example_grammars:%.peg=%/parser.go)`.
    - Add a new pattern rule: `%/parser.go: %.peg`. The recipe will run `./bin/canopy --lang go $<`.
    - Add a `test-go` target that depends on `$(test_grammars_go)` and runs the Go test command.
    - Add `test-go` to the `test-all` target.
    - Add `$(example_grammars_go)` to the `examples` target.
- **Update the compiler:** Modify `bin/canopy` to recognize `--lang go`. It should create the output directory (e.g., `test/grammars/choices/`) and a `go.mod` file inside it.
- **Create a skeleton builder:** Add a new `src/builders/go.js` with the minimal logic to create the output directory and placeholder files.
- **Create the test module:** Add the `test/go` directory and a `test/go/go.mod` file.
- **Use the template files:** the content is static for now, just has to compile.
- **Do not run go test yet in the Makefile**

**Files to modify:**
- `Makefile`
- `bin/canopy`
- `src/builders/go.js` (new)
- `test/go/go.mod` (new)
- `templates/go/actions.go.tpl` (new)
- `templates/go/parser.go.tpl` (new)
- `templates/go/treenode.go.tpl` (new)

**Testable Outcome:** Running `make test-go` triggers the creation of all test grammar directories (e.g., `test/grammars/choices/`), each containing a `go.mod` file. The command then attempts to run tests, which does nothing as none exist. Running `make examples` does the same for the example grammars.

---

## Phase 2: Code Generation and Compilation

This phase focuses on implementing the code generation logic to produce valid, compilable Go code from the grammar files.

- **Implement the Go builder:** Flesh out `src/builders/go.js` to translate the Canopy AST into Go source code.
- **Create Go templates:** Create and populate `actions.go.tpl`, `parser.go.tpl`, and `treenode.go.tpl` in `templates/go/` with the necessary template logic to generate the Go files as defined in the design document.
- **Ensure compilability:** The primary goal is to ensure that all generated Go parsers compile successfully when requested by `make`.

**Files to modify:**
- `src/builders/go.js`
- `templates/go/actions.go.tpl`
- `templates/go/parser.go.tpl`
- `templates/go/treenode.go.tpl`

**Testable Outcome:** Running `make test-go` should compile all test grammars into valid Go packages without any compilation errors before attempting to run the (still non-existent) tests.

---

## Phase 3: Test Framework and Test Porting

With a working code generator, the next step is to build the test suite to verify the correctness of the generated parsers.

- **Create a parse helper:** Implement `test/go/parse_helper.go` to provide common setup and assertion functions.
- **Port all tests:** Translate the tests from an existing language suite to Go. This involves creating `*_test.go` files that import the compiled grammars and assert their correctness.

**Files to modify:**
- `test/go/parse_helper.go` (new)
- `test/go/choices_test.go` (new)
- `test/go/extensions_test.go` (new)
- `test/go/node_actions_test.go` (new)
- `test/go/predicates_test.go` (new)
- `test/go/quantifiers_test.go` (new)
- `test/go/sequences_test.go` (new)
- `test/go/terminals_test.go` (new)

**Testable Outcome:** Running `make test-go` now executes the full test suite. Tests are expected to fail, but they must compile and run without panicking.

---

## Phase 4: Test Execution and Implementation Polishing

This is the main implementation loop where the generated code is tested and refined until it is fully correct.

- **Fix bugs:** Run `make test-go` repeatedly. Debug and fix issues in the `src/builders/go.js` builder and the Go templates until all tests pass.

**Files to modify:**
- `src/builders/go.js`
- `templates/go/*.tpl` (as needed)

**Testable Outcome:** Running `make test-go` completes successfully with all tests passing.

---

## Phase 5: CI, Documentation, and Cleanup

The final phase is to integrate the Go tests into the CI pipeline, provide documentation, and update cleanup scripts.

- **Create CI workflow:** Add a `.github/workflows/go.yml` file that runs `make test-go` on pushes and pull requests.
- **Write documentation:** Create `site/langs/go.md` explaining how to use the Go language target.
- **Update site navigation:** Modify `site/_config.yml` to include the new Go documentation page.
- **Update README:** Add Go to the list of supported languages in `README.md`.
- **Update clean rule:** Modify the `clean-test` and `clean-examples` rules in the `Makefile` to remove the generated Go grammar directories.

**Files to modify:**
- `.github/workflows/go.yml` (new)
- `site/langs/go.md` (new)
- `site/_config.yml`
- `README.md`
- `Makefile`

**Testable Outcome:** The `go` workflow passes in GitHub Actions. The Go documentation is rendered on the project's website. Running `make clean` removes all generated Go files and directories.
