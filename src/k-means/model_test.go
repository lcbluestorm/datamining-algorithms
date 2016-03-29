package kmeans

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	var m, m1 KeansModel
	var mm, mm1 SimpleKeansModel = []float64{3, 4}, []float64{0, 0}
	m = mm
	m1 = mm1
	d := m.GetDistance(mm1)
	//models := []KeansModel{m, m1}
	//r := mm.GetCenter(models)
	var a SimpleKeansModel
	fmt.Println(a, d, m, m1)
}
