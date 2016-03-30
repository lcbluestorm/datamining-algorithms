package kmeans

import (
	//	"fmt"
	"math"
)

type Distance interface {
	GetDistance(d Distance) float64
}

type KeansModel interface {
	Distance
	GetCenter(models *[]KeansModel) KeansModel
}

type SimpleKeansModel []float64

func NewSimpleKeansModel(size int) SimpleKeansModel {
	return make([]float64, size)
}

func (m SimpleKeansModel) GetDistance(d Distance) float64 {
	var distance float64 = 0
	simpleModel := d.(SimpleKeansModel)
	for i, _ := range m {
		distance += math.Pow(m[i]-simpleModel[i], 2)
	}
	return math.Sqrt(distance)
}

func (m SimpleKeansModel) GetCenter(models *[]KeansModel) KeansModel {
	length := len(*models)
	if length == 0 {
		return nil
	}
	first := (*models)[0].(SimpleKeansModel)
	size := len(first)
	if size == 0 {
		return nil
	}
	chs := make([]chan float64, size)
	var center SimpleKeansModel = NewSimpleKeansModel(size)
	for i := 0; i < size; i++ {
		chs[i] = make(chan float64, 1)
		sum(models, i, chs[i])
	}
	for j := 0; j < size; j++ {
		sum := <-chs[j]
		center[j] = sum
	}
	return center
}

func sum(models *[]KeansModel, i int, ch chan float64) {
	size := len(*models)
	var sum float64
	for j := 0; j < size; j++ {
		model := (*models)[j]
		simpleModel := model.(SimpleKeansModel)
		sum += simpleModel[i]
	}
	ch <- sum / float64(size)
}
