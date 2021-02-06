package handler

import (
	"encoding/json"
	"net/http"
)

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

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	resp := Response{
		Success: true,
		Data:    data,
		Errors:  []string{},
	}
	writeResponse(w, http.StatusOK, resp)
}

func writeErrorResponse(w http.ResponseWriter, errors []error) {
	messages := make([]string, len(errors))
	for i, v := range errors {
		messages[i] = v.Error()
	}

	resp := Response{
		Success: false,
		Data:    struct{}{},
		Errors:  messages,
	}
	writeResponse(w, http.StatusBadRequest, resp)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
