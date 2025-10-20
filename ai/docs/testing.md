# Testing and CI Strategy

This document outlines the testing and Continuous Integration (CI) strategy for the Canopy project.

## Continuous Integration (CI)

The project uses [GitHub Actions](https://github.com/features/actions) for its CI pipeline. The configuration files are located in the `.github/workflows` directory.

Four separate workflows are defined:

- `java.yml`: Tests the Java code.
- `node.yml`: Tests the JavaScript code.
- `python.yml`: Tests the Python code.
- `ruby.yml`: Tests the Ruby code.

These workflows are triggered on every `push` and `pull_request` to the repository. Each workflow runs the test suite against multiple versions of the corresponding language to ensure broad compatibility.

The CI jobs are kicked off using `make` with a specific target for each language: `test-java`, `test-js`, `test-python`, and `test-ruby`.

## Test Setup

A core part of the testing strategy involves compiling `.peg` grammar files into parsers for each target language. This process is managed by the `Makefile` and the `bin/canopy` compiler.

1.  **Grammar Files**: The base test grammars are located in `test/grammars` and have a `.peg` extension.

2.  **Compilation**: Before running the tests for a specific language, the `Makefile` uses the `bin/canopy` executable to compile the `.peg` files into language-specific modules. For example, `test/grammars/choices.peg` is compiled into:
    - `test/grammars/choices.js` for JavaScript
    - `test/grammars/choices/Grammar.java` for Java
    - `test/grammars/choices.py` for Python
    - `test/grammars/choices.rb` for Ruby

3.  **Usage in Tests**: The test files for each language then import these generated modules to parse input strings and assert that the resulting syntax tree is correct. For example, in `test/python/choices_test.py`, the compiled grammar is imported with `from grammars import choices`.

This ensures that the Canopy compiler is producing correct and usable parsers for all supported languages.

## Language-Specific Testing

The testing strategy is tailored for each of the languages supported by Canopy.

### Java

- **Location**: `test/java`
- **Framework**: [JUnit 5](https://junit.org/junit5/)
- **Runner**: [Apache Maven](https://maven.apache.org/)
- **Setup**: The `pom.xml` file in the test directory declares the necessary dependencies.
- **Run Command**:
  ```bash
  make test-java
  ```

### JavaScript

- **Location**: `test/javascript`
- **Framework**: [jstest](https://github.com/jcoglan/jstest)
- **Runner**: [npm](https://www.npmjs.com/)
- **Setup**: Run `npm install` in the `test/javascript` directory to install dependencies.
- **Run Command**:
  ```bash
  make test-js
  ```

### Python

- **Location**: `test/python`
- **Framework**: [pytest](https://docs.pytest.org/)
- **Runner**: [pipenv](https://pipenv.pypa.io/)
- **Setup**: Run `pipenv install --dev` from the project root to create a virtual environment and install dependencies listed in the `Pipfile`.
- **Run Command**:
  ```bash
  make test-python
  ```

### Ruby

- **Location**: `test/ruby`
- **Framework**: [Test::Unit](https://www.rubydoc.info/stdlib/test-unit) (via Rake)
- **Runner**: [Rake](https://ruby.github.io/rake/)
- **Setup**: Run `bundle install` from the project root to install dependencies from the `Gemfile`.
- **Run Command**:
  ```bash
  make test-ruby
  ```
