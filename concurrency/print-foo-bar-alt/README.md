# Print "Foo" and "Bar" Alternately

## Problem Statement:
You are given two functions, `Foo()` and `Bar()`. The task is to print `"foo"` and `"bar"` alternately `n` times.

## Intuition:
To alternate between printing `"foo"` and `"bar"`, we can use a synchronization mechanism. By using a channel to synchronize the execution of the `Foo()` and `Bar()` methods, we can ensure that they execute in the desired sequence.

## Approach:
1. Create a struct `FooBar` with fields `n` (the number of times to print) and a channel `ch` to synchronize the execution of `Foo()` and `Bar()`.
2. Implement the `NewFooBar` function to initialize a `FooBar` object with the given number `n` and a buffered channel.
3. Define the `Foo()` method to print `"foo"` and signal the channel to allow the `Bar()` method to execute.
4. Define the `Bar()` method to wait for the signal from `Foo()`, print `"bar"`, and signal the channel to allow the `Foo()` method to execute.
5. Use goroutines to concurrently execute the `Foo()` and `Bar()` methods.

## Time Complexity:
The time complexity of this approach is `O(n)`, where `n` is the number of times to print `"foo"` and `"bar"` alternately.

## Space Complexity:
The space complexity is `O(1)` because the additional space used does not depend on the input size.