# Command

YOU SHALL NOT PASS!

## One sentence description

An object which represents an instruction to perform a particular action. Contains all the information necessary for the action to be taken.

## Motivation

- Ordinary statements are perishable.
  - Cannot undo field assignment.
  - Cannot directly serialize a sequence of actions (calls).
- Want an object that represents an operation.
  - `person` should change it `age` to value `22`.
  - `car` should do `explode()`.
- Uses: GUI commands, multi-level undo/redo, macro recording and more!

## Summary

- Encapsulate all details of an operation in a separate object.
- Define functions for applying the command (either in the command itself, or eleswhere).
- Optionally define instructions for undoing the command.
- Can create composite commands (a.k.a macros).
