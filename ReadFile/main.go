package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// số lượng go routine chạy
const numWorkers = 4

// hàm duyệt file trong thư mục
func listFiles(dir string, fileChan chan<- string) {
	defer close(fileChan)           // Đóng channel khi hoàn thành
	entries, err := os.ReadDir(dir) // đọc danh sách trong đường dẫn dir
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() { // Chỉ xử lý file, bỏ qua thư mục con
			fileChan <- filepath.Join(dir, entry.Name()) // gửi đường dẫn file vào channel
		}
	}
}

func worker(fileChan <-chan string, wg *sync.WaitGroup) { // nhận file từ filechan, lấy kích thước và in ra
	defer wg.Done() // đóng khi hoàn thành
	for filePath := range fileChan {
		info, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}
		fmt.Printf("File: %s, Size: %d bytes\n", filePath, info.Size()) // in ra thông tin file
	}
}

func main() {
	dir := "/home/anhdc/web-service-gin/Text" // thư mục cần duyệt
	// tạo channel chứa đường dẫn file
	fileChan := make(chan string, 10)
	// tạo waitgroup để đợi tất cả các go routine hoàn thành
	var wg sync.WaitGroup

	// Goroutine duyệt file
	go listFiles(dir, fileChan)

	// 4 go routinwe xử lý file
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)                // tăng bộ đếm
		go worker(fileChan, &wg) // gọi hàm worker
	}

	wg.Wait() // Đợi tất cả workers hoàn thành
}
