# CP2 — Structured log contract

Same JSON fields in `drp-api-gateway`, `drp-payment-api`, `drp-reservation-api`
and `drp-notification-api`. Filter by `trace_id` to rebuild one pay-and-confirm
attempt. `correlation_id` is the D-C16 uuid (`X-Correlation-Id` / error `traceId`).
`trace_id` is the 32-hex part of W3C `traceparent`.

## Fields

| Name | Type | Req | Example | Why |
| --- | --- | --- | --- | --- |
| timestamp | string ISO-8601 UTC | yes | 2026-10-21T22:41:01.120Z | order the hops |
| level | string | yes | ERROR | filter noise |
| service | string | yes | drp-payment-api | which repo wrote the line |
| version | string | yes | 2.0.0 | match a git tag |
| environment | string | yes | qa | do not mix env queries |
| trace_id | string hex32 | yes | 7a3c91e2b04d4f1a8c6e5d9f2b1a0843 | join HTTP and outbox |
| span_id | string hex16 | yes | a11b22c33d44e55f | one hop |
| correlation_id | uuid | yes | 9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20 | D-C16 header / body |
| event | string | yes | payment.provider.timeout | machine name |
| message | string | yes | provider did not answer | human line |
| duration_ms | int | yes | 3004 | golden signal latency |
| reservation_id | uuid | no | 2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c | domain |
| payment_id | uuid | no | 8a71d0e4-3c5b-4f9a-91d2-6b7c8d9e0f11 | domain |
| space_id | uuid | no | 0b1a2c3d-4e5f-6789-abcd-ef0123456789 | domain |
| reservation_state | enum | no | PAYMENT_PENDING | domain; only the three live states |

## Never log

| Data | Mask |
| --- | --- |
| password / `Spacehub1!` | never write |
| JWT and `Authorization` | `Authorization: Bearer ***` |
| email (`member@spacehub.local`) | drop; use `sub` if needed |
| card / PAN / CVV | MVP has none; still forbidden |
| raw `Idempotency-Key` | hash prefix only, `idem_9f3a…` |

Do not put `user_id` or raw URL in metric labels later (CP4). Same idea here.

## Levels

| Level | When | SpaceHub example |
| --- | --- | --- |
| ERROR | timeout or 5xx | `drp-payment-api` waits 3 s for reservation lookup and cuts |
| WARN | 422 business | currency not `COP`, or reservation already has a payment |
| INFO | start/end of a hop | gateway accepted `POST /api/v1/payments` |
| DEBUG | local detail | Flyway bind on `payment.outbox`, no secrets |

## Seven lines — one failed pay (same `trace_id`)

```json
{"timestamp":"2026-10-21T22:41:01.120Z","level":"INFO","service":"drp-api-gateway","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"1111111111111111","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"gateway.payments.accepted","message":"POST /api/v1/payments","duration_ms":2,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
{"timestamp":"2026-10-21T22:41:01.131Z","level":"INFO","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"2222222222222222","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.attempt.started","message":"E-16 create payment","duration_ms":4,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c","reservation_state":"PAYMENT_PENDING"}
{"timestamp":"2026-10-21T22:41:04.140Z","level":"ERROR","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"3333333333333333","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.reservation.timeout","message":"drp-reservation-api did not answer GET reservation","duration_ms":3004,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
{"timestamp":"2026-10-21T22:41:04.188Z","level":"INFO","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"4444444444444444","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.reservation.retry","message":"second lookup after timeout","duration_ms":41,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c","reservation_state":"CONFIRMED"}
{"timestamp":"2026-10-21T22:41:04.190Z","level":"WARN","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"5555555555555555","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.rejected.invalid_status","message":"reservation is not PAYMENT_PENDING","duration_ms":2,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c","reservation_state":"CONFIRMED"}
{"timestamp":"2026-10-21T22:41:04.195Z","level":"WARN","service":"drp-api-gateway","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"6666666666666666","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"gateway.payments.rejected","message":"upstream 422 INVALID_STATUS_TRANSITION","duration_ms":3075,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
{"timestamp":"2026-10-21T22:41:04.201Z","level":"DEBUG","service":"drp-payment-api","version":"2.0.0","environment":"qa","trace_id":"7a3c91e2b04d4f1a8c6e5d9f2b1a0843","span_id":"7777777777777777","correlation_id":"9c2e1a70-4b8f-4d21-9a06-3e5c7d8b1f20","event":"payment.outbox.skipped","message":"no PaymentConfirmed row on 422","duration_ms":1,"reservation_id":"2f0c6a18-9b44-4e71-8c3a-0d1e2f3a4b5c"}
```

Line 3 is the timeout. Line 5 is the business 422. No `PaymentConfirmed` was written.
