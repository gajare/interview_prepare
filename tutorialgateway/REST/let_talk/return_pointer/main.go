package main

import "fmt"

type Person struct {
	Name string
}

// by value (returns a copy)
func (p Person) SetNameCopy(name string) Person {
	p.Name = name
	return p
}

// by pointer (modifies same object)
func (p *Person) SetNamePtr(name string) *Person {
	p.Name = name
	return p
}

func main() {
	p1 := Person{}
	p1.SetNameCopy("Amol")
	fmt.Println("After SetNameCopy:", p1.Name) // ❌ empty

	p2 := &Person{}
	p2.SetNamePtr("Rahul")
	fmt.Println("After SetNamePtr:", p2.Name) // ✅ Rahul
}
