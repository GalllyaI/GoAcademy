package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func multiping(urls []string, c int) {

	processQueue := make(chan string, c)
	done := make(chan string)

	go func() {
		for _, urlToProcess := range urls {
			processQueue <- urlToProcess
			go func(url string) {
				time.Sleep(2 * time.Second)
				pingURL(url)
				<-processQueue
				done <- url
			}(urlToProcess)
		}

	}()

	for range urls {
		<-done
	}

}

func main() {

	var c int
	var url string
	flag.StringVar(&url, "url", "", "url to fetch")
	flag.IntVar(&c, "c", 2, "number concurrent urls")
	flag.Parse()

	if url == "" {
		flag.PrintDefaults()
		return
	}

	if c == 0 {
		flag.PrintDefaults()
		return
	}

	urls := flag.Args()
	multiping(urls, c)
}

// -c 3 -url /home https://www.ccc.eu https://www.obuvki.bg https://www.google.com https://www.yahoo.com https://google.com https://facebook.com
