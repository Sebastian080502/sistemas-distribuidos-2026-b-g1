package observability

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

const HeaderTraceparent = "traceparent"
const HeaderCorrelation = "X-Correlation-Id"

type Context struct {
	TraceID       string
	SpanID        string
	ParentSpanID  string
	CorrelationID string
}

func ParseTraceparent(h string) (traceID, parentSpan string, ok bool) {
	parts := strings.Split(strings.TrimSpace(h), "-")
	if len(parts) != 4 || parts[0] != "00" || len(parts[1]) != 32 || len(parts[2]) != 16 {
		return "", "", false
	}
	return parts[1], parts[2], true
}

func FormatTraceparent(c Context) string {
	return "00-" + c.TraceID + "-" + c.SpanID + "-01"
}

func NewID(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func NewContext(incomingTraceparent, incomingCorr string) Context {
	c := Context{SpanID: NewID(8), CorrelationID: incomingCorr}
	if tid, parent, ok := ParseTraceparent(incomingTraceparent); ok {
		c.TraceID = tid
		c.ParentSpanID = parent
	} else {
		c.TraceID = NewID(16)
	}
	if c.CorrelationID == "" {
		c.CorrelationID = NewID(16)
	}
	return c
}
