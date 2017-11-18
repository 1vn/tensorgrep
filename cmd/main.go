package main

import (
	"bufio"
	"flag"
	"log"

	"github.com/1vn/tensorgrep"
)

var (
	checkpointDir = flag.String("checkpointDir", "", "path to checkpoint directory")
)

func main() {
	flag.Parse()

	eventLoader, err := tensorgrep.NewEventLoaderFromDir(*checkpointDir)
	if err != nil {
		panic(err)
	}

	eventScanner := bufio.NewScanner(eventLoader.File)

	for eventScanner.Scan() {
		log.Println(eventScanner.Text())
	}
}
