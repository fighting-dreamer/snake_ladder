package reader

import (
	"fmt"
	"log"
)

type ConsoleReader struct {
}

func NewConsoleExecutor() *ConsoleReader {
	return &ConsoleReader{}
}

func (cr *ConsoleReader) ReadInt() int {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalf("got error while trying to read integer : ", err)
	}
	return n
}

func (cr *ConsoleReader) ReadDouble() float64 {
	var n float64
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalf("got error while trying to read float : ", err)
	}
	return n
}

func (cr *ConsoleReader) ReadString() string {
	var n string
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalf("got error while trying to read string : ", err)
	}
	return n
}
