package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FileExecutor struct {
	filePath string
}

func NewFileExecutor(filePath string) *FileExecutor {
	return &FileExecutor{
		filePath: filePath,
	}
}

func (fe *FileExecutor) Execute() error {
	fmt.Print("--------- File Executor -------")

	file, err := os.OpenFile(fe.filePath, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			log.Fatalln("Error : Read Permission denied")
		}
		return err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for {
		success := scanner.Scan()
		if success == false {
			err = scanner.Err()
			return err
		}

		text := scanner.Text()
		fmt.Println(text)
	}

	return nil
}
