package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestStruct struct {
	ID    int
	Name  string
	Email string
}

func TestCsvDataResponse(t *testing.T) {
	tests := []struct {
		name           string
		dataList       any
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid data list",
			dataList:       []TestStruct{{ID: 1, Name: "John Doe", Email: "john@example.com"}, {ID: 2, Name: "Jane Doe", Email: "jane@example.com"}},
			expectedStatus: http.StatusOK,
			expectedBody:   "ID,Name,Email\n1,John Doe,john@example.com\n2,Jane Doe,jane@example.com\n",
		},
		{
			name:           "Empty data list",
			dataList:       []TestStruct{},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Empty data list\n",
		},
		{
			name:           "Invalid data type",
			dataList:       "invalid data",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid data type: expected slice\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			csvDataResponse(rr, req, tt.dataList)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(tt.expectedBody) {
				t.Errorf("expected body %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}
