package observability

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIncomingTraceIDEqualsDownstream(t *testing.T) {
	var got string
	down := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, _, _ = ParseTraceparent(r.Header.Get(HeaderTraceparent))
		w.WriteHeader(http.StatusAccepted)
	}))
	defer down.Close()

	in := "00-7a3c91e2b04d4f1a8c6e5d9f2b1a0843-1111111111111111-01"
	var buf strings.Builder
	log := Logger{Out: &buf, Service: "drp-reservation-api", Version: "2.0.0", Environment: "qa"}
	mux := http.NewServeMux()
	mux.Handle("/metrics", http.HandlerFunc(MetricsHandler))
	mux.Handle("/consume", TraceMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := FromContext(r.Context())
		_ = (Client{HTTP: down.Client(), Log: log}).PostEvent(down.URL, ctx)
		w.WriteHeader(http.StatusNoContent)
	})))

	req := httptest.NewRequest(http.MethodPost, "/consume", nil)
	req.Header.Set(HeaderTraceparent, in)
	req.Header.Set(HeaderCorrelation, "9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if got != "7a3c91e2b04d4f1a8c6e5d9f2b1a0843" {
		t.Fatalf("downstream trace_id=%s", got)
	}
	line, _ := io.ReadAll(strings.NewReader(buf.String()))
	if !strings.Contains(string(line), `"trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843"`) {
		t.Fatalf("log missing trace: %s", line)
	}
	mrec := httptest.NewRecorder()
	mux.ServeHTTP(mrec, httptest.NewRequest(http.MethodGet, "/metrics", nil))
	if !strings.Contains(mrec.Body.String(), "reservation_requests_total") {
		t.Fatal("metrics missing")
	}
}

