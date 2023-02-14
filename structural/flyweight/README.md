# Flyweight

Space Optimization!

## One sentence description

A space optimization technique that lets us use less memory by storing externally the data associated with similar objects.

## Motivation

- Avoid redundancy when storing data.
- E.g., MMORPG.
  - Plenty of users with identical first/last names.
  - No sense in storing same first/last name over and over again.
  - Store a list of names and references to them (indices, pointers, etc.).
- E.g., bold or italic text formatting
  - Don't want each character to have a formatting character.
  - Operate on ranges (e.g., line number, start/end positions).

## Summary

- Store common data externally.
- Specify an index or a pointer into the external data store.
- Define the idea of 'ranges' on homogeneous collections and store date related to those ranges.
