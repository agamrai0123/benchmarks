package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	bufferSizes := []int{
		// 1024,             // 1 KB
		// 2 * 1024,         // 2 KB
		// 4 * 1024,         // 4 KB
		// 8 * 1024,         // 8 KB
		// 16 * 1024,        // 16 KB
		// 32 * 1024,        // 32 KB
		// 64 * 1024,        // 64 KB
		// 128 * 1024,       // 128 KB
		// 256 * 1024,       // 256 KB
		512 * 1024,      // 512 KB
		1024 * 1024,     // 1 MB
		2 * 1024 * 1024, // 2 MB
		4 * 1024 * 1024, // 4 MB
		8 * 1024 * 1024, // 8 MB
		// 16 * 1024 * 1024, // 16 MB
		// 32 * 1024 * 1024, // 32 MB
		// 64 * 1024 * 1024, // 64 MB
	}
	for _, bufferSize := range bufferSizes {
		started := time.Now()
		run(bufferSize)
		duration := time.Since(started).Seconds()
		fmt.Printf("Buffer size: %d bytes, Time: %0.6f seconds\n", bufferSize, duration)
	}
}

func run(bufferSize int) {
	file, err := os.Open("../1brcgolang/measurements.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// scanner.Buffer(make([]byte, bufferSize), bufferSize)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
	}
}
