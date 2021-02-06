package domain

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore_Validate(t *testing.T) {
	tests := []struct {
		id    int
		score float64
		want  error
	}{
		{1, 1, nil},
		{1, 5, nil},
		{1, 0, errors.New("invalid score detected at id 1")},
		{1, 6, errors.New("invalid score detected at id 1")},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("score.Validate, %d %f", tt.id, tt.score)
		t.Run(name, func(t *testing.T) {
			s := &Score{
				UserId: tt.id,
				Score:  tt.score,
			}

			err := s.Validate()
			if tt.want != nil {
				assert.EqualError(t, err, tt.want.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestScores_Error(t *testing.T) {
	scores := &ScoreValidationError{
		Errors: []error{
			errors.New("empty managers score"),
			errors.New("empty team score"),
			errors.New("empty others score"),
		},
	}

	actual := scores.Error()
	expected := "empty managers score\nempty team score\nempty others score\n"
	assert.Equal(t, expected, actual)
}

func TestScores_ValidateEmptyScore(t *testing.T) {
	scores := Scores{}

	actual := scores.Validate()
	expected := &ScoreValidationError{
		Errors: []error{
			errors.New("empty managers score"),
			errors.New("empty team score"),
			errors.New("empty others score"),
		},
	}
	assert.Equal(t, expected, actual)
}

func TestScores_ValidateDuplicateID(t *testing.T) {
	scores := Scores{
		Manager: []Score{
			{1, 1},
			{2, 2},
		},
		Team: []Score{
			{3, 3},
			{4, 4},
		},
		Others: []Score{
			{5, 5},
			{1, 1},
		},
	}

	actual := scores.Validate()
	expected := &ScoreValidationError{
		Errors: []error{
			errors.New("duplicate id found"),
		},
	}
	assert.Equal(t, expected, actual)
}

func TestScores_ValidateInvalidScore(t *testing.T) {
	scores := Scores{
		Manager: []Score{
			{0, 0},
			{1, 1},
			{2, 2},
		},
		Team: []Score{
			{3, 3},
			{4, 4},
		},
		Others: []Score{
			{5, 5},
			{6, 6},
		},
	}

	actual := scores.Validate()
	expected := &ScoreValidationError{
		Errors: []error{
			errors.New("invalid score detected at id 0"),
			errors.New("invalid score detected at id 6"),
		},
	}
	assert.Equal(t, expected, actual)
}

func TestScores_ValidateNoError(t *testing.T) {
	scores := Scores{
		Manager: []Score{
			{1, 1},
		},
		Team: []Score{
			{2, 2},
		},
		Others: []Score{
			{3, 3},
		},
	}

	actual := scores.Validate()
	expected := &ScoreValidationError{
		Errors: []error{},
	}
	assert.Equal(t, expected, actual)
}
