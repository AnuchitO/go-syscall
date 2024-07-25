package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	// Open the file using syscall.Open
	fd, _, errno := syscall.Syscall(syscall.SYS_OPEN, uintptr(unsafe.Pointer(syscall.StringBytePtr("file.txt"))), uintptr(syscall.O_RDONLY), 0)
	if errno != 0 {
		fmt.Printf("Error opening file: %s\n", syscall.Errno(errno))
		return
	}

	// Create a buffer to read the file contents
	buf := make([]byte, 1024)

	// Read from the file using syscall.Read
	n, _, errno := syscall.Syscall(syscall.SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if errno != 0 {
		fmt.Printf("Error reading file: %s\n", syscall.Errno(errno))
		return
	}

	// Print the read content
	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])

	// Close the file using syscall.Close
	syscall.Syscall(syscall.SYS_CLOSE, uintptr(fd), 0, 0)
}
