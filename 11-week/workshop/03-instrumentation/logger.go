package observability

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type Line struct {
	Timestamp        string `json:"timestamp"`
	Level            string `json:"level"`
	Service          string `json:"service"`
	Version          string `json:"version"`
	Environment      string `json:"environment"`
	TraceID          string `json:"trace_id"`
	SpanID           string `json:"span_id"`
	CorrelationID    string `json:"correlation_id"`
	Event            string `json:"event"`
	Message          string `json:"message"`
	DurationMS       int64  `json:"duration_ms"`
	ReservationID    string `json:"reservation_id,omitempty"`
	PaymentID        string `json:"payment_id,omitempty"`
	SpaceID          string `json:"space_id,omitempty"`
	ReservationState string `json:"reservation_state,omitempty"`
}

type Logger struct {
	Out         io.Writer
	Service     string
	Version     string
	Environment string
}

func NewLogger() Logger {
	return Logger{Out: os.Stdout, Service: "drp-reservation-api", Version: "2.0.0", Environment: "qa"}
}

func (l Logger) Write(ctx Context, level, event, message string, durationMS int64, extra Line) {
	if l.Out == nil {
		l.Out = os.Stdout
	}
	row := extra
	row.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	row.Level = level
	row.Service = l.Service
	row.Version = l.Version
	row.Environment = l.Environment
	row.TraceID = ctx.TraceID
	row.SpanID = ctx.SpanID
	row.CorrelationID = ctx.CorrelationID
	row.Event = event
	row.Message = message
	row.DurationMS = durationMS
	b, _ := json.Marshal(row)
	_, _ = l.Out.Write(append(b, '\n'))
}
