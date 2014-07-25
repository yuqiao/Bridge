package main

import (
	"fmt"
)

type Base struct {
}

func (b Base) info() string {
	return fmt.Sprintf(" id=%d", b.id)
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main() {
	rect1 := NewRect(1, 2, 4, 5)
	fmt.Println(rect1.info())
}
