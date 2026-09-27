package main

import (
	"fmt"
	"math/rand"
	"time"
)

func s(name string, r chan<- string) {
	d := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(d)
	r <- fmt.Sprintf("результат из %s (за %v)", name, d)
}

func main() {
	l := []string{"источник1", "источник2", "источник3"}
	r := make(chan string, len(l))

	for _, x := range l {
		go s(x, r)
	}
	f := <-r
	fmt.Println("Первый успешный результат:", f)
}
