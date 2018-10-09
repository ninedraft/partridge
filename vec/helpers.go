package vec

import (
	"math"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func Range(start, end int) Vector {
	return MakeVec(end-start, func(i int) float64 {
		return float64(i) + float64(start)
	})
}

func Square(x float64) float64 {
	return math.Pow(x, 2)
}

func Random(n int, values ...float64) Vector {
	if len(values) > 0 {
		var valuesN = len(values)
		return MakeVec(n, func(int) float64 {
			return values[rnd.Intn(valuesN)]
		})
	}
	return MakeVec(n, func(int) float64 {
		return rnd.Float64()
	})
}
