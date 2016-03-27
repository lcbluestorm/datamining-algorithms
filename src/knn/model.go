package knn

import (
	"math"
)

type Model struct {
	attr1 float32
	attr2 float32
	label string
}

func NewModel(attr1, attr2 float32, label string) *Model {
	return &Model{
		attr1: attr1,
		attr2: attr2,
		label: label}
}

func (m Model) GetLabel() string {
	return m.label
}

func (m Model) GetDistance(m1 Model) float32 {
	r1 := math.Pow(float64((m.attr1 - m1.attr1)), 2)
	r2 := math.Pow(float64((m.attr2 - m1.attr2)), 2)
	r := math.Sqrt(r1 + r2)
	return float32(r)
}
