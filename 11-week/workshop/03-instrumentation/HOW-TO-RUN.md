# How to run (CP3)

Windows PowerShell, Go 1.22+. No extra modules.

```
cd 11-week/workshop/03-instrumentation
go test -count=1 -v
```

Output on this machine (2026-10-21):

```
=== RUN   TestIncomingTraceIDEqualsDownstream
--- PASS: TestIncomingTraceIDEqualsDownstream (0.03s)
PASS
ok  	spacehub/drp-reservation-api/observability	2.401s
```

Same `trace_id` on consume and outbound `ReservationConfirmed`. `/metrics` has the counter and histogram.
