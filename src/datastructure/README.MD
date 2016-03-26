# Introduction
    some data structure for the mining algorithms.
    
# Heap
    It contains max-heap and min-heap which supports float/int. And the heap is mainly used 
    for the top k storage. 
## usage
    1. create a heap 
        maxheap := []interface{}{1.3, 3.3, 4.3, 5.3, 8.3, 9.3} 
        minheap := []interface{}{1.3, 3.3, 4.3, 5.3, 8.3, 9.3}
        maxHeap := NewMaxHeap(maxheap)
        minHeap := NewMinHeap(minheap)
    2. add a item to heap
        maxHeap.Add(8)
    3. insert a item to heap if the value of the item is upper than the value of heap's root for min-heap
    and conversely for max-heap
        minHeap.InsertValue(16.8)

