package main

import (
	"fmt"
	"log"
	"syscall"
)

const (
	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
)

func main() {
	fd, err := syscall.Open("/dev/tty", syscall.O_RDWR, 0777)
	if err != nil {
		return
	}

	fmt.Println("OpenPort")
	buf := make([]byte, 128)
	read, err := syscall.Read(fd, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:read]))
}
