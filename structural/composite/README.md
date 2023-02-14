# Composite

Treatng individual and aggragate objects uniformly.

## One sentence description

A mechanism for treating individual (scalar) objects and compositions of objects in a uniform manner.

## Motivation

- Object use other objects' fields/methods through embedding.
- Composition lets us make compound objects.
  - E.g., a mathematical expression composed of simple expressions; or
  - A shape group made of serveral different shapes.
- Composite design pattern is used to treat both single (scalar) and composite object uniformly.
  - I.e., `Foo` and `[]Foo` have common APIs.

## Summary

- Objects can use other objects via composition.
- Some composed and singular objects need similar/identical behaviors.
- Composite design pattern lets us threat both types of objects uniformly.
- Iteration supported with the Iterator design pattern.
