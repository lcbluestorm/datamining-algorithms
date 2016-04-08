package c45

import (
	"bufio"
	"os"
	"strings"
)

type Parse interface {
	Parse() []C45Model
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
func (f FileParse) Parse() []C45Model {
	var result = make([]C45Model, 0, 1000)
	var size = len(f.path)
	chs := make([]chan []C45Model, size)
	for i := 0; i < size; i++ {
		path := f.path[i]
		chs[i] = make(chan []C45Model, 1)
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

func parse(path string, ch chan []C45Model) {
	var result []C45Model
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
			attributes, decision := values[:length], values[length]
			m := NewSimpleC45Model(decision, attributes)
			result = append(result, m)
		}
	}
	ch <- result
}
