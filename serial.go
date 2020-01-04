package main

import (
	"fmt"
	"log"
	"time"

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
	n, err := port.Write([]byte("#___CNCON000\n"))
	if err != nil {
		log.Fatal(err)
	}
	buff := make([]byte, 100)

	n, err = port.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	resp := fmt.Sprintf("%v", string(buff[:n]))
	fmt.Println(resp)
	if resp == "ok" {
		n, err := port.Write([]byte("#CNCGCODE000\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n)
		time.Sleep(200* time.Millisecond)
		gc := make([]string,5)
		gc[0] = "G0X20Y20"
		gc[1] = "G0X-20Y20"
		gc[2] = "G0X-20Y-20"
		gc[3] = "G0X20Y-20"
		gc[4] = "G0X0Y0"

		printLine := true
		i := 0

		for {
			if printLine {
				port.Write([]byte(gc[i]))
				fmt.Println(gc[i])
				printLine = false
				i++
			}
			n, err := port.Read(buff)
			if err != nil {
				log.Fatal(err)
				break
			}
			if n == 0 {
				fmt.Println("\nEOF")
				break
			}
			resp := fmt.Sprintf("%v", string(buff[:n]))
			fmt.Println(resp)
			if resp == "ok" {
				printLine = true
			}
			if i > 5 {
				break
			}
		}
	}
	fmt.Println("Salut")

}
