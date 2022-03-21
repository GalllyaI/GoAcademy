package main

import (
	"fmt"
	"sync"
)

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8}
	evenChan := processEven(inputs)
	oddChan := processOdd(inputs)

	// Can you please explain why it is not reciving all the values from the channel?
	// The length of the chanel is 4 and it prints only two values

	// fmt.Println(len(evenChan))
	// for range evenChan {
	// 	j, more := <-evenChan
	// 	if more {
	// 		fmt.Print(j, " ")
	// 	}
	// }

	for i := range evenChan {
		fmt.Print(i, " ")
	}

	fmt.Println()

	for i := range oddChan {
		fmt.Print(i, " ")
	}

}

func processEven(inputs []int) chan int {
	var wg sync.WaitGroup

	result := make(chan int, len(inputs))

	evenCh := make(chan int)

	func() {
		for _, n := range inputs {
			go func(n int) {
				evenCh <- n
				if n%2 == 0 {
					result <- n
				}
			}(n)
			<-evenCh
		}

		wg.Wait()

	}()

	close(evenCh)
	close(result)

	return result
}

func processOdd(inputs []int) chan int {
	var wg sync.WaitGroup

	result := make(chan int, len(inputs))
	oddCh := make(chan int)

	func() {
		for _, n := range inputs {
			go func(n int) {
				oddCh <- n
				if n%2 != 0 {
					result <- n
				}
			}(n)
			<-oddCh
		}

		wg.Wait()

	}()

	close(oddCh)
	close(result)

	return result
}
