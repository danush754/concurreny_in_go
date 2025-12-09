package main

import (
	"context"
	"time"
)

// generator sends numbers 0..n-1 and stops when ctx is done.
// it returns a read-only channel.
func generator(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
			// simulate work to produce each value (e.g., reading, computing)
			time.Sleep(20 * time.Millisecond)
		}
	}()
	return out
}

// square receives from in, squares, and sends to out until in is closed or ctx done.
func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v * v:
			}
			// simulate heavier CPU work
			time.Sleep(40 * time.Millisecond)
		}
	}()
	return out
}

// filterEven collects only even numbers from in and returns a slice of results.
func collect(ctx context.Context, in <-chan int) []int {
	var res []int
	for v := range in {
		select {
		case <-ctx.Done():
			return res
		default:
			// filter: only even squares
			if v%2 == 0 {
				res = append(res, v)
			}
			// simulate small I/O/write cost
			time.Sleep(10 * time.Millisecond)
		}
	}
	return res
}
