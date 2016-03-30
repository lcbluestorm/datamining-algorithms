package kmeans

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	var mm, mm1 SimpleKeansModel = []float64{3, 4}, []float64{0, 0}
	var mmm KeansModel
	mmm = mm
	d := mmm.GetDistance(mm1)
	models := []KeansModel{mm, mm1}
	r := mm.GetCenter(models)
	fmt.Println(d, r)
}
