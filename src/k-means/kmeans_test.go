package kmeans

import (
	"fmt"
	"testing"
)

func TestKmeans(t *testing.T) {

	var mm, mm1 SimpleKeansModel = []float64{3, 4}, []float64{0, 0}
	models := []KeansModel{mm, mm1}
	kmeans := NewKMeans(models, 1)
	kk := kmeans.RandomK()
	result := make(map[string][]KeansModel, 1)
	for key, value := range result {
		fmt.Println("aaa", key, value)
	}
	r := make([]int, 3, 5)
	r = append(r, 8)
	fmt.Println(r, len(r), cap(r))
	result["aa"] = models
	result["bb"] = models
	result["cc"] = models
	result["dd"] = models
	fmt.Println(kk, len(result))
}
