package main

import (
	"fmt"
	"syscall"
)

func main() {
	handle, err := syscall.Open("/dev/qmux_gps/qmux_connect_socket", syscall.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("device open failed")
		syscall.Close(handle)
		fmt.Println(handle, err)
	} else {
		fmt.Println("device open success")
		buf := make([]byte, 128)
		read, err := syscall.Read(handle, buf)
		fmt.Println("device read success")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:read]))
	}
}
