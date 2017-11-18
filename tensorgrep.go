package tensorgrep

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// this should be the protobuf struct for event.proto
type EventLoader struct {
	File *os.File
}

type Event struct {
}

func NewEventLoaderFromDir(dir string) (*EventLoader, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read directory %s", dir)
	}

	var eventFileName string
	for _, file := range files {
		fileName := file.Name()
		if isTensorFlowEventsFile(fileName) {
			eventFileName = fileName
			break
		}
	}

	if eventFileName == "" {
		return nil, errors.New("event file not found")
	}

	eventFile, err := os.Open(fmt.Sprintf("%s/%s", dir, eventFileName))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open event file %s", eventFileName)
	}

	return &EventLoader{
		eventFile,
	}, nil
}

func (e *EventLoader) NextEvent() *Event {
	return &Event{}
}

// inspired from https://github.com/tensorflow/tensorboard/blob/master/tensorboard/backend/event_processing/event_accumulator.py#L101
func isTensorFlowEventsFile(fileName string) bool {
	return strings.Contains(fileName, "tfevents")
}
