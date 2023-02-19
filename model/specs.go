package model

type Specs struct {
	SpecsGroups []SpecGroup
}

type SpecGroup struct {
	Name  string
	Specs []Spec
}

type Spec struct {
	Name  string
	Value any
}
