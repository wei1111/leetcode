package _chan

import (
	"fmt"
	"sync"
)

// OrderPrint
// golang顺序打印数字
// n个协程序，打印到m
func OrderPrint(n, m int) {
	var wg sync.WaitGroup
	ch := make(chan int)
	closeCh := make(chan int, n)
	num := 0
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-closeCh:
					return
				case <-ch:
					if num > m {
						for i := 0; i < n; i++ {
							closeCh <- num
						}
						return
					}
					fmt.Printf("%d\n", num)
					num++
					ch <- num
				}
			}
		}()
	}
	ch <- num
	wg.Wait()
	fmt.Println("done")
}
