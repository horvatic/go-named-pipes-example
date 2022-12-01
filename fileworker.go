package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var fileWorkerInPipe = "fileWorkerInPipe"
var controlInPipe = "controlInPipe"

func main() {
	writeFile, _ := os.OpenFile(controlInPipe, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	readFile, _ := os.OpenFile(fileWorkerInPipe, os.O_CREATE, os.ModeNamedPipe)
	reader := bufio.NewReader(readFile)

	i := 0
	for {
		writeFile.WriteString(fmt.Sprintf("Message:%d\n", i))
		i++
		time.Sleep(time.Second)

		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("Message from controller:" + string(line))
		}
	}
}
