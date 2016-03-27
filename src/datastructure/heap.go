package datastructure

import (
	"reflect"
)

const (
	MAXHEAP = "maxHeap"
	MINHEAP = "minHeap"
)

type Item struct {
	value interface{}
	item  interface{}
}

type Heap struct {
	heap []interface{}
	attr string
}

func NewItem(value, item interface{}) Item {
	return Item{value: value, item: item}
}
func NewMaxHeap(heap []interface{}) Heap {
	buildMaxHeap(heap)
	return Heap{heap: heap, attr: MAXHEAP}
}

func NewMinHeap(heap []interface{}) Heap {
	buildMinHeap(heap)
	return Heap{heap: heap, attr: MINHEAP}
}

func (h Heap) GetSize() int {
	return len(h.heap)
}

func (h Heap) GetRoot() interface{} {
	if len(h.heap) == 0 {
		return nil
	}
	return h.heap[0]
}

func (h Heap) GetValues() []interface{} {
	size := len(h.heap)
	values := make([]interface{}, size)
	copy(values, h.heap)
	return values
}

func (h Heap) GetAttr() string {
	return h.attr
}

func (h Heap) FindBottom() (index int, value interface{}) {
	size := len(h.heap)
	if size == 0 {
		return -1, nil
	}
	startIndex := size/2 + 1
	max := toFloat64(h.heap[startIndex-1])
	min := max
	maxIndex, minIndex := startIndex-1, startIndex-1
	for i := startIndex + 1; i <= size; i++ {
		cur := toFloat64(h.heap[i-1])
		if cur > max {
			max = cur
			maxIndex = i - 1

		}
		if cur < min {
			min = cur
			minIndex = i - 1
		}
	}
	var resultIndex int
	var resultValue interface{}
	if h.attr == MAXHEAP {
		resultIndex = minIndex
	} else if h.attr == MINHEAP {
		resultIndex = maxIndex
	}
	resultValue = h.heap[resultIndex]
	return resultIndex, resultValue
}

func (h *Heap) InsertValue(value interface{}) bool {
	size := len(h.heap)
	if size == 0 {
		h.heap = append(h.heap, value)
		return true
	}
	if h.attr == MAXHEAP {
		if toFloat64(value) < toFloat64(h.heap[0]) {
			h.heap[0] = value
			adjustMaxHeapDown(h.heap, 1, size)
			return true
		}
	} else if h.attr == MINHEAP {
		if toFloat64(value) > toFloat64(h.heap[0]) {
			h.heap[0] = value
			adjustMinHeapDown(h.heap, 1, size)
			return true
		}
	} else {
		panic("heap type error(not maxHeap/minHeap)")
	}
	return false
}

// add a value to the heap
func (h *Heap) Add(value interface{}) {
	h.heap = append(h.heap, value)
	if h.attr == MAXHEAP {
		adjustMaxHeapUp(h.heap, len(h.heap), 1)
	} else if h.attr == MINHEAP {
		adjustMinHeapUp(h.heap, len(h.heap), 1)
	} else {
		panic("heap type error(not maxHeap/minHeap)")
	}
}

///////////////////////////////////////////////////////
/////////////// build max heap ////////////////////////
func buildMaxHeap(array []interface{}) {
	size := len(array)
	for i := size / 2; i > 0; i -= 1 {
		adjustMaxHeapDown(array, i, size)
	}
}

func adjustMaxHeapDown(array []interface{}, start, end int) {
	tmp := array[start-1]
	for i := 2 * start; i <= end; i *= 2 {
		if i < end {
			leftChildValue := toFloat64(array[i-1])
			rightChildValue := toFloat64(array[i])
			if rightChildValue > leftChildValue {
				i++
			}
		}
		if toFloat64(tmp) > toFloat64(array[i-1]) {
			break
		}
		array[start-1] = array[i-1]
		start = i
	}
	array[start-1] = tmp
}
func adjustMaxHeapUp(array []interface{}, start, end int) {
	tmp := array[start-1]
	tmpFloat := toFloat64(tmp)
	for i := start / 2; i >= end; i /= 2 {
		parentValue := toFloat64(array[i-1])
		if tmpFloat < parentValue {
			break
		}
		array[start-1] = array[i-1]
		start = i
	}
	array[start-1] = tmp
}

///////////////////////////////////////////////////////
/////////////// build min heap ////////////////////////
func buildMinHeap(array []interface{}) {
	size := len(array)
	for i := size / 2; i > 0; i -= 1 {
		adjustMinHeapDown(array, i, size)
	}
}

func adjustMinHeapDown(array []interface{}, start, end int) {
	tmp := array[start-1]
	for i := 2 * start; i <= end; i *= 2 {
		if i < end {
			leftChildValue := toFloat64(array[i-1])
			rightChildValue := toFloat64(array[i])
			if rightChildValue < leftChildValue {
				i++
			}
		}
		if toFloat64(tmp) < toFloat64(array[i-1]) {
			break
		}
		array[start-1] = array[i-1]
		start = i
	}
	array[start-1] = tmp
}
func adjustMinHeapUp(array []interface{}, start, end int) {
	tmp := array[start-1]
	tmpFloat := toFloat64(tmp)
	for i := start / 2; i >= end; i /= 2 {
		parentValue := toFloat64(array[i-1])
		if tmpFloat > parentValue {
			break
		}
		array[start-1] = array[i-1]
		start = i
	}
	array[start-1] = tmp
}

//// convert interface{}(int/float) to float64
func toFloat64(i interface{}) float64 {
	if i == nil {
		panic("param in func toFloat64() is nil")
	}
	value := reflect.ValueOf(i)
	var result float64
	var k = value.Kind()
	if k == reflect.Struct {
		ii := i.(Item)
		v := ii.value
		value = reflect.ValueOf(v)
		k = value.Kind()
	}
	if k == reflect.Float64 || k == reflect.Float32 {
		result = value.Float()
	} else {
		result = float64(value.Int())
	}
	return result
}
