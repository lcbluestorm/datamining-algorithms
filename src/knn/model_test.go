package knn

import (
	"testing"
)

func TestModel(t *testing.T) {
	m1 := NewSimpleModel(0, 0, "a")
	m2 := NewSimpleModel(1, 0, "b")

	ret := m1.GetDistance(m2)
	if ret != 1 {
		t.Errorf("ret=%s", ret)
	}
}
