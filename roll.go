package dice

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

type Roll []DieRoll

type DieRoll struct {
	Type  string
	Side  int
	Joker bool
}

var pow2 = [7]int{0, 0, 0, 1, 2, 4, 8}

// Pre-calculate base scores for sets
var baseScores = [7]int{0, 1000, 200, 300, 400, 500, 600}

// Pre-calculate straight patterns
var straights = [3]struct {
	score int
	sides []int
}{
	{1500, []int{1, 2, 3, 4, 5, 6}},
	{750, []int{2, 3, 4, 5, 6}},
	{500, []int{1, 2, 3, 4, 5}},
}

// Pre-calculate single die scores
var singleScores = [6]int{100, 0, 0, 0, 50, 0}

func (r Roll) Score() int {
	var jokers int
	var counts [6]int

	// separate jokers and sides
	for i := range r {
		if r[i].Joker {
			jokers++
		} else {
			counts[r[i].Side-1]++
		}
	}

	var total int
	for {
		var usedSides [6]int
		var usedJokers int
		var best int

		// count sets first, keeping the best only
		for i := 1; i <= 6; i++ {
			c := counts[i-1]
			t := c + jokers
			if t < 3 {
				continue
			}
			if s := baseScores[i] * pow2[t]; s > best {
				best, usedJokers = s, jokers
				clear(usedSides[:])
				usedSides[i-1] = c
			}
		}
		// check straights against the set, keeping the best
		for _, s := range straights {
			if score, j, u := checkStraight(counts, jokers, s.score, s.sides); score > best {
				best, usedJokers, usedSides = score, j, u
			}
		}
		// if no sets or straights, count singles and return
		if best == 0 {
			for i := range 6 {
				total += counts[i] * singleScores[i]
			}
			return total
		}
		// otherwise, remove used dice and repeat
		total += best
		jokers -= usedJokers
		for i := range 6 {
			counts[i] -= usedSides[i]
		}
	}
}

func checkStraight(counts [6]int, jokers, score int, sides []int) (int, int, [6]int) {
	var used [6]int
	var missing int
	// see what's missing from the stgraight
	for _, s := range sides {
		if counts[s-1] == 0 {
			missing++
		} else {
			used[s-1] = 1
		}
	}
	// score it if we can use jokers
	if missing <= jokers {
		return score, missing, used
	}
	// otherwise return 0 score
	return 0, 0, used
}

func (r Roll) Reroll() {
	for i := range r {
		r[i].Reroll()
	}
}

func (r Roll) String() string {
	var sides []string
	slices.SortFunc(r, func(a, b DieRoll) int {
		return cmp.Compare(a.Side, b.Side)
	})
	for _, d := range r {
		var s string
		if d.Joker {
			s = "J"
		} else {
			s = fmt.Sprintf("%d", d.Side)
		}
		sides = append(sides, s)
	}
	return strings.Join(sides, " ")
}

func (d *DieRoll) Reroll() {
	ws := Types[d.Type]
	var i, w, n, s int
	for _, w = range ws {
		n += w
	}
	r := rand.IntN(n)
	for i, w = range ws {
		s += w
		if r < s {
			break
		}
	}
	if i == 0 && d.Type == "Devil's Head" {
		d.Side, d.Joker = 0, true
	} else {
		d.Side, d.Joker = i+1, false
	}
}
