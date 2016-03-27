package datastructure

import (
	"fmt"
	//	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {

	heap := []interface{}{Item{"a", 1.3}, Item{"b", 2.3}, Item{"c", 1}, Item{"c", 5.8}}
	maxHeap := NewMinHeap(heap)
	maxHeap.Add(Item{"e", 13})
	maxHeap.InsertValue(Item{"f", 5})
	index, value := maxHeap.FindBottom()
	fmt.Println("%s,%s,%s,%s", maxHeap, index, value, maxHeap.GetSize())

}
