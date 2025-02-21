package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const workerCount = 3

// Sinh số ngẫu nhiên gửi vào channel
func generateNumbers(n int, out chan<- int) {
	for i := 0; i < n; i++ {
		num := rand.Intn(100)
		fmt.Println("Tạo số:", num)
		out <- num
	}
	close(out)
}

// Worker xử lý số
func worker(id int, in <-chan int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		result := fmt.Sprintf("Worker %d xử lý: %d -> %d", id, num, num*2)
		out <- result
	}
}

// Tổng hợp kết quả từ nhiều worker
func fanIn(in <-chan string, done chan<- bool) {
	for msg := range in {
		fmt.Println(msg)
	}
	done <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := make(chan int, 10)
	results := make(chan string, 10)
	done := make(chan bool)

	// Tạo số
	go generateNumbers(10, nums)

	// Tạo worker xử lý số
	var wg sync.WaitGroup
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, nums, results, &wg)
	}

	// Goroutine đóng channel kết quả khi tất cả worker hoàn thành
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-in: Nhận dữ liệu từ worker và in kết quả
	go fanIn(results, done)

	// Chờ tất cả hoàn thành
	<-done
	fmt.Println("Hoàn thành tất cả công việc!")
}
