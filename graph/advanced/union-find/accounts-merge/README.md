# Merge accounts with common emails


## Intuition
When facing the problem of merging accounts based on common emails, we need to identify groups of accounts that are connected through shared emails. This is a classic union-find problem where each account is initially considered as a separate entity, and accounts are merged as we discover shared emails. The Union-Find algorithm is used to group accounts based on common emails. Each account is represented by a node, and if two accounts share a common email, they are merged into the same group.

## Approach:
1. **UnionFind Struct**: This struct maintains the parent array to track the parent of each node and the rank array for union-by-rank optimization.
2. **newUnionFind Function**: Initializes a new Union-Find data structure for the given accounts. It assigns each account as its own parent and initializes the rank based on the number of emails in each account.
3. **find Method**: Finds the root (parent) of a node in the Union-Find structure using path compression to optimize the search.
4. **union Method**: Performs union operation between two nodes. It checks their ranks and assigns the parent based on rank to optimize the tree structure.
5. **accountsMerge Function**: 
   - Initializes a Union-Find data structure and a map to store the index of each email.
   - Iterates through each account and email, merging accounts with common emails using Union-Find.
   - Collects merged accounts by finding the root for each email index and appending emails to the corresponding root.
   - Formats the merged accounts by sorting emails and appending them to the account's name.
   - Returns the formatted merged accounts.

## Time Complexity
- Constructing the Union-Find data structure: O(n), where n is the number of accounts.
- Merging accounts and finding roots: O(nα(n)), where α(n) is the inverse Ackermann function (nearly constant time).
- Sorting emails for each account: O(nlogn) in the worst-case scenario.
Overall, the time complexity is dominated by the sorting step.

## Space Complexity:
- Union-Find data structure: O(n), where n is the number of accounts.
- Additional space for maps and slices: O(n).
Overall, the space complexity is O(n).

This approach efficiently merges accounts with common emails while maintaining the integrity of the Union-Find data structure.