package main

import "container/heap"

type Order struct {
	price  int
	amount int
}

type MinHeap []Order

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Order)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	o := old[n-1]
	*h = old[:n-1]
	return o
}
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h MinHeap) Top() Order         { return h[0] }
func (h *MinHeap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

type MaxHeap []Order

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Order)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	o := old[n-1]
	*h = old[:n-1]
	return o
}
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h MaxHeap) Top() Order         { return h[0] }
func (h *MaxHeap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

const MOD = 1e9 + 7

func getNumberOfBacklogOrders(orders [][]int) int {

	// KEY IDEA: To execute Buy order => Look at the sell order's backlog and compare it with the min available in
	// the sell order backlog. All sell orders need to be stored in => Min Heap
	// Vice versa, to execute Sell order => Look at the buy order's backlog and compare it with the max available price in the backlog
	// All buy orders need to be stored in => Max Heap

	buyHeap := &MaxHeap{}
	sellHeap := &MinHeap{}

	for _, ord := range orders {
		price, amount := ord[0], ord[1]

		switch ord[2] {
		// Buy Order
		case 0:
			// Try matching the buy order as long as this condition holds
			for amount > 0 && (*sellHeap).Len() > 0 && (*sellHeap).Top().price <= price {
				minSellOrder := heap.Pop(sellHeap).(Order)
				sellAmount := minSellOrder.amount
				txAmount := min(amount, sellAmount)
				minSellOrder.amount -= txAmount
				amount -= txAmount

				// Partially filled sell order
				if minSellOrder.amount > 0 {
					heap.Push(sellHeap, minSellOrder)
				}
			}

			// If no sell orders are available or partially-filled buy order add it to the buy heap
			if amount > 0 {
				heap.Push(buyHeap, Order{price: price, amount: amount})
			}

		// Sell order
		case 1:
			// Try matching the sell order as long as this condition holds
			for amount > 0 && (*buyHeap).Len() > 0 && (*buyHeap).Top().price >= price {
				maxBuyOrder := heap.Pop(buyHeap).(Order)
				buyAmount := maxBuyOrder.amount
				txAmount := min(amount, buyAmount)
				maxBuyOrder.amount -= txAmount
				amount -= txAmount

				// Partially filled buy order
				if maxBuyOrder.amount > 0 {
					heap.Push(buyHeap, maxBuyOrder)
				}
			}

			// If no buy orders are available or partially-filled sell order add it to the sell heap
			if amount > 0 {
				heap.Push(sellHeap, Order{price: price, amount: amount})
			}
		}
	}

	var backlog int
	for _, buyBacklog := range *buyHeap {
		backlog += buyBacklog.amount
		if backlog > MOD {
			backlog %= MOD
		}
	}

	for _, sellBacklog := range *sellHeap {
		backlog += sellBacklog.amount
		if backlog > MOD {
			backlog %= MOD
		}
	}

	return backlog
}
