package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/iio:device0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("1")
	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("2")
	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("3")
	log.Printf("%q", buf[:n])
}
