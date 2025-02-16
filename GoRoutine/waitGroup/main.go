package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Khai báo WaitGroup để đồng bộ hóa goroutines
	var mu sync.Mutex     // Khai báo Mutex để bảo vệ việc truy cập vào biến chung
	counter := 0          // Biến chia sẻ giữa các goroutine

	fmt.Println("Ứng dụng bắt đầu")

	// Khởi tạo và chạy 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Thêm một goroutine vào WaitGroup
		go func(id int) {
			defer wg.Done() // Khi goroutine hoàn thành, giảm bộ đếm của WaitGroup

			//chờ để đảm bảo các goroutine khác có thể chạy
			time.Sleep(time.Millisecond * time.Duration(id*100))

			mu.Lock() // Đảm bảo chỉ một goroutine có thể truy cập và thay đổi counter tại một thời điểm
			counter++ // Thực hiện công việc chia sẻ: tăng giá trị counter
			fmt.Printf("Goroutine %d: counter = %d\n", id, counter)
			mu.Unlock() // Giải phóng mutex để các goroutine khác có thể truy cập

		}(i) // Truyền id của goroutine vào closure để biết đâu là goroutine nào
	}

	// Goroutine chính không cần đợi
	fmt.Println("Đang chạy goroutines...")

	// Đợi tất cả goroutines hoàn thành
	wg.Wait()

	// Khi tất cả goroutines hoàn thành in ra giá trị của counter
	fmt.Printf("Ứng dụng kết thúc, giá trị final của counter: %d\n", counter)
}
