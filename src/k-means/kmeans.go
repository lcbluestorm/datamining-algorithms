package kmeans

import (
	"math/rand"
	"time"
)

type KMeans struct {
	dataSet []KmeansModel
	k       int
}

func NewKMeans(dataSet []KmeansModel, k int) KMeans {
	if k < 1 {
		panic("the value of k must be positive")
	}
	if len(dataSet) < k {
		panic("the size of dataSet must be bigger than k")
	}
	return KMeans{dataSet: dataSet, k: k}
}

func (kmeans KMeans) Clustering() (map[int][]KmeansModel, int) {
	result := make(map[int][]KmeansModel, kmeans.k)
	var kCenter map[int]KmeansModel
	kCenter = kmeans.RandomK()
	capacity := len(kmeans.dataSet) / kmeans.k
	iters := 0
	for {
		iters++
		// init k clusters
		for j, _ := range kCenter {
			result[j] = make([]KmeansModel, 0, capacity)
		}
		// clustering each item
		for _, item := range kmeans.dataSet {
			var minDistance float64
			var index int
			var flag bool = true
			for j, center := range kCenter {
				if flag {
					minDistance = item.GetDistance(center)
					index = j
					flag = false
					continue
				}
				curDistance := item.GetDistance(center)
				if curDistance < minDistance {
					index, minDistance = j, curDistance
				}
			}
			result[index] = append(result[index], item)
		}
		//recalculate the k-center
		var tmp = kmeans.dataSet[0]
		var newKCenter map[int]KmeansModel = make(map[int]KmeansModel, kmeans.k)
		for key, value := range result {
			newKCenter[key] = tmp.GetCenter(value)
		}
		var counts int = 0
		for _, value1 := range newKCenter {
			flag := false
			for j, value2 := range kCenter {
				if value1.Equal(value2) {
					kCenter[j] = nil
					flag = true
					break
				}
			}
			if !flag {
				counts++
			}
		}
		if counts == 0 {
			break
		}
		kCenter = newKCenter
	}
	return result, iters
}

func (kmeans KMeans) RandomK() map[int]KmeansModel {
	size := len(kmeans.dataSet)
	result := make(map[int]KmeansModel, kmeans.k)
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
