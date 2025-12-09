package main

func generatorWithDone(done <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 0; i < 10; i++ {

			select {
			case <-done:
				return
			case out <- i:

			}

		}
	}()

	return out
}

func bufferedGeneratorWithDone(done <-chan struct{}) <-chan int {
	outChan := make(chan int, 10)

	go func() {
		defer close(outChan)

		for i := 0; i < 50; i++ {
			select {
			case <-done:
				return
			case outChan <- i:

			}

		}
	}()

	return outChan
}
