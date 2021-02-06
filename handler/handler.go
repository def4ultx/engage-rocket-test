package handler

import (
	"encoding/json"
	"errors"
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

type CalculatedScoreResponse struct {
	Scores CalculatedScore `json:"scores"`
}

func CalculateScoreHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := []error{errors.New("cannot decode request body")}
		writeErrorResponse(w, msg)
		return
	}

	errs, ok := req.Scores.Validate().(*domain.ScoreValidationError)
	if ok && len(errs.Errors) != 0 {
		writeErrorResponse(w, errs.Errors)
		return
	}

	minManagerSize, minTeamSize, minOtherSize := 0, 3, 3
	var (
		managerScore = domain.AverageScore(req.Scores.Manager, minManagerSize)
		teamScore    = domain.AverageScore(req.Scores.Team, minTeamSize)
		othersScore  = domain.AverageScore(req.Scores.Others, minOtherSize)
	)

	scores := CalculatedScore{
		Manager: &managerScore,
		Team:    &teamScore,
		Others:  &othersScore,
	}
	body := CalculatedScoreResponse{scores}
	writeSuccessResponse(w, &body)
}
