package main

import (
	"fmt"
	"strconv"
	"sync"
)

type ZEO struct {
	wg sync.WaitGroup
	mu sync.Mutex

	isZero bool
	i      int
	n      int
	result string
}

func initZEO(n int) *ZEO {
	return &ZEO{
		wg: sync.WaitGroup{},
		mu: sync.Mutex{},

		isZero: true,
		n:      n,
	}
}

func (zeo *ZEO) zero() {
	for {
		zeo.mu.Lock()

		if zeo.i > zeo.n {
			fmt.Println(zeo.i)
			zeo.mu.Unlock()
			break
		}
		if zeo.isZero {
			//fmt.Println("calling zero")
			zeo.result += "0"
			zeo.i++
			zeo.isZero = false
		}
		zeo.mu.Unlock()
	}

	zeo.wg.Done()
}

func (zeo *ZEO) one() {
	for {
		zeo.mu.Lock()

		if zeo.i > zeo.n {
			zeo.mu.Unlock()
			break
		}
		if !zeo.isZero && zeo.i%2 != 0 {
			//fmt.Println("calling one")
			zeo.result += strconv.Itoa(zeo.i)
			zeo.isZero = true
		}
		zeo.mu.Unlock()
	}

	zeo.wg.Done()
}

func (zeo *ZEO) two() {
	for {
		zeo.mu.Lock()

		if zeo.i > zeo.n {
			zeo.mu.Unlock()
			break
		}
		if !zeo.isZero && zeo.i%2 == 0 {
			//fmt.Println("calling two")
			zeo.result += strconv.Itoa(zeo.i)
			zeo.isZero = true
		}
		zeo.mu.Unlock()
	}

	zeo.wg.Done()
}

func main() {

	n := 50
	zeo := initZEO(n)

	zeo.wg.Add(3)
	go zeo.zero()
	go zeo.one()
	go zeo.two()

	zeo.wg.Wait()

	fmt.Println(zeo.result[:len(zeo.result)-1])
}
