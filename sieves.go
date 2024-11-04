package primes

func Eratosthenes(end int) []int {
	composites := make([]bool, end)
	p := 2

outer:
	for {
		// Mark all multiples of p as composite.
		for i := p * 2; i < end; i += p {
			composites[i] = true
		}

		// Find the next value of `p` by checking for the next unmarked number in the sequence,
		// and exit the entire thing if we exceed `end`.
		for i := p + 1; ; i++ {
			if i >= end {
				break outer
			}

			if !composites[i] {
				p = i
				break
			}
		}
	}

	// Convert the composites to actual numbers and return.
	var out []int
	for i, composite := range composites {
		if !composite {
			out = append(out, i)
		}
	}

	return out
}
