# Introduction

    some data structure for the mining algorithms.
    
# Heap

    It contains max-heap and min-heap which supports float/int/Item{}. And the heap is mainly used for the top k storage. 
    
## usage

*1. create a heap* 
    
    maxheap := []interface{}{3.3, 5, 8.8,9}
    minheap := []interface{}{3.3,5,8.8,9}
    items := []interface{Item{"a", 1.3}, Item{"b", 2.3}}
    maxHeap := NewMaxHeap(maxheap)
    minHeap := NewMinHeap(minheap)
    itemHeap := NewMaxHeap(items)

*2. add a item to heap*

    maxHeap.Add(8)
    itemHeap.Add{Item{"c", 8}}

*3. insert a item to heap if the value of the item is upper than the value of heap's root for min-heap and conversely for max-heap*

    maxHeap.InsertValue(13)
    itemHeap.InsertValue{Item{"d", 6.8}}
    
*4. get the maximum item in min-heap(minimun item for max-heap)*

    maxHeap.FindBottom()
    
