package stringOrFloat

import "fmt"

type StringOrFloat64 struct {
	value interface{}
}

func (sof StringOrFloat64) IsString() bool {
	var _, isString = sof.value.(string)
	return isString
}

func (sof StringOrFloat64) IsFloat64() bool {
	var _, isFloat64 = sof.value.(float64)
	return isFloat64
}

func (sof StringOrFloat64) IsNil() bool {
	return sof.value == nil
}

func (sof StringOrFloat64) AsString() string {
	return sof.value.(string)
}

func (sof StringOrFloat64) AsFloat64() float64 {
	return sof.value.(float64)
}

func (sof StringOrFloat64) Interface() interface{} {
	return sof.value
}

func (sof StringOrFloat64) String() string {
	return fmt.Sprint(sof.value)
}

func String(str string) StringOrFloat64 {
	return StringOrFloat64{str}
}

func Float64(v float64) StringOrFloat64 {
	return StringOrFloat64{v}
}
