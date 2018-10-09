package vec

type Points struct {
	xx Vector
	yy Vector
}

func NewPoints(xx, yy Vector) Points {
	if xx.Len() != yy.Len() {
		panic("NewPoints expect two vectors of equal length")
	}
	return Points{
		xx: xx,
		yy: yy,
	}
}

func (points Points) Len() int {
	return points.xx.Len()
}

func (points Points) XY(i int) (x, y float64) {
	return points.xx[i], points.yy[i]
}
