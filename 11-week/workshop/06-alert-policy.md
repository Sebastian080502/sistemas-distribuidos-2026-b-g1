# CP7 — Burn-rate policy

Availability SLO is **99.5 %** over 30 days (CP6). Allowed error
ratio is **0.005**. Monthly request budget is **18** failed
`POST /api/v1/payments`. Time budget is **216** minutes of full outage.

Burn rate `1` spends the month exactly on pace. The rules in
`06-alerts.yml` use `payment_requests_total` from CP4, not CPU.

## Fast alert — burn 14.4 for 1 hour (confirmed 5 min)

14.4 × (1 / 720 hours) = **2 %** of the 30-day budget before it fires.

That is **0.36** failed payments (2 % of 18) or about **4 minutes** of
full outage (2 % of 216). In a small coworking that is a short desk
outage: gateway or `drp-payment-api` 5xx, or a timeout storm.

`severity: page`. I am the only member, so the page is me. I open the
CP8 runbook.

## Slow alert — burn 6 for 6 hours (confirmed 30 min)

6 × (6 / 720) = **5 %** of the budget.

That is **0.9** failed payments or about **11 minutes** of full outage
spread over six hours. A leak, not a hard crash: publisher stuck,
`payment_outbox_pending` growing, `PaymentConfirmed` not leaving.

`severity: ticket`. I do not wake up for this if the fast alert is
quiet, but I stop new HUs (50 % policy) if it stays red.

## Two cause-based alerts I will not create

1. **CPU > 80 % on the `drp-payment-api` container.** A user at the
   SpaceHub desk does not feel CPU. The fast burn already pages when
   E-16 fails. CPU without failed payments is noise.

2. **Postgres connections on `drp-infra-postgres`.** Same reason. The
   composition root can look busy while E-16 still returns 201. I watch
   `payment_outbox_pending` and the burn rules. I do not page on a
   cause I cannot map to `PAYMENT_PENDING` stuck.

Camunda is not an alert source. It is not in the running path.
