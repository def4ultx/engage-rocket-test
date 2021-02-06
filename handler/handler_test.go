package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateScoreHandler_Success(t *testing.T) {

	h := http.HandlerFunc(CalculateScoreHandler)
	r := httptest.NewRecorder()

	body := `
	{
		"scores": {
			"managers": [
				{ "userId": 1, "score": 1 },
				{ "userId": 2, "score": 5 }
			],
			"team": [
				{ "userId": 4, "score": 1 },
				{ "userId": 5, "score": 5 },
				{ "userId": 6, "score": 3 },
				{ "userId": 7, "score": 2 }
			],
			"others": [
				{ "userId": 8, "score": 1 },
				{ "userId": 9, "score": 5 }
			]
		}
	}
	`
	req := httptest.NewRequest(http.MethodPost, "/score", strings.NewReader(body))
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	expected := `
	{
		"success": true,
		"data": {
			"scores": {
				"manager": 3,
				"team": 2.75,
				"others": 0
			}
		},
		"errors": []
	}
	`
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func TestCalculateScoreHandler_InvalidRequest(t *testing.T) {

	h := http.HandlerFunc(CalculateScoreHandler)
	r := httptest.NewRecorder()

	body := ``
	req := httptest.NewRequest(http.MethodPost, "/score", strings.NewReader(body))
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)

	expected := `
	{
		"success": false,
		"data": {},
		"errors": [
			"cannot decode request body"
		]
	}
	`
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(actual))
}
