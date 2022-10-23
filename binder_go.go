package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var fd int
var err error

const DEVICE string = "/dev/binder" /* 设备文件*/

type BinderVersion struct {
	protocol_version int32
}

func main() {
	fd, err = syscall.Open(DEVICE, syscall.O_RDWR, 0777)
	if err != nil {
		fmt.Println("device open failed")
		syscall.Close(fd)
		fmt.Println(fd, err)
	} else {
		fmt.Println("device open success")
		var bv BinderVersion
		bv.protocol_version = 0
		//定义了当前Binder版本号
		//http://androidxref.com/kernel_3.18/xref/drivers/staging/android/uapi/binder.h
		_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), 3221512713, uintptr(unsafe.Pointer(&bv)))

		if err != 0 {
			fmt.Println("Syscall failed")
		} else {
			fmt.Printf("binder version is %d\n", bv.protocol_version)
			fmt.Println("Syscall success")

			//https://geektutu.com/post/quick-go-mmap.html
			_, err := syscall.Mmap(int(fd), 0, 128*1024, 1, 2)
			if err != nil {
				fmt.Println("Mmap failed")
			} else {
				fmt.Println("Mmap success")
			}
		}
		syscall.Close(fd)
	}
}
