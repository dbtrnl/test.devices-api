//go:build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

const baseURL = "http://localhost:8080"

func TestListDevices_NoFilters(t *testing.T) {
	res, err := http.Get(baseURL + "/devices")
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	var devices []map[string]any
	err = json.NewDecoder(res.Body).Decode(&devices)

	require.NoError(t, err)
	require.Len(t, devices, 7)
}

func TestListDevices_WithBothFilters(t *testing.T) {
	queryParams := url.Values{}
	queryParams.Add("brand", "Brand number one")
	queryParams.Add("state", "available")

	res, err := http.Get(baseURL + "/devices?" + queryParams.Encode())
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		getResponse(t, res)
		t.FailNow()
	}

	var devices []map[string]any
	err = json.NewDecoder(res.Body).Decode(&devices)

	require.Nil(t, err)
	require.Len(t, devices, 1)
}

func TestListDevices_WithBrandFilter(t *testing.T) {
	queryParams := url.Values{}
	queryParams.Add("brand", "Brand number one")

	res, err := http.Get(baseURL + "/devices?" + queryParams.Encode())
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		getResponse(t, res)
		t.FailNow()
	}

	var devices []map[string]any
	err = json.NewDecoder(res.Body).Decode(&devices)

	require.Nil(t, err)
	require.Len(t, devices, 3)
}

func TestListDevices_WithNonExistentBrand(t *testing.T) {
	queryParams := url.Values{}
	queryParams.Add("brand", "This Brand doesn't exist")

	res, err := http.Get(baseURL + "/devices?" + queryParams.Encode())
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		getResponse(t, res)
		t.FailNow()
	}

	var devices []map[string]any
	err = json.NewDecoder(res.Body).Decode(&devices)

	require.Nil(t, err)
	require.Len(t, devices, 0)
}

func TestListDevice_NonExistent(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/devices/3c46c5a3-4fa9-4ce1-8693-12e467f03731", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestListDevice_Success(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/devices/3c46c5a3-4fa9-4ce1-8693-12e467f03730", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusOK, res.StatusCode)
}
func TestCreateDevice_InvalidNameInvalidBrand(t *testing.T) {
	payload := map[string]any{
		"name":  "a",
		"brand": "a",
	}

	body, _ := json.Marshal(payload)

	res, err := http.Post(baseURL+"/devices", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateDevice_Success(t *testing.T) {
	payload := map[string]any{
		"name":  "Integration Test Create Device",
		"brand": "Integration Test Create Test Brand",
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

func TestDeleteDevice_InvalidUUID(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/devices/invalid-uuid", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestDeleteDevice_InUse(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/devices/d3f32eba-fc67-4440-b054-1cf821042610", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusConflict, res.StatusCode)
}

func TestDeleteDevice_AlreadyDeleted(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/devices/3c46c5a3-4fa9-4ce1-8693-12e467f03730", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestDeleteDevice_Success(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/devices/77b85ecb-e767-4f2d-a3cf-be53dae49274", nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		getResponse(t, res)
		t.FailNow()
	}

	require.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestUpdateDevice_FullUpdate(t *testing.T) {
	var respBody []byte
	externalID := "f3b4ea8f-1e69-4736-9cc7-f8775761fefd"

	payload := map[string]any{
		"name":  "Updated Name",
		"brand": "Updated Brand",
		"state": "available",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPatch, baseURL+"/devices/"+externalID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		respBody = getResponse(t, res)
		t.FailNow()
	} else {
		respBody = getResponse(t, res)
	}

	var updated map[string]any
	require.NoError(t, json.Unmarshal(respBody, &updated))

	require.Equal(t, "Updated Name", updated["name"])
	require.Equal(t, "Updated Brand", updated["brand"])
	require.Equal(t, "available", updated["state"])
}

func TestUpdateDevice_PartialUpdate(t *testing.T) {
	var respBody []byte

	externalID := "c343dabd-0afd-4f5b-a632-42be259df112"

	payload := map[string]any{
		"name": "Partial Update Name",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPatch, baseURL+"/devices/"+externalID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		respBody = getResponse(t, res)
		t.FailNow()
	} else {
		respBody = getResponse(t, res)
	}

	var updated map[string]any
	require.NoError(t, json.Unmarshal(respBody, &updated))

	require.Equal(t, "Partial Update Name", updated["name"])
	require.Equal(t, "Brand number one", updated["brand"])
	require.Equal(t, "inactive", updated["state"])
}

func TestUpdateDevice_InUse_UpdateName(t *testing.T) {
	externalID := "d3f32eba-fc67-4440-b054-1cf821042610"

	payload := map[string]any{
		"name": "Update Name",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPatch, baseURL+"/devices/"+externalID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusConflict {
		_ = getResponse(t, res)
		t.FailNow()
	}
}

func TestUpdateDevice_InUse_UpdateBrand(t *testing.T) {
	externalID := "d3f32eba-fc67-4440-b054-1cf821042610"

	payload := map[string]any{
		"brand": "Update Brand",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPatch, baseURL+"/devices/"+externalID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusConflict {
		_ = getResponse(t, res)
		t.FailNow()
	}
}

func getResponse(t *testing.T, res *http.Response) []byte {
	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	t.Logf("StatusCode: %d", res.StatusCode)
	t.Logf("Body: %s", string(body))

	return body
}
