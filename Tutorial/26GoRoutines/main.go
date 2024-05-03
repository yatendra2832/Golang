package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"Test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://go.dev",
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Something Went Wrong in the Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
