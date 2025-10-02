package main

import "fmt"

type Amol interface {
	show()
}
type Gajare struct {
	sirname string
}

func (g *Gajare) show1() {
	fmt.Println("sirname:", g.sirname)
}

func main() {
	g := Gajare{sirname: "bhoi"}
	g.show1()

}
