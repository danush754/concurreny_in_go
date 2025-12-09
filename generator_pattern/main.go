package main

import (
	"fmt"
)

func main() {

	for value := range numberGenerator(5) {

		fmt.Println("val", value)
	}

	// generator test with done channel

	done := make(chan struct{})
	nums := generatorWithDone(done)

	for i := 0; i < 5; i++ {
		value, ok := <-nums
		if !ok {
			fmt.Println("channel closed")
			return
		}

		fmt.Println("value:", value)

	}

	close(done)

	for v := range nums {
		fmt.Println("extra values", v)
	}

	// generator with buffered channel

	doneForBuffered := make(chan struct{})

	result := bufferedGeneratorWithDone(doneForBuffered)
	for value := range result {
		if value == 10 {
			break
		}
		fmt.Println("valFromBuffered", value)
	}

	close(doneForBuffered)
}
