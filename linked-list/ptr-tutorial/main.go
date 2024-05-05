package main

import "fmt"

type Listnode struct {
	val  int
	next *Listnode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	// a: &{val:0 next:<nil>} and b:&{val:0 next:<nil>}
	// a: &{val:0 next:0xc0000140b0} and b:&{val:0 next:0xc0000140b0}
	// a: &{val:0 next:0xc0000140b0} and b:&{val:1 next:0xc0000140c0}
	// a: &{val:0 next:0xc0000140b0} and b:&{val:1 next:<nil>}

	var a *Listnode = &Listnode{}
	b := a
	fmt.Printf("a: %+v and b:%+v\n", a, b)

	b.next = &Listnode{val: 1, next: &Listnode{}}
	fmt.Printf("a: %+v and b:%+v\n", a, b)

	b = b.next
	fmt.Printf("a: %+v and b:%+v\n", a, b)

	b.next = b.next.next
	fmt.Printf("a: %+v and b:%+v\n", a, b)

	// p= 0xc000092010
	// q= 0xc000092010
	// p= 0xc000092010
	// q= 0xc000092030

	var i int = 42
	p := &i
	q := p
	fmt.Println("p=", p)
	fmt.Println("q=", q)

	var j int = 7
	q = &j
	fmt.Println("p=", p)
	fmt.Println("q=", q)

	// FOLLOWING CAUSES RUNTIME panic: runtime error: index out of range [0] with length 0

	// var result []int
	// result[0] = 1

	// var results [][]int
	// results[0] = append(results[0], 1)

	// CORRECT WAY
	// OUTPUT results [[0]] result[0]

	var result []int
	result = append(result, 0)

	var results [][]int
	results = append(results, result)

	fmt.Println("results", results)
	fmt.Println("result", result)

	// ALTERNATIVE
	// OUTPUT: resultsAlt [[0]] resultAlt [0]
	resultsAlt := make([][]int, 1)
	resultAlt := make([]int, 1)

	resultAlt[0] = 0
	resultsAlt[0] = resultAlt

	fmt.Println("resultsAlt", resultsAlt)
	fmt.Println("resultAlt", resultAlt)

	// Trees

	// Form the binary tree
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	fmt.Printf("Root of the tree is %+v \n", root)

	// Create a temp variable pointing to the root
	curr := root
	curr = curr.Right
	// Progress along the right subtree path
	fmt.Printf("Curr of the tree  after progressing to the right is %+v \n", curr)

	// Keep the root intact
	fmt.Printf("Root of the tree is %+v \n", root)

	root = root.Right
	fmt.Printf("Root of the tree using original reference is %+v \n", root)

	// Output

	//	Root of the tree is &{Val:3 Left:0xc0000b20d8 Right:0xc0000b2108}
	// Curr of the tree  after progressing to the right is &{Val:4 Left:0xc0000b2120 Right:0xc0000b2138}
	// Root of the tree is &{Val:3 Left:0xc0000b20d8 Right:0xc0000b2108}
	// Root of the tree is &{Val:4 Left:0xc0000b2120 Right:0xc0000b2138}
}
