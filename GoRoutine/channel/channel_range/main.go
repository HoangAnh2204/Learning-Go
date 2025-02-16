package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// truyền dữ liệu vào chan
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // Đóng channel

	for val := range ch {
		fmt.Println(val) // 1, 2, 3
	}
}
