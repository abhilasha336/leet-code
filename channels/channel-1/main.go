package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	for val := range ch {
		fmt.Println("val", val)
	}
	wg.Wait()

}
