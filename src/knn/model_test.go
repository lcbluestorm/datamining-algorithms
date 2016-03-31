package knn

import (
	"testing"
)

func TestModel(t *testing.T) {
	data1 := []float64{0, 0}
	data2 := []float64{1, 0}
	m1 := NewSimpleModel(data1, "a")
	m2 := NewSimpleModel(data2, "b")
	ret := m1.GetDistance(m2)
	if ret != 1 {
		t.Errorf("ret=%s", ret)
	}
}
