package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nishchaydeep15/go-task-api/model"
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

func TestGetTask(t *testing.T) {
	AddTask(httptest.NewRecorder(), httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name":"GetMe","category":"testing"}`)))

	req := httptest.NewRequest("GET", "/tasks?name=GetMe", nil)
	rr := httptest.NewRecorder()
	GetTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", rr.Code)
	}

	var task model.Task
	if err := json.NewDecoder(rr.Body).Decode(&task); err != nil {
		t.Fatal("Invalid JSON response")
	}
	if task.Name != "GetMe" {
		t.Errorf("Expected task name GetMe, got %s", task.Name)
	}
}
