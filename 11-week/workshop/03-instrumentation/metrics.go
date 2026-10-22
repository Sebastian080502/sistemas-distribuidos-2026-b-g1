package observability

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mu      sync.Mutex
	total   = map[string]int{"ok": 0, "error": 0}
	samples []int64
)

func Observe(outcome string, durationMS int64) {
	mu.Lock()
	defer mu.Unlock()
	total[outcome]++
	samples = append(samples, durationMS)
}

func MetricsHandler(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "reservation_requests_total{route=\"on_payment_confirmed\",outcome=\"ok\"} %d\n", total["ok"])
	fmt.Fprintf(w, "reservation_requests_total{route=\"on_payment_confirmed\",outcome=\"error\"} %d\n", total["error"])
	le := []int64{50, 100, 300, 1000, 3000}
	counts := make([]int, len(le))
	var sum int64
	for _, s := range samples {
		sum += s
		for i, b := range le {
			if s <= b {
				counts[i]++
			}
		}
	}
	for i, b := range le {
		fmt.Fprintf(w, "reservation_request_duration_milliseconds_bucket{le=\"%d\"} %d\n", b, counts[i])
	}
	fmt.Fprintf(w, "reservation_request_duration_milliseconds_bucket{le=\"+Inf\"} %d\n", len(samples))
	fmt.Fprintf(w, "reservation_request_duration_milliseconds_sum %d\n", sum)
	fmt.Fprintf(w, "reservation_request_duration_milliseconds_count %d\n", len(samples))
}
