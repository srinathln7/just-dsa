package main 


func findMin(nums []int) int {
    
    // Key Idea: Since the original array is sorted and we need a O(log.n) algorithm, we can think
    // of using binary search. Indeed, if the array is rotated, we can check if the mid. element 
    // is lesser than the right element. If it was, it means the right half is already sorted,
    // we move our search to the left. If otherwise, we move our search to the right half.

    l, r := 0, len(nums)-1
    for l <= r {
        mid := l + (r-l)/2
        switch {
            // Right half is sorted and move the search to the left.
            // NOTE: `mid` can still be the min. element
            case nums[mid] < nums[r]: r = mid
            
            // Since nums[mid] > nums[r], we know for sure, `mid` is not the minimum
            // Hence, we move the search to `mid+1`
            default: l = mid+1
        }
    }

    return nums[r]
}
