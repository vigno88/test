package main

import (
	"fmt"
	"log"

	"go.bug.st/serial.v1"
)

func main() {

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	mode := &serial.Mode{
		BaudRate: 9600,
	}
	_, err = serial.Open("/dev/ttyAMA0", mode)
	if err != nil {
		log.Fatal(err)
	}

}
