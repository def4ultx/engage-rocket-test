package handler

type CalculatedScore struct {
	Manager *float64 `json:"manager,omitempty"`
	Team    *float64 `json:"team,omitempty"`
	Others  *float64 `json:"others,omitempty"`
}
type Request struct {
	Scores Scores `json:"scores"`
}

type ResponseData struct {
	Scores CalculatedScore `json:"scores"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}
