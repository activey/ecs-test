package main

import (
	"fmt"
	"net"
)

const (
	port       = ":1331" // Server listening port
	bufferSize = 1024
)

func main() {
	// Start listening on the UDP port
	serverAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server listening on", port)

	buffer := make([]byte, bufferSize)
	for {
		// Read UDP message
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP message:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", clientAddr, message)

		// Respond to client
		response := "Server received: " + message
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
