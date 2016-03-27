package knn

import (
	"math"
)

type Model struct {
	attr1 float64
	attr2 float64
	label string
}

func NewModel(attr1, attr2 float64, label string) *Model {
	return &Model{
		attr1: attr1,
		attr2: attr2,
		label: label}
}

func (m Model) GetLabel() string {
	return m.label
}
func (m *Model) SetLabel(label string) {
	m.label = label
}

func (m Model) GetDistance(m1 Model) float64 {
	r1 := math.Pow((m.attr1 - m1.attr1), 2)
	r2 := math.Pow((m.attr2 - m1.attr2), 2)
	r := math.Sqrt(r1 + r2)
	return r
}
