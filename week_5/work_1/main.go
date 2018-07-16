package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	// close(ch)  // 錯誤程式碼
	// 錯誤原因: 在資料尚未傳入 channel 前就關閉 channel，導致資料傳入時發生錯誤
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
