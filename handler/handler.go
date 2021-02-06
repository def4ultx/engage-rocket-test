package handler

import (
	"net/http"

	"engage-rocket-test/domain"
)

type CalculatedScore struct {
	Manager *float64 `json:"manager,omitempty"`
	Team    *float64 `json:"team,omitempty"`
	Others  *float64 `json:"others,omitempty"`
}

type Request struct {
	Scores domain.Scores `json:"scores"`
}

type ResponseData struct {
	Scores CalculatedScore `json:"scores"`
}

func CalculateScoreHandler(w http.ResponseWriter, r *http.Request) {
}
