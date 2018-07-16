package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	num := 20 // 階乘數量

	result := 0
	ch := make(chan int, num)

	for i := num; i > 0; i-- {
		// 進行階乘計算
		go factorial(i, ch)
	}

	for i := 1; i <= num; i++ {
		// 將計算結果進行相加
		v := <-ch
		result += v
	}

	close(ch)
	elapsed := time.Since(start)

	fmt.Println("計算結果:", result)
	fmt.Println("花費時間:", elapsed)
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}

	ch <- result
}
