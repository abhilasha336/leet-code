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
	recieve(ev, od)

	wg.Wait()
}

func recieve(evr, odr <-chan int) {
	for {
		select {
		case v, ok := <-evr:
			if !ok {
				evr = nil

			} else {
				fmt.Println("even", v)

			}

		case m, ok := <-odr:
			if !ok {
				odr = nil
			} else {
				fmt.Println("odd", m)
			}

		}

		if odr == nil && evr == nil {
			break
		}
	}

}
