package vec

import "testing"

func TestMatrixString(test *testing.T) {
	var metrix = NewMatrix(3, 3)
	metrix.Vector = metrix.Vector.Map(func(x float64) float64 {
		return 1.23
	})
	test.Log("\n", metrix)
}
