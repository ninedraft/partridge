package vec

type Labels []string

func (labels Labels) Len() int {
	return len(labels)
}

func (labels Labels) New() Labels {
	return make(Labels, 0, labels.Len())
}

func (labels Labels) Copy() Labels {
	return append(labels.New(), labels...)
}
