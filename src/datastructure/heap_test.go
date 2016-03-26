package datastructure

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	// MaxHeap Test
	heap := []interface{}{1.3, 3.3, 4.3, 5.3, 8.3, 9.3}
	maxHeap := NewMaxHeap(heap)
	index, value := maxHeap.FindBottom()
	maxHeap.Add(8)
	maxHeap.InsertValue(5)
	fmt.Println("%s,%s,%s", maxHeap, index, value)

	heap1 := []interface{}{6.8, 1.3, 3.3, 4.3, 5.3, 8.3, 9.3}
	minHeap := NewMinHeap(heap1)
	index1, value1 := minHeap.FindBottom()
	fmt.Println("%s,%s,%s", minHeap, index1, value1)
	//minHeap.Add(1)
	minHeap.InsertValue(16.8)
	fmt.Println("%s,%s,%s", minHeap, index1, value1)
}
