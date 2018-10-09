package vec

import (
	"bytes"
	"fmt"
	sof "misc/partridge/stringOrFloat"
)

type Frame struct {
	strings []string
	vector  Vector
	indexes []int
}

func FrameFromValues(values ...interface{}) Frame {
	var frame = Frame{}
	for _, value := range values {
		var index int
		switch value := value.(type) {
		case string:
			index = -1 - len(frame.strings)
			frame.strings = append(frame.strings, value)
		case float64:
			index = frame.vector.Len()
			frame.vector = append(frame.vector, value)
		case int:
			index = frame.vector.Len()
			frame.vector = append(frame.vector, float64(value))
		default:
			panic(fmt.Sprintf("[vec.FrameFromValues] expected string or float64, got %T", value))
		}
		frame.indexes = append(frame.indexes, index)
	}
	return frame
}

func (frame Frame) Value(column int) sof.StringOrFloat64 {
	var index = frame.indexes[column]
	if index < 0 {
		return sof.String(frame.strings[-1-index])
	}
	return sof.Float64(frame.vector.Value(index))
}

func (frame Frame) String() string {
	var buf = bytes.NewBufferString("[")
	for i, item := range frame.Values() {
		switch {
		case item.IsString():
			fmt.Fprintf(buf, "%q", item.AsString())
		default:
			fmt.Fprintf(buf, "%v", item.Interface())
		}
		if (i + 1) != frame.Len() {
			fmt.Fprintf(buf, ", ")
		}
	}
	return buf.String() + "]"
}

func (frame Frame) Copy() Frame {
	frame.vector = frame.vector.Copy()
	var newStrings = make([]string, 0, len(frame.strings))
	frame.strings = append(newStrings, frame.strings...)
	return frame
}

func (frame Frame) Len() int {
	return len(frame.indexes)
}

func (frame Frame) Values() []sof.StringOrFloat64 {
	var values = make([]sof.StringOrFloat64, 0, frame.Len())
	for index := 0; index < frame.Len(); index++ {
		values = append(values, frame.Value(index))
	}
	return values
}

func (frame Frame) Slice() []interface{} {
	var values = make([]interface{}, 0, frame.Len())

	for index := 0; index < frame.Len(); index++ {
		values = append(values, frame.Value(index).Interface())
	}
	return values
}
