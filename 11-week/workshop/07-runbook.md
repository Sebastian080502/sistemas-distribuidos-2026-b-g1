# CP8 — Runbook: SpaceHubPaymentFastBurn

Page for the fast burn on `POST /api/v1/payments` (E-16).
I am the only member. The page is mine. Confirm is still D-C14,
not HTTP and not Camunda.

## Symptoms

- Fast alert red: burn 14.4 for 1 h, still true for 5 min.
- At the SpaceHub desk, pay stays spinning or returns 5xx.
- Reservations stay `PAYMENT_PENDING`. No `PaymentConfirmed`.
- `payment_outbox_pending` may climb if HTTP 201 is lying and
  the publisher is stuck. If HTTP itself fails, the counter
  `payment_requests_total{outcome="error"}` climbs.

## Confirm it (do this before any change)

PromQL (same metric as CP7):

```
sum(rate(payment_requests_total{route="/api/v1/payments",outcome="error"}[1h]))
/
sum(rate(payment_requests_total{route="/api/v1/payments"}[1h]))
```

If that ratio is above `14.4 * 0.005` = **0.072**, the page is real.

Log query, one `trace_id` from a failed pay (CP2 contract):

```
{service="drp-payment-api"} |= "7a3c91e2b04d4f1a8c6e5d9f2b1a0843"
```

Also grep `X-Correlation-Id` / `correlation_id` (D-C16). You should
see gateway accept, then `payment.reservation.timeout` or a 5xx,
and **no** `PaymentConfirmed` outbox row.

Look at `gateway_upstream_inflight{route="/api/v1/payments"}` and
`drp-payment-api` logs. Do **not** open `reservation` tables from
payment.

## First mitigations (lowest risk first)

1. **Flip `drp-front` to `failover`.** The desk stops sending live
   E-16. Writes stay fail-closed. This is the closest thing we have
   to a feature flag. It does not invent money. It stops the bleed.

2. **Rollback the last `drp-payment-api` or `drp-api-gateway` image**
   on the env that is burning (`qa` / `main`). Cherry-pick stays the
   promotion path. Do not force-push. Do not retag `v2.0.0`.

3. **Do not scale Postgres and do not start Camunda.** Extra
   replicas of an empty publisher will not emit `PaymentConfirmed`.
   If `payment_outbox_pending` is the only red gauge and HTTP is
   201, the fix is the publisher, not CPU.

Circuit breaker is not in the product today. I will not pretend it
is. Failover + rollback is what the repos actually have.

## Escalation

| Who | When |
| --- | --- |
| Me (only member) | First page. I run this list. |
| `ariel5253` (instructor / PO) | If I must change a contract (G-01, a new event) or touch `main` without a review. I do not self-merge `main`. |

There is no second on-call.

## Recovery

The page is clear when the 1 h error ratio is back under `0.072`
for 15 min **and** a real `POST /api/v1/payments` returns 201,
`payment.outbox` gets `PaymentConfirmed`, and a reservation moves
`PAYMENT_PENDING` → `CONFIRMED`. A 200 on `/health` is not enough.
