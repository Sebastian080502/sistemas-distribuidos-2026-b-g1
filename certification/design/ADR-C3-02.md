# ADR-C3-02: Cross-AZ warm standby for RDS (RPO 5 min, RTO 15 min)

## Status

Accepted

## Context

Section 4 of the SpaceHub design note. The critical store is the single RDS that holds overlap (BR-001) and payment idempotency (BR-004). `04-requirements/non-functional.md` NFR-002 only asks for local Compose health; it does not give RTO or RPO. Huawei chapter 11 forces a pattern with numbers. A member who just paid cannot lose the last afternoon of reservations because the primary AZ failed.

## Decision

RDS for PostgreSQL uses **warm standby across two availability zones**: primary in AZ-1, synchronous standby in AZ-2, CCE node pools in both zones, one ELB. RPO is **5 minutes** (standby is seconds behind; we budget 5 minutes for detect + DNS/ELB). RTO is **15 minutes** (promote standby, ELB health, CCE already running in AZ-2). MongoDB/DDS for `notification` may use the same layout but is **not** on the RTO clock (BR-006).

## Alternatives considered

- **Cold standby / daily OBS backup only:** rejected because a daily dump cannot meet a 5-minute RPO; every `POST /reservations` after the last snapshot is gone.
- **Hot-hot multi-region active-active:** rejected for the MVP. We have one composition root, unanswered Q-001, and a one-developer team. Dual-write across regions would invent a conflict model that `-docs` does not have.

## Consequences

- Positive: AZ failure does not require Flyway replay from OBS; members keep the last minutes of confirmed bookings.
- Negative: synchronous replication adds commit latency on `POST /reservations` and `POST /payments`; we still have to stay under NFR-001 (800 ms and 1 s) and will watch RDS write lag as a first alert.
