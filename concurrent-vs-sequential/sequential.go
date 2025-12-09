package main

import (
	"time"
)

// doAllSequential runs the same logic but sequentially in one goroutine.
func doAllSequential(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		// generator work
		time.Sleep(20 * time.Millisecond) // produce i

		// square work
		sq := i * i
		time.Sleep(40 * time.Millisecond) // compute

		// filter & collect
		if sq%2 == 0 {
			res = append(res, sq)
		}
		time.Sleep(10 * time.Millisecond) // small I/O/write
	}
	return res
}
