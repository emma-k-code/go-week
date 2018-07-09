package main

import (
	"fmt"
)

func scoreCalculate() {
	fmt.Println("練習1")
	scoreList := [5]int{37, 45, 67, 32, 11}
	fmt.Printf("分數列表：%v\n", scoreList)

	sum := 0
	for _, score := range scoreList {
		sum += score
	}
	fmt.Printf("總和：%d\n", sum)

	avg := sum / len(scoreList)
	fmt.Printf("平均：%d\n", avg)
}

func findMinScore() {
	var minScore int
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minScore = x[0]
	for _, score := range x {
		if minScore > score {
			minScore = score
		}
	}

	fmt.Println("練習2")
	fmt.Printf("最小值：%d\n", minScore)
}

func main() {
	scoreCalculate()
	fmt.Println()
	findMinScore()
}
