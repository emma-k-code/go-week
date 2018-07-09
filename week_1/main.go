package main

import (
	"strconv"

	Show "./show"
)

func main() {
	// 變數宣告
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love Golang_"

	ab := a + int(b)
	Show.Print("a + b = %d\n", ab)

	abc := a + int(b) + int(c)
	Show.Print("a + b + c = %d\n", abc)

	fe := f / float64(e)
	Show.Print("f / e = %f\n", fe)

	ad := a + strToInt(d)
	Show.Print("a + d = %d\n", ad)

	xa := x + strconv.Itoa(a)
	Show.Print("x & a = %s\n", xa)
}

// string 轉 int
func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
