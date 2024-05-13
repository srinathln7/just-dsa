package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	CHF
)

var exchange = map[Currency]map[Currency]float64{
	USD: {USD: 1.00, EUR: 0.92, GBP: 0.79, CHF: 0.89},
	EUR: {USD: 1.09, EUR: 1.00, GBP: 0.86, CHF: 0.96},
	GBP: {USD: 1.28, EUR: 1.19, GBP: 1.00, CHF: 1.13},
	CHF: {USD: 1.13, EUR: 1.04, GBP: 0.88, CHF: 1.00},
}

func main() {

	// Key Idea: For an arbitrage opportunity to be present we need to check if there is a triangular cycle resulting in the product of weights greater than 1.0
	// We can modify Floyd Warshall algorithm (which in the essence is used to compute the shortest distance between all pair of vertices) to compute the product
	// from i->j, j->k, k->i and check if that results in 1.0

	currencies := []Currency{USD, EUR, GBP, CHF}
	n := len(currencies)

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if i == k {
				continue
			}
			for j := 0; j < n; j++ {
				if i == j || j == k {
					continue
				}

				rate1 := exchange[currencies[i]][currencies[j]]
				rate2 := exchange[currencies[j]][currencies[k]]
				rate3 := exchange[currencies[k]][currencies[i]]
				pdt := rate1 * rate2 * rate3
				if pdt > 1.01 {
					fmt.Printf("arbitrage @ %d -> %d -> %d -> %d  with rate %f \n", currencies[i], currencies[j], currencies[k], currencies[i], pdt)
				}
			}
		}
	}

}
