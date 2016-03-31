package knn

import (
	"bufio"
	"fmt"
	"os"
)

type Parse interface {
	Parse() []KnnModel
}

type FileParse struct {
	path string
}

func NewFileParse(path string) FileParse {
	return FileParse{path: path}
}

func (f FileParse) GetPath() string {
	return f.path
}

func (f *FileParse) SetPath(path string) {
	f.path = path
}
func (f FileParse) Parse() []KnnModel {
	var result = make([]KnnModel, 0, 30)
	readFile, err := os.Open(f.path)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(readFile)
	for {
		buf := make([]byte, 1024)
		n, _ := reader.Read(buf)
		if n == 0 {
			break
		}
		fmt.Printf("%d,%s", n, string(buf))
	}
	return result
}
