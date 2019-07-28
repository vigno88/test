package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "dev/ttyUSB1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		_, err = s.Write([]byte("test"))
		if err != nil {
			log.Fatal(err)
		}
	}

}

