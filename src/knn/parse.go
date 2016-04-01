package knn

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Parse interface {
	Parse() []KnnModel
}

type FileParse struct {
	path []string
}

func NewFileParse(path []string) FileParse {
	return FileParse{path: path}
}

func (f FileParse) GetPath() []string {
	return f.path
}

func (f *FileParse) SetPath(path []string) {
	f.path = path
}
func (f FileParse) Parse() []KnnModel {
	var result = make([]KnnModel, 0, 1000)
	var size = len(f.path)
	chs := make([]chan []KnnModel, size)
	for i := 0; i < size; i++ {
		path := f.path[i]
		chs[i] = make(chan []KnnModel, 1)
		go parse(path, chs[i])
	}
	for i := 0; i < size; i++ {
		ret := <-chs[i]
		if ret != nil {
			result = append(result, ret...)
		}
	}
	return result
}

func parse(path string, ch chan []KnnModel) {
	var result []KnnModel
	readFile, err := os.Open(path)
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
		lines := strings.Split(string(buf[:n-1]), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			values := strings.Fields(line)
			length := len(values) - 1
			data := make([]float64, length)
			label := values[length]
			var err error
			for j := 0; j < length; j++ {
				data[j], err = strconv.ParseFloat(values[j], 64)
				if err != nil {
					panic(err)
				}
			}
			m := NewSimpleKnnModel(data, label)
			result = append(result, m)
		}
	}
	ch <- result
}
