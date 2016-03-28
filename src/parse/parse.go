package parse

type BaseParse interface {
	Parse(interface{}) interface{}
}

type KnnParse struct{}

func (p KnnParse) Parse(param interface{}) interface{} {
	var result interface{}

	return result

}
