package vec

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFrame(test *testing.T) {
	var values = []interface{}{1, 2, "three"}
	var frame = FrameFromValues(values...)
	fmt.Println(frame.indexes)
	fmt.Println(frame.Values())
	reflect.DeepEqual(frame.Slice(), values)
}
