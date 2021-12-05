package data

import "log"

func NewLinePoints() LinePoints {
	var lp LinePoints = map[int]map[int]int{}
	return lp
}

func (l LinePoints) GetCount(x, y int) int {
	ymap, ok := l[x]
	if !ok {
		l[x] = map[int]int{y: 0}
		return 0
	}
	yval, ok := ymap[y]
	if !ok {
		ymap[y] = 0
		return 0
	}
	return yval
}

func (l LinePoints) AddCount(x, y int) {
	l.AddCountN(x, y, 1)
}

func (l LinePoints) AddCountN(x, y, delta int) {
	val := l.GetCount(x, y)
	l[x][y] = val + delta
}

func (l LinePoints) Merge(other LinePoints) {
	for ox, oys := range other {
		if oys == nil {
			continue
		}
		for oy, val := range oys {
			l.AddCountN(ox, oy, val)
		}
	}
}

func (l LinePoints) OverlappingPoints() int {
	var overlapping int
	log.Printf("Length of l is: %d\n", len(l))
	for x, ys := range l {
		if ys != nil {
			for y, val := range ys {
				if val >= 2 {
					overlapping++
					log.Printf("Point (%d,%d) overlapped %d times\n", x, y, val)
				}
			}
		}
		log.Printf("'ys' is nil for x: %d\n", x)
	}
	return overlapping
}
