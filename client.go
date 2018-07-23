package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", ":8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("Sending to server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}
	log.Println("Receiving from server")
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:len]))
}
