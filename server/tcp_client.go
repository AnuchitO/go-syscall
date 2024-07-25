package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Specify address details
	var serverAddr syscall.SockaddrInet4
	serverAddr.Port = 8080
	serverAddr.Addr = [4]byte{127, 0, 0, 1} // Localhost IP

	// Create a socket
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		os.Exit(1)
	}
	defer syscall.Close(sockfd)

	// Connect to the server
	err = syscall.Connect(sockfd, &serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}

	// Send data to the server
	data := []byte("Hello from client!\n")
	n, err := syscall.Write(sockfd, data)
	if err != nil {
		fmt.Println("Error writing to server:", err)
		os.Exit(1)
	}

	// Read data from the server
	buffer := make([]byte, 1024)
	n, err = syscall.Read(sockfd, buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		log.Fatal("closing port:", syscall.Close(sockfd))
		os.Exit(1)
	}

	fmt.Println("Received response from server:", string(buffer[:n]))

	// Close the socket explicitly
	log.Fatal("closing port:", syscall.Close(sockfd))
}
