package kmeans

type KMeans struct {
	dataSet []KeansModel
	k       int
}

func NewKMeans(dataSet []KeansModel, k int) KMeans {
	return KMeans{dataSet: dataSet, k: k}
}

func (kmeans KMeans) Clustering() map[string][]KeansModel {
	var result = make(map[string][]KeansModel, kmeans.k)
	return result
}
