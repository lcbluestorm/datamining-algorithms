package knn

import (
	//"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	p := NewFileParse("data.txt")
	p.Parse()
}
