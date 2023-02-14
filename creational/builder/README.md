# Builder

When construction gets a little bit too complicated.

## One sentence description

When piecewise object construction is complicated, provide an API for doing it succinctly.

## Motivation

- Some objects ara simple and can be created in single constructor call, other objects require a lot of ceremony to create.
- Having a factory function with 10 arguments is not productive. Instead, opt for piecewise (piece-by-piece) construction.
- Builder provides an API for constructing an object step-by-step.

## Summary

- A builder is a separate component used for building an object.
- To make builder fluent, return the receiver - allowing chaining.
- Different facets of an object can be built with different builders working in tandem via a common struct.
