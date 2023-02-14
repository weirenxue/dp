# Decorator

Adding behavior without altering the type itself.

## One sentence description

Facilitates the addition of behaviors to individual objects through embedding.

## Motivation

- Want to augment an object with addtional functionality.
- Do not want to rewrite or alter existing code (OCP).
- Want to keep new functionality separate (SRP).
- Need to be able to interact with existing structures.
- Solution: embed the decorated object and provide addtional functionality.

## Summary

- A decorator embeds the decorated object.
- Adds utility fields and methods to augment the object's features.
- Often used to emulate multiple inheritance (may require extra work).
