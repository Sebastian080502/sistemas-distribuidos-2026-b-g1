package authz

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mirrors drp-reservation-api POST /internal/events (E-23, not on the gateway).
// Story: [security] TB6 rejects a confirm hop without X-Internal-Token.

func internalEvents(token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/internal/events" {
			http.NotFound(w, r)
			return
		}
		if token == "" {
			http.Error(w, `{"error":"SERVICE_UNAVAILABLE","message":"El consumo interno no está configurado"}`, http.StatusServiceUnavailable)
			return
		}
		got := strings.TrimSpace(r.Header.Get("X-Internal-Token"))
		if got == "" || got != token {
			http.Error(w, `{"error":"UNAUTHORIZED","message":"Token de autenticación requerido"}`, http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
}

func TestInternalEventsWithoutTokenIs401(t *testing.T) {
	srv := httptest.NewServer(internalEvents("spacehub-dev-internal"))
	defer srv.Close()
	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/internal/events", strings.NewReader(
		`{"eventId":"77777777-7777-4777-8777-777777777777","eventType":"PaymentConfirmed","payload":{"reservationId":"22222222-2222-2222-2222-222222222222"}}`,
	))
	req.Header.Set("Authorization", "Bearer user-jwt")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", res.StatusCode)
	}
}
