package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPUs := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPUs)
}

