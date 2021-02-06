package domain

type Score struct {
	UserId int     `json:"userId"`
	Score  float64 `json:"score"`
}

type Scores struct {
	Manager []Score `json:"managers"`
	Team    []Score `json:"team"`
	Others  []Score `json:"others"`
}
