package main

import "fmt"

type Listnode struct {
	val  int
	next *Listnode
}

func main() {
	// 	a: &{val:0 next:<nil>} and b:&{val:0 next:<nil>}
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
}
