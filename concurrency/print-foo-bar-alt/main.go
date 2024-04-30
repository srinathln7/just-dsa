package main

import (
	"fmt"
	"sync"
)

type FooBar struct {
	n  int
	ch chan int

	wg sync.WaitGroup
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:  n,
		ch: make(chan int),

		wg: sync.WaitGroup{},
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.ch <- 0
		<-fb.ch
	}

	fb.wg.Done()
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.ch <- 0
	}

	fb.wg.Done()
}

func printFoo() {
	print("foo")
}

func printBar() {
	print("bar")
}

func main() {
	fb := NewFooBar(2)

	fb.wg.Add(2)
	go fb.Foo(printFoo)
	go fb.Bar(printBar)

	fb.wg.Wait()
	fmt.Println()
}
