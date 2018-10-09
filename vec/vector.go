package vec

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

type Vector []float64

func MakeVec(len int, source func(i int) float64) Vector {
	var vector = make(Vector, 0, len)
	for i := 0; i < len; i++ {
		vector = append(vector, source(i))
	}
	return vector
}

func (vector Vector) New() Vector {
	return make(Vector, 0, len(vector))
}

func (vector Vector) Len() int {
	return len(vector)
}

func (vector Vector) Copy() Vector {
	return append(vector.New(), vector...)
}

func (vector Vector) AddVector(xx Vector) Vector {
	if vector.Len() != xx.Len() {
		panic("add expects vectors with equal lengths")
	}
	var result = vector.Copy()
	for i, x := range xx {
		result[i] += x
	}
	return result
}

func (vector Vector) AddScalar(k float64) Vector {
	var result = vector.Copy()
	for i := range vector {
		result[i] += k
	}
	return result
}

func (vector Vector) MulVector(xx Vector) Vector {
	if vector.Len() != xx.Len() {
		panic("mul expect vectors with equal lengths")
	}
	var result = vector.Copy()
	for i, x := range xx {
		result[i] *= x
	}
	return result
}

func (vector Vector) MulScalar(k float64) Vector {
	var result = vector.Copy()
	for i := range vector {
		result[i] *= k
	}
	return result
}

func (vector Vector) MulDot(another Vector) float64 {
	if vector.Len() != another.Len() {
		panic("mul expect vectors with equal lengths")
	}
	var result float64
	for i, x := range another {
		result += vector[i] * x
	}
	return result
}

func (vector Vector) Append(xx Vector) Vector {
	return append(vector.Copy(), xx...)
}

func (vector Vector) Sum() float64 {
	var accum float64
	for _, x := range vector {
		accum += x
	}
	return accum
}

func (vector Vector) Map(op func(x float64) float64) Vector {
	var result = vector.Copy()
	for i, x := range vector {
		result[i] = op(x)
	}
	return result
}

func (vector Vector) ReduceLeft(start float64, op func(x, y float64) float64) float64 {
	var accum = start
	for _, x := range vector {
		accum = op(accum, x)
	}
	return accum
}

func (vector Vector) ReduceRight(start float64, op func(x, y float64) float64) float64 {
	var accum = start
	var vLen = vector.Len()
	for i := vLen - 1; i >= 0; i-- {
		var x = vector[i]
		accum = op(accum, x)
	}
	return accum
}

func (vector Vector) Filter(pred func(x float64) bool) Vector {
	var filtered = vector.New()
	for _, x := range vector {
		if pred(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func (vector Vector) Mean() float64 {
	return vector.Sum() / float64(vector.Len())
}

func (vector Vector) SumOfSquares() float64 {
	return vector.Map(Square).Sum()
}

func (vector Vector) XY(i int) (x, y float64) {
	return float64(i), vector[i]
}

func (vector Vector) Value(i int) float64 {
	return vector[i]
}

func (vector Vector) Max() float64 {
	if vector.Len() == 0 {
		return 0
	}
	return vector.ReduceLeft(vector[0], math.Max)
}

func (vector Vector) Min() float64 {
	if vector.Len() == 0 {
		return 0
	}
	return vector.ReduceLeft(vector[0], math.Min)
}

func (vector Vector) Less(i, j int) bool {
	return vector[i] < vector[j]
}

func (vector Vector) Sorted() Vector {
	var sorted = vector.Copy()
	sort.Float64s(sorted)
	return sorted
}

func (vector Vector) IsSorted() bool {
	return sort.Float64sAreSorted(vector)
}

func (vector Vector) Median() float64 {
	var vLen = vector.Len()
	switch vLen {
	case 0:
		return 0
	case 1:
		return vector[0]
	case 2:
		return (vector[0] + vector[1]) / 2
	default:
		var sorted = vector.Sorted()
		var midpoint = vLen / 2
		if vLen%2 != 0 {
			return sorted[midpoint]
		}
		return (sorted[midpoint] + sorted[midpoint-1]) / 2
	}
}

func (vector Vector) Quantile(p int) float64 {
	if p < 0 || p > 100 {
		panic(fmt.Sprintf("[vec.Quantile] expected p in range 0..100, got %d", p))
	}
	return vector.Sorted()[vector.Len()*p/100]
}

func (vector Vector) Dot(xx Vector) float64 {
	if vector.Len() != xx.Len() {
		panic(fmt.Sprintf("[vec.Dot] expected vectors of equal length, got %d and %d",
			vector.Len(), xx.Len()))
	}
	var result float64
	for i, x := range xx {
		result += vector[i] * x
	}
	return result
}

func (vector Vector) DeMean() Vector {
	var mean = vector.Mean()
	return vector.Map(func(x float64) float64 {
		return x - mean
	})
}

func (vector Vector) Indexes() Vector {
	var indexes = vector.New()
	for index := range vector {
		indexes = append(indexes, float64(index))
	}
	return indexes
}

func (vector Vector) Mode() Vector {
	return vector.Count().Mode()
}

func (vector Vector) Variance() float64 {
	var vLenSubOne = float64(vector.Len() - 1)
	var mean = vector.Mean()
	return vector.Map(func(x float64) float64 {
		return Square(mean - x)
	}).Sum() / vLenSubOne
}

func (vector Vector) StdDeviation() float64 {
	return math.Sqrt(vector.Variance())
}

func (vector Vector) InterquartileRange() float64 {
	return vector.Quantile(75) - vector.Quantile(25)
}

func (vector Vector) Covariance(xx Vector) float64 {
	if vector.Len() != xx.Len() {
		PanicF("[vec.Covariance] expected vector of length %d, got %d",
			vector.Len(), xx.Len())
	}
	var vLenSubOne = float64(vector.Len() - 1)
	return vector.DeMean().Dot(xx.DeMean()) / vLenSubOne
}

func (vector Vector) Correlation(xx Vector) float64 {
	var devV = vector.StdDeviation()
	var devX = xx.StdDeviation()
	if devX > 0 && devV > 0 {
		return vector.Covariance(xx) / (devX * devV)
	}
	return 0
}

func (vec Vector) WithLabels(labels Labels) LabeledVec {
	return VecWithLabels(vec, labels)
}

func (vector Vector) Stats() Stats {
	var stats Stats
	if vector.Len() == 0 {
		return stats
	}
	stats.Sum = vector[0]
	stats.Max = vector[0]
	stats.Min = vector[0]
	for _, x := range vector[1:] {
		if x > stats.Max {
			stats.Max = x
		}
		if x < stats.Min {
			stats.Min = x
		}
		stats.Sum += x
	}
	return stats
}

type Stats struct {
	Max float64
	Min float64
	Sum float64
}

func (stats Stats) String() string {
	return fmt.Sprintf(
		"max: %v\n"+
			"min: %v\n"+
			"sum: %v",
		stats.Max,
		stats.Min,
		stats.Sum)
}

func (stats Stats) Range() float64 {
	return stats.Max - stats.Min
}

func (vector Vector) CSVLine(delim string) string {
	var buf = bytes.NewBuffer(make([]byte, 0, 10*vector.Len()))
	var vLen = vector.Len()
	for i, x := range vector {
		if i == vLen-1 {
			delim = ""
		}
		fmt.Fprintf(buf, "%g%s", x, delim)
	}
	return buf.String()
}
