**Explanation:**

1. The `Serialize` function takes the root of the N-ary tree and recursively encodes each node's value and its children into a string. The string representation follows the format `value(child1,child2,...,childN)`, where each child is represented recursively.

2. The `Deserialize` function takes the serialized string and reconstructs the N-ary tree. It first parses the root value, then recursively parses and constructs the children nodes based on the parentheses and commas in the string.

3. The `parseValue` and `parseNode` helper functions are used to extract values and child nodes from the serialized string.

4. In the `main` function, an example N-ary tree is created, serialized, and then deserialized to demonstrate the correctness of the algorithm.

When you run this code, it will output:

```
Serialized tree: 1(3(5,6),2,4)
Deserialized tree: 1(3(5,6),2,4)
```

