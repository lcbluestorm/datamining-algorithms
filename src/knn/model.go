package knn

import (
	"math"
)

type KnnModel interface {
	GetLabel() string
	SetLabel(label string)
	GetDistance(m1 KnnModel) float64
}

type SimpleModel struct {
	data  []float64
	label string
}

func NewSimpleModel(data []float64, label string) *SimpleModel {
	return &SimpleModel{
		data:  data,
		label: label}
}

func (m SimpleModel) GetLabel() string {
	return m.label
}
func (m *SimpleModel) SetLabel(label string) {
	m.label = label
}

func (m SimpleModel) GetDistance(m1 KnnModel) float64 {
	simpleModel := m1.(*SimpleModel)
	mData, m1Data := m.data, simpleModel.data
	mDataSize, m1DataSize := len(mData), len(m1Data)
	if mDataSize != m1DataSize {
		panic("the size of two SimpleModel must be same")
	}
	var sum float64 = 0
	for i, _ := range mData {
		sum += math.Pow(mData[i]-m1Data[i], 2)
	}
	return math.Sqrt(sum)
}
