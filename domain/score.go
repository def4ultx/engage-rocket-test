package domain

import (
	"errors"
	"fmt"
)

type Score struct {
	UserId int     `json:"userId"`
	Score  float64 `json:"score"`
}

type Scores struct {
	Manager []Score `json:"managers"`
	Team    []Score `json:"team"`
	Others  []Score `json:"others"`
}

type ScoreValidationError struct {
	Errors []error
}

func (e *ScoreValidationError) Error() string {
	var str string
	for _, v := range e.Errors {
		str += v.Error()
		str += "\n"
	}
	return str
}

func (s *Score) Validate() error {
	if s.Score < 1 || s.Score > 5 {
		msg := fmt.Sprintf("invalid score detected at id %d", s.UserId)
		return errors.New(msg)
	}
	return nil
}

func (s *Scores) Validate() error {
	errs := ScoreValidationError{
		make([]error, 0),
	}

	if len(s.Manager) == 0 {
		errs.Errors = append(errs.Errors, errors.New("empty managers score"))
	}
	if len(s.Team) == 0 {
		errs.Errors = append(errs.Errors, errors.New("empty team score"))
	}
	if len(s.Others) == 0 {
		errs.Errors = append(errs.Errors, errors.New("empty others score"))
	}

	var scores []Score
	scores = append(scores, s.Manager...)
	scores = append(scores, s.Team...)
	scores = append(scores, s.Others...)

	var err error
	for _, v := range scores {
		err = v.Validate()
		if err != nil {
			errs.Errors = append(errs.Errors, err)
		}
	}

	ids := map[int]struct{}{}
	for _, v := range scores {
		_, ok := ids[v.UserId]
		if ok {
			errs.Errors = append(errs.Errors, errors.New("duplicate id found"))
			break
		}
		ids[v.UserId] = struct{}{}
	}

	return &errs
}
