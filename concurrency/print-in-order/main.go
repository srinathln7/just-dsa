package main

import "fmt"

type Foo struct {
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
}

func (f *Foo) first() {
	fmt.Println("first")
	<-f.ch1
}

func (f *Foo) second() {
	f.ch1 <- struct{}{}
	fmt.Println("second")
	<-f.ch2
}

func (f *Foo) third() {
	f.ch2 <- struct{}{}
	fmt.Println("third")
	<-f.ch3
}

func main() {
	f := &Foo{
		ch1: make(chan struct{}),
		ch2: make(chan struct{}),
		ch3: make(chan struct{}),
	}

	nums := []int{1, 2, 3}
	for idx, thread := range nums {
		switch idx {
		case 0:
			fmt.Printf("Thread %d  calling first()\n", thread)
			go f.first()
		case 1:
			fmt.Printf("Thread %d  calling second()\n", thread)
			go f.second()
		case 2:
			fmt.Printf("Thread %d  calling third()\n", thread)
			go f.third()
		}
	}

	// Signal third channel to start
	f.ch3 <- struct{}{}

}
