package c45

import ()

type C45Model interface {
	GetDecision() string
	GetAttributes() []string
}

type SimpleC45Model struct {
	decision   string
	attributes []string
}

func NewSimpleC45Model(decision string, attributes []string) *SimpleC45Model {
	return &SimpleC45Model{decision: decision, attributes: attributes}
}
func (m SimpleC45Model) GetDecision() string {
	return m.decision
}

func (m SimpleC45Model) GetAttributes() []string {
	return m.attributes
}

func (m *SimpleC45Model) SetDecision(decision string) {
	m.decision = decision
}

func (m *SimpleC45Model) SetAttributes(attrs []string) {
	m.attributes = attrs
}
