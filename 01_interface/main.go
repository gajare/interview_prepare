package main

import "fmt"

type Shape interface {
	Area()
}
type Tringle struct {
	Length float64
	Height float64
}
type Circle struct {
	reduis float64
}

type Reactangle struct {
	Length float64
	Width  float64
	Height float64
}

func (r *Reactangle) Area() {
	fmt.Println(r.Height * r.Length * r.Width)
}

func (c *Circle) Area() {
	fmt.Println(3.14 * (c.reduis * c.reduis))
}
func (t *Tringle) Area() {
	fmt.Println(t.Length * t.Height)
}
func Calculate(s Shape) {
	s.Area()
}
func main() {
	t := Tringle{Length: 10, Height: 20}
	c := Circle{reduis: 10}
	r := Reactangle{Length: 10, Height: 10, Width: 10}
	Calculate(&t)
	Calculate(&c)
	Calculate(&r)
}
