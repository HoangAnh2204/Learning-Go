package main

import (
	"fmt"
)

func sum(arr []int, result chan int) {
	total := 0
	for _, num := range arr {
		total += num
	}
	result <- total // gửi kết quả vào channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Chia mảng thành 2 phần
	mid := len(numbers) / 2
	//từ vị tris 0 đến mid
	part1 := numbers[:mid]
	//từ vị trí mid đến hết
	part2 := numbers[mid:]

	// Tạo channel để nhận kết quả từ Goroutines
	result := make(chan int, 2)

	// Chạy 2 Goroutines để tính tổng từng phần
	go sum(part1, result)
	go sum(part2, result)

	// Nhận kết quả từ 2 Goroutines
	sum1 := <-result
	sum2 := <-result

	// Tổng cuối cùng
	finalSum := sum1 + sum2
	// In ra tổng từng phần và tổng cuối cùng
	fmt.Println("Tổng của phần 1:", sum1)
	fmt.Println("Tổng của phần 2:", sum2)
	fmt.Println("Tổng của mảng là:", finalSum)
}
