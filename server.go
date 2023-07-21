package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("New client connected:", conn.RemoteAddr().String())

	for {
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from client:", err)
			break
		}

		message := string(buffer[:bytesRead])
		fmt.Println("Received from", conn.RemoteAddr().String(), ":", message)
	}

	conn.Close()
	fmt.Println("Client disconnected:", conn.RemoteAddr().String())
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Println("Server started and listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
