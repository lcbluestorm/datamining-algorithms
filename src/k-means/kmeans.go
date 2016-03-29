package kmeans

type KMeans struct {
	dataSet []KeansModel
	k       int
}

func NewKMeans(dataSet []KeansModel, k int) KMeans {
	return KMeans{dataSet: dataSet, k: k}
}

func Clustering() map[string][]KeansModel {

}
