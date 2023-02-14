# Mediator

Facilitates communication between components.

## One sentence description

A component that facilitates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other.

## Motivation

- Componenets may go in and out of a system at any time.
  - Chat room participants.
  - Players in an MMORPG.
- It makes no sense for them to have direct references (pointers) to one another.
  - Those references may go dead.
- Solution: have then all refer to some central component that facilitates communication.

## Summary

- Create the mediator and have each object in the system point to it.
  - E.g., assign a field in factory function.
- Mediator engages in bidrectional communication with its connected components.
- Mediator has methods the components can call.
- Components have methods the mediator can call.
- Event processing (e.g., Rx) libraries make communication easier to implement.
