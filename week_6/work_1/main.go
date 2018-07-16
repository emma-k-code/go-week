package main

import (
	"fmt"
	"time"
)

type CellPhone interface {
	Name() string
	Size() int
	TalkTime() time.Duration
}

type Iphone struct {
	version       string
	width, height int
	battery       time.Duration
}

func (i Iphone) Name() string {
	return i.version
}

func (i Iphone) Size() int {
	return i.height * i.width
}

func (i Iphone) TalkTime() time.Duration {
	return i.battery * time.Hour
}

type Pixel struct {
	version       string
	width, height int
	battery       time.Duration
}

func (p Pixel) Size() int {
	return p.height * p.width
}

func (p Pixel) TalkTime() time.Duration {
	return p.battery * time.Minute
}

//他是圓形的
type IWatch struct {
	version string
	radius  int
	battery time.Duration
}

func (i IWatch) Name() string {
	return i.version
}

func (i IWatch) Size() int {
	size := float64(i.radius) * float64(i.radius) * 3.14
	return int(size)
}

func (i IWatch) TalkTime() time.Duration {
	return i.battery * time.Hour
}

func main() {
	iwatch := IWatch{radius: 15, battery: 12, version: "iwatch-XX"}

	ShowPhone(iwatch)
}

func ShowPhone(c CellPhone) {
	fmt.Printf("Product %v \n", c.Name())
	fmt.Printf("Size %v \n", c.Size())
	fmt.Printf("Talk time %v \n", c.TalkTime())
	fmt.Println()
}
