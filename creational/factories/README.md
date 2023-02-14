# Factories

Ways of controlling how an object is constructed.

## One sentence description

A component responsible solely for the wholesale (not piecewise) creation of objects.

## Motivation

- Object creation logic becomes too convoluted.
- Struct has too many fields, need to initialize all correctly.
- Wholesale object creation (non-piecewise, unlike [Builder](../builder/)) can be outsourced to

  - A seperate function (Factory Function, a.k.a. Constructor).
  - That may exist in a separate struct (Factory).

## Summary

- A factory function (a.k.a. constructor) is a helper function for making struct instances.
- A factory is any entity that can take care of object creation.
- Can be a function or a dedicated struct.
