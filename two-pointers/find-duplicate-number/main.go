package main

func findDuplicate(nums []int) int {

	// Key Idea: NOTICE the integer ranges from [1...n]. This is a good clue that we might want to use the
	// range itself as an index of the array. To find the duplicate number that occurs only once => We can frame
	// it as a linked list problem We'll treat the array as a linked list where each element points to the value at its index.
	// Since there is one repeated number, there will be a cycle in this linked list. Notice that`0` is never
	// included in the range which means no node will point to the zeroth node. This is IMPORTANT since this gaurantees we
	// start traversing the linked list at a non-cyclic portion.

	// For Ex: [1,3,4,2, 2] => Node 0 with Val 1 points to Node 1 with Val 3. Node 1 with Val 3 points Node 3 with Val 2 etc.

	// Initialize both `slow` and `fast` pointers at the non-cyclic head portion
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		// Intersecting point
		if slow == fast {
			break
		}
	}

	//  Move slow back to the head and start traversing until `slow`
	// and `fast` meet which represents the start of the cyclic portion => Duplicate number
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}
