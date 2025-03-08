package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	Value int
}

type Result struct {
	Value int
	Error error
}

var (
	sharedCounter int
	mu            sync.Mutex
)

func generateData(dataChan chan Data, numItems int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		dataChan <- Data{Value: rand.Intn(100)}

		mu.Lock()
		sharedCounter++
		mu.Unlock()
	}
	close(dataChan) // Tutup channel setelah selesai mengirim semua data
}

func processData(dataChan chan Data, resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataChan {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		if rand.Float64() < 0.1 {
			resultChan <- Result{Error: fmt.Errorf("random processing error")}
			continue
		}
		resultChan <- Result{Value: data.Value * rand.Intn(10)}

		mu.Lock()
		sharedCounter++
		mu.Unlock()
	}
}

func consumeResults(resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range resultChan {
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
		} else {
			fmt.Println("Result:", result.Value)
		}

		mu.Lock()
		sharedCounter++
		mu.Unlock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numItems := 50
	dataChan := make(chan Data, 10)
	resultChan := make(chan Result, 10)

	var wg sync.WaitGroup
	wg.Add(3)

	go generateData(dataChan, numItems, &wg)
	go processData(dataChan, resultChan, &wg)
	go consumeResults(resultChan, &wg)

	// Gunakan goroutine untuk menutup `resultChan` setelah semua worker selesai
	go func() {
		wg.Wait()
		close(resultChan) // Pastikan resultChan hanya ditutup setelah worker selesai
	}()

	wg.Wait() // Tunggu semua goroutine selesai
	fmt.Println("Program finished.")
	fmt.Println("Shared Counter:", sharedCounter)
}
