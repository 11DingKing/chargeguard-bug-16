package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	ResetTaskHTTPState()
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("POST", "/task", nil))
	if rr.Code != http.StatusBadRequest || rr.Body.String() != "{\"active\":0,\"stored\":0}\n" {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
}
