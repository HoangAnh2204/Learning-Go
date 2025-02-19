package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const numWorkers = 4 // Số lượng worker goroutine

// listFiles gửi danh sách file vào channel
func listFiles(dir string, fileChan chan<- string) {
	defer close(fileChan) // Đóng channel sau khi hoàn thành

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() { // Chỉ lấy file, không lấy thư mục con
			fileChan <- filepath.Join(dir, entry.Name()) // Gửi file vào channel
		}
	}
}

// worker đọc nội dung file và in ra màn hình
func worker(fileChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for filePath := range fileChan {
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		// Giới hạn in ra tối đa 200 ký tự để tránh quá dài
		output := string(content)
		if len(output) > 200 {
			output = output[:200] + "..." // Cắt bớt và thêm dấu "..."
		}

		fmt.Printf("File: %s\nContent:\n%s\n\n", filePath, output)
	}
}

func main() {
	dir := "/home/anhdc/web-service-gin/Text" // Thư mục chứa các file
	fileChan := make(chan string, 10)         // Channel chứa danh sách file
	var wg sync.WaitGroup

	// Khởi động goroutine để duyệt file
	go listFiles(dir, fileChan)

	// Tạo 4 worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(fileChan, &wg)
	}

	wg.Wait() // Đợi tất cả worker hoàn thành trước khi kết thúc chương trình
}
