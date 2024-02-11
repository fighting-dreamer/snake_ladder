package reader

import (
	"fmt"
	"log"
	"os"
)

type FileReader struct {
	filePath string
	file     *os.File
}

func NewFileReader(filePath string) (*FileReader, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			log.Fatal("Error : read permission denied : ", err)
		}
		return nil, err
	}
	return &FileReader{
		file:     file,
		filePath: filePath,
	}, nil
}

func (fr *FileReader) ReadInt() int {
	var n int
	if _, err := fmt.Fscan(fr.file, &n); err != nil {
		log.Fatalf("got error while trying to read integer : ", err)
	}
	return n
}

func (fr *FileReader) ReadDouble() float64 {
	var n float64
	if _, err := fmt.Fscan(fr.file, &n); err != nil {
		log.Fatalf("got error while trying to read float : ", err)
	}
	return n
}

func (fr *FileReader) ReadString() string {
	var n string
	if _, err := fmt.Fscan(fr.file, &n); err != nil {
		log.Fatalf("got error while trying to read string : ", err)
	}
	return n
}
