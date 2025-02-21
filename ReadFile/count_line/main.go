package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const numWorkers = 4 // Số worker goroutine

// listFiles gửi danh sách file .txt vào channel
func listFiles(dir string, fileChan chan<- string) {
	defer close(fileChan) // Đóng channel khi hoàn thành

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") { // Chỉ lấy file .txt
			fileChan <- filepath.Join(dir, entry.Name()) // Gửi file vào channel
		}
	}
}

// worker nhận file từ channel, đếm số dòng và in kết quả
func worker(fileChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for filePath := range fileChan {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}

		// Đếm số dòng
		scanner := bufio.NewScanner(file)
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}

		file.Close()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning file:", err)
			continue
		}

		// In kết quả
		fmt.Printf("File: %s, Lines: %d\n", filePath, lineCount)
	}
}

func main() {
	dir := "/home/anhdc/web-service-gin/Text" // Thư mục cần duyệt
	fileChan := make(chan string, 10)         // Channel chứa danh sách file
	var wg sync.WaitGroup

	// Goroutine duyệt file
	go listFiles(dir, fileChan)

	// Tạo 4 worker goroutine
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(fileChan, &wg)
	}

	wg.Wait() // Đợi tất cả workers hoàn thành trước khi kết thúc chương trình
}
