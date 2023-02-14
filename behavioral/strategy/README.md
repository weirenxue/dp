# Strategy

System behavior partially specified at runtime.

## One sentence description

Seperates an algorithm into its 'skeleton' and concrete implementation steps, which can be varied at run-time.

## Motivaiton

- Many algorithms can be decomposed into higher- and lower-level parts.
- Making tea can be decomposed into
  - The process of making a hot beverage (boil water, pour into cup); and
  - Tea-specific things (put teabag into water).
- The high-level algorithm can then be reused for making coffee or hot chocolate
  - Supported by beverage-specific strategies.

## Summary

- Define an algorithm at a high level.
- Define the interface you expect each strategy to follow.
- Support the injection of the strategy into the high-level algorithm.
