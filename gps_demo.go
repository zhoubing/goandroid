package main

import (
	"fmt"
	"io"
	"net"
)

const (
	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
)

func main() {
	//fd, err := syscall.Open("/dev/iio\\:device0", syscall.O_RDWR, 0777)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("OpenPort")
	//buf := make([]byte, 128)
	//read, err := syscall.Read(fd, buf)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(buf[:read]))
	//addr, err := net.ResolveUnixAddr("unix", "/dev/socket/location/socket_hal")
	unix, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: "/dev/socket/location/mq/location-mq-s", Net: "unix"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		headData := make([]byte, 512)
		_, err = io.ReadFull(unix, headData)
		if err != nil {
			fmt.Println(err)
			return
		} //ReadFull
	}
}
