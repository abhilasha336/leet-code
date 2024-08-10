package main

import (
	"fmt"
	"sync"
)

func main() {

	ev := make(chan int)
	od := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 15; i++ {
			if i%2 == 0 {
				ev <- i
			} else {
				od <- i
			}
		}

		close(ev)
		close(od)

	}()

	for {
		select {
		case v, ok := <-ev:
			if !ok {
				ev = nil

			} else {
				fmt.Println("even", v)

			}

		case m, ok := <-od:
			if !ok {
				od = nil
			} else {
				fmt.Println("odd", m)
			}

		}

		if od == nil && ev == nil {
			break
		}
	}

	wg.Wait()
}
