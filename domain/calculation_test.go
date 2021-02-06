package domain

import (
	"testing"
)

func TestAverageScore(t *testing.T) {
	tests := []struct {
		name string
		want float64

		scores  []Score
		minSize int
	}{
		{
			name:    "nil score, return 0",
			want:    0.0,
			scores:  nil,
			minSize: 1,
		},
		{
			name:    "empty score, return 0",
			want:    0.0,
			scores:  []Score{},
			minSize: 1,
		},
		{
			name:    "empty score and -1 minSize, return 0",
			want:    0.0,
			scores:  []Score{},
			minSize: -1,
		},
		{
			name: "size of score less than minSize, return 0",
			want: 0.0,
			scores: []Score{
				{0, 1},
			},
			minSize: 2,
		},
		{
			name: "size of score equal minSize, return 0",
			want: 0,
			scores: []Score{
				{0, 1},
				{0, 5},
				{0, 3},
			},
			minSize: 3,
		},
		{
			name: "size of score more than minSize, return 2.75",
			want: 2.75,
			scores: []Score{
				{0, 1},
				{0, 5},
				{0, 3},
				{0, 2},
			},
			minSize: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AverageScore(tt.scores, tt.minSize); got != tt.want {
				t.Errorf("AverageScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
