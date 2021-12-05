package process

import (
	"github.com/drognisep/aoc2021/aoc5/data"
	"log"
)

func ExtractPoints(in <-chan *data.Line, out chan<- data.LinePoints) {
	lp := data.NewLinePoints()
	go func() {
		for line := range in {
			points := line.LinePoints()
			log.Printf("Merging points: %#v\n", points)
			lp.Merge(points)
		}
		out <- lp
		close(out)
	}()
}
