package vec

import "fmt"

const (
	ColumnFloat64 ColumnKind = "float64"
	ColumnString  ColumnKind = "string"
)

type ColumnKind string

func (kind ColumnKind) String() string {
	return string(kind)
}

type Scheme []Column

func (scheme Scheme) Columns() []string {
	var columns = make([]string, 0, len(scheme))
	for _, column := range scheme {
		columns = append(columns, column.Name)
	}
	return columns
}

func (scheme Scheme) Column(name string) (int, bool) {
	for i, column := range scheme {
		if column.Name == name {
			return i, true
		}
	}
	return -1, false
}

func SchemeFromFrame(columns []string, frame Frame) Scheme {
	if len(columns) != frame.Len() {
		panic(fmt.Sprintf("[vec.SchemeFromFrame] scheme %s expects %d columns, got %d",
			columns, len(columns), frame.Len()))
	}
	var scheme = make(Scheme, 0, frame.Len())
	for i, item := range frame.Values() {
		switch {
		case item.IsFloat64():
			scheme = append(scheme, Column{
				Name: columns[i],
				Kind: ColumnFloat64,
			})
		case item.IsString():
			scheme = append(scheme, Column{
				Name: columns[i],
				Kind: ColumnString,
			})
		default:
			panic(fmt.Sprintf("[vec.SchemeFromFrame] expects only non-nil items in dataframe"))
		}
	}
	return scheme
}

func (scheme Scheme) MatchFrame(frame Frame) bool {
	if len(scheme) != frame.Len() {
		return false
	}
	for i, item := range frame.Values() {
		switch {
		case scheme[i].Kind == ColumnString && item.IsString():
			continue
		case scheme[i].Kind == ColumnFloat64 && item.IsFloat64():
			continue
		default:
			return false
		}
	}
	return true
}

type Column struct {
	Name string
	Kind ColumnKind
}

func (column Column) String() string {
	return column.Name + ":" + column.Kind.String()
}
