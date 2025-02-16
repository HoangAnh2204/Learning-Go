package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, ch chan<- string) {

	workTime := rand.Intn(5) + 1
	// Thời gian làm việc ngẫu nhiên 1 đến 5 giây
	time.Sleep(time.Duration(workTime) * time.Second)

	// Sau khi hoàn thành công việc, gửi kết quả vào channel
	ch <- fmt.Sprintf("Worker %d hoàn thành công việc sau %d giây", id, workTime)
}

// Hàm xử lý công việc với timeout
func manageJob(id int, ch chan<- string, timeout chan<- string) {
	// Tạo kênh kết quả từ worker
	resultCh := make(chan string)

	// Chạy worker Goroutine
	go worker(id, resultCh)

	// Sử dụng select để chờ kết quả hoặc timeout
	select {
	case res := <-resultCh:
		// Nhận kết quả từ worker
		ch <- res
	case <-time.After(3 * time.Second): // Timeout sau 3 giây
		// Nếu công việc quá lâu, gửi thông báo timeout
		timeout <- fmt.Sprintf("Worker %d timeout!", id)
	}
}

func main() {
	// Các kênh để nhận kết quả và thông báo timeout
	resultCh := make(chan string)
	timeoutCh := make(chan string)

	// Bắt đầu quản lý công việc
	for i := 1; i <= 5; i++ {
		go manageJob(i, resultCh, timeoutCh)
	}

	// Chờ đợi và xử lý kết quả hoặc thông báo timeout
	for i := 1; i <= 5; i++ {
		select {
		case result := <-resultCh:
			fmt.Println(result)
		case timeoutMsg := <-timeoutCh:
			fmt.Println(timeoutMsg)
		}
	}
}
