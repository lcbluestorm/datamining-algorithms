package knn

import (
	"testing"
)

func TestKnn(t *testing.T) {
	var dataSet []KnnModel = make([]KnnModel, 3)
	var k int
	k = 3
	dataSet[0] = NewSimpleKnnModel([]float64{1, 1.1}, "A")
	dataSet[1] = NewSimpleKnnModel([]float64{1, 1}, "A")
	dataSet[2] = NewSimpleKnnModel([]float64{0, 0}, "B")
	newModel := NewSimpleKnnModel([]float64{0, 0.1}, "None")

	knnClassifier := NewKnnClassifier(dataSet, k)
	if knnClassifier.Classify(newModel) != "A" {
		t.Errorf("KnnClassifier failed")
	}
	parse := NewFileParse([]string{"data/data1.txt", "data/data2.txt"})
	ret := parse.Parse()
	pKnnClassifier := NewKnnClassifier(ret, k)
	newModel1 := NewSimpleKnnModel([]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "None")
	label := pKnnClassifier.Classify(newModel1)
	if label != "B" {
		t.Errorf("KnnClassifier failed")
	}
}
