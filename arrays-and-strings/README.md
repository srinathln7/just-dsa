# Arrays


## Static Array

Insertion on end  - `O(1)`

Insertion on middle - `O(n)`

Deletion on end - `O(1)`

Deletion on middle - `O(n)`

Reading - `O(1)`


## Slice (Dynamic Array)

Insertion on end  - `O(1)`

Insertion on middle - `O(n)`, amotized time complexity is `O(1)`

Deletion on end - `O(1)`

Deletion on middle - `O(n)`

Reading - `O(1)`


## How Slices work in Go?

In Go, a slice is a lightweight data structure that provides a dynamic view into an underlying array. It consists of three components:

1. **Pointer to the underlying array**: This pointer points to the first element of the underlying array that the slice refers to. It specifies the starting point of the slice.

2. **Length**: The length of the slice represents the number of elements present in the slice. It specifies how many elements are accessible within the slice.

3. **Capacity**: The capacity of the slice represents the maximum number of elements that the slice can hold without resizing the underlying array. It specifies the maximum length that the slice can grow to without allocating additional memory.

The slice data structure is defined in Go's runtime and is not directly exposed in user code. However, the behavior and characteristics of slices are well-defined by the Go language specification.

Here's a visual representation of a slice:

```
+---------------------+
|   Pointer to Array  | --> underlying array
+---------------------+
|        Length       | --> number of elements accessible
+---------------------+
|       Capacity      | --> maximum length without reallocation
+---------------------+
```

In Go, slices are created using the slice expression syntax `a[start:end]`, where `start` is the starting index and `end` is the exclusive ending index of the slice. If `start` is omitted, it defaults to `0`, and if `end` is omitted, it defaults to the length of the array or slice being sliced.

Slices can be dynamically resized by using the `append()` function, which may allocate a new underlying array if necessary and copy existing elements to the new array.

### Passing Slice by Value and Ref to other functions

Go is a pass-by-value language, which means that when you pass an argument to a function, a copy of the argument is made and passed to the function. If you pass a slice to a function, you're passing a copy of the slice header, which includes the slice's length, capacity, and a pointer to the underlying array. Modifying the slice header inside the function (such as appending elements to it) will not affect the original slice outside the function. By passing a reference to the slice (For eg: `*[]int`), we're essentially passing a pointer to the slice header. This allows us to modify the original slice directly inside the function by dereferencing the pointer. Any modifications made to the slice inside the function will be visible outside of it, as we're working with the original slice. If we were to pass the slice itself (`[]int`), modifications made to the slice inside the function would not be reflected outside of it, as we would be working with a copy of the slice header. Therefore, passing a reference to the slice ensures that we can update the slice's contents and see those changes outside of the calling function.


### But.. Aren't slice itself a pointer to the underlying array?
While a slice in Go is essentially a data structure that includes a reference to an underlying array, along with the length and capacity of the slice. When you pass a slice to a function, you are indeed passing a reference to the underlying array and other metadata (length and capacity). However, it's important to note that while the reference to the underlying array is shared between the original slice and the slice passed to the function, the slice header itself is still passed by value. This means that modifications to the slice header, such as appending elements or resizing the slice, are not reflected outside of the function when you pass the slice directly (without using a pointer). So, passing a reference to the slice (`*[]int`) allows you to modify the slice header (such as appending elements) inside the function and see those changes reflected outside of it. Passing the slice directly (`[]int`) would not allow such modifications to be visible outside the function.


```
package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a)

	changeMe(a)
	fmt.Println(a)

	changeMeRef(&a)
	fmt.Println(a)

	// Output
	// [1 2]
	// [3 4 5 6]
	// [3 4]
	// [7 8 9 10]
	// [7 8 9 10]
}

func changeMe(a []int) {
	a[0], a[1] = 3, 4  // Changing the underlying ref.
	a = append(a, 5) // modifying slice header refleced only inside the scope (not in `main`)
	a = append(a, 6) // modifying slice header refleced only inside the scope (not in `main`)	
	fmt.Println(a)
}

func changeMeRef(a *[]int) {
	(*a)[0], (*a)[1] = 7, 8 // Changing the underlying ref.
	*a = append(*a, 9) // modifying pointer to the slice header refleced outside the scope
	*a = append(*a, 10)// modifying pointer to the slice header refleced outside the scope 
	fmt.Println(*a)
}


```
