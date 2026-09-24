# Service Communication — REST, gRPC and Messaging — SpaceHub / DRP

**Unit 2 · Weekly · Corte 2**

This activity defines how SpaceHub services communicate: synchronous REST (default), optional internal gRPC later, and asynchronous events for facts that already happened. Delivery is at-least-once; consumers must be idempotent.

## 1. Communication choice

| Interaction | Mode | Contract | Reason |
|---|---|---|---|
| Browser to api-gateway | REST | OpenAPI `/api/v1` | User is waiting; HTTP tooling. |
| Gateway to identity / space / reservation / payment | REST | OpenAPI `/api/v1` | Immediate login, availability, create reservation, start payment. |
| reservation-service to space-service (slot allowed by catalog) | REST | OpenAPI `/api/v1` | Create reservation needs an immediate yes/no. |
| payment-service to reservation-service and notification-service | Async event | `PaymentConfirmed.v1` / `PaymentFailed.v1` | Payment result must not couple to mail or a second HTTP hop. |
| reservation-service to notification-service | Async event | `ReservationConfirmed.v1` / `ReservationCancelled.v1` | Notification downtime must not un-confirm a booking. |
| Internal high-throughput later | gRPC | Versioned `.proto` | Only if measured latency justifies it. Not MVP. |

## 2. REST — synchronous

Default for public APIs and any call where the caller needs the answer now. Trade-off: temporal coupling. Timeouts belong on remaining sync calls.

## 3. gRPC — not the SpaceHub default

gRPC is typed and efficient over HTTP/2. SpaceHub Corte 2 stays on REST/JSON. A `.proto` is a later ADR if an internal path is proven hot.

## 4. Messaging — asynchronous

Local broker (AMQP) when Compose exists. Publisher emits `PaymentConfirmed`; reservation and notification consume independently.

- Queue: work distribution (notification workers).
- Pub/sub: one fact, several contexts (payment result → reservation + notification).

## 5. Delivery semantics and idempotency

Exactly-once **delivery** is not guaranteed. Exactly-once **processing**:

```text
at-least-once delivery + eventId + store processed ids (or unique constraint)
```

`PaymentConfirmed.v1` includes `eventId`, `schemaVersion`, `reservationId`, `paymentId`, `correlationId`, `occurredAt`. `notification-service` ignores a repeated `eventId` (TD-009). `payment-service` already requires HTTP `Idempotency-Key` (FR-006).

## 6. A path we will face

Sync chain `Angular -> gateway -> reservation -> space -> payment -> notification` would freeze the booking if mail is down. Mitigation: keep REST until payment intent; publish events after money/state changes; notification is async.

## 7. Common mistakes

- Long synchronous chains.
- Assuming at-least-once means no duplicates.
- Using events when the user is waiting for availability.
- No timeout on remaining REST calls.
- No idempotency on notification/payment consumers.

## Self-check

### Question 1
**Main benefit of async messaging?**  
**Answer:** Decoupling; the producer survives consumer downtime.

### Question 2
**gRPC is a strong choice for...**  
**Answer:** Internal high-throughput contract-first calls (not SpaceHub MVP).

### Question 3
**At-least-once delivery requires consumers to be...**  
**Answer:** Idempotent.

### Question 4
**A pub/sub topic lets...**  
**Answer:** One event reach several independent subscribers.

### Question 5
**Exactly-once delivery over a network is...**  
**Answer:** Not guaranteed; engineer exactly-once processing.

### Question 6
**A long synchronous chain under load tends to...**  
**Answer:** Cascade (threads wait, pools exhaust).

## This week

For each SpaceHub interaction, sync vs async is recorded in `drp-docs/09-microservices/communication-patterns.md`. Consumers in code wait for instructor-created service repos (TD-003 outbox, TD-007 idempotent notification).
