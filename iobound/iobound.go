package iobound

import (
	"log"
	"net/http"
	"runtime"
	"sync"
)

// Make HTTP GET requests to a set of URLs sequentially
func getURLSequential(urls []string) {
	for _, url := range urls {
		if _, err := http.Get(url); err != nil {
			log.Println(err)
		}
	}
}

// Make HTTP GET requests to a set of URLs concurrently
func getURLConcurrent(urls []string) {
	goroutines := runtime.NumCPU()
	var waitGroup sync.WaitGroup
	waitGroup.Add(goroutines)

	ch := make(chan string, len(urls))
	for _, u := range urls {
		ch <- u
	}
	close(ch)

	for i := 0; i < goroutines; i++ {
		go func() {
			for u := range ch {
				if _, err := http.Get(u); err != nil {
					log.Println(err)
				}
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}