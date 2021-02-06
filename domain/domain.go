package main

type Score struct {
	UserId int     `json:"userId"`
	Score  float64 `json:"score"`
}

type Scores struct {
	Manager []Score `json:"managers"`
	Team    []Score `json:"team"`
	Others  []Score `json:"others"`
}

type Request struct {
	Scores Scores `json:"scores"`
}

type CalculatedScore struct {
	Manager *float64 `json:"manager,omitempty"`
	Team    *float64 `json:"team,omitempty"`
	Others  *float64 `json:"others,omitempty"`
}

type ResponseData struct {
	Scores CalculatedScore `json:"scores"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}
