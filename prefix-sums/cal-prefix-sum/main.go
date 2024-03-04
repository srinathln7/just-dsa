package main

type NumArray struct {
	Val []int
}

func Constructor(nums []int) NumArray {
	val := make([]int, len(nums))
	val[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		val[i] = val[i-1] + nums[i]
	}
	return NumArray{Val: val}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.Val[right] - this.Val[left-1]
	}
	return this.Val[right]
}
