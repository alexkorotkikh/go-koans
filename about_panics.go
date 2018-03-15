package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	defer func() {
		if err := recover(); err != nil {
			// recovered
		}
	}()

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
