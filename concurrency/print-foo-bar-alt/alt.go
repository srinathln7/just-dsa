package main

import (
	"fmt"
)

type foobar struct {
	n      int
	fooCh  chan struct{}
	barCh  chan struct{}
	doneCh chan struct{}
}

func Newfoobar(n int) *foobar {
	return &foobar{
		n:      n,
		fooCh:  make(chan struct{}, 1), // Buffered channel of size 1
		barCh:  make(chan struct{}, 1), // Buffered channel of size 1
		doneCh: make(chan struct{}, 0), // Unbuffered channel
	}
}

func (fb *foobar) Foo() {
	for i := 0; i < fb.n; i++ {

		// Wait until you get the signal from `Bar`
		// Initial signal is sent from `main`
		<-fb.fooCh

		fmt.Print("foo")

		// Send signal to Bar
		fb.barCh <- struct{}{}
	}
}

func (fb *foobar) Bar() {
	for i := 0; i < fb.n; i++ {

		// Wait until you get the signal from `Foo`
		<-fb.barCh

		fmt.Print("bar")

		// Send signal to Foo
		fb.fooCh <- struct{}{}
	}

	fb.doneCh <- struct{}{}
}

func alt() {
	obj := Newfoobar(2)

	// Signal buffered `fooCh` to start
	obj.fooCh <- struct{}{}

	go obj.Foo()
	go obj.Bar()

	// Wait until done channel is unblocked
	<-obj.doneCh

	fmt.Println()
}
