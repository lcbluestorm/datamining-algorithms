package knn

import (
	"testing"
)

func TestModel(t *testing.T) {
	m1 := NewModel(0, 0, "a")
	m2 := NewModel(1, 0, "b")

	ret := m1.GetDistance(*m2)
	if ret != 1 {
		t.Errorf("ret=%s", ret)
	}
}
