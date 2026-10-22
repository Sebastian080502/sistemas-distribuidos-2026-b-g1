# CP9 — Incident drill: hop 2 hangs 3 s

For 40 minutes, `drp-payment-api` (callee of CP1 hop 2) answers in
**3 seconds on 25 %** of `POST /api/v1/payments`. The other 75 % stay
fast. I treat 3 s as a **timeout / `outcome="error"`**, same as the
CP2 line `payment.reservation.timeout`. Desk hours: **15 pays / hour**
(120 / 8 h day).

## One slow trace (same `trace_id`)

- `front.startPayment` **3100 ms**
  - `gateway.proxyPayments` **3050 ms**
    - `payment.createAttempt` **3004 ms** ← hop 2
      - `payment.reservation.lookup` **3004 ms** (timeout)
      - `payment.writeOutbox` **not run** (no `PaymentConfirmed`)

```json
{"timestamp":"2026-10-21T20:00:12.010Z","level":"INFO","service":"drp-api-gateway","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"1111111111111111","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"gateway.payments.accepted","message":"POST /api/v1/payments","duration_ms":3,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
{"timestamp":"2026-10-21T20:00:12.020Z","level":"INFO","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"2222222222222222","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.attempt.started","message":"E-16 create payment","duration_ms":4,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c","reservation_state":"PAYMENT_PENDING"}
{"timestamp":"2026-10-21T20:00:15.030Z","level":"ERROR","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"3333333333333333","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.reservation.timeout","message":"drp-reservation-api did not answer GET reservation","duration_ms":3004,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
{"timestamp":"2026-10-21T20:00:15.040Z","level":"WARN","service":"drp-api-gateway","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"6666666666666666","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"gateway.payments.rejected","message":"upstream timeout","duration_ms":3075,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
```

## SLI values in the 40 min window

| SLI | During the 40 min | 30-day after |
| --- | --- | --- |
| Availability E-16 | **75 %** (25 % `error`) | ~3 extra fails / 3600 → still near 99.5 % |
| Latency ≤ 300 ms | **75 %** good | 3 slow / 3600, SLO 99 % still holds |
| Confirm after event | **quiet** — no `PaymentConfirmed`, rows stay `PAYMENT_PENDING` | 3 missed confirms |

## Which alert fires, and when

Error ratio during the hang is **0.25**. Fast threshold is
`14.4 * 0.005 = 0.072`.

- **5 m window** is over 0.072 from minute **5**.
- **1 h window** reaches 0.072 only around minute **18**
  (`3.75 errors/h × t/60 / 15 > 0.072`).
- Rule `for: 5m` → **`SpaceHubPaymentFastBurn` pages at ~minute 23**.

`SpaceHubPaymentSlowBurn` does **not** fire: 40 min of this leak in a
6 h window is ~0.028, under `6 * 0.005 = 0.030`.

## Runbook step that stops the bleed

**Mitigation 1:** set `drp-front` to `failover`. Live E-16 stops.
Then rollback payment/gateway if a bad image is the hang.

## Blameless postmortem

**Timeline (desk, qa)**

| t | What |
| --- | --- |
| 0 | Hop 2 starts answering 3 s on 25 % of pays |
| 5 | 5 m error ratio 0.25; 1 h window still low |
| 18 | 1 h window crosses 0.072 |
| 23 | Fast burn pages. I open CP8 |
| 25 | Front → `failover`. New live pays stop |
| 40 | Hang would have ended; bleed already stopped at 25 |

**Impact:** ~10 attempts in 40 min, **~3 failed**. Request budget 18 →
**17 %** spent if I had let it run. Time-equivalent **10 min** of full
outage (`0.25 × 40`). After failover at t=25: **~2 fails**, ~11 % of
the 18.

**Root cause:** `drp-payment-api` waited on a reservation lookup and
hit 3 s. Not a person. Not Camunda (it is not on the path).

**Contributing:** no latency burn rule (CP7 only watches `outcome=
error` if we had returned 201 slowly, nobody would page). No publisher
yet, so a 201 with a stuck outbox would also hide D-C14. Front was
`live` in this drill; today the product default is still `failover`.

### Action items (user stories)

1. **HU-OBS-001 — Latency burn on E-16**  
   As the only operator, I want a burn rule on the 300 ms SLI so a
   3 s hang pages even when the status is 201.  
   **AC:** a 15 min synthetic 3 s delay on 25 % of pays fires a page;
   `promtool check rules` passes; no `user_id` label.

2. **HU-PAY-HTTP-001 — Timeout and fail-closed on E-16**  
   As a member at the desk, I want a 201 or a 422/5xx in **under 1 s**
   if reservation lookup hangs, and **no** `PaymentConfirmed` on
   timeout.  
   **AC:** a 3 s stub on GET reservation returns 5xx/timeout, zero
   outbox rows, Anexo A sonda 16 still 400 without key.

3. **HU-OUTBOX-002 — Publisher copies `traceparent`**  
   As the only operator, I want `PaymentConfirmed` to leave
   `payment.outbox` with the same `trace_id` as E-16.  
   **AC:** one `go`/`JUnit` test: inbound `trace_id` equals the event
   header; no Camunda; no `SELECT` on `reservation.*` from payment.
