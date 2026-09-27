package main

import "fmt"

func w(n int, j <-chan int, r chan<- int) {
	for x := range j {
		r <- x * x
	}
}

func main() {
	j := make(chan int, 10)
	r := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go w(i, j, r)
	}

	for i := 1; i <= 10; i++ {
		j <- i
	}
	close(j)

	for i := 1; i <= 10; i++ {
		fmt.Println(<-r)
	}
}
