package knn

import (
	"fmt"
	"testing"
)

func TestKnn(t *testing.T) {
	var dataSet []*Model = make([]*Model, 3)
	var k int
	k = 3
	dataSet[0] = NewModel(1, 1.1, "A")
	dataSet[1] = NewModel(1, 1, "A")
	dataSet[2] = NewModel(0, 0, "B")
	newModel := NewModel(0, 0.1, "None")

	knnClassifier := NewKnnClassifier(dataSet, k)

	//fmt.Println("%s", knnClassifier)
	fmt.Println("%s", knnClassifier.Classify(*newModel))
}
