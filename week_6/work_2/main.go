package main

import (
	"fmt"
	"time"

	Rocket "rocket"
)

type A struct {
	count int
}

func (r A) Launch() {
	for i := r.count; i > 0; i-- {
		fmt.Printf("A 倒數 %d \n", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Println("A 發射")
}

type B struct {
	count int
}

func (r B) Launch() {
	for i := r.count; i > 0; i-- {
		fmt.Printf("B 倒數 %d \n", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Println("B 發射")
}

func main() {
	a := A{count: 3}
	b := B{count: 5}

	ShowPhone(a)
	fmt.Println()
	ShowPhone(b)
}

func ShowPhone(r Rocket.Rocket) {
	r.Launch()
}
