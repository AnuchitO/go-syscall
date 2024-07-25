package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Create a TCP socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer syscall.Close(fd)

	// Bind the socket to an address
	addr := &syscall.SockaddrInet4{Port: 8080}
	err = syscall.Bind(fd, addr)
	if err != nil {
		fmt.Println("Error binding socket:", err)
		return
	}

	// Listen for incoming connections
	err = syscall.Listen(fd, 10)
	if err != nil {
		fmt.Println("Error listening on socket:", err)
		return
	}

	// Accept a connection
	conn, _, err := syscall.Accept(fd)
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer syscall.Close(conn)

	// Read data from the client
	buf := make([]byte, 1024)
	n, _, err := syscall.Recvfrom(conn, buf, 0)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	fmt.Printf("Received request %d bytes: %s\n", n, buf[:n])

	// Create a destination address
	destAddr := &syscall.SockaddrInet4{
		Port: 8080,                  // Replace with the desired port
		Addr: [4]byte{127, 0, 0, 1}, // Replace with the desired IP address
	}

	// Send a response to the client
	data := []byte("Hello from server!\n")
	err = syscall.Sendto(conn, data, 0, destAddr)
	if err != nil {
		fmt.Println("Error sending to client:", err)
		return
	}
}
