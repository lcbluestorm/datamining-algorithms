package knn

import (
	"github.com/lcbluestorm/datamining-algorithms/src/datastructure"
)

type KnnClassifier struct {
	dataSet *[]*Model
	k       int
}

func NewKnnClassifier(dataSet *[]*Model, k int) KnnClassifier {
	return KnnClassifier{dataSet: dataSet, k: k}
}

func (knn KnnClassifier) Classify(m *Model) string {

	maxHeap := datastructure.NewMaxHeap(interface{})
	for i, model := range knn.dataSet {
		distance := m.GetDistance(*model)
		if maxHeap.GetSize() < knn.k {
			item := datastructure.NewItem(distance, model)
			maxHeap.Add(item)
		} else {
			maxHeap.InsertValue(distance, model)
		}
	}
	//statistic the num of labels
	values := maxHeap.GetValues()
	statics := make(map[string]int, knn.k)
	for _, value := range values {
		v := value.(Item)
		item := v.item
		label := item.GetLabel()
		vv, ok := statics[label]
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
