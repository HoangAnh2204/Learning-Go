package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	// Tạo một Goroutine khác để gửi dữ liệu vào channel
	go func() {
		// thời gian chờ 3 giây
		time.Sleep(3 * time.Second)
		ch <- "Hello"
	}()
	// Sử dụng select để chờ kết quả hoặc timeout
	select {
	// Nhận dữ liệu từ channel
	case msg := <-ch:
		fmt.Println("Received:", msg)
	// Timeout sau 2 giây
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}
