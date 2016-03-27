package datastructure

import (
	"fmt"
	//	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	// MaxHeap Test
	heap := []interface{}{Item{1.3, "a"}, Item{2.3, "b"}, Item{1, "c"}, Item{5.8, "c"}}
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

	//var item interface{}
	//item = Item{1, "a"}
	//v := reflect.ValueOf(item)
	//ii := item.(Item)
	//fmt.Println("%s,%s", v, ii.value)

	//	heap1 := []interface{}{6.8, 1.3, 3.3, 4.3, 5.3, 8.3, 9.3}
	//	minHeap := NewMinHeap(heap1)
	//	index1, value1 := minHeap.FindBottom()
	//	fmt.Println("%s,%s,%s", minHeap, index1, value1)
	//	minHeap.Add(1)
	//	minHeap.InsertValue(16.8)
	//	fmt.Println("%s,%s,%s", minHeap, index1, value1)
}
