//go:build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const baseURL = "http://localhost:8080"

func TestCreateDevice(t *testing.T) {
	payload := map[string]any{
		"name":  "Integration Device",
		"brand": "Test Brand",
	}

	body, _ := json.Marshal(payload)

	res1, err := http.Post(baseURL+"/devices", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer res1.Body.Close()

	require.Equal(t, http.StatusCreated, res1.StatusCode)

	var created map[string]any
	err = json.NewDecoder(res1.Body).Decode(&created)
	require.NoError(t, err)

	// Testing idempotency. If same payload is sent, 2nd status code must be 200 and the 2nd response must be the same as the 1st.
	res2, err := http.Post(baseURL+"/devices", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer res2.Body.Close()

	require.Equal(t, http.StatusOK, res2.StatusCode)

	var existing map[string]any
	err = json.NewDecoder(res2.Body).Decode(&existing)
	require.NoError(t, err)

	require.Equal(t, created, existing)
}