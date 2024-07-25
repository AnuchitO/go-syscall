package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	_, err = conn.Write([]byte("Hello from client!\n"))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	// Read data from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	fmt.Println("Received response from server:", string(buffer[:n]))
}
