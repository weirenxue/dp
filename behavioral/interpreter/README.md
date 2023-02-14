# Interpreter

Interpreters are all around us. Even now, in this very room.

## One sentence description

A component that processes structured text data. Does so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).

## Motivation

- Textual input needs to be processed
  - E.g., turned into linked structures.
  - AST = Abstract Syntax Tree.
- Some examples.
  - Programming language compilers, interpreters and IDEs.
  - HTML, XML and similar.
  - Numeric expression (3+4/5).
  - Regular expressions.
- Turning strings into linked structures in a complicated process.

## Summary

- Barring simple cases, an interpreter acts in two stages.
- Lexing turns text into a set of tokens, e.g.
  `3*(4+5) -> Lit[3] Star Lparen Lit[4] Plus Lit[5] Rparen`
- Parsing tokens into meaningful constructs (AST = Abstract Syntax Tree)  
  <!-- markdownlint-disable MD038 -->
  `-> MultiplicationExpression[`  
  `　　　Integer[3],`  
  `　　　AdditionExpression[`  
  `　　　　　Integer[4], Integer[5]`  
  `　　　]`  
  `]`
  <!-- markdownlint-restore MD038 -->
- Parsed data can then be traversed using the Visitor pattern.
