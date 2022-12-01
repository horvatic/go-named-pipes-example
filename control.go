package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

var fileWorkerInPipe = "fileWorkerInPipe"
var controlInPipe = "controlInPipe"

func main() {
	os.Remove(fileWorkerInPipe)
	os.Remove(controlInPipe)

	_ = unix.Mkfifo(fileWorkerInPipe, 0666)
	_ = unix.Mkfifo(controlInPipe, 0666)

	readFile, _ := os.OpenFile(controlInPipe, os.O_CREATE, os.ModeNamedPipe)
	writeFile, _ := os.OpenFile(fileWorkerInPipe, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	reader := bufio.NewReader(readFile)

	for {
		writeFile.WriteString("Message\n")
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("Message from file worker:" + string(line))
		}
	}
}
