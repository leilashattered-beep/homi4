package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync"
)

func h(n int, ch <-chan string, w *sync.WaitGroup) {
	defer w.Done()
	for p := range ch {
		f, err := os.Open(p)
		if err != nil {
			fmt.Println(p, "-> ошибка открытия:", err)
			continue
		}

		m := md5.New()
		if _, err := io.Copy(m, f); err != nil {
			fmt.Println(p, "-> ошибка чтения:", err)
			f.Close()
			continue
		}
		f.Close()

		fmt.Printf("%s -> %x\n", p, m.Sum(nil))
	}
}

func main() {
	l := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	ch := make(chan string, len(l))
	var w sync.WaitGroup
	for i := 1; i <= 3; i++ {
		w.Add(1)
		go h(i, ch, &w)
	}

	for _, p := range l {
		ch <- p
	}
	close(ch)

	w.Wait()
}
