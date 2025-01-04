package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

const (
	memName = "/share_mem1" // name of the shared memory object
	memSize = 1024          // size of the shared memory object in bytes
)

func main() {

	// Open the shared memory object
	fd, err := unix.Open(memName, unix.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error open memory:", err)
		os.Exit(1)
	}
	defer unix.Close(fd)

	// Map the shared memory
	mem, err := unix.Mmap(fd, 0, memSize, unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		fmt.Println("Error maping the memory:", err)
		os.Exit(1)
	}
	defer unix.Munmap(mem)

	// Read the shared memory
	fmt.Println("Readed:", string(mem))
}
