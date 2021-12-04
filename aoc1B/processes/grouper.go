package processes

func Grouper(input <-chan int, output chan<- int) {
	var group [3]int
	var setNum int
	go func() {
		for i := range input {
			group[setNum] = i
			if setNum < 2 {
				setNum++
				continue
			}
			s := sum(group)
			output <- s
			shift(&group)
		}
		close(output)
	}()
}

func shift(buf *[3]int) {
	buf[0] = buf[1]
	buf[1] = buf[2]
	buf[2] = 0
}

func sum(buf [3]int) int {
	return buf[0] + buf[1] + buf[2]
}
