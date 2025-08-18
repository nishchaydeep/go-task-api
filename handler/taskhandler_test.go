package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddTaskValidation(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		wantStatus int
	}{
		{"Valid Task", `{"name":"Test Task","category":"work"}`, http.StatusCreated},
		{"Empty Name", `{"name":" ","category":"work"}`, http.StatusBadRequest},
		{"Invalid JSON", `{name:Invalid}`, http.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/tasks", strings.NewReader(tc.payload))
			rr := httptest.NewRecorder()
			AddTask(rr, req)

			if rr.Code != tc.wantStatus {
				t.Errorf("Expected status %d, got %d", tc.wantStatus, rr.Code)
			}
		})
	}
}

func TestGetTaskNotFound(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
	}{
		{"Task Not Found", "name=TaskDoesnotExist", http.StatusNotFound},
		{"Empty Query", "", http.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/tasks?"+tc.query, nil)
			rr := httptest.NewRecorder()
			GetTask(rr, req)

			if rr.Code != tc.wantStatus {
				t.Errorf("Expected status %d, got %d", tc.wantStatus, rr.Code)
			}
		})
	}
}
