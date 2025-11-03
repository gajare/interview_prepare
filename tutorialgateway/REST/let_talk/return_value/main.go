package main

import "fmt"

// Define a struct
type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) SetLength(length int) Rectangle {
	r.Length = length
	return r // returns a copy
}

func (r Rectangle) SetWidth(width int) Rectangle {
	r.Width = width
	return r // returns another copy
}

func (r Rectangle) Area() {
	fmt.Println("Area:", r.Length*r.Width)
}

func main() {
	// Each method returns a *copy*, so the chain works with temporary values
	rect := Rectangle{}
	r1 := rect.SetLength(10)
	r2 := r1.SetWidth(5)
	r2.Area()
	rect.SetLength(5).SetWidth(5).Area()
}
