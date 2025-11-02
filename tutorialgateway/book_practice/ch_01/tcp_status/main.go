package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "golang.org:80") //Connects to the web server over TCP
	fmt.Println("conn :", conn)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") //Sends HTTP GET request
	fmt.Println("conn :", conn)

	status, _ := bufio.NewReader(conn).ReadString('\n') //Reads the first line of the server’s response
	fmt.Println("status", status)
}
