package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleIncomingMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Print("Received: ", message)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Start typing messages:")
	go handleIncomingMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Fprintln(conn, message)
	}

	if scanner.Err() != nil {
		log.Println("Error reading input:", scanner.Err())
	}
}
