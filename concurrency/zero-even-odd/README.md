# Print Zero Odd Even 

We have a struct `ZEO` with three methods: `zero`, `one`, and `two`. Each method is responsible for generating a specific pattern of numbers or digits based on certain conditions.

### Intuition

The `ZEO` struct uses a mutex to ensure that only one goroutine modifies the shared state at a time. Each method is called concurrently as goroutines. The goal is to print a sequence of numbers and zeros, where the pattern depends on the method.

### Approach

1. Create a `ZEO` struct with fields for managing synchronization, tracking state, and storing the result.
2. Implement three methods: `zero`, `one`, and `two`, each representing a specific pattern.
3. Use a mutex to control access to shared data to prevent data races.
4. Inside each method, use a loop to generate the desired sequence of numbers and zeros based on the method's conditions.
5. Break out of the loop when the terminating condition (`i > n`) is met.
6. Add each generated digit or number to the `result` string.
7. Release the mutex lock after each critical section to allow other goroutines to access shared data.
8. In the `main` function, initialize the `ZEO` struct, start the goroutines for each method, wait for all goroutines to finish using `sync.WaitGroup`, and then print the final result.

### Time and Space Complexity

- **Time Complexity:** The time complexity of each method is O(n), where n is the input number. The overall time complexity depends on the number of goroutines and the execution time of each method.
- **Space Complexity:** The space complexity is O(n) due to the `result` string, which stores the generated sequence.