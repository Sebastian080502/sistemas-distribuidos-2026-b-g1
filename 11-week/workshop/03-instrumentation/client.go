package observability

import (
	"net/http"
	"time"
)

func Inject(req *http.Request, ctx Context) {
	req.Header.Set(HeaderTraceparent, FormatTraceparent(ctx))
	req.Header.Set(HeaderCorrelation, ctx.CorrelationID)
}

type Client struct {
	HTTP *http.Client
	Log  Logger
}

func (c Client) PostEvent(url string, ctx Context) error {
	start := time.Now()
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	Inject(req, ctx)
	if c.HTTP == nil {
		c.HTTP = http.DefaultClient
	}
	res, err := c.HTTP.Do(req)
	ms := time.Since(start).Milliseconds()
	if err != nil {
		Observe("error", ms)
		c.Log.Write(ctx, "ERROR", "reservation.outbox.publish.fail", err.Error(), ms, Line{})
		return err
	}
	defer res.Body.Close()
	Observe("ok", ms)
	c.Log.Write(ctx, "INFO", "reservation.outbox.publish", "ReservationConfirmed headers copied", ms, Line{})
	return nil
}
