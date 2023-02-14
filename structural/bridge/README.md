# Bridge

Connecting components together through abstractions.

## One sentence description

A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy).

## Motivation

- Bridge prvents a `Cartesian product` complexity explosion.
- Example:
  - Common type ThreadScheduler.
  - Can be preemptive or cooperative.
  - Can run on Windows or Unix
  - End up with a 2x2 scenario: WindowsPTS, UnixPTS, WindowsCTS, UnixCTS
- Bridge pattern avoids the entity explosion.

## Summary

- Decouple abstraction from implementation.
- Both can exist as hierarchies.
- A stronger form of encapsulation.
