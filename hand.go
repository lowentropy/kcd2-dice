package dice

type Pool map[string]int

type Hand [6]string

func (h Hand) Score(n int) int {
	roll := Roll{}
	for _, t := range h {
		roll = append(roll, DieRoll{Type: t})
	}
	var total int
	for range n {
		roll.Reroll()
		score := roll.Score()
		total += score
	}
	return total / n
}

func (p Pool) Best(n int) (Hand, int) {
	var flat []string
	for t, k := range p {
		for range k {
			flat = append(flat, t)
		}
	}
	var best Hand
	var max int
	eachCombination(flat, func(h Hand) {
		if s := h.Score(n); s > max {
			best, max = h, s
		}
	})
	return best, max
}

// Call the given function for each selection of
// 6 dice from the given list.
func eachCombination(list []string, f func(Hand)) {
	var combine func([]string, int, Hand, int)
	combine = func(remaining []string, pos int, current Hand, start int) {
		// Base case: if we've filled all 6 positions
		if pos == 6 {
			f(current)
			return
		}

		// Try each remaining die starting from 'start' index
		for i := start; i < len(remaining); i++ {
			next := current
			next[pos] = remaining[i]
			combine(remaining, pos+1, next, i+1)
		}
	}

	// Start the recursion with all dice, position 0, and start index 0
	combine(list, 0, Hand{}, 0)
}
