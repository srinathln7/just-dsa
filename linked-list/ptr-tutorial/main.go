package main

import "fmt"

type Listnode struct {
	val  int
	next *Listnode
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
}
