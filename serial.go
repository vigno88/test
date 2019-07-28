package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/serial1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("port open")

	_, err = s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message sent")

}
