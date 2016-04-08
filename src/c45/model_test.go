package c45

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	m := NewSimpleC45Model("dd", []string{"aa", "bb", "cc"})
	m.SetDecision("ee")
	fmt.Println(m, m.GetDecision(), m.GetAttributes())
}
