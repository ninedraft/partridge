package vec

import (
	"bytes"
	"fmt"
)

type Matrix struct {
	Vector
	shape Shape
}

func (matrix Matrix) Shape() Shape {
	return matrix.shape
}

func NewMatrix(width, height int) Matrix {
	return Matrix{
		shape: Shape{
			Width:  width,
			Height: height,
		},
		Vector: make(Vector, width*height),
	}
}

type Shape struct {
	Width  int
	Height int
}

func (matrix Matrix) Get(x, y int) float64 {
	if x < 0 || x >= matrix.shape.Width {
		panic(fmt.Sprintf("[Matrix.Get] index 'x'(%d) out of range 0..%d",
			x, matrix.shape.Width))
	}
	if y < 0 || y >= matrix.shape.Height {
		panic(fmt.Sprintf("[Matrix.Get] index 'y'(%d) out of range 0..%d",
			y, matrix.shape.Height))
	}
	return matrix.Vector[x*matrix.shape.Width+y]
}

func (matrix Matrix) String() string {
	var buf = &bytes.Buffer{}
	for i, x := range matrix.Vector {
		fmt.Fprintf(buf, " %.2g", x)
		if (i+1)%matrix.shape.Width == 0 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
