package vec

import (
	"fmt"
)

type DataFrame struct {
	frames []Frame
	scheme Scheme
}

func MakeDataFrame(n int, factory func(index int) Frame, columns ...string) DataFrame {
	var frames = make([]Frame, 0, n)
	var scheme Scheme
	for i := 0; i < n; i++ {
		var frame = factory(i)
		if i == 0 {
			scheme = SchemeFromFrame(columns, frame)
		}
		if !scheme.MatchFrame(frame) {
			panic(fmt.Sprintf("[vec.MakeDataFrame] iteration %d: frame %v doesn't match scheme %s",
				i, frame.Slice(), scheme))
		}
		frames = append(frames, frame)
	}
	return DataFrame{
		frames: frames,
		scheme: scheme,
	}
}

func (dataframe DataFrame) ColumnIndex(name string) (index int, column bool) {
	return dataframe.scheme.Column(name)
}

func (dataframe DataFrame) ColumnFloats(name string) Vector {
	var column, ok = dataframe.ColumnIndex(name)
	if !ok {
		panic(fmt.Sprintf("[vec.DataFrame.ColumnFloats] column %q not found", name))
	}
	var columnData = make(Vector, 0, len(dataframe.frames))
	for i, frame := range dataframe.frames {
		if len(frame.indexes) <= column {
			panic(fmt.Sprintf("[vec.DataFrame.ColumnFloats] frame %d column %q:%d not found",
				i,
				name, column))
		}
		columnData = append(columnData, frame.Value(column).AsFloat64())
	}
	return columnData
}

func (dataframe DataFrame) ColumnStrings(name string) []string {
	var column, ok = dataframe.ColumnIndex(name)
	if !ok {
		panic(fmt.Sprintf("[vec.DataFrame.ColumnStrings] column %q not found", name))
	}
	var columnData = make([]string, 0, len(dataframe.frames))
	for i, frame := range dataframe.frames {
		if len(frame.indexes) <= column {
			panic(fmt.Sprintf("[vec.DataFrame.ColumnStrings] frame %d column %q:%d not found",
				i, name, column))
		}
		columnData = append(columnData, frame.Value(column).AsString())
	}
	return columnData
}

func (dataframe DataFrame) Columns() []string {
	return dataframe.scheme.Columns()
}

func (dataframe DataFrame) Filter(pred func(frame Frame) bool) DataFrame {
	var filtered = dataframe
	filtered.frames = make([]Frame, 0, len(dataframe.frames))
	for _, frame := range dataframe.frames {
		if pred(frame.Copy()) {
			filtered.frames = append(filtered.frames, frame.Copy())
		}
	}
	return filtered
}

func (dataframe DataFrame) Frame(index int) Frame {
	return dataframe.frames[index].Copy()
}

func (dataframe DataFrame) Scheme() Scheme {
	return append(Scheme{}, dataframe.scheme...)
}

func (dataframe DataFrame) Map(op func(frame Frame) Frame, columns ...string) DataFrame {
	return MakeDataFrame(len(dataframe.frames), func(index int) Frame {
		var frame = dataframe.frames[index].Copy()
		return op(frame)
	}, columns...)
}

func (dataframe DataFrame) Len() int {
	return len(dataframe.frames)
}

func ConcatDataFrames(dataframes ...DataFrame) DataFrame {
	var framesN = 0
	for _, dataframe := range dataframes {
		framesN += dataframe.Len()
	}
	var frames = make([]Frame, 0, framesN)
	for _, dataframe := range dataframes {
		frames = append(frames, dataframe.frames...)
	}
	return MakeDataFrame(framesN, func(index int) Frame {
		return frames[index].Copy()
	}, dataframes[0].Columns()...)
}

func (dataframe DataFrame) GroupByFloat(split func(frame Frame) float64, apply func(dataframe DataFrame) DataFrame) DataFrame {
	var groups = make(map[float64][]Frame, len(dataframe.frames))
	for _, frame := range dataframe.frames {
		var key = split(frame.Copy())
		var group = groups[key]
		group = append(group, frame.Copy())
		groups[key] = group
	}
	var dataframes = make([]DataFrame, 0, len(groups))
	var scheme = dataframe.Columns()
	for _, group := range groups {
		var dataframe = MakeDataFrame(len(group), func(index int) Frame {
			return group[index]
		}, scheme...)
		dataframes = append(dataframes, apply(dataframe))
	}
	return ConcatDataFrames(dataframes...)
}
