package data

func (l *Line) Direction() LineDirection {
	switch {
	case l.X1 == l.X2:
		return Vertical
	case l.Y1 == l.Y2:
		return Horizontal
	default:
		return Diagonal
	}
}

func (l *Line) LinePoints() LinePoints {
	lp := NewLinePoints()
	switch l.Direction() {
	case Vertical:
		x := l.X1
		var y, y2 int
		if l.Y1 < l.Y2 {
			y = l.Y1
			y2 = l.Y2
		} else {
			y = l.Y2
			y2 = l.Y1
		}

		for ; y <= y2; y++ {
			lp.AddCount(x, y)
		}
	case Horizontal:
		y := l.Y1
		var x, x2 int
		if l.X1 < l.X2 {
			x = l.X1
			x2 = l.X2
		} else {
			x = l.X2
			x2 = l.X1
		}
		for ; x <= x2; x++ {
			lp.AddCount(x, y)
		}
	}
	return lp
}
