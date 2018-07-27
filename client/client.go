package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	n := 400
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				panic(fmt.Errorf("%d %v", i, err))
			}
			if resp.StatusCode != 200 {
				panic(fmt.Errorf("not 200"))
			}
			resp.Body.Close()
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
