package main

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	bits      []byte
	hashFuncs []func(data []byte) int
}

func calcoptimalParams(n int, p float64) (m, k int) {

	// Calc. `m` and `k` based on the derived formula
	// m = - n * ln(p) /(ln(2))^2  => bit size
	// p = m/n * ln(2) => num of hash functions

	m = int(math.Ceil(-float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)))
	k = int(math.Round(float64(m) * math.Ln2 / float64(n)))
	return m, k
}

func genHashFunc(seed, size int) func([]byte) int {
	return func(data []byte) int {
		h1, h2 := murmur3.Sum128WithSeed(data, uint32(seed))
		//MISTAKE: return int((h1 + h2)) % size
		return int((h1 + h2) % uint64(size))
	}
}

func newBloomFilter(n int, p float64) *BloomFilter {
	m, k := calcoptimalParams(n, p)

	hashFuncs := make([]func([]byte) int, k)
	for i := 0; i < k; i++ {
		hashFuncs[i] = genHashFunc(i, m)
	}

	return &BloomFilter{
		bits:      make([]byte, (m+7)/8),
		hashFuncs: hashFuncs,
	}
}

func (bf *BloomFilter) Add(data []byte) {
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)
		bf.bits[idx/8] |= 1 << (idx % 8)
	}
}

func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hashFunc := range bf.hashFuncs {
		idx := hashFunc(data)

		if bf.bits[idx/8]&(1<<(idx%8)) == 0 {
			return false
		}
	}

	return true
}

func main() {
	n, p := 10000, 0.01

	obf := newBloomFilter(n, p)
	datas := [][]byte{
		[]byte("apple"),
		[]byte("banana"),
		[]byte("chiros"),
		[]byte("doctor"),
	}

	for _, data := range datas {
		obf.Add(data)
	}

	tests := [][]byte{
		[]byte("apple"),
		[]byte("eagle"),
		[]byte("hen"),
		[]byte("banana"),
		[]byte("chiros"),
		[]byte("doctor"),
		[]byte("mango"),
	}

	for _, test := range tests {
		if obf.Contains(test) {
			fmt.Printf("%s might be in the set\n", string(test))
		} else {
			fmt.Printf("%s definitely not in the set\n", string(test))
		}
	}
}
