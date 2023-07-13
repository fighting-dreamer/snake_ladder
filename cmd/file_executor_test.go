package cmd

import (
	"fmt"
	"testing"
)

func TestFileExecuteFileRead(t *testing.T) {
	fileExecutor := NewFileExecutor("./input.dat")
	err := fileExecutor.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
