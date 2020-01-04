package main

import (
	"bufio"
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
	serialReader := bufio.NewReader(port)

	buff := make([]byte, 100)
	n, err = serialReader.Read(buff)
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
		gc := make([]string,10)
		gc[0] = "G0X30Y30\n"
		gc[1] = "G0X-30Y30\n"
		gc[2] = "G0X-30Y-30\n"
		gc[3] = "G0X30Y-30\n"
		gc[4] = "G0X0Y0\n"
		gc[5] = "G0X30Y30\n"
		gc[6] = "G0X-30Y30\n"
		gc[7] = "G0X-30Y-30\n"
		gc[8] = "G0X30Y-30\n"
		gc[9] = "G0X0Y0\n"

		printLine := true
		i := 0

		for {
			fmt.Println("Wait to print")
			if printLine {
				port.Write([]byte(gc[i]))
				fmt.Println(gc[i])
				i++
			}
			n, err := serialReader.Read(buff)
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
			if resp[0] == 19 {
				printLine = false
				fmt.Println("XOFF")
			}
			if resp[0] == 17 {
				printLine = true
				fmt.Println("XOFF")
			}
			if i >= len(gc) {
				port.Write([]byte("%\n"))
				break
			}
		}
	}
	fmt.Println("Salut")

}
