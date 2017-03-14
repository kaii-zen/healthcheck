package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: healthcheck [cpu|mem|disk]")
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "cpu":
		checkCPU()
	case "mem", "memory":
		checkMem()
	case "disk":
		checkDisk()
	default:
		usage()
	}
}

// Convert bytes to Gibibytes so that things look pretty
func toGiB(b uint64) float64 {
	return float64(b) / 1073741824
}
