package knn

import (
	"math"
)

type Model interface {
	GetLabel() string
	SetLabel(label string)
	GetDistance(m1 Model) float64
}

type SimpleModel struct {
	attr1 float64
	attr2 float64
	label string
}

func NewSimpleModel(attr1, attr2 float64, label string) *SimpleModel {
	return &SimpleModel{
		attr1: attr1,
		attr2: attr2,
		label: label}
}

func (m SimpleModel) GetLabel() string {
	return m.label
}
func (m *SimpleModel) SetLabel(label string) {
	m.label = label
}

func (m SimpleModel) GetDistance(m1 Model) float64 {
	simpleModel1 := m1.(*SimpleModel)
	r1 := math.Pow((m.attr1 - simpleModel1.attr1), 2)
	r2 := math.Pow((m.attr2 - simpleModel1.attr2), 2)
	r := math.Sqrt(r1 + r2)
	return r
}
