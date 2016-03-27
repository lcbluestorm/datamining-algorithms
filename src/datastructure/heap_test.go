package datastructure

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	// MaxHeap Test
	heap := []interface{}{Item{"a", 1.3}, Item{"b", 2.3}, Item{"c", 1}, Item{"c", 5.8}}
	//heap := []interface{}{}
	maxHeap := NewMaxHeap(heap)
	index, value := maxHeap.FindBottom()
	//maxHeap.Add(8)
	//maxHeap.InsertValue(5)
	fmt.Println("%s,%s,%s,%s", maxHeap, index, value, maxHeap.GetSize())

	test := maxHeap.GetValues()
	test[0] = "aaa"
	fmt.Println("%s", maxHeap)
	fmt.Println("%s", test)

	var item interface{}
	item = Item{"a", 3}
	v := reflect.ValueOf(item)
	ii := item.(Item)
	fmt.Println("%s,%s", v, ii.GetItem())

	//	heap1 := []interface{}{6.8, 1.3, 3.3, 4.3, 5.3, 8.3, 9.3}
	//	minHeap := NewMinHeap(heap1)
	//	index1, value1 := minHeap.FindBottom()
	//	fmt.Println("%s,%s,%s", minHeap, index1, value1)
	//	minHeap.Add(1)
	//	minHeap.InsertValue(16.8)
	//	fmt.Println("%s,%s,%s", minHeap, index1, value1)
}
