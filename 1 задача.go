package main

import (
	"fmt"
	"sync"
	"time"
)

func p(w *sync.WaitGroup) {
	defer w.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go p(&w)
	w.Wait()
	fmt.Printin("Гортуна завершилась")
}
