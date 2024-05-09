package main

type Price struct {
	val  int
	span int
}

type StockSpanner struct {
	stack []Price
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {

	// Key Idea: We need to find how many elements consecutive  to my left including myself are lesser than or equal to myself
	span := 1
	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1].val {
		top := this.stack[len(this.stack)-1]

		// Accumulate the span from the top
		span += top.span

		// Pop the element from the stack
		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, Price{val: price, span: span})
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
