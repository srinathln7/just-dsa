# Gas Station
You are given two integer arrays `gas` and `cost` representing the amount of gas at each station and the cost of traveling from one station to the next, respectively. You need to find the starting gas station's index from where you can travel around the circuit once in the clockwise direction without running out of gas. If there is no such starting station, return -1. If there exists a solution, it is guaranteed to be unique.

## Intuition
To solve this problem efficiently, we can use a greedy approach. We iterate through the gas stations, accumulating the incoming gas (from refueling at stations) and outgoing gas (from traveling to the next station). If the total incoming gas is less than the total outgoing gas, it implies that there is not enough gas to complete the circuit. Otherwise, we iterate through the stations again, maintaining a running sum of the net gas (incoming - outgoing). If the net gas ever becomes negative, we reset the starting index and the net gas to 0, as starting from the current index would not yield a successful circuit.

## Approach
1. Initialize variables `incoming` and `outgoing` to 0 to track the total incoming and outgoing gas, respectively.
2. Iterate through the gas stations using a loop:
    - Add the gas at the current station to `incoming`.
    - Add the cost of traveling to the next station to `outgoing`.
3. If `incoming` is less than `outgoing`, return -1 since there is not enough gas to complete the circuit.
4. Initialize variables `startIdx` and `net` to 0 to track the starting index and the net gas, respectively.
5. Iterate through the gas stations again using a loop:
    - Update `net` by adding the gas at the current station and subtracting the cost of traveling to the next station.
    - If `net` becomes negative, reset `startIdx` to the next station's index and reset `net` to 0.
6. Return `startIdx` as the starting gas station's index.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of gas stations. This is because we iterate through the gas stations twice.

## Space Complexity
The space complexity is O(1) since we use only a constant amount of extra space regardless of the size of the input arrays.