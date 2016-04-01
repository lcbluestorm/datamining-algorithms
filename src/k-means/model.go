package kmeans

import (
	"math"
)

type Distance interface {
	GetDistance(d Distance) float64
}

type KmeansModel interface {
	Distance
	GetCenter(models []KmeansModel) KmeansModel
	Equal(model KmeansModel) bool
}

type SimpleKmeansModel []float64

func NewSimpleKmeansModel(size int) SimpleKmeansModel {
	return make([]float64, size)
}

func (m SimpleKmeansModel) GetDistance(d Distance) float64 {
	var distance float64 = 0
	simpleModel := d.(SimpleKmeansModel)
	for i, _ := range m {
		distance += math.Pow(m[i]-simpleModel[i], 2)
	}
	return math.Sqrt(distance)
}
func (m SimpleKmeansModel) Equal(model KmeansModel) bool {
	if model == nil {
		return false
	}
	simpleModel := model.(SimpleKmeansModel)
	modelSize := len(simpleModel)
	mSize := len(m)
	if mSize != modelSize {
		return false
	}
	for i := 0; i < mSize; i++ {
		if m[i] != simpleModel[i] {
			return false
		}

	}
	return true

}
func (m SimpleKmeansModel) GetCenter(models []KmeansModel) KmeansModel {
	length := len(models)
	if length == 0 {
		return nil
	}
	first := models[0].(SimpleKmeansModel)
	size := len(first)
	if size == 0 {
		return nil
	}
	chs := make([]chan float64, size)
	var center SimpleKmeansModel = NewSimpleKmeansModel(size)
	for i := 0; i < size; i++ {
		chs[i] = make(chan float64, 1)
		go sum(models, i, chs[i])
	}
	for j := 0; j < size; j++ {
		sum := <-chs[j]
		center[j] = sum
	}
	return center
}

func sum(models []KmeansModel, i int, ch chan float64) {
	size := len(models)
	var sum float64
	for j := 0; j < size; j++ {
		model := models[j]
		simpleModel := model.(SimpleKmeansModel)
		sum += simpleModel[i]
	}
	ch <- sum / float64(size)
}
