package dice

import (
	"fmt"
	"sort"
	"testing"
)

func TestHandScore(t *testing.T) {
	p := Pool{
		"Fair":               2,
		"Odd":                2,
		"Even":               1,
		"King's":             1,
		"Lu":                 1,
		"Lucky":              1,
		"Painted":            1,
		"Strip":              1,
		"Cautious Cheater's": 1,
		"Favourable":         1,
		"Saint Antiochus'":   1,
	}
	h, score := p.Best(1000)
	fmt.Printf("best hand: %v, score: %d\n", h, score)
}

func TestEachCombination(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []Hand
	}{
		{
			name:  "exactly 6 elements",
			input: []string{"A", "B", "C", "D", "E", "F"},
			expected: []Hand{
				{"A", "B", "C", "D", "E", "F"},
			},
		},
		{
			name:  "7 elements",
			input: []string{"A", "B", "C", "D", "E", "F", "G"},
			expected: []Hand{
				{"A", "B", "C", "D", "E", "F"},
				{"A", "B", "C", "D", "E", "G"},
				{"A", "B", "C", "D", "F", "G"},
				{"A", "B", "C", "E", "F", "G"},
				{"A", "B", "D", "E", "F", "G"},
				{"A", "C", "D", "E", "F", "G"},
				{"B", "C", "D", "E", "F", "G"},
			},
		},
		{
			name:  "8 elements",
			input: []string{"A", "B", "C", "D", "E", "F", "G", "H"},
			expected: []Hand{
				{"A", "B", "C", "D", "E", "F"},
				{"A", "B", "C", "D", "E", "G"},
				{"A", "B", "C", "D", "E", "H"},
				{"A", "B", "C", "D", "F", "G"},
				{"A", "B", "C", "D", "F", "H"},
				{"A", "B", "C", "D", "G", "H"},
				{"A", "B", "C", "E", "F", "G"},
				{"A", "B", "C", "E", "F", "H"},
				{"A", "B", "C", "E", "G", "H"},
				{"A", "B", "C", "F", "G", "H"},
				{"A", "B", "D", "E", "F", "G"},
				{"A", "B", "D", "E", "F", "H"},
				{"A", "B", "D", "E", "G", "H"},
				{"A", "B", "D", "F", "G", "H"},
				{"A", "B", "E", "F", "G", "H"},
				{"A", "C", "D", "E", "F", "G"},
				{"A", "C", "D", "E", "F", "H"},
				{"A", "C", "D", "E", "G", "H"},
				{"A", "C", "D", "F", "G", "H"},
				{"A", "C", "E", "F", "G", "H"},
				{"A", "D", "E", "F", "G", "H"},
				{"B", "C", "D", "E", "F", "G"},
				{"B", "C", "D", "E", "F", "H"},
				{"B", "C", "D", "E", "G", "H"},
				{"B", "C", "D", "F", "G", "H"},
				{"B", "C", "E", "F", "G", "H"},
				{"B", "D", "E", "F", "G", "H"},
				{"C", "D", "E", "F", "G", "H"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []Hand
			eachCombination(tt.input, func(h Hand) {
				result = append(result, h)
			})

			if len(result) != len(tt.expected) {
				t.Errorf("got %d combinations, want %d", len(result), len(tt.expected))
			}

			// Sort both slices to ensure consistent ordering for comparison
			sortHands(result)
			sortHands(tt.expected)

			// Compare each hand
			for i := range result {
				if !handsEqual(result[i], tt.expected[i]) {
					t.Errorf("combination %d:\ngot  %v\nwant %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// Helper function to sort a slice of Hands
func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		for k := 0; k < 6; k++ {
			if hands[i][k] != hands[j][k] {
				return hands[i][k] < hands[j][k]
			}
		}
		return false
	})
}

// Helper function to compare two Hands
func handsEqual(a, b Hand) bool {
	for i := 0; i < 6; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
