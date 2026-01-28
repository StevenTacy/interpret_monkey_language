# Monkey Interpreter in Go

A precise and idiomatic implementation of the Monkey programming language interpreter, featuring a complete toolchain from lexical analysis to AST evaluation.

## Features

- **Lexical Analysis (Lexer)**: Tokenizes raw source code into structured tokens.
- **Recursive Descent Parsing**: A Pratt parser that handles operator precedence and transforms tokens into an Abstract Syntax Tree (AST).
- **Evaluation (Evaluator)**: Traverses the AST and executes the logic using an internal object system.
- **Environment**: Support for variable bindings and scoped execution.
- **REPL**: An interactive Read-Eval-Print Loop for real-time code execution.

## Project Structure

- `lexer/`: Implementation of the lexical scanner.
- `parser/`: The Pratt parser for building the AST.
- `ast/`: Definition of the Abstract Syntax Tree nodes.
- `evaluator/`: The evaluation logic that breathes life into the AST.
- `object/`: The object system used for internal value representation.
- `token/`: Token definitions and keyword mapping.
- `repl/`: Interactive shell implementation.

## Monkey Language Syntax

Monkey supports:
- **Data Types**: Integers, Booleans, Strings, Arrays, and Hashes.
- **Expressions**: Arithmetic (`+`, `-`, `*`, `/`), Comparisons (`==`, `!=`, `<`, `>`), and Prefix operators (`!`, `-`).
- **Statements**: `let` for bindings, `return` for function exit.
- **Functions**: First-class functions with parameters and closures.
- **Control Flow**: `if-else` expressions.

### Built-in Functions

- `len(item)`: Returns the length of a string or array.
- `first(array)`: Returns the first element of an array.
- `last(array)`: Returns the last element of an array.
- `rest(array)`: Returns a new array containing all elements except the first.
- `put(args...)`: Prints the inspection of the provided arguments to stdout.

## Getting Started

### Prerequisites

- Go 1.25.1 or later.

### Installation

Clone the repository and build the project:

```bash
go build -o monkey main.go
```

### Running the REPL

Start the interactive interpreter:

```bash
./monkey
```

Example usage:

```monkey
>> let add = fn(a, b) { a + b };
>> add(10, 5);
15
>> let arr = [1, 2, 3];
>> len(arr);
3
```

## Implementation Details

### Lexer
The lexer (`lexer/lexer.go`) performs a single pass over the input string, identifying characters and grouping them into tokens defined in `token/token.go`. It handles identifiers, numbers, and multi-character operators like `==` and `!=`.

### Parser
The parser (`parser/parser.go`) implements a **Pratt Parser** (Top Down Operator Precedence). 

Key fields in the `Parser` struct:
- `l`: A pointer to the `Lexer` instance.
- `curToken` / `peekToken`: Tracking the current and next tokens in the stream.
- `prefixParseFn`: A map of token types to functions that handle prefix expressions (e.g., `-5`, `!true`).
- `infixParseFn`: A map of token types to functions that handle infix expressions (e.g., `5 + 5`).

### Evaluator
The evaluator (`evaluator/evaluator.go`) implements a tree-walking strategy. It recursively processes AST nodes, maintaining state within an `Environment` to track variable assignments and function scopes. Values are represented using an internal object system (`object/object.go`), supporting `Integer`, `Boolean`, `String`, `Array`, `Hash`, and `Function` types.
