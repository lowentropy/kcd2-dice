package dice

import (
	"testing"
)

func TestRollScore(t *testing.T) {
	tests := []struct {
		name     string
		dice     []int
		expected int
	}{
		{"Single 1", []int{1}, 100},
		{"Single 5", []int{5}, 50},
		{"Three 1s", []int{1, 1, 1}, 1000},
		{"Four 1s", []int{1, 1, 1, 1}, 2000},
		{"Three 3s", []int{3, 3, 3}, 300},
		{"Four 3s", []int{3, 3, 3, 3}, 600},
		{"Five 3s", []int{3, 3, 3, 3, 3}, 1200},
		{"Straight 1-6", []int{1, 2, 3, 4, 5, 6}, 1500},
		{"Straight 2-6", []int{2, 3, 4, 5, 6}, 750},
		{"Straight 1-5", []int{1, 2, 3, 4, 5}, 500},
		{"Mixed combo", []int{1, 1, 1, 5, 5, 6}, 1100},
		{"With joker straight", []int{1, 2, 3, 4, 5, 0}, 1500},
		{"With joker set", []int{3, 3, 0}, 300},
		{"Two sets", []int{3, 3, 3, 5, 5, 5}, 800},
		{"Two jokers", []int{1, 0, 0, 5, 5, 5}, 2100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roll := make(Roll, len(tt.dice))
			for i, d := range tt.dice {
				if d == 0 {
					roll[i] = DieRoll{Joker: true}
				} else {
					roll[i] = DieRoll{Side: d}
				}
			}
			if got := roll.Score(); got != tt.expected {
				t.Errorf("Roll.Score() = %v, want %v", got, tt.expected)
			}
		})
	}
}
