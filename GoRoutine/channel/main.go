package main

import (
	"fmt"
	"time"
)

const (
	numJobs   = 10 // Số lượng công việc
	numWorker = 3  // Số worker chạy đồng thời
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs { // Nhận công việc từ channel
		fmt.Printf("[Worker %d] Xử lý công việc %d\n", id, job)
		results <- fmt.Sprintf("Worker %d đã hoàn thành công việc %d", id, job)
	}
}

func main() {
	// Bỏ rand.Seed và rand.Intn vì không cần tính ngẫu nhiên nữa
	//rand.Seed(time.Now().UnixNano())

	jobs := make(chan int, numJobs)       // Channel chứa công việc
	results := make(chan string, numJobs) // Channel chứa kết quả

	// Khởi tạo worker Goroutines
	for i := 1; i <= numWorker; i++ {
		go worker(i, jobs, results)
	}

	// Đẩy công việc vào channel
	go func() {
		for j := 1; j <= numJobs; j++ {
			fmt.Printf("Đẩy công việc %d vào hàng đợi\n", j)
			jobs <- j
		}
		close(jobs) // Đóng jobs để workers biết không còn công việc
	}()

	// Nhận kết quả từ workers (không dùng vòng lặp cố định)
	go func() {
		for res := range results { // Đọc dữ liệu đến khi channel đóng
			fmt.Println(res)
		}
	}()

	// Đợi một chút để đảm bảo tất cả Goroutine hoàn thành trước khi thoát
	time.Sleep(1 * time.Second)
	fmt.Println("Tất cả công việc đã hoàn thành!")
}

// các kết quả mỗi lần chạy khác nhau là do có nhiều go routine chạy song song
// -> các công việc được xử lí theo thứ tự khác nhau trong mỗi lần chạy
