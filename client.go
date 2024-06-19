package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const (
	// Base URL of the server
	baseURL = "http://localhost:8089"
	// Total number of requests to send
	totalRequests = 100)

func main() {
	// Create a WaitGroup to wait for all requests to finish
	var wg sync.WaitGroup

	// Send totalRequests number of requests to /hello
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go sendHelloRequest(i, &wg)
	}

	// Send totalRequests number of requests to /form
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go sendFormRequest(i, &wg)
	}

	// Send totalRequests number of requests to /time
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go sendTimeRequest(i, &wg)
	}

	// Wait for all requests to finish
	wg.Wait()

	fmt.Println("All requests completed.")
}

func sendHelloRequest(requestID int, wg *sync.WaitGroup) {
	defer wg.Done()
	url := baseURL + "/hello"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request %d to /hello failed: %v\n", requestID, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Request %d to /hello completed: %s\n", requestID, resp.Status)
}

func sendFormRequest(requestID int, wg *sync.WaitGroup) {
	defer wg.Done()
	url := baseURL + "/form"
	data := "name=JohnDoe&address=123MainStreet"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Printf("Request %d to /form failed: %v\n", requestID, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Request %d to /form completed: %s\n", requestID, resp.Status)
}

func sendTimeRequest(requestID int, wg *sync.WaitGroup) {
	defer wg.Done()
	url := baseURL + "/time"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request %d to /time failed: %v\n", requestID, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Request %d to /time completed: %s\n", requestID, resp.Status)
}
