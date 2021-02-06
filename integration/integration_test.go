// +build integration

package integration

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	data, err := ioutil.ReadFile("request.json")

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/score", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	expected, err := ioutil.ReadFile("response.json")
	assert.NoError(t, err)

	assert.JSONEq(t, string(expected), string(actual))
}
