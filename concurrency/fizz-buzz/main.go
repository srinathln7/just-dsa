package main

import (
	"fmt"
	"sync"
)

type FizzBuzz struct {
	wg sync.WaitGroup
	mu sync.Mutex

	i int
	n int

	result []interface{}
}

func initFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{
		wg: sync.WaitGroup{},
		mu: sync.Mutex{},
		n:  n,
		i:  1,
	}
}

func (fb *FizzBuzz) printFizz() {
	for fb.i <= fb.n {
		if fb.i%3 == 0 && fb.i%5 != 0 {
			fb.mu.Lock()
			fb.result = append(fb.result, "fizz")
			fb.i++
			fb.mu.Unlock()
		}
	}

	fb.wg.Done()
}

func (fb *FizzBuzz) printBuzz() {
	for fb.i <= fb.n {
		if fb.i%3 != 0 && fb.i%5 == 0 {
			fb.mu.Lock()
			fb.result = append(fb.result, "buzz")
			fb.i++
			fb.mu.Unlock()
		}
	}

	fb.wg.Done()
}

func (fb *FizzBuzz) printFizzBuzz() {
	for fb.i <= fb.n {
		if fb.i%15 == 0 {
			fb.mu.Lock()
			fb.result = append(fb.result, "fizzbuzz")
			fb.i++
			fb.mu.Unlock()
		}
	}

	fb.wg.Done()
}

func (fb *FizzBuzz) printNumber() {
	for fb.i <= fb.n {
		if fb.i%3 != 0 && fb.i%5 != 0 {
			fb.mu.Lock()
			fb.result = append(fb.result, fb.i)
			fb.i++
			fb.mu.Unlock()
		}
	}

	fb.wg.Done()
}

func main() {

	n := 15
	fb := initFizzBuzz(n)

	fb.wg.Add(4)
	go fb.printFizz()
	go fb.printBuzz()
	go fb.printFizzBuzz()
	go fb.printNumber()
	fb.wg.Wait()

	fmt.Println(fb.result)
}
