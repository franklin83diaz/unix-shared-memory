package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

const (
	memName = "/shm_example1" // name of the shared memory object
	memSize = 1024            // size of the shared memory object (bytes)
)

func main() {

	// create a new shared memory object
	fd, err := unix.Open(memName, unix.O_CREAT|unix.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error creating shared memory", err)
		os.Exit(1)
	}

	defer unix.Close(fd)
	defer unix.Unlink(memName)

	// Set the size of the shared memory object
	if err := unix.Ftruncate(fd, memSize); err != nil {
		fmt.Println("Error redimensionando memoria:", err)
		os.Exit(1)
	}

	// Map the shared memory object to the process's address space
	mem, err := unix.Mmap(fd, 0, memSize, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		fmt.Println("Error mapeando memoria:", err)
		os.Exit(1)
	}
	defer unix.Munmap(mem)

	// Write data to the shared memory object
	message := "Hello from writer!"
	copy(mem, message)

	fmt.Println("Data Write in the shared memory:", message)

	// Wait for the reader to read the data
	fmt.Println("Press ENTER to exit")
	fmt.Scanln()
}
