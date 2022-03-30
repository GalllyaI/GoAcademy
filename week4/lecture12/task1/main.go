package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {

	channel := make(chan string, bufferLimit)
	result := make(chan string)

	go func() {

		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case channel <- data:

			case <-ticker.C:
				for i := 0; i <= len(channel); i++ {
					result <- <-channel
				}
			}
		}
	}()

	return result
}

func main() {

	out := generateThrottled("foo", 2, 2*time.Second)

	for f := range out {

		log.Println(f)
	}
}
