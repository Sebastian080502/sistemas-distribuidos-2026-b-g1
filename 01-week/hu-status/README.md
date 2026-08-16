<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers.

     Your weekly grade is read AUTOMATICALLY from this file:
       01-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 01

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->

- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Group 1 - Distributed Reservation Platform
- SPRINT_GOAL: Prepare the product, domain, backlog, consistency strategy, and project organization required to start the first formal development Sprint.
<!-- CONFIG-END -->

## 1. User stories worked this week

| HU ID      | Title                                                        | Status (todo/doing/done) | Evidence (PR or commit URL)                                    |
| ---------- | ------------------------------------------------------------ | ------------------------ | -------------------------------------------------------------- |
| HU-RES-001 | Consult available spaces by date and time                    | doing                    | Pending - Week 01 recovery commit                              |
| HU-RES-002 | Consult the availability of a selected space                 | todo                     | Pending - implementation planned for a future Sprint           |
| HU-RES-003 | Create a reservation for an available space                  | todo                     | Pending - implementation planned for a future Sprint           |
| HU-RES-004 | Consult the status of a reservation                          | todo                     | Pending - implementation planned for a future Sprint           |
| HU-RES-005 | Process the payment associated with a reservation            | todo                     | Pending - implementation planned for a future Sprint           |
| HU-RES-006 | Confirm a reservation after successful payment               | todo                     | Pending - distributed workflow to be implemented               |
| HU-RES-007 | Receive a notification when a reservation changes its status | todo                     | Pending - asynchronous integration planned for a future Sprint |
| HU-RES-008 | Manage reservable spaces and their availability rules        | doing                    | Pending - domain definition in progress                        |

## 2. My individual contribution

- Defined the project domain as a **Distributed Platform for Space Management and Reservations**, focused on managing reservable spaces, availability, reservations, payments, and notifications.
- Identified the main problem to solve: avoiding conflicting reservations, providing reliable availability information, coordinating reservation and payment processes, and maintaining a traceable reservation lifecycle.
- Defined the initial product scope and established boundaries to prevent uncontrolled growth during the remaining development period.
- Identified the main actors and business operations involved in the initial MVP: users, space administrators, availability queries, reservation creation, payment processing, reservation confirmation, and notifications.
- Defined the initial product backlog through HU-RES-001 to HU-RES-008, deriving the stories from business capabilities rather than technical implementation details.
- Analyzed the consistency requirements of the main operations. Reservation creation and payment confirmation require strong consistency and idempotent processing, while notifications and future reporting capabilities can use eventual consistency.
- Defined the initial delivery semantics strategy. Critical commands and events will use at-least-once delivery combined with idempotent consumers where duplicate processing is possible. User-facing queries will use request/response semantics.
- Established the architectural direction that the project will use **microservices from the beginning**, with service boundaries derived from business capabilities and bounded contexts rather than technical layers.
- Established that each microservice must own its data and that a shared database between services will be avoided in order to prevent a distributed monolith.
- Defined scalability as a project requirement: the initial MVP will contain only the services justified by the domain, while the architecture will allow future bounded contexts such as reviews, analytics, promotions, search, loyalty, or auditing to be introduced when justified.
- Defined the intention to reuse the project as a foundation for the Software Architecture course, keeping the architecture compatible with a future AWS deployment involving EC2, RDS, IAM, VPC, and additional infrastructure requirements that may be introduced later.
- Established documentation as a first-class project deliverable, including the Product Brief, domain documentation, C4 diagrams, ADRs, API contracts, testing strategy, security documentation, and deployment documentation.
- Defined the project management approach as **Scrumban**, using weekly Sprints for the development process and Kanban through GitHub Projects for continuous work visualization.
- Prepared the initial project organization required to complete the recovery activities and begin the first formal development Sprint.

## 3. Blockers and risks

- The project definition was established after the first weeks of the course, so the Week 01 and Week 02 activities are being recovered before starting the first formal development Sprint.
- The project had to be selected carefully to satisfy the requirements of Distributed Systems while also remaining reusable for the Software Architecture course and its future AWS deployment.
- The exact future requirements of the Software Architecture course are not fully known yet. The architecture therefore needs to remain extensible without introducing unnecessary infrastructure or services prematurely.
- The final number of microservices has not been fixed arbitrarily. Service boundaries will be validated through domain analysis and bounded contexts to avoid unnecessary service fragmentation.
- The project is being developed individually, which creates a scope and time risk. The initial MVP must therefore remain controlled and additional microservices will only be introduced when they represent a justified business capability.
- The first formal development Sprint has not started yet. The current work corresponds to project initialization and recovery of the first academic activities.
- AWS deployment is planned as a later stage. Infrastructure decisions will be documented so that the system can evolve from the local containerized environment to AWS without redesigning the business domain.

## 4. Plan for next week

- Complete the Domain-Driven Design analysis of the reservation platform.
- Identify and validate the bounded contexts and their responsibilities.
- Create the initial Context Map and define the relationships between bounded contexts.
- Define the initial microservice boundaries based on business capabilities.
- Document the architectural decision in ADR-001.
- Define the initial communication strategy between services, including synchronous REST communication and asynchronous messaging where appropriate.
- Define the ownership and persistence strategy for each service.
- Refine the initial backlog and add testable acceptance criteria to the main user stories.
- Configure the GitHub Projects board and prepare the first formal weekly Sprint.
- Start implementation of the first microservice only after the domain and architectural boundaries have been validated.

## 5. Compliance self-check

- [x] Project problem and initial scope defined
- [x] Initial backlog created from business needs
- [x] Consistency strategy defined for critical operations
- [x] Delivery semantics considered for distributed operations
- [x] Microservices selected as the architectural direction
- [x] Initial scalability strategy defined
- [x] No secrets or credentials committed
- [ ] Bounded Contexts and Context Map completed
- [ ] ADR-001 published
- [ ] Automated tests added
- [ ] Per-environment HU branch + PR completed
- [ ] First formal Sprint started

Notes on the unchecked items:

- Bounded Contexts, Context Map, and ADR-001 belong to the architectural work being completed during the Week 02 recovery.
- No production microservice implementation has been committed yet because the project is currently in the initialization and architecture-definition stage.
- The first formal Sprint will begin after the recovery activities and GitHub Projects setup are completed.
- Branch and PR evidence will be added once the first implementation HU enters development following the repository workflow.

## 6. Evidence links

- Product brief: [`prd.md`](./prd.md) - Distributed Platform for Space Management and Reservations.
- Course learning material (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repository: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
- GitHub Project: Pending - project board to be created during project initialization.

The project follows the principle of **splitting services for a reason rather than for fashion**. Microservices will be derived from meaningful business boundaries, explicit contracts, independent data ownership, and justified scalability or deployment needs.
