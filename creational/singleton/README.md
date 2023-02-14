# Singleton

A design pattern everyone loves to hate... but is it really that bad?

## One sentence description

A component which is instantiated only once.

## Motivation

- For some components it only makes sense to have one in the system.
  - Database repository.
  - Object factory.
- E.g., the construction call is expensive.
  - We only do it once.
  - We give everyone the same instance.
- Want to prevent anyone creating additional copies.
- Need to take care of lazy instantiation.

## Summary

- Lazy one-time initialization using `sync.Once`.
- Adhere to DIP: depend on interfaces, not concreate types.
- Singleton is not scary.
