package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	TOTAL_PHILOSOPHERS = 5
	TOTAL_FORKS        = 5
	TOTAL_ITERATIONS   = 100
)

type Philosopher struct {
	id        int
	leftFork  *Fork
	rightFork *Fork
}

func newPhilosopher(id int, leftFork, rightFork *Fork) *Philosopher {
	return &Philosopher{
		id:        id,
		leftFork:  leftFork,
		rightFork: rightFork,
	}
}

type Fork struct {
	id int
	mu sync.Mutex
}

func newFork(id int) *Fork {
	return &Fork{
		id: id,
		mu: sync.Mutex{},
	}
}

func (p *Philosopher) pickLeftFork() {
	p.leftFork.mu.Lock()
}

func (p *Philosopher) putLeftFork() {
	p.leftFork.mu.Unlock()
}

func (p *Philosopher) pickRightFork() {
	p.rightFork.mu.Lock()
}

func (p *Philosopher) putRightFork() {
	p.rightFork.mu.Unlock()
}

func (p *Philosopher) pickUpForks() {
	p.pickLeftFork()
	p.pickRightFork()
}

func (p *Philosopher) putDownForks() {
	p.putLeftFork()
	p.putRightFork()
}

func (p *Philosopher) eat() {
	fmt.Printf("Philosopher%d is eating now...\n", p.id)
	time.Sleep(100 * time.Millisecond)
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher%d is thinking now...\n", p.id)
	time.Sleep(100 * time.Millisecond)
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < TOTAL_ITERATIONS; i++ {
		p.think()

		p.pickUpForks()
		p.eat()
		p.putDownForks()
	}

}

func main() {

	// Create 5 forks
	forks := make([]*Fork, TOTAL_FORKS)
	for i := 0; i < len(forks); i++ {
		forks[i] = newFork(i)
	}

	// Create 5 philosophers
	philosophers := make([]*Philosopher, TOTAL_PHILOSOPHERS)

	rightFork := forks[TOTAL_FORKS-1]
	for i := 0; i < len(philosophers); i++ {
		philosophers[i] = newPhilosopher(i, forks[i], rightFork)

		// Right becomes left fork for the next person in circle
		rightFork = forks[i]
	}

	// IMPORTANT: Since we do not care who gets to eat first i.e. order of computation does not matter,
	// we can use waitgroups here

	var wg sync.WaitGroup
	wg.Add(TOTAL_PHILOSOPHERS)

	// Let the philosophers dine concurrently
	for _, philosopher := range philosophers {
		go philosopher.dine(&wg)
	}

	wg.Wait()

	fmt.Println("ALL PHILOSOPHERS FINISHED DINING!!!")
}
