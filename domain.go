package main

type Score struct {
	UserId int     `json:"userId"`
	Score  float64 `json:"score"`
}

type Request struct {
	Scores struct {
		Manager []Score `json:"managers"`
		Team    []Score `json:"team"`
		Others  []Score `json:"others"`
	} `json:"scores"`
}

type Data struct {
	Scores struct {
		Manager float32 `json:"manager"`
		Team    float32 `json:"team"`
		Others  float32 `json:"others"`
	} `json:"scores"`
}

type Response struct {
	Success bool     `json:"success"`
	Data    Data     `json:"data"`
	Errors  []string `json:"errors"`
}
