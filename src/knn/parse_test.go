package knn

import (
	"testing"
)

func TestParse(t *testing.T) {
	p := NewFileParse([]string{"data/data1.txt", "data/data2.txt"})
	p.Parse()
}
