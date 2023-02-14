# Iterator

How traversal of data structures happens and who makes it happen.

## One sentence description

An object that facilitates the traversal of a data structure.

## Motivation

- Iteration (traversal) is a core functionality of various data structures.
- An `iterator` is a type that facilitates the traversal.
  - Keeps a pointer to the current element.
  - Knows how to move to a different element.
- Go allows iteration with `range`.
  - Built-in support in many object (arrays, slices, etc.).
  - Can be supported in our own structs.

## Summary

- An iterator specifies how you can traverse an object.
- Moves along the iterated collection, indicating when last element has been reached.
- Not idiomatic in Go (no standard Iterable interface).
