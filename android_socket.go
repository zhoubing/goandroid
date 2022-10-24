package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

var conn net.Conn

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:26001")
	lis, _ := net.ListenTCP("tcp4", addr)
	fmt.Println("Server started")
	for {
		fmt.Println("before: ", "Accept")
		con, _ := lis.Accept()
		b := make([]byte, 256)
		count, _ := con.Read(b)
		fmt.Println("Received: ", count)
		fmt.Println("Received: ", string(b[:count]))

		var res Result
		errs := json.Unmarshal(b[:count], &res)
		if errs != nil {
			fmt.Println("json error")
		}
		if res.Code == 1 {
			fmt.Println("ftp event", res.Message)
			if conn != nil {
				fmt.Println("before write to android")
				_, err := conn.Write([]byte(res.Message + "\n"))
				if err != nil {
					fmt.Println("write to android error ", err)
				}
				fmt.Println("after write to android")
			}
			//con.Write([]byte("Bye"))
			con.Close()
		} else if res.Code == 2 {
			conn = con
			fmt.Println("android client registered", res.Message)
		}
	}
}
