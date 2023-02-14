# Template Method

A high-level bluprint for an algorithm to be completed by inheritors.

## One sentence description

A skeleton algorithm defined in a function. Function can either use an interface (like Strategy) or can take serveral functions as arguments.

## Motivation

- Algorithms can be decomposed into common parts + specifics.
- Strategy pattern does this through composition.
  - High-level algorithm uses an interface.
  - Concrete implementations implement the interface.
  - We keep a pointer to the interface; provide concrete implementations.
- Template Method performs a similar operation, but
  - It's typically just a function, not a struct with a reference to the implementation.
  - Can still use interfaces (just like Strategy); or
  - Can be functional (take several functions as parameters).

## Summary

- Very similar to Strategy.
- Typically implementation:
  - Define an interface with common operations.
  - Make use of those operations inside a function.
- Alternative functional approach:
  - Make a function that takes several functions.
  - Can pass in functions that capture local state.
  - No need for either structs or interfaces!
