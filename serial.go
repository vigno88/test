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
	port, err := serial.Open("/dev/ttyAMA0", mode)
	if err != nil {
		log.Fatal(err)
	}

	n, err := port.Write([]byte("dsadsadasdasdas\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

}
