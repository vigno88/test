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
	fmt.Println([]byte(resp))
	if resp == "ok\r\n" {
		n, err := port.Write([]byte("#CNCGCODE000\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", n)
		time.Sleep(200* time.Millisecond)
		gc := make([]string,5)
		gc[0] = "G0X20Y20\n"
		gc[1] = "G0X-20Y20\n"
		gc[2] = "G0X-20Y-20\n"
		gc[3] = "G0X20Y-20\n"
		gc[4] = "G0X0Y0\n"

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
			if resp == "ok\r\n" {
				printLine = true
			}
			if i > 4 {
				break
			}
		}
	}
	fmt.Println("Salut")

}
