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
