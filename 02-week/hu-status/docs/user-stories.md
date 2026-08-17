# User Story Backlog

## Distributed Reservation Platform

This document contains the initial user-story backlog of the Distributed Reservation Platform.

The stories are derived from the needs identified in the PRD and are related to the bounded contexts defined in the Context Map.

Each story must have testable acceptance criteria before it is considered ready for implementation.

---

## HU-ADR-001 — Formalize the architectural decision

### Story

As a student of the Distributed Systems course, I want to document the initial architectural decision of the Distributed Reservation Platform, so that the bounded contexts, candidate microservices, and rejected alternatives are justified.

### Bounded Context

Architecture / cross-cutting.

### Priority

High.

### Acceptance criteria

```gherkin
Scenario: Bounded contexts identified
  Given the initial domain analysis
  When the architecture is documented
  Then the main bounded contexts of the platform must be identified using business language

Scenario: Academic constraint vs business need
  Given that the project belongs to Distributed Systems
  When the architecture is selected
  Then the use of microservices must be justified
  And a clear distinction is made between the academic constraint and the actual business need

Scenario: Alternatives considered
  When the architectural decision is documented
  Then the considered alternatives and the reasons for rejecting them must be recorded

Scenario: Responsibility of each service
  When the initial architecture is defined
  Then each microservice must have a clearly bounded business responsibility

Scenario: Synchronous and asynchronous communication
  When communication between services is documented
  Then the planned synchronous and asynchronous interactions must be identified

Scenario: Data ownership
  When data ownership is documented
  Then it must be established that each microservice is responsible for its own data

Scenario: ADR immutability
  When the ADR is accepted
  Then any later change to the architectural decision must be recorded through a new ADR
```

### Evidence

- [`adr-001-architecture.md`](./adr-001-architecture.md)
- [`context-map.md`](./context-map.md)

---

## HU-ARC-001 — Identify bounded contexts

### Story

As the team, I want to define the bounded contexts of the platform so that clear business boundaries are established.

### Priority

High.

### Acceptance criteria

```gherkin
Scenario: Bounded contexts are identified in business language
  Given the initial domain analysis of the Distributed Reservation Platform
  When the bounded contexts are documented
  Then Identity and Access, Space Management, Reservation Management, Payment Management, and Notification Management are identified
  And each context is described with business language, not database tables
  And each context has a clearly documented responsibility
```

### Evidence

- [`context-map.md`](./context-map.md)

---

## HU-ARC-002 — Define the initial Context Map

### Story

As the team, I want to document the Context Map so that the relationships between bounded contexts are established.

### Priority

High.

### Acceptance criteria

```gherkin
Scenario: Context Map records relationships and data ownership
  Given the bounded contexts of the Distributed Reservation Platform
  When the Context Map is documented
  Then the relationships between contexts are recorded
  And the owner of each data set is identified
  And there is no direct access between databases
  And the planned synchronous and asynchronous communication contracts are identified
```

### Evidence

- [`context-map.md`](./context-map.md)

---

## HU-ARC-003 — Document the architectural decision in ADR-001

### Story

As the team, I want to document the architectural decision through ADR-001.

### Priority

High.

### Acceptance criteria

```gherkin
Scenario: ADR-001 captures the architecture decision
  Given the need to justify the initial architecture
  When ADR-001 is written
  Then the Context, Decision, considered alternatives, rejection reasons, and consequences are documented
```

### Evidence

- [`adr-001-architecture.md`](./adr-001-architecture.md)

---

## HU-ARC-004 — Define testable acceptance criteria

### Story

As the team, I want to define testable Given/When/Then acceptance criteria for the architectural and product backlog, so that each story can be verified before implementation.

### Priority

High.

### Acceptance criteria

```gherkin
Scenario: Backlog stories are testable
  Given the architectural and product stories of Week 02
  When the acceptance criteria are written
  Then each story includes Given/When/Then scenarios that can be verified
```

### Evidence

- [`user-stories.md`](./user-stories.md)

---

## HU-RES-001 — Consult available spaces

### Story

As a user, I want to consult the spaces available for a date and time range, so that I can select a space I can reserve.

### Bounded Context

Space Management.

### Priority

High.

### Acceptance criteria

#### AC-001 — Valid query

Given a valid date and time range,
when the user queries availability,
then the system must show the spaces available for that period.

#### AC-002 — Space already reserved

Given a space that already has a confirmed reservation during the requested period,
when its availability is queried,
then that space must not appear as available.

#### AC-003 — Invalid range

Given an invalid date and time range,
when the query is made,
then the system must reject the request and indicate the error.

#### AC-004 — Query without side effects

When availability is queried,
then the query must not modify any existing reservation.

---

## HU-RES-002 — Consult the availability of a selected space

### Story

As a user, I want to consult the availability of a specific space, so that I know whether I can reserve it on a given date and time.

### Bounded Context

Space Management.

### Priority

High.

### Acceptance criteria

#### AC-001 — Space available

Given an existing space and a valid period with no overlapping confirmed reservations,
when the user queries that space,
then the system must indicate that the space is available.

#### AC-002 — Space unavailable

Given a space with a confirmed reservation that overlaps the requested period,
when the user queries that space,
then the system must indicate that it is not available.

#### AC-003 — Space does not exist

Given a space identifier that does not exist,
when its availability is queried,
then the system must reject the request indicating that the space was not found.

#### AC-004 — Blocked period

Given a space with a period blocked by the administrator,
when that period is queried,
then the system must indicate that the space is not available.

---

## HU-RES-003 — Create a reservation for an available space

### Story

As a user, I want to create a reservation for an available space, so that I can secure the space for a specific date and time.

### Bounded Context

Reservation Management.

### Priority

High.

### Acceptance criteria

#### AC-001 — Valid reservation

Given an authenticated user, an existing space, and an available period,
when the user creates a reservation,
then the system must create the reservation in a valid initial state and return its identifier.

#### AC-002 — Availability conflict

Given a space that is not available in the requested period,
when the user tries to create the reservation,
then the system must reject the operation and must not create the reservation.

#### AC-003 — Overlapping confirmed reservations

Given that a confirmed reservation already exists for the same space and an overlapping period,
when another user tries to create a reservation,
then the system must reject it to satisfy BR-001.

#### AC-004 — Incomplete data

Given that required reservation data is missing,
when the request is submitted,
then the system must reject it and indicate the invalid fields.

---

## HU-RES-004 — Consult the status of a reservation

### Story

As a user, I want to consult the status of a reservation, so that I know where it is in its lifecycle.

### Bounded Context

Reservation Management.

### Priority

High.

### Dependency

HU-RES-004 may be implemented after HU-RES-003.

### Acceptance criteria

#### AC-001 — Existing reservation query

Given an existing reservation,
when the authorized user queries its status,
then the system must return the identifier, space, period, and current state.

#### AC-002 — Reservation does not exist

Given a reservation identifier that does not exist,
when it is queried,
then the system must indicate that the reservation was not found.

#### AC-003 — Unauthorized access

Given a user who is not the reservation owner and is not an authorized administrator,
when they try to query the reservation,
then the system must reject access.

#### AC-004 — Visible states

When a reservation is queried,
then the returned state must be one of the defined states: PENDING, PAYMENT_PENDING, CONFIRMED, or CANCELLED.

---

## HU-RES-005 — Process the payment associated with a reservation

### Story

As a user, I want to process the payment associated with a reservation, so that I can confirm the use of the space.

### Bounded Context

Payment Management.

### Priority

High.

### Acceptance criteria

#### AC-001 — Payment started

Given a reservation in a state that allows payment,
when the user requests the payment,
then the system must create a payment operation and associate it with the reservation.

#### AC-002 — Idempotence

Given the same payment operation submitted more than once,
when it is processed,
then the system must not generate multiple financial effects.

#### AC-003 — Payment result

When payment processing finishes,
then the Payment Service must publish a result event, for example PaymentConfirmed or PaymentFailed.

#### AC-004 — Reservation missing or not payable

Given a reservation that does not exist or does not allow payment,
when payment is requested,
then the system must reject the operation.

---

## HU-RES-006 — Confirm a reservation after successful payment

### Story

As a user, I want my reservation to be confirmed after a successful payment, so that the space is assigned definitively.

### Bounded Context

Reservation Management.

### Priority

High.

### Acceptance criteria

#### AC-001 — Confirmation after successful payment

Given a reservation pending payment,
when the PaymentConfirmed event is received,
then the reservation must move to CONFIRMED.

#### AC-002 — Failed payment

Given a reservation pending payment,
when the PaymentFailed event is received,
then the reservation must not move to CONFIRMED.

#### AC-003 — Duplicate event

Given that PaymentConfirmed is received twice,
when the second event is processed,
then the system must not generate a second confirmation or duplicate effects.

#### AC-004 — Invalid transition

Given a CANCELLED reservation,
when a payment event is received,
then the system must not confirm it.

---

## HU-RES-007 — Receive a notification when a reservation changes its status

### Story

As a user, I want to receive a notification when my reservation changes its status, so that I stay informed without constantly querying the system.

### Bounded Context

Notification Management.

### Priority

Medium.

### Acceptance criteria

#### AC-001 — Notification on confirmation

Given a reservation that moves to CONFIRMED,
when ReservationConfirmed is published,
then the Notification Service must create and deliver a notification to the user.

#### AC-002 — Notification on cancellation

Given a cancelled reservation,
when ReservationCancelled is published,
then the user must receive a cancellation notification.

#### AC-003 — Notification failure

Given a failure in notification delivery,
when the reservation is already confirmed,
then the reservation must not be invalidated.

#### AC-004 — Duplicate events

Given a duplicate notification event,
when it is processed,
then the system must not generate uncontrolled duplicate deliveries.

---

## HU-RES-008 — Manage reservable spaces and their availability rules

### Story

As a space administrator, I want to register and update spaces and their availability rules, so that I can control which resources can be reserved and in which periods.

### Bounded Context

Space Management.

### Priority

High.

### Acceptance criteria

#### AC-001 — Register space

Given an authenticated administrator and valid space data,
when they register a space,
then the system must persist it and make it queryable.

#### AC-002 — Update space

Given an existing space,
when the administrator updates its information,
then the changes must be persisted without inconsistently affecting confirmed reservations.

#### AC-003 — Block period

Given an existing space,
when the administrator blocks a period,
then that period must not appear as available for new reservations.

#### AC-004 — Invalid data

Given a registration or update with invalid data,
when the request is submitted,
then the system must reject it and indicate the error.

---

## Definition of Ready

A story may enter the Sprint when:

- The business goal is clearly defined.
- The responsible bounded context is identified.
- The acceptance criteria are testable.
- Known dependencies are identified.
- The scope is small enough for the Sprint.
- There are no unresolved critical questions that block implementation.

## Definition of Done

A story is considered done when:

- The functionality meets its acceptance criteria.
- The code is implemented.
- The corresponding tests are added.
- There are no known critical errors.
- The change has been reviewed through a Pull Request.
- The necessary documentation has been updated.
- The change is integrated into the corresponding branch.
- The story evidence is recorded in the weekly status.

Architectural stories HU-ADR-001 and HU-ARC-001 to HU-ARC-004 support the architecture definition and are not developed as independent product features.
