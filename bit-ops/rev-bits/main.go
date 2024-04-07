package main

func reverseBits(num uint32) uint32 {
	var result uint32
	// Since it is a 32-bit integer, we need to left shift by 32 times
	for i := 0; i < 32; i++ {
		result <<= 1      // Left shift the result by 1
		result |= num & 1 // Set the result as the rightmost bit of num
		num >>= 1         // divide the num by 2
	}
	return result
}
