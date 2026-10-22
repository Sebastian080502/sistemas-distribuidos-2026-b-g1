# CP4 — Golden signals for the SpaceHub pay-and-confirm path

Four signals, four services. Names follow Prometheus: `snake_case`,
unit suffix, `_total` on counters. Labels stay bounded. No `user_id`,
no email, no raw URL, no `reservation_id`.

Users of SpaceHub expect `POST /api/v1/payments` under 300 ms on the
happy path (MVP provider confirms at once). The incident drill uses a
3 s hang, so histogram buckets stop at 3 s.

Buckets (seconds): `0.05`, `0.1`, `0.3`, `1`, `3`.

## Cardinality

| Label | Bound | Replaces |
| --- | --- | --- |
| `route` | `/api/v1/payments`, `on_payment_confirmed`, `on_reservation_confirmed` | raw URL |
| `method` | `POST`, `GET`, `CONSUME` | none |
| `outcome` | `ok`, `error`, `reject_422` | status code spray |
| `environment` | `dev`, `qa`, `main` | none |

A `user_id` or `space_id` label would explode. Those stay in the CP2 log.

## `drp-api-gateway`

| Signal | Metric | Type | Labels | PromQL |
| --- | --- | --- | --- | --- |
| Latency | `gateway_request_duration_seconds` | histogram | route, method | `histogram_quantile(0.95, sum by (le) (rate(gateway_request_duration_seconds_bucket{route="/api/v1/payments"}[5m])))` |
| Traffic | `gateway_requests_total` | counter | route, method, outcome | `sum(rate(gateway_requests_total{route="/api/v1/payments"}[5m]))` |
| Errors | `gateway_requests_total` | counter | route, outcome | `sum(rate(gateway_requests_total{route="/api/v1/payments",outcome="error"}[5m]))` |
| Saturation | `gateway_upstream_inflight` | gauge | route | `gateway_upstream_inflight{route="/api/v1/payments"}` |

## `drp-payment-api`

| Signal | Metric | Type | Labels | PromQL |
| --- | --- | --- | --- | --- |
| Latency | `payment_request_duration_seconds` | histogram | route, method | `histogram_quantile(0.95, sum by (le) (rate(payment_request_duration_seconds_bucket{route="/api/v1/payments"}[5m])))` |
| Traffic | `payment_requests_total` | counter | route, outcome | `sum(rate(payment_requests_total{route="/api/v1/payments"}[5m]))` |
| Errors | `payment_requests_total` | counter | route, outcome | `sum(rate(payment_requests_total{outcome="error"}[5m]))` |
| Saturation | `payment_outbox_pending` | gauge | environment | `payment_outbox_pending` |

`payment_outbox_pending` is rows in `payment.outbox` not published. If it
grows, `PaymentConfirmed` never reaches reservation.

## `drp-reservation-api`

| Signal | Metric | Type | Labels | PromQL |
| --- | --- | --- | --- | --- |
| Latency | `reservation_request_duration_seconds` | histogram | route | `histogram_quantile(0.95, sum by (le) (rate(reservation_request_duration_seconds_bucket{route="on_payment_confirmed"}[5m])))` |
| Traffic | `reservation_requests_total` | counter | route, outcome | `sum(rate(reservation_requests_total{route="on_payment_confirmed"}[5m]))` |
| Errors | `reservation_requests_total` | counter | route, outcome | `sum(rate(reservation_requests_total{outcome="error"}[5m]))` |
| Saturation | `reservation_outbox_pending` | gauge | environment | `reservation_outbox_pending` |

CP3 already exposes the counter and a millisecond histogram in the
workshop module. In product code the unit becomes `_seconds`.

## `drp-notification-api`

| Signal | Metric | Type | Labels | PromQL |
| --- | --- | --- | --- | --- |
| Latency | `notification_consume_duration_seconds` | histogram | route | `histogram_quantile(0.95, sum by (le) (rate(notification_consume_duration_seconds_bucket[5m])))` |
| Traffic | `notification_events_total` | counter | route, outcome | `sum(rate(notification_events_total{route="on_reservation_confirmed"}[5m]))` |
| Errors | `notification_events_total` | counter | outcome | `sum(rate(notification_events_total{outcome="error"}[5m]))` |
| Saturation | `notification_mongo_inflight` | gauge | environment | `notification_mongo_inflight` |

Mongo `notifications` is the write. Duplicate `sourceEventId` is not an
error (BR-006); do not count it as `error`.
