package knn

import (
	//	"fmt"

	"github.com/lcbluestorm/datamining-algorithms/src/datastructure/heap"
)

type KnnClassifier struct {
	dataSet []Model
	k       int
}

func NewKnnClassifier(dataSet []Model, k int) KnnClassifier {
	return KnnClassifier{dataSet: dataSet, k: k}
}

func (knn KnnClassifier) Classify(m Model) string {

	h := make([]heap.Item, 0)
	maxHeap := heap.NewMaxHeap(h)
	for _, model := range knn.dataSet {
		distance := m.GetDistance(model)
		item := heap.NewSimpleItem(model, float64(distance))
		if maxHeap.GetSize() < knn.k {
			maxHeap.Add(item)
		} else {
			maxHeap.InsertValue(item)
		}
	}
	//statistic the counts of labels
	items := maxHeap.GetItems()
	statics := make(map[string]int, knn.k)
	for _, value := range items {
		v := value.(heap.SimpleItem)
		item := v.GetItem()
		ii := item.(Model)
		label := ii.GetLabel()
		_, ok := statics[label]
		if ok {
			statics[label]++
		} else {
			statics[label] = 1
		}

	}
	// find the maxinum(counts) label
	max := 0
	var result string
	for key, value := range statics {
		if value > max {
			result = key
		}
	}
	return result
}
