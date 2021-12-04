package processes

func Comparer(input <-chan int, output chan<- int) {
	go func() {
		var (
			prev          = 0
			prevSet       = false
			increaseCount = 0
		)
		for i := range input {
			if !prevSet {
				prevSet = true
				prev = i
				continue
			}
			if i > prev {
				increaseCount++
			}
			prev = i
		}
		output <- increaseCount
		close(output)
	}()
}
