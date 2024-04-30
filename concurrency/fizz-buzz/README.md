# FizzBuzz Mlti Threaded
You are given a class `FizzBuzz` with methods to print "fizz", "buzz", "fizzbuzz", and numbers. The task is to print numbers from 1 to `n` following certain rules: 
- If the number is divisible by 3 but not by 5, print "fizz".
- If the number is divisible by 5 but not by 3, print "buzz".
- If the number is divisible by both 3 and 5, print "fizzbuzz".
- Otherwise, print the number itself.

## Intuition
The given problem can be solved using concurrency and synchronization techniques to print the desired output sequentially without any data race.

## Approach
1. Initialize a `FizzBuzz` struct with a mutex, a wait group, the current number `i`, and the upper limit `n`.
2. Define methods `printFizz`, `printBuzz`, `printFizzBuzz`, and `printNumber` to handle printing of "fizz", "buzz", "fizzbuzz", and numbers, respectively.
3. In each method, use a loop to print the corresponding value based on the given conditions. Use the mutex to prevent data race when accessing shared resources.
4. In the `main` function, create a `FizzBuzz` object with the desired limit `n`.
5. Spawn four goroutines using the `go` keyword to call the respective methods concurrently.
6. Wait for all goroutines to finish using the `Wait` method of the wait group.
7. Finally, print the result.

## Time Complexity
- The time complexity of this solution is O(n), where n is the given limit.

## Space Complexity
- The space complexity is O(n) due to the result array storing the output.