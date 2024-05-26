package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

const (
	mb = 1024 * 1024
	gb = 1024 * mb
)

func main() {
	slog.Info("1 Billion Row Challenge Go!")
	start := time.Now()

	doneChannel := make(chan bool)
	bufferChannel := make(chan []byte)
	go fileGenerator(bufferChannel, "../data/measurements.txt")
	go processChunks(bufferChannel, doneChannel)

	<-doneChannel
	duration := time.Since(start)
	slog.Info("Done reading file")
	fmt.Println(duration)
}

func fileGenerator(bufferChannel chan []byte, filename string) {

	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		buffer := make([]byte, 1*mb)
		chunk, err := r.Read(buffer)
		buffer = buffer[:chunk]
		if chunk == 0 {
			if err == io.EOF {
				fmt.Println("Reached EOF")
				slog.Info("Reached EOF")
				break
			}
			if err != nil {
				slog.Error("error while reading chunks", err)
				break
			}
		}
		bufferChannel <- buffer
	}
	close(bufferChannel)
}

func processChunks(bufferChannel chan []byte, doneChannel chan bool) {
	for buffer := range bufferChannel {
		// to stuff
		if len(buffer) > 0{
			continue
		}
	}
	doneChannel <- true
}
