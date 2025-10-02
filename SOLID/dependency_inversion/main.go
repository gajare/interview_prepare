package main

import "fmt"

type Database interface {
	SaveOrder()
}

type MySQL struct{}

func (m MySQL) SaveOrder() { fmt.Println("save database : mysql") }

type Postgres struct{}

func (p Postgres) SaveOrder() { fmt.Println("save databse : postrges") }

type OrderService struct{ db Database }

func (m OrderService) PlaceOrder() { m.db.SaveOrder() }

func main() {
	mySQL := MySQL{}
	service := OrderService{db: mySQL}
	service.PlaceOrder()

	postrges := Postgres{}
	service = OrderService{db: postrges}
	service.PlaceOrder()

}
