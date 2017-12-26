package main

import (
	"bufio"
	"flag"
	"log"

	"github.com/eclesh/recordio"
	"github.com/golang/protobuf/proto"

	"github.com/1vn/tensorgrep"
	"github.com/1vn/tensorgrep/protobuf"
)

var (
	checkPointDir = flag.String("d", "", "path to checkpoint directory")
)

// type Record struct {
// 	   uint64    length
//     uint32    masked crc of length
//     byte      data[length]
//     uint32    masked crc of dat
// }
func main() {
	flag.Parse()

	eventLoader, err := tensorgrep.NewEventLoaderFromDir(*checkPointDir)
	if err != nil {
		panic(err)
	}

	s := recordio.NewScanner(bufio.NewReader(eventLoader.File))
	for s.Scan() {
		data := s.Bytes()
		if len(data) < 2 {
			continue
		}
		log.Println(data[0], data[1], data[len(data)-1], data[len(data)-2], len(data))
		count := 0
		for _, b := range data {
			if b > 0 {
				count++
			}
		}
		log.Println(count)

		ev := protobuf.Event{}
		proto.Unmarshal(data[3:len(data)-1], &ev)
		log.Println(ev)
	}
	log.Println("over")
}
