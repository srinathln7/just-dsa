# Print in order 
Implement a concurrency model using Go channels that allows three different goroutines to execute three methods of a `Foo` struct (`first()`, `second()`, and `third()`) sequentially.

## Intuition
We can use three channels (`ch1`, `ch2`, and `ch3`) to coordinate the execution of the three methods. Each method waits to receive a signal from the previous method before executing.

## Approach
1. Create a struct `Foo` with three channels `ch1`, `ch2`, and `ch3`.
2. Implement the `first()`, `second()`, and `third()` methods of `Foo`, where each method prints a message and waits to receive a signal from the previous channel before proceeding.
3. In the `main()` function, create an instance of `Foo`.
4. Start three goroutines, each calling one of the `first()`, `second()`, or `third()` methods of `Foo`.
5. Signal the `ch3` channel to start the execution sequence.

## Time Complexity
The time complexity is constant `O(1)` because the number of operations does not depend on the input size.

## Space Complexity
The space complexity is constant `O(1)` because the memory usage does not depend on the input size.