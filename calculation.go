package main

func AverageScore(scores []Score, minSize int) float64 {
	if len(scores) == 0 {
		return 0
	}
	if minSize >= len(scores) {
		return 0
	}

	var sum float64
	for _, v := range scores {
		sum += v.Score
	}

	average := sum / float64(len(scores))
	return average
}
