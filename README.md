# Design Pattern

---

1. Creational patterns

    Provide object creation mechanisms that increase flexibility and reuse of exisiting code.

    1. [Builder](./creational/builder/): When piecewise object construction is complicated, provide an API for doing it succinctly.
    1. [Factory](./creational/factories/): A component responsible solely for the wholesale (not piecewise) creation of objects.
    1. [Prototype](./creational/prototype/): A partially or fully initialized object that you copy (clone) and make use of.
    1. [Singleton](./creational/singleton/): A component which is instantiated only once.

1. Structural patterns

    Explain how to assemble objects and classes into larger structures, while keeping these structures flexible and efficient.

    1. [Adapter](./structural/adapter/): A construct which adapts an existing interface X to conform to the required interface Y.
    1. [Bridge](./structural/bridge/): A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy).
    1. [Composite](./structural/composite/): A mechanism for treating individual (scalar) objects and compositions of objects in a uniform manner.
    1. [Decorator](./structural/decorator/): Facilitates the addition of behaviors to individual objects through embedding.
    1. [Facade](./structural/facade/): Provides a simple, easy to understand/use interface over a large and sophisticated body of code.
    1. [Flyweight](./structural/flyweight/): A space optimization technique that lets us use less memory by storing externally the data associated with similar objects.
    1. [Proxy](./structural/proxy/): A type that functions as an interface to a particular resource. That resource may be remote, expensive to construct, or may require logging or some other added functionality.

1. Behavioral patterns

    Take care of effective communication and the assignment of responsibilities between objects.

    1. [Chain of Responsibility](./behavioral/chain_of_responsibility/): A chain of components who all get a chance to process a command or a query, optionally having default processing implementation and an ability to terminate the processing chain.
    1. [Command](./behavioral/command/): An object which represents an instruction to perform a particular action. Contains all the information necessary for the action to be taken.
    1. [Interpreter](./behavioral/interpreter/): A component that processes structured text data. Does so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).
    1. [Iterator](./behavioral/iterator/): An object that facilitates the traversal of a data structure.
    1. [Mediator](./behavioral/mediator/): A component that facilitates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other.
    1. [Memento](./behavioral/memento/): A token representing the system state. Lets us roll back to the state when the token was generated. May or may not directly expose state information.
    1. [Observer](./behavioral/observer/): An observer is an object that wishes to be informed about events happening in the system. The entity generating the events is an observable.
    1. [State](./behavioral/state/): A pattern in which the object's behavior is determined by its state. An object transitions from one state to another (something needs to trigger a transition). A formalized construct which manages state and transitions is called a state machine.
    1. [Strategy](./behavioral/strategy/)
    1. [Template Method](./behavioral/template_method/)
    1. [Visitor](./behavioral/visitor/)
