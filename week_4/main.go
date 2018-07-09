package main

import (
	"fmt"
	"time"
)

func main() {
	// 建立紙漿傳輸通道
	pulpChan := make(chan int, 10)
	// 建立紙張傳輸通道
	paperChan := make(chan int, 10)
	// 完成通知
	doneChan := make(chan string)

	// 目標印刷 60000 張紙
	goal := 60000

	// 每秒生產 1000 公斤的紙漿
	go loggingShed(1000, &goal, pulpChan)
	// 每秒生產 5000 張紙
	go paperMill(5000, &goal, pulpChan, paperChan)
	// 每秒生產 3000 張紙
	go paperMill(3000, &goal, pulpChan, paperChan)
	// 每秒印刷 6000 張紙
	go printingFactory(6000, &goal, paperChan, doneChan)

	done := <-doneChan

	fmt.Printf("%s 已印刷 %d 張紙\n", done, 60000)

	fmt.Println("====工作已完成====")
}

/* 伐木工廠 */
func loggingShed(power int, goal *int, c chan int) {
	for i := *goal; i > 0; i = *goal {
		// 生產 %d 公斤紙漿
		fmt.Printf("生產 %d 公斤紙漿\n", power)
		// 將產能切格為最小單位 (100公斤) 傳入通道
		for p := 0; p < power/100; p++ {
			// 傳入紙漿通道中
			c <- 100
		}
		// 休息 1s
		time.Sleep(time.Duration(1) * time.Second)
	}
}

/* 造紙⼯廠 */
func paperMill(power int, goal *int, c1 chan int, c2 chan int) {
	// 需要花費的紙漿量 (公斤)
	spend := power / 10
	// 已取得的紙漿量 (公斤)
	take := 0

	for i := *goal; i > 0; i = *goal {
		// 從通道取得所需的紙漿量
		for take < spend {
			take += <-c1
		}
		// 生產 %d 張紙
		fmt.Printf("生產 %d 張紙\n", power)
		// 將生產的紙張傳入通道
		c2 <- power
		// 紙漿量 歸零
		take = 0
		// 休息 1s
		time.Sleep(time.Duration(1) * time.Second)
	}
}

/* 印刷廠 */
func printingFactory(power int, goal *int, c chan int, down chan string) {
	// 已取得的紙張數
	take := 0

	for i := *goal; i > 0; i = *goal {
		// 從通道取得紙張
		for take < power {
			take += <-c
		}
		// 減少目標數量
		*goal -= power
		// 印刷 %d 張紙
		fmt.Printf("印刷 %d 張紙\n", power)
		// 已取得的紙張數 - 花費的紙張數
		take -= power
		// 休息 1s
		time.Sleep(time.Duration(1) * time.Second)
	}

	down <- "done"
}
