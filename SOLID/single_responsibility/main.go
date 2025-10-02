package main

import "fmt"

type User struct{}

//single method/function, class, struct should have single responisibility

func (u *User) PrintUser(name string) {
	fmt.Println("user :", name)
}

type Email struct{}

func (e *Email) PrintEmail(sendToUser string) {
	fmt.Println("sending mail to :", sendToUser)
}

func main() {
	u := User{}
	e := Email{}
	u.PrintUser("amol")
	e.PrintEmail("gajare")
}

// this is wrong function which has two responsibilties
func (u *User) PrintAll(name, sendTo string) {
	fmt.Println("user is :", name)
	fmt.Println("sending to :", sendTo)
}
