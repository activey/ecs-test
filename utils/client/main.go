package main

import (
	"fmt"
	"net"
)

const serverAddress = "127.0.0.1:1331" // Server address and port

func main() {
	// Resolve UDP address of the server
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at", serverAddress)

	// Send a message to the server
	message := "Hello from the client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Receive response from server
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return
	}

	fmt.Printf("Received response from server: %s\n", string(buffer[:n]))
}
