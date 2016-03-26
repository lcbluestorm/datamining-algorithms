# Introduction

    some data structure for the mining algorithms.
    
# Heap

    It contains max-heap and min-heap which supports float/int. And the heap is mainly used for the top k storage. 
    
## usage

    create a heap 
    
    maxheap := []interface{}{3.3, 5, 8.8,9}
    minheap := []interface{}{3.3,5,8.8,9}
    maxHeap := NewMaxHeap(maxheap)
    minHeap := NewMinHeap(minheap)

    add a item to heap

    maxHeap.Add(8)

    insert a item to heap if the value of the item is upper than the value of heap's root for min-heap and conversely 
for max-heap

    maxHeap.InsertValue(13)
