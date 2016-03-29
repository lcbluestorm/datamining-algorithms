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
	GetCenter(models []KeansModel) KeansModel
}

type SimpleKeansModel []float64

func NewSimpleKeansModel() SimpleKeansModel {
	return make([]float64, 0)
}

func (m SimpleKeansModel) GetDistance(d Distance) float64 {
	var distance float64 = 0
	simpleModel := d.(SimpleKeansModel)
	for i, _ := range m {
		distance += math.Pow(m[i]-simpleModel[i], 2)
	}
	return math.Sqrt(distance)
}

func (m SimpleKeansModel) GetCenter(models []KeansModel) KeansModel {
	len := len(models)
	if len == 0 {
		return nil
	}
	var center SimpleKeansModel
	first := models[0].(SimpleKeansModel)
	size := len(first)
	if size == 0 {
		return nil
	}
	for i := 0; i < size; i++ {
	}
	return center
}

func sum(i int, model []KeansModel) {

}
