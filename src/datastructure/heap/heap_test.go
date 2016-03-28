package heap

import (
	"fmt"
	//	"reflect"
	"testing"
)

type Float float64

func (x Float) Less(than Item) bool {
	return x < than.(Float)
}

func TestHeap(t *testing.T) {
	var i1, i2, i3, i4, i5 Float = 1.3, 3.3, 5, 8, 5.8
	heap := []Item{i1, i2, i3, i4, i5}
	maxHeap := NewMaxHeap(heap)
	//	minHeap := NewMinHeap(heap)
	var ii1 Float
	ii1 = 1
	//maxHeap.Add(ii1)
	maxHeap.InsertValue(ii1)
	index, value := maxHeap.FindBottom()
	fmt.Println("%s,%s,%s", maxHeap, index, value)

}
