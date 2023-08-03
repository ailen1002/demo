package service

import (
	"go.bug.st/serial"
	"log"
)

func Portlist() (ports []string) {
	ports, err := serial.GetPortsList()
	if ports == nil {
		log.Println("serial port not found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return ports
}
