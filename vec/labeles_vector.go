package vec

import (
	"fmt"
)

type LabeledVec struct {
	Vector
	labels Labels
}

func VecWithLabels(vec Vector, labels Labels) LabeledVec {
	if vec.Len() != labels.Len() {
		panic(fmt.Sprintf("[vec.VecWithLabel] expected Vector and Labels of equal length, got %d and %d",
			vec.Len(), labels.Len()))
	}
	return LabeledVec{
		Vector: vec.Copy(),
		labels: labels.Copy(),
	}
}

func (vec LabeledVec) Get(i int) (float64, string) {
	return vec.Vector[i], vec.labels[i]
}

func (vec LabeledVec) Labels() Labels {
	return vec.labels.Copy()
}

func (vec LabeledVec) Copy() LabeledVec {
	return LabeledVec{
		Vector: vec.Vector.Copy(),
		labels: vec.labels.Copy(),
	}
}
