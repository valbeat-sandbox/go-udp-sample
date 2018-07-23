package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", ":8888")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		// 接続してきた相手のアドレスを受け取る
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		log.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
