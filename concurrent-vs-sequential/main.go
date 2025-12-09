package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	const n = 20

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Goroutines before start:", runtime.NumGoroutine())
	start := time.Now()

	gen := generator(ctx, n)
	sq := square(ctx, gen)
	fmt.Println("Goroutines after end:", runtime.NumGoroutine())
	results := collect(ctx, sq)

	elapsed := time.Since(start)

	fmt.Printf("Results (%d items): %v\n", len(results), results)
	fmt.Println("Elapsed (concurrent pipeline):", elapsed)

	// sequential code
	// const n = 20
	// start := time.Now()
	// results := doAllSequential(n)
	// elapsed := time.Since(start)

	// fmt.Printf("Results (%d items): %v\n", len(results), results)
	// fmt.Println("Elapsed (sequential):", elapsed)
}
