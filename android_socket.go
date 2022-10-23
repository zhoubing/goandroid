package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:26001")
	lis, _ := net.ListenTCP("tcp4", addr)
	fmt.Println("Server started")
	con, _ := lis.Accept()
	b := make([]byte, 256)
	count, _ := con.Read(b)
	fmt.Println("Received: ", count)
	fmt.Println("Received: ", string(b[:count]))
	con.Close()
	fmt.Println("Server end")
}
