# CP6 — SLOs and error budget (SpaceHub pay-and-confirm)

Journey: `POST /api/v1/payments` (E-16) then consume `PaymentConfirmed`
(D-C14). Window: **30 days**. I am the only member, so the policy is
what I stop doing, not a committee.

## Traffic I use for the budget

SpaceHub is a small coworking, not a marketplace. About **40 spaces**,
weekday bookings of 1–2 h. I count **120** `POST /api/v1/payments` per
day (one payment after each `PAYMENT_PENDING` create). Classroom traffic
is lower; 120 is the number I would run the product with.

30 days × 120 = **3 600** valid payment requests per month.

A 422 `INVALID_STATUS_TRANSITION` (already `CONFIRMED`) is a **valid**
event that is not "good", but I keep it out of the availability SLI:
availability only counts `5xx` and timeouts. Business 422 is the
latency/confirm SLIs, not a hole in uptime.

## SLO 1 — Availability of E-16

| | |
| --- | --- |
| SLI | good / valid = `outcome!="error"` on `POST /api/v1/payments` |
| PromQL | `1 - (sum(rate(payment_requests_total{route="/api/v1/payments",outcome="error"}[30d])) / sum(rate(payment_requests_total{route="/api/v1/payments"}[30d])))` |
| Target | **99.5 %** in 30 days |
| Why | If payment is down, D-C14 never fires and the row stays `PAYMENT_PENDING` |

**Budget in requests:** 0.5 % × 3 600 = **18** failed payments / month.

**Budget in minutes:** 0.5 % × 30 × 24 × 60 = **216** minutes of full
outage (gateway or `drp-payment-api` returning `error` for everyone).

## SLO 2 — Latency of E-16

| | |
| --- | --- |
| SLI | share of requests with duration ≤ 300 ms |
| PromQL | `sum(rate(payment_request_duration_seconds_bucket{route="/api/v1/payments",le="0.3"}[30d])) / sum(rate(payment_request_duration_seconds_count{route="/api/v1/payments"}[30d]))` |
| Target | **99 %** ≤ 300 ms in 30 days |
| Why | CP4: MVP provider confirms at once; users wait at the desk |

**Budget:** 1 % × 3 600 = **36** slow payments / month. A 3 s hang (the
CP9 drill) burns this SLO first, then availability if it becomes 5xx.

## SLO 3 — Confirm after `PaymentConfirmed`

This is the SpaceHub-specific one. Confirm is not HTTP.

| | |
| --- | --- |
| SLI | consume `on_payment_confirmed` with `outcome="ok"` / all consumes |
| PromQL | `sum(rate(reservation_requests_total{route="on_payment_confirmed",outcome="ok"}[30d])) / sum(rate(reservation_requests_total{route="on_payment_confirmed"}[30d]))` |
| Target | **99 %** in 30 days |
| Why | BR-003 / D-C14: `PAYMENT_PENDING` must become `CONFIRMED` when the event arrives. If `reservation.outbox` has no publisher, this SLI never even has events |

If `payment_outbox_pending` grows, this SLI goes quiet (no consumes),
which I treat as **burn**: no confirms is as bad as errors.

**Budget:** 1 % of 3 600 events ≈ **36** missed confirms / month.

## Error budget policy (solo)

| Spent | What I do |
| --- | --- |
| **50 %** | I freeze new HUs (no front polish, no extra reports). I only work payment HTTP, outbox publisher, or consume. |
| **75 %** | No deploys except a fix. I page myself (I am on-call). I follow the CP8 runbook. |
| **100 %** | I stop shipping. Incident + postmortem. The next commit is a fix or a rollback, not a feature. Camunda stays out: it is not in the contract. |

G-01 (`amountCents`) does **not** pause this policy. An open PO issue
is not an error-budget exception.
