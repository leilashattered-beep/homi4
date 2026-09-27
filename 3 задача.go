package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 15; i++ {
		fmt.Println("запрос", i, "выполнен в", time.Now().Format("15:04:05.000"))
	}
}
