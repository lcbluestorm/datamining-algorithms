package knn

import (
	"fmt"
)

type KnnClassifier struct {
	data []*Model
	k    int
}

func NewKnnClassifier(data []*Model, k int) *KnnClassifier {
	return &KnnClassifier{data: data, k: k}
}

func (k KnnClassifier) Classify(m Model) string {
	topK := make([]Model, k.k)
	var label string
	for i, model := range k.data {
		fmt.Println("i=%s,model=%s", i, model)

	}
	return label
}
