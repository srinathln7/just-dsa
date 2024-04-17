package main

type FooBar struct {
	n  int
	ch chan int
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:  n,
		ch: make(chan int),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.ch <- 0
		<-fb.ch
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.ch <- 0
	}
}
