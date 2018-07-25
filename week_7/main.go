package main

import (
	"./phone"
)

type A struct {
	Name  string
	Brand string
}

func (p A) GetName() string {
	return p.Name
}
func (p A) Camera() string {
	return p.Brand
}
func (p A) Microphone() string {
	return p.Brand
}

type B struct {
	Name  string
	Brand string
}

func (p B) GetName() string {
	return p.Name
}
func (p B) Camera() string {
	return p.Brand
}
func (p B) Microphone() string {
	return p.Brand
}

func main() {
	a := A{Name: "手機A", Brand: "Sony"}
	b := B{Name: "手機B", Brand: "Samsung"}

	phone.ClickShutter(a)
	phone.Recording(a)
	phone.ClickShutter(b)
	phone.Recording(b)
}
