package main

import (
	"bufio"
	"fmt"
	"log"

	"go.bug.st/serial.v1"
)

func main() {

	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyAMA0", mode)
	if err != nil {
		log.Fatal(err)
	}
	serialWriter := bufio.NewWriter(port)

	n, err := serialWriter.Write([]byte("#___CNCON000\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))

	}
	fmt.Println("Salut")

}
