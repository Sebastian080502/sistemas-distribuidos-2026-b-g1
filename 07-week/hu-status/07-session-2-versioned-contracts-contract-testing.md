# Planning — Versioned Contracts and Contract Testing — SpaceHub / DRP

**Unit 2 · Planning · Corte 2**

This activity formalizes the contracts that let SpaceHub services evolve independently: OpenAPI as source of truth, compatibility rules, and consumer-driven contract testing when CI exists.

## 1. The contract is the source of truth

| Kind | Artifact | Location |
|------|----------|----------|
| REST | OpenAPI 3.0 | `drp-docs/07-api/contracts/openapi/*.yaml` |
| Events | Fields in `event-catalog.md` + service `events.md` | JSON payload with `eventId` / `schemaVersion` |
| gRPC | `.proto` | Not in MVP |

When the instructor creates the service repositories, those codebases must follow these files. If code and contract diverge, the **contract** wins until an ADR changes it. Students do not create GitHub repositories.

## 2. Every endpoint or event declares

1. Method + path, or event name / routing key.
2. Request or event schema (including `Idempotency-Key` and `X-Correlation-Id` where required).
3. Response schema.
4. Error envelope (`_shared.yaml` `ErrorResponse`: `error`, `message`, `details`, `traceId`).
5. Version (`/api/v1` or `PaymentConfirmed.v1`).

Conventions: UUID ids, ISO-8601 UTC, pagination `?page=&limit=`.

## 3. Versioning and backward compatibility

Services deploy independently.

**Compatible:** add optional response field; add new endpoint.

**Breaking:** rename, remove, or retype a field; add a required request field. Then:

```text
/api/v1/reservations
/api/v2/reservations
```

or `PaymentConfirmed.v2`. Keep v1 until consumers migrate. HTTP APIs may send `Sunset`.

## 4. Consumer-driven contract testing (Pact)

Unit tests do not catch producer field renames. Pact flow:

```text
Angular or gateway consumer test
        -> pact file
        -> producer verification in CI
        -> block QA promotion if incompatible
```

Until services exist, this is TD-008 (`todo`). OpenAPI files are already the published contract (TD-004 depends on running producers).

## 5. SpaceHub contract examples

### Availability (sync)

```http
GET /api/v1/spaces/available?from=...&to=...
```

- `200` list of spaces
- `400` `VALIDATION_ERROR`

### Payment (sync command + async fact)

```http
POST /api/v1/payments
Idempotency-Key: uuid
```

Then event `PaymentConfirmed.v1` (`eventId`, `reservationId`, `paymentId`, `correlationId`).

## 6. Contract testing backlog (after instructor repos exist)

| Priority | Work item | Gate | Status |
|---|---|---|---|
| Must | Keep OpenAPI per service in `drp-docs` | Schemas validate | doing |
| Must | Implement APIs against those files | Producer matches OpenAPI | todo |
| Must | Pact or equivalent: web-app/gateway -> space-service | CI fails on breaking response | todo |
| Must | Version events `*.v1` | Consumers agree on `eventId` | doing (docs) |
| Should | Duplicate-event test on notification | Same `eventId` -> one inbox row | todo |
| Should | Compatibility check: optional fields pass; rename fails | CI | todo |

## 7. A path we will face

If payment-service renames `amount` to `total` under `/api/v1`, Angular can show empty totals while both unit-test suites stay green. Pact/OpenAPI verification in CI must fail that producer build.

## 8. Common mistakes

- No machine-readable contract.
- Breaking changes under the same version.
- Only unit tests for cross-service compatibility.
- Changing event fields without consumers.
- Contract tests that never run in CI.

## Self-check

### Question 1
**Source of truth for an API should be...**  
**Answer:** A versioned machine-readable contract (OpenAPI / proto / schema).

### Question 2
**Which change is backward-compatible?**  
**Answer:** Adding an optional field.

### Question 3
**A breaking change requires...**  
**Answer:** A new version and deprecation of the old one.

### Question 4
**Pact makes...**  
**Answer:** Producer CI fail when a consumer expectation breaks.

### Question 5
**A standard error envelope should include...**  
**Answer:** `error`/`code`, `message`, `details`, `traceId`.

### Question 6
**A silent field rename is prevented by...**  
**Answer:** Contracts, consumer pacts, and CI verification.

## This week

Contracts stay in `drp-docs`. Implement and verify them in the service repositories **when the instructor creates them** (TD-004).
