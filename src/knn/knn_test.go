package knn

import (
	"fmt"
	"testing"
)

func TestKnn(t *testing.T) {
	var dataSet []KnnModel = make([]KnnModel, 3)
	var k int
	k = 2
	dataSet[0] = NewSimpleModel([]float64{1, 1.1}, "A")
	dataSet[1] = NewSimpleModel([]float64{1, 1}, "A")
	dataSet[2] = NewSimpleModel([]float64{0, 0}, "B")
	newModel := NewSimpleModel([]float64{0, 0.1}, "None")

	knnClassifier := NewKnnClassifier(dataSet, k)

	fmt.Println("%s", knnClassifier.Classify(newModel))
}
