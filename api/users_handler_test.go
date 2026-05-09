package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsers(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		wantStatus     int
		wantLimit      int
		checkLimitBody bool
	}{
		{
			name:           "no limit param defaults to 20",
			query:          "",
			wantStatus:     http.StatusOK,
			wantLimit:      20,
			checkLimitBody: true,
		},
		{
			name:           "limit=5 returns 5",
			query:          "?limit=5",
			wantStatus:     http.StatusOK,
			wantLimit:      5,
			checkLimitBody: true,
		},
		{
			name:           "limit=101 returns 400",
			query:          "?limit=101",
			wantStatus:     http.StatusBadRequest,
			checkLimitBody: false,
		},
		{
			name:           "limit=0 returns 400",
			query:          "?limit=0",
			wantStatus:     http.StatusBadRequest,
			checkLimitBody: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/users"+tc.query, nil)
			rr := httptest.NewRecorder()

			ListUsers(rr, req)

			if rr.Code != tc.wantStatus {
				t.Errorf("status = %d, want %d", rr.Code, tc.wantStatus)
			}

			if tc.checkLimitBody {
				var body map[string]any
				if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
					t.Fatalf("failed to decode response body: %v", err)
				}
				// JSON numbers decode as float64
				gotLimit, ok := body["limit"].(float64)
				if !ok {
					t.Fatalf("limit field missing or wrong type in body: %v", body)
				}
				if int(gotLimit) != tc.wantLimit {
					t.Errorf("limit = %d, want %d", int(gotLimit), tc.wantLimit)
				}
			}
		})
	}
}
