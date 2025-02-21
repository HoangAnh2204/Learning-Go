package main

import (
	"fmt"
	"sync"
	"time"
)

const workerCount = 3 // Giới hạn số goroutine chạy đồng thời

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d xử lý công việc %d\n", id, job)
		time.Sleep(1 * time.Second) // Giả lập xử lý lâu
	}
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Tạo worker
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Gửi công việc vào channel
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs) // Đóng channel để thông báo không còn công việc mới

	wg.Wait() // Đợi tất cả worker hoàn thành
	fmt.Println("Hoàn thành tất cả công việc!")
}
