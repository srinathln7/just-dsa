package main

type Element struct {
	val int
	min int
}

func find132pattern(nums []int) bool {

	// Key Idea: To find the 132 pattern, we need to maintain a monotonic decreasing stack technique to efficiently
	// find the `k` candidate who is lesser than the top of the stack but greater than all of the other elements in the
	// stack. To check if it is greater than all of the other elements except top, we keep track of the left minimum of
	// each element before pushing it to stack

	n := len(nums)
	if n < 3 {
		return false
	}

	var stack []Element
	currLeftMin := nums[0]
	for i := 0; i < n; i++ {

		// We need to maintain a stack that is monotonically decreasing in order
		// Hence pop if the incoming number is greater than the top of the stack
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1].val {
			stack = stack[:len(stack)-1]
		}

		// Check for an ideal `k` candidate
		if len(stack) > 0 && (nums[i] > stack[len(stack)-1].min) && (nums[i] < stack[len(stack)-1].val) {
			return true
		}

		stack = append(stack, Element{val: nums[i], min: currLeftMin})
		currLeftMin = min(currLeftMin, nums[i])
	}

	return false
}
