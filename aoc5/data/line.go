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
	x, y := l.X1, l.Y1
	dx, dy := l.stepSlope()
	for x != l.X2 || y != l.Y2 {
		lp.AddCount(x, y)
		x += dx
		y += dy
	}
	lp.AddCount(l.X2, l.Y2)
	return lp
}

func (l *Line) stepSlope() (dx, dy int) {
	switch {
	case l.X2 > l.X1:
		dx = 1
	case l.X2 < l.X1:
		dx = -1
	default:
		dx = 0
	}
	switch {
	case l.Y2 > l.Y1:
		dy = 1
	case l.Y2 < l.Y1:
		dy = -1
	default:
		dy = 0
	}
	return dx, dy
}
