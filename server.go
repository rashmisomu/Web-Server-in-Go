package main
//10,5,20
//40,25,100 some req not ,if client req 75 all pass
//35,20,70
//20,10,200 ok 
import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	// ch1 is a channel to send request timestamps to workers
	ch1 = make(chan int64,2)
	// workers is the number of concurrent workers to spawn
	workers = 10
	wg      sync.WaitGroup
)

// workerFxn is a function that runs in a goroutine and processes
func workerFxn() {
	defer wg.Done()
	for timestamp := range ch1 {
		log.Printf("Worker received+ timestamp: %d\n", timestamp)
		// Simulate some work by sleeping for 5 seconds
		time.Sleep(time.Second * 5)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	time.Sleep(5 * time.Millisecond)
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	go func() {
		ch1 <- time.Now().UnixNano()
	}()
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current time is: %s", currentTime)
	go func() {
		ch1 <- time.Now().UnixNano()
	}()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Accepted"))

	// Send the request timestamp to the channel
	go func() {
		ch1 <- time.Now().UnixNano()
	}()
}

func main() {
	// Spawn the workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go workerFxn()
	}

	// Register the handler functions
	file := http.FileServer(http.Dir("./static"))
	http.Handle("/", file)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	fmt.Println("Starting server on port 8089")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

	// Wait for all workers to finish
	close(ch1)
	wg.Wait()
}
