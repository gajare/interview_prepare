package main

import "fmt"

type Walker interface {
	Walk()
}

type Flyer interface {
	Fly()
}

type Swimmer interface {
	Swim()
}

type Dog struct{}

func (d Dog) Walk() { fmt.Println("Dog is walking") }

type Duck struct{}

func (d Duck) Walk() { fmt.Println("Duck is walking") }
func (d Duck) Fly()  { fmt.Println("Duck is flying") }
func (d Duck) Swim() { fmt.Println("Duck is swimming") }

func main() {
	var w Walker = Dog{}
	w.Walk()

	var f Flyer = Duck{}
	f.Fly()
}
