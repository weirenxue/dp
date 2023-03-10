# Facade

Exposing several components through a single interface.

## One sentence description

Provides a simple, easy to understand/use interface over a large and sophisticated body of code.

## Motivation

- Balancing complexity and presentation/usability.
- Typical home
  - Many subsystem (electrical, sanitation).
  - Complex internal structure (e.g., floor layers).
  - End user is not exposed to internals.
- Same with software.
  - Many systems working to provide flexibility, but...
  - API consumers want it to 'just work'.

## Summary

- Build a Facade to provide a simplified API over a set of components.
- May wish to (optionally) expose internals through the facade.
- May allow users to 'escalate' to use more complex APIs if they need to.
