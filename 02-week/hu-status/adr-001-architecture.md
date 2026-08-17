# ADR-001: Microservices architecture

- Status: Accepted
- Date: 2026-08-17
- Decision maker: Juan Sebastian Osorio Fierro
- Scope: Overall architecture of the Distributed Reservation Platform
- Story: HU-ADR-001 — Formalize the architectural decision

## 1. Context

The Distributed Reservation Platform manages reservable spaces, availability, reservations, payments, and notifications.

The project belongs to the Distributed Systems course and must demonstrate distributed-system concepts: communication between services, component independence, fault tolerance, consistency, scalability, and architectural evolution.

The project will initially be developed by a single student, so the architecture must stay within a controlled scope that can be implemented, tested, and documented during the academic term.

At the same time, the architecture must allow the system to grow by adding new business capabilities without rebuilding the whole solution.

The initial domain analysis identified these main capabilities:

- Identity and access.
- Space management.
- Reservation management.
- Payment management.
- Notification management.

These capabilities have different responsibilities and can evolve independently.

## 2. Problem

An architecture must be selected that demonstrates distributed-systems principles without adding unnecessary complexity.

The architecture must:

- Separate business responsibilities.
- Keep clear boundaries between capabilities.
- Avoid a shared database between services.
- Allow synchronous and asynchronous communication.
- Isolate failures between components.
- Allow independent scalability when needed.
- Make it easier to add future capabilities.
- Remain implementable by a single developer.
- Keep architectural documentation clear.

## 3. Decision

A microservices architecture will be used.

Each microservice will be responsible for a defined business capability and will have an explicit responsibility boundary.

The initial microservices will be:

1. Identity Service.
2. Space Service.
3. Reservation Service.
4. Payment Service.
5. Notification Service.

These services correspond to the bounded contexts identified in the initial Context Map.

The number of microservices is not a goal by itself.

A new microservice may be added only when it has a clear business responsibility, independent data ownership, and a real need for independent evolution, scalability, or deployment.

## 4. Rationale

Microservices are selected because the project must demonstrate characteristics of distributed systems, including:

- Communication between independent processes.
- Synchronous communication through APIs.
- Asynchronous communication through events.
- Independent data ownership.
- Distributed consistency.
- Idempotence.
- Failure isolation.
- Independent scalability.
- Independent deployment.
- Independent evolution of capabilities.

The architecture also allows the project to grow without turning the system into a single component that is hard to maintain.

## 5. Alternatives considered

### 5.1 Traditional monolith

One alternative is to implement the whole system as a single application.

#### Advantages

- Lower initial complexity.
- Faster development.
- A single running process.
- A single database.
- Lower communication complexity.

#### Disadvantages

- Does not adequately demonstrate communication between microservices.
- There is no real isolation between capabilities.
- Failures can affect the whole application.
- Scaling must be applied to the entire system.
- There is no independent deployment per capability.
- It does not adequately demonstrate distributed consistency and asynchronous communication.

#### Decision

Rejected.

Although it would be simpler for a single developer, it does not meet the main academic goal of the project.

### 5.2 Modular monolith

Another alternative is a single deployment with clearly separated internal modules.

#### Advantages

- Keeps domain boundaries.
- Lower operational complexity.
- Easier testing.
- Supports an orderly initial evolution.
- Can be suitable for small systems.

#### Disadvantages

- Modules still share the same process.
- There is no real deployment independence.
- Distributed communication is not fully demonstrated.
- Infrastructure failures can affect the whole process.
- Independent scalability of each capability is limited.

#### Decision

Rejected as the final architecture.

Principles of the modular monolith, such as clear boundaries and separation of responsibilities, will still be used inside each microservice.

### 5.3 Microservices

The third alternative separates the main business capabilities into independent services.

#### Advantages

- Allows independent deployment.
- Allows independent scalability.
- Makes failure isolation easier.
- Allows synchronous and asynchronous communication.
- Each service can own its data.
- Demonstrates distributed-systems concepts.
- Makes it easier to add new capabilities.

#### Disadvantages

- Higher operational complexity.
- Distributed communication.
- Possible network failures.
- Harder integration testing.
- Need for observability.
- Need to handle distributed consistency.
- More components to maintain.

#### Decision

Selected.

The extra complexity is justified by the academic and architectural goals of the project.

## 6. Microservice boundaries

The initial boundaries come from the bounded contexts.

### Identity Service

Responsibility:

- Users.
- Authentication.
- Authorization.
- Roles and permissions.

It will not be responsible for reservations, payments, or spaces.

### Space Service

Responsibility:

- Reservable spaces.
- Availability.
- Blocked periods.

It will not be responsible for creating or managing reservations.

### Reservation Service

Responsibility:

- Reservation creation.
- States.
- Confirmation.
- Cancellation.
- Reservation lifecycle.

It will be the main coordinator of the business flow.

### Payment Service

Responsibility:

- Payment operations.
- Payment states.
- Transaction identifiers.
- Payment-related events.

It will not be responsible for the full reservation lifecycle.

### Notification Service

Responsibility:

- Receiving events.
- Creating notifications.
- Delivery.
- Delivery status.

It will not be responsible for modifying reservations or payments directly.

## 7. Data ownership

Each microservice will own its data.

The initial distribution is:

| Service        | Main data                                |
| -------------- | ---------------------------------------- |
| Identity       | Users, credentials, roles, and permissions |
| Spaces         | Spaces, availability, and blocks         |
| Reservations   | Reservations and states                  |
| Payments       | Payments and states                      |
| Notifications  | Notifications and deliveries             |

A microservice must not query another microservice's tables directly.

Information between services will be exchanged through:

- APIs.
- Events.
- Explicit contracts.

## 8. Communication strategy

Two main mechanisms will be used.

### Synchronous communication

It will be used when the consuming service needs an immediate response.

Example:

Reservation Service
|
| Availability query
v
Space Service

Another example:

Reservation Service
|
| Payment request
v
Payment Service

The initial synchronous communication will be implemented with REST APIs.

### Asynchronous communication

It will be used when an operation can continue in a decoupled way.

Example:

Payment Service
|
| PaymentConfirmed
v
Reservation Service

Another example:

Reservation Service
|
| ReservationConfirmed
v
Notification Service

The messaging mechanism will be defined in a later architectural decision when the infrastructure requirements are established.

## 9. Consistency

The architecture will use different consistency strategies.

### Critical operations

Operations related to:

- Availability.
- Reservation creation.
- Reservation confirmation.
- Critical payment states.

must prevent serious inconsistencies such as:

- Double booking.
- Duplicate payment.
- Impossible states.

### Non-critical operations

Notifications, reports, and analytical capabilities may use eventual consistency.

For example:

A reservation may be confirmed even if the notification is still pending delivery.

## 10. Idempotence

Distributed operations must consider duplicate requests or events.

Processing of important events must be idempotent when applicable.

For example:

If PaymentConfirmed is received twice, the system must not generate two financial confirmations or duplicate effects.

The concrete implementation of idempotence will be defined while developing the corresponding service.

## 11. Failure isolation

Services must minimize the impact of failures on each other.

For example:

- A failure in the Notification Service must not cancel a confirmed reservation.
- A failed payment event must be processable without affecting Space Service availability.
- A temporarily unavailable service must not cause another service to access its database directly.
- Event consumers must handle duplicate events when necessary.

The exact failure behavior will be documented and tested during implementation.

## 12. Scalability

The architecture will allow services to scale independently.

For example:

If availability queries receive more load than reservation creation, the Space Service can be scaled independently.

Likewise, if notification processing grows significantly, the Notification Service can add instances without scaling every other service.

Scalability does not mean adding unnecessary microservices.

Each new service must be justified by a business capability or a real technical need.

## 13. Extensibility

The architecture will allow new capabilities to be added as new services when needed.

Potential examples:

- Review Service.
- Promotion Service.
- Analytics Service.
- Search Service.
- Loyalty Service.
- Audit Service.

These services are not part of the initial MVP.

The current architecture must allow them to be added without directly modifying the databases of existing services.

## 14. Positive consequences

The decision provides:

- Alignment with the Distributed Systems course.
- Clearly defined business boundaries.
- Independent deployment.
- Independent scalability.
- Failure isolation.
- Synchronous and asynchronous communication.
- Independent data ownership.
- Easier addition of new capabilities.
- The ability to demonstrate real distributed-systems patterns and problems.
- An architecture prepared for future evolution.

## 15. Negative consequences

The decision also introduces costs:

- Higher development complexity.
- Higher deployment complexity.
- Need to operate several services.
- Higher testing complexity.
- Network latency.
- Partial failures.
- Need for observability.
- Need to handle distributed consistency.
- Higher documentation effort.

Because the project is developed individually, the MVP will be limited to keep complexity under control.

## 16. Scope constraints

To keep the architecture manageable, the following constraints apply:

- The MVP will initially use five candidate microservices.
- Services will not be created only to increase the number of microservices.
- Future capabilities remain outside the MVP.
- Each service must have automated tests.
- Each service must have minimum documentation.
- Each service must have a clearly bounded responsibility.
- Shared databases will not be used.
- Important architectural decisions must be documented with ADRs.

## 17. Criteria for adding a new microservice

A new microservice may be added only when it meets one or more justified conditions:

1. It represents an independent business capability.
2. It has a clear domain boundary.
3. It requires independent scalability.
4. It requires independent deployment.
5. It has data that must be managed separately.
6. It significantly reduces system coupling.
7. It isolates a responsibility with a different evolution cycle.

The number of microservices will not be used as a success metric.

## 18. Decision status

Accepted.

The microservices architecture will be used as the architectural direction for the Distributed Reservation Platform.

The boundaries defined in this ADR correspond to the initial design and may be refined through future ADRs when new relevant information appears.

## 19. Immutability rule

Once accepted, this ADR **must not be modified**.

Any later change to this architectural decision must be documented in a new file (`adr-002-*.md`) that:

- Explicitly references ADR-001 as the record being replaced.
- States what changed and why.
- Documents the context, decision, alternatives, and consequences of the new decision.

## 20. Related documents

This decision is related to:

- [`context-map.md`](./context-map.md)
- [`user-stories.md`](./user-stories.md)
- [`../../01-week/hu-status/prd.md`](../../01-week/hu-status/prd.md)
- [`README.md`](./README.md)
- [`README.es.md`](./README.es.md)

The Context Map defines the relationships between bounded contexts.

This ADR documents the decision to use those boundaries as the basis for the microservices architecture.
