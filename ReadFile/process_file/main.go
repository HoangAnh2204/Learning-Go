package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Hàm xử lý từng dòng (đếm số ký tự)
func processLine(line string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	if line == "" {
		fmt.Println("Dòng rỗng, bỏ qua...")
		return
	}

	result := fmt.Sprintf("Dòng: %s -> Số ký tự: %d\n", line, len(line))
	fmt.Println("Đang xử lý:", result) // Thêm log
	ch <- result
}

func main() {
	// Mở file để đọc
	file, err := os.Open("/home/anhdc/web-service-gin/Text/text1.txt")
	if err != nil {
		fmt.Println("Lỗi mở file:", err)
		return
	}
	defer file.Close()

	// Tạo channel để nhận kết quả
	ch := make(chan string, 100)
	var wg sync.WaitGroup

	// Đọc file từng dòng
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go processLine(line, ch, &wg) // Chạy goroutine xử lý
	}

	// Đóng channel sau khi tất cả goroutine kết thúc
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Mở file để ghi kết quả
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Lỗi tạo file:", err)
		return
	}
	defer outputFile.Close()

	// Nhận dữ liệu từ channel và ghi vào file
	writer := bufio.NewWriter(outputFile)
	for result := range ch {
		writer.WriteString(result)
	}
	fmt.Println("Hoàn thành xử lý file! Kết quả được ghi vào output.txt")
}
