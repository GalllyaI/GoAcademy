package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	lastElement string
	count       int
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for cp.count < times*2-1 {
			cp.Lock()
			if cp.lastElement == "Bar" {
				fmt.Print("Foo")
				cp.lastElement = "Foo"
				cp.count++
			}
			cp.Unlock()
		}
	}()

}

func (cp *ConcurrentPrinter) PrintBar(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for cp.count < times*2 {
			cp.Lock()
			if cp.lastElement == "Foo" {
				fmt.Print("Bar")
				cp.lastElement = "Bar"
				cp.count++
			}
			cp.Unlock()
		}
	}()

}

func main() {
	times := 10
	cp := &ConcurrentPrinter{lastElement: "Bar"}
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.Wait()
}
