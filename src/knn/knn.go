package knn

import (
	"../datastructure"

	"fmt"
)

type KnnClassifier struct {
	dataSet *[]*Model
	k       int
}

func NewKnnClassifier(dataSet *[]*Model, k int) KnnClassifier {
	return KnnClassifier{dataSet: data, k: k}
}

func (knn KnnClassifier) Classify(m *Model) string {

	maxHeap := datastructure.NewMaxHeap(interface{})
	var label string
	for i, model := range knn.dataSet {
		distance := m.GetDistance(*model)
		if maxHeap.GetSize() < knn.k {
			item := datastructure.NewItem(distance, model)
			maxHeap.Add(item)
		} else {
			maxHeap.InsertValue(distance, model)
		}
	}
	values := maxHeap.GetValues()
	counts := make(map[string]int, knn.k)
	for _, value := range values {
		v := value.(Item)
	}

	return label
}
