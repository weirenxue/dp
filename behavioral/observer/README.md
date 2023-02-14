# Observer

I am watching you!

## One sentence description

An observer is an object that wishes to be informed about events happening in the system. The entity generating the events is an observable.

## Motivation

- We need to be informed when certain things happen.
  - Object's field changes.
  - Object does something.
  - Some external event occurs.
- We want to listen to events and notified when they occur.
- Two participants: *observable* and *observer*.

## Summary

- Observer is an intrusive approach.
- Must provide a way of clients to subscribe.
- Event data sent from observable to all observers.
- Data represented as `interface{}`.
- Unsubscription is possible.
