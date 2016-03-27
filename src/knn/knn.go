package knn

import (
	//	"fmt"

	"github.com/lcbluestorm/datamining-algorithms/src/datastructure"
)

type KnnClassifier struct {
	dataSet []*Model
	k       int
}

func NewKnnClassifier(dataSet []*Model, k int) KnnClassifier {
	return KnnClassifier{dataSet: dataSet, k: k}
}

func (knn KnnClassifier) Classify(m Model) string {

	heap := []interface{}{}
	maxHeap := datastructure.NewMaxHeap(heap)
	for _, model := range knn.dataSet {
		distance := m.GetDistance(*model)
		item := datastructure.NewItem(model, distance)
		if maxHeap.GetSize() < knn.k {
			maxHeap.Add(item)
		} else {
			maxHeap.InsertValue(item)
		}
	}
	//statistic the num of labels
	values := maxHeap.GetValues()
	statics := make(map[string]int, knn.k)
	for _, value := range values {
		v := value.(datastructure.Item)
		item := v.GetItem()
		ii := item.(*Model)
		label := ii.GetLabel()
		_, ok := statics[label]
		if ok {
			statics[label]++
		} else {
			statics[label] = 1
		}

	}

	// find the maxinum label's num
	max := 0
	var result string
	for key, value := range statics {
		if value > max {
			result = key
		}
	}
	return result
}
