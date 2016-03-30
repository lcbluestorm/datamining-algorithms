package kmeans

import (
	//	"fmt"
	"math/rand"
	"time"
)

type KMeans struct {
	dataSet []KeansModel
	k       int
}

func NewKMeans(dataSet []KeansModel, k int) KMeans {
	if k < 1 {
		panic("the value of k must be positive")
	}
	if len(dataSet) < k {
		panic("the size of dataSet must be bigger than k")
	}
	return KMeans{dataSet: dataSet, k: k}
}

func (kmeans KMeans) Clustering() map[int][]KeansModel {
	result := make(map[int][]KeansModel, kmeans.k)
	kCenter := kmeans.RandomK()
	capacity := kmeans.k / len(kmeans.dataSet)
	for {
		for _, item := range kmeans.dataSet {
			for i := 0; i < kmeans.k; i++ {
				result[i] = make([]KeansModel, 0, capacity)
			}

			var minDistance float64
			var index int
			for j, center := range kCenter {
				curDistance := item.GetDistance(center)
				if curDistance < minDistance {
					index, minDistance = j, curDistance
				}
			}
			result[index] = append(result[index], item)
		}
		var tmp = kmeans.dataSet[0]
		for i := 0; i < kmeans.k; i++ {
		}
		var continu bool
		if true {
			break
		}
	}
	return result
}

func (kmeans KMeans) RandomK() map[int]KeansModel {
	size := len(kmeans.dataSet)
	result := make(map[int]KeansModel, kmeans.k)
	for i := 0; i < kmeans.k; {
		rand.Seed(time.Now().UnixNano())
		index := rand.Int63n(int64(size))
		ii := int(index)
		if _, ok := result[ii]; !ok {
			result[ii] = kmeans.dataSet[index]
			i++
		}
	}
	return result
}
