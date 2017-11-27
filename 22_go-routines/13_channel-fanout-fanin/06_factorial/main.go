package main

import (
	"fmt"
	"math/big"
	"sync"
)

func main() {
	n := 1000
	bc := fact(n)

	cs := make([]chan *big.Int, n/2)
	for i := 0; i < n/2; i++ {
		cs[i] = calc(bc)
	}

	// fan in
	for n := range merge(cs...) {
		fmt.Println(n)
	}
}

func fact(n int) chan int {
	out := make(chan int)
	go func() {
		for i := n; i > 0; i -= 2 {
			x := i * (i - 1)
			if x == 0 {
				x++
			}
			out <- x
		}
		close(out)
	}()
	return out
}

func calc(in chan int) chan *big.Int {
	out := make(chan *big.Int)
	go func() {
		for n := range in {
			res := big.NewInt(int64(n))
			if res == nil {
				fmt.Println("HAYOH")
			}
			out <- res
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan *big.Int) chan *big.Int {
	out := make(chan *big.Int)
	res := big.NewInt(1)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan *big.Int) {
			mutex.Lock()
			in := <-ch
			if in == nil {
				fmt.Println("HAYOH")
			}
			res.Mul(res, in)
			mutex.Unlock()
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		out <- res
		close(out)
	}()
	return out
}
