package main

func numberGenerator(maxVal int) <-chan int {

	outChan := make(chan int) //unbuffered channel

	go func() {
		defer close(outChan)
		for i := 0; i < maxVal; i++ {
			outChan <- i

		}
	}()

	return outChan
}
