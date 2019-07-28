package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyAMA0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("port open")
	for i := 0; i < 100; i++ {
		_, err = s.Write([]byte("test"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("message sent")
	}

}

