package main

import "fmt"

func s(n int, ch <-chan int, d chan<- struct{}) {
	for x := range ch {
		fmt.Printf("сервис %d обработал запрос %d\n", n, x)
	}
	d <- struct{}{}
}

func main() {
	const n = 3
	const total = 10

	c := make([]chan int, n)
	d := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		c[i] = make(chan int)
		go s(i+1, c[i], d)
	}

	for i := 1; i <= total; i++ {
		t := (i - 1) % n
		c[t] <- i
	}

	for i := 0; i < n; i++ {
		close(c[i])
	}

	for i := 0; i < n; i++ {
		<-d
	}
}
