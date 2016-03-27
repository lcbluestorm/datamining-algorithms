package knn

import (
	"fmt"
	"testing"
)

func TestKnn(t *testing.T) {
	var dataSet []*Model
	var k int
	k = 3
	for i := 0; i < 10; i++ {
		//	New
	}
	knnClassifier := NewKnnClassifier(dataSet, k)

	fmt.Println("%s", knnClassifier)
}
