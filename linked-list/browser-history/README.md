# Browser History

## Problem

Design a browser history implementation that supports the following operations:
- `Constructor(homepage string)`: Initializes the browser history with the given homepage.
- `Visit(url string)`: Visits the `url`, moving it to the top of the history.
- `Back(steps int)`: Moves `steps` back in history. If you can only return `x` steps in the history and `steps > x`, you will return only `x` steps. Return the current `url` after moving back in history.
- `Forward(steps int)`: Moves `steps` forward in history. If you can only forward `x` steps in the history and `steps > x`, you will forward only `x` steps. Return the current `url` after moving forward in history.

## Approach

We can implement this using a doubly linked list where each node represents a URL. The `BrowserHistory` struct maintains a pointer to the current node in the history. When visiting a new URL, we create a new node and add it to the list, updating pointers accordingly. Backward and forward navigation involves traversing the list in the respective direction by adjusting pointers.

## Implementation

We define a `Node` struct representing a URL node with a string value and pointers to the previous and next nodes. The `BrowserHistory` struct holds a pointer to the current node in the history.

- **Constructor(homepage string):** Initializes the browser history with the given homepage.
- **Visit(url string):** Creates a new node for the visited URL and adds it to the history, updating pointers.
- **Back(steps int):** Moves backward in history by `steps` nodes, updating the current node pointer.
- **Forward(steps int):** Moves forward in history by `steps` nodes, updating the current node pointer.

## Time Complexity

- Constructor: O(1)
- Visit: O(1)
- Back: O(steps)
- Forward: O(steps)

## Space Complexity

O(1) for each operation.
