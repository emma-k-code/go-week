package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	// "io/ioutil"
	// "encoding/json"
)

type APIResult struct {
	Result bool
	Spend  time.Duration
}

var Path string
var Method string
var Data string
var Num int

func main() {
	flag.StringVar(&Path, "p", "", "API路徑")
	flag.StringVar(&Method, "m", "GET", "GET/POST/PUT/DETELE")
	flag.StringVar(&Data, "d", "", "API參數")
	flag.IntVar(&Num, "n", 5, "測試連線次數")

	flag.Parse()

	// API 路徑不可為空
	if Path == "" {
		fmt.Println("請輸入 API Path")
		os.Exit(0)
	}
	// Method 一律調整為大寫
	Method = strings.ToUpper(Method)
	// Method 驗證
	if Method != "GET" && Method != "POST" && Method != "PUT" && Method != "DETELE" {
		fmt.Println("Method 錯誤")
		os.Exit(0)
	}

	// === API 測試 ===

	// 接收 API 執行結果
	ch := make(chan APIResult, Num)
	successTimes := 0
	failTimes := 0
	spendTimes := []time.Duration{}

	for i := Num; i > 0; i-- {
		// 建立連線
		go getAPI(ch)
	}

	// 取得連線時間與結果
	for i := Num; i > 0; i-- {
		result := <-ch

		// 紀錄執行成功/失敗次數
		if result.Result {
			successTimes++
		} else {
			failTimes++
		}

		// 紀錄花費時間
		spendTimes = append(spendTimes, result.Spend)
	}

	close(ch)
	// === API 測試 ===

	// === 結果計算 ===

	var max time.Duration
	var min time.Duration
	var total time.Duration
	// 取得最大、小花費時間 與 總花費時間
	for i, spend := range spendTimes {
		total += spend

		if i == 0 {
			max = spend
			min = spend
			continue
		}
		if spend > max {
			max = spend
		}
		if spend < min {
			min = spend
		}
	}

	// 將花費時間轉為數字進行平均計算 (單位為 ns)
	avgNs := total.Nanoseconds() / int64(Num)
	// 取得平均花費時間
	avg := time.Duration(avgNs)
	// === 結果計算 ===

	fmt.Printf("單一最大花費時間: %+v\n", max)
	fmt.Printf("單一最小花費時間: %+v\n", min)
	fmt.Printf("總花費時間: %+v\n", total)
	fmt.Printf("平均花費時間: %+v\n", avg)

	fmt.Println("連線成功次數: ", successTimes)
	fmt.Println("連線失敗次數: ", failTimes)

}

/*
 call API
*/
func getAPI(ch chan APIResult) {
	start := time.Now()

	// === 連線 ===
	req, _ := http.NewRequest(Method, Path, strings.NewReader(Data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, _ := client.Do(req)

	elapsed := time.Since(start)

	// // === 顯示回傳 ===
	// body, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()
	// var data map[string]interface{}
	// json.Unmarshal([]byte(body), &data)
	// fmt.Printf("data: %+v\n", data)

	// === 判斷連線是否成功 ===
	result := false
	if resp.StatusCode == 200 {
		result = true
	}

	fmt.Printf("單一花費時間: %v\n", elapsed)

	ch <- APIResult{result, elapsed}
}
