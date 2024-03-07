package main

import "fmt"

func experiement() {
	// Original array
	arr := [6]int{0, 1, 2, 3, 5, 0}
	fmt.Println("Original array:", arr)

	// Element to insert
	element := 4
	// Index to insert the element
	index := 4

	// Shift elements to the right starting from the specified index
	//copy(arr[index+1:], arr[index:len(arr)-1])
	copy(arr[index+1:], arr[index:]) // Both the above are same

	// Insert the new element at the specified index
	arr[index] = element

	// Print the modified array with the inserted element
	// 	Original array: [0 1 2 3 5 0]
	// Array with inserted element: [0 1 2 3 4 5]
	fmt.Println("Array with inserted element:", arr)

	// Since the array's size is unchanged, if we want to insert say a specified value 4 at index 4, then we have to first
	// move all elements to the right of index 4 by 1 unit to the right (which means the last element would be popped out by default)
	// and then insert the character.

}
