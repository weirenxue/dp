# Proxy

An interface for accessing a particular resource.

## One sentence description

A type that functions as an interface to a particular resource. That resource may be remote, expensive to construct, or may require logging or some other added functionality.

## Motivation

- You are calling `foo.Bar()`. This assumes that `foo` is in the same process as `Bar()`.
- What if, later on, you want to put all `Foo`-related operations into a separate process?
  - Can you avoid changing your code?
- Proxy to the rescue!
  - Same interface, entirely different behavior.
- This is called a *communication proxy*
  - Other type: logging, virtual, guarding, ...

## Summary

- A proxy has the same interface as the underlying object.
- To create a proxy, simply replicate the existing interface of an object.
- Add relevant functionality to the redefined methods.
- Different proxies (communication, logging, caching, etc.) have completely different behaviors.

## Proxy vs. Decorator

- Proxy tries to provide an identical interface; decorator provides an enhanced interface.
- Decorator typically aggregates (or has pointer to) what it is decorating; proxy doesn't have to.
- Proxy might not event be working with materialized object.
