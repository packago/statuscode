package statuscode

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestReturn400BadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	Return400BadRequest(w, "Bad Request Msg")
	resp := w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("Return400BadRequest() StatusCode = %v want 400", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Bad Request Msg" {
		t.Errorf("Return400BadRequest() Message = %v want Bad Request Msg", respBody.Message)
	}
}

func TestReturn401Unauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	Return401Unauthorized(w, "Unauthorized Msg")
	resp := w.Result()
	if resp.StatusCode != 401 {
		t.Errorf("Return401Unauthorized() StatusCode = %v want 401", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Unauthorized Msg" {
		t.Errorf("Return401Unauthorized() Message = %v want Unauthorized Msg", respBody.Message)
	}
}

func TestReturn403Forbidden(t *testing.T) {
	w := httptest.NewRecorder()
	Return403Forbidden(w, "Fodbidden Msg")
	resp := w.Result()
	if resp.StatusCode != 403 {
		t.Errorf("Return403Forbidden() StatusCode = %v want 403", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Fodbidden Msg" {
		t.Errorf("Return403Forbidden() Message = %v want Fodbidden Msg", respBody.Message)
	}
}

func TestReturn404NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	Return404NotFound(w, "Not Found Msg")
	resp := w.Result()
	if resp.StatusCode != 404 {
		t.Errorf("Return404NotFound() StatusCode = %v want 404", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Not Found Msg" {
		t.Errorf("Return404NotFound() Message = %v want Not Found Msg", respBody.Message)
	}
}

func TestReturn409Conflict(t *testing.T) {
	w := httptest.NewRecorder()
	Return409Conflict(w, "Conflict Msg")
	resp := w.Result()
	if resp.StatusCode != 409 {
		t.Errorf("Return409Conflict() StatusCode = %v want 409", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Conflict Msg" {
		t.Errorf("Return409Conflict() Message = %v want Conflict Msg", respBody.Message)
	}
}

func TestReturn422UnprocessableEntity(t *testing.T) {
	w := httptest.NewRecorder()
	Return422UnprocessableEntity(w, "Unprocessable Entity Msg")
	resp := w.Result()
	if resp.StatusCode != 422 {
		t.Errorf("Return422UnprocessableEntity() StatusCode = %v want 422", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Unprocessable Entity Msg" {
		t.Errorf("Return422UnprocessableEntity() Message = %v want Unprocessable Entity Msg", respBody.Message)
	}
}

func TestReturn429TooManyRequests(t *testing.T) {
	w := httptest.NewRecorder()
	Return429TooManyRequests(w, "Too Many Requests Msg")
	resp := w.Result()
	if resp.StatusCode != 429 {
		t.Errorf("Return429TooManyRequests() StatusCode = %v want 429", resp.StatusCode)
	}
	var respBody struct {
		Message string `json:"message"`
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &respBody)
	if respBody.Message != "Too Many Requests Msg" {
		t.Errorf("Return429TooManyRequests() Message = %v want Too Many Requests Msg", respBody.Message)
	}
}
