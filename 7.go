package main

import (
	"fmt",
	"sync"
)

func manager(i chan int, g chan chan int) {
	c := 0
	for {
		select {
		case n := <-i:
			c += n
		case r := <-g:
			r <- c
		}
	}
}

func main() {
	i := make(chan int)
	g := make(chan chan int)

	go manager(i,g)

	var a sync.WaitGroup
	for k := 0; k < 100; k++ {
		a.Add(1)
		go func() {
			defer a.Done()
			i <- 1
		}()
	}
	a.Wait()

	r := make(chan int)
	g <- r 
	fmt.Println(<-r)

}

