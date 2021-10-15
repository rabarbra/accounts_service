package logg

import (
	"github.com/segmentio/kafka-go"
)

type Logger struct {
	kafka *kafka.Writer
}

func New() *Logger {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Async: true,
	}

	return &Logger{
		kafka: w,
	}
}
