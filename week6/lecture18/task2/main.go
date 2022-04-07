package main

import (
	"fmt"
	"sync"

	"time"
)

func PrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			time.Sleep(sleep)
			if k%i == 0 {
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}

func goPrimesAndSleep(n int, sleep time.Duration) []int {
	var wg sync.WaitGroup
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			time.Sleep(sleep)
			if k%i == 0 {
				wg.Add(1)
				go func(vk int, vi int) {
					defer wg.Done()
					if vk == vi {
						res = append(res, vk)

					}
				}(k, i)
				break
			}
		}
	}
	wg.Wait()
	return res
}

func main() {
	result := goPrimesAndSleep(20, 1*time.Millisecond)

	for _, v := range result {
		fmt.Print(v, " ")

	}

	fmt.Println()

	r2 := PrimesAndSleep(20, 100*time.Millisecond)

	for _, v := range r2 {
		fmt.Print(v, " ")
	}
}
