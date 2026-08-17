<!-- PLANTILLA HU-STATUS - no elimine los marcadores <!-- ... -->.

     La calificación semanal se lee AUTOMÁTICAMENTE desde este archivo:
       02-week/hu-status/README.md  (dentro de SU fork). Inglés. -->

# Estado Semanal - Semana 02

<!-- CONFIG-START - debe coincidir con el CONFIG de su repositorio de perfil (username/username) -->

- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Grupo 1 - Distributed Reservation Platform
- SPRINT_GOAL: Definir los límites iniciales del dominio, la dirección arquitectónica, el flujo de trabajo Git y el backlog necesarios para preparar el primer Sprint formal de desarrollo.
<!-- CONFIG-END -->

## 1. Historias de usuario trabajadas esta semana

| HU ID      | Título                                                                              | Estado (todo/doing/done) | Evidencia (URL de PR o commit)                               |
| ---------- | ----------------------------------------------------------------------------------- | ------------------------ | ------------------------------------------------------------ |
| HU-ADR-001 | Formalizar la decisión arquitectónica de la Plataforma Distribuida de Reservas      | done                     | [`adr-001-architecture.md`](./adr-001-architecture.md)       |
| HU-ARC-001 | Identificar los bounded contexts de la plataforma de reservas                       | done                     | [`context-map.md`](./context-map.md)                         |
| HU-ARC-002 | Definir el Context Map inicial y las relaciones entre bounded contexts              | done                     | [`context-map.md`](./context-map.md)                         |
| HU-ARC-003 | Documentar la decisión arquitectónica inicial de microservicios                     | done                     | [`adr-001-architecture.md`](./adr-001-architecture.md)       |
| HU-ARC-004 | Definir criterios de aceptación verificables para el backlog arquitectónico inicial | done                     | [`user-stories.md`](./user-stories.md)                       |

## 2. Mi contribución individual

- Realicé el ejercicio práctico de Git y GitHub correspondiente a la Sesión 2 de la Semana 02.
- Practiqué el flujo de desarrollo basado en ramas utilizando ramas de desarrollo y QA.
- Creé y utilicé ramas asociadas a historias de usuario durante el ejercicio.
- Practiqué la creación y gestión de Pull Requests entre ramas.
- Practiqué la integración de cambios entre ramas mediante cherry-pick.
- Revisé la relación entre ramas de historias de usuario, ramas de entorno, Pull Requests e integración.
- Apliqué los conceptos del flujo Git a la organización que utilizará el proyecto de la plataforma de reservas.
- Analicé el dominio de la plataforma de reservas utilizando el enfoque Domain-Driven Design presentado durante la sesión.
- Identifiqué los bounded contexts iniciales a partir de capacidades de negocio y no de capas técnicas.
- Identifiqué Identidad y Acceso, Gestión de Espacios, Gestión de Reservas, Gestión de Pagos y Gestión de Notificaciones como candidatos iniciales a bounded contexts.
- Definí las responsabilidades iniciales de cada bounded context.
- Establecí que cada futuro microservicio deberá ser propietario de los datos requeridos por su capacidad de negocio.
- Establecí que los servicios no deberán acceder directamente a la base de datos de otro servicio.
- Definí las relaciones iniciales entre los bounded contexts para el Context Map.
- Identifiqué la comunicación síncrona como apropiada para consultas inmediatas y determinados comandos.
- Identifiqué la comunicación asíncrona como apropiada para eventos desacoplados como resultados de pagos, cambios de estado de reservas y notificaciones.
- Definí la dirección arquitectónica inicial como una arquitectura de microservicios basada en capacidades de negocio y bounded contexts.
- Formalicé la decisión arquitectónica en el ADR-001, incluyendo contexto, alternativas, consecuencias, propiedad de datos, estrategia de comunicación y la regla de inmutabilidad.
- Convertí la decisión arquitectónica en historias de usuario verificables con criterios de aceptación Given/When/Then.
- Preparé la estructura necesaria para iniciar el primer Sprint formal de desarrollo.

## 3. Bloqueantes y riesgos

- Las actividades de la Semana 02 se están recuperando después de que la definición inicial del proyecto se completara posteriormente a lo previsto.
- La tecnología de comunicación y el message broker todavía no han sido seleccionados porque esas decisiones de implementación corresponden a un ADR posterior.
- El ciclo de vida completo de las reservas y los estados de pago todavía requieren validación antes de la implementación.
- El proyecto está siendo desarrollado individualmente, por lo que la cantidad de servicios y funcionalidades debe mantenerse controlada.
- El ejercicio de ramas de Git fue completado como actividad de aprendizaje, pero el flujo completo todavía debe aplicarse a las HUs reales del proyecto.
- El primer Sprint formal de desarrollo todavía no ha comenzado porque las actividades actuales corresponden a la preparación y definición arquitectónica del proyecto.

## 4. Plan para la próxima semana

- Configurar el tablero de GitHub Projects utilizando Backlog, Ready, In progress, In review y Done.
- Definir el Sprint Goal y el Sprint Backlog del primer Sprint formal de desarrollo.
- Seleccionar un conjunto reducido de historias del MVP que puedan terminarse en una semana.
- Definir los contratos iniciales de API y eventos entre servicios.
- Crear la primera rama de desarrollo de una HU siguiendo el flujo Git validado.
- Iniciar la implementación del primer microservicio después de validar los límites arquitectónicos.

## 5. Autoverificación de cumplimiento

- [x] Flujo de ramas e integración Git practicado
- [x] Bounded contexts identificados
- [x] Límites iniciales de servicios analizados
- [x] Principio de propiedad de datos definido
- [x] Comunicación síncrona y asíncrona identificadas
- [x] Dirección arquitectónica inicial definida
- [x] Backlog inicial refinado
- [x] Context Map publicado
- [x] ADR-001 publicado
- [x] Criterios de aceptación verificables completados
- [ ] Tablero de GitHub Projects configurado
- [ ] Primer Sprint formal iniciado

Notas sobre los elementos pendientes:

- GitHub Projects será configurado antes del primer Sprint formal de desarrollo.
- El primer Sprint comenzará únicamente después de completar las actividades de preparación y definición arquitectónica.

## 6. Enlaces de evidencia

- Ejercicio de ramas Git/GitHub: Pendiente - agregar evidencia del repositorio del ejercicio.
- Brief del producto: [`../../01-week/hu-status/prd.md`](../../01-week/hu-status/prd.md).
- Context Map: [`context-map.md`](./context-map.md).
- ADR-001: [`adr-001-architecture.md`](./adr-001-architecture.md).
- Historias de usuario y criterios de aceptación: [`user-stories.md`](./user-stories.md).
- Material de aprendizaje del curso (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repositorio: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1

El proyecto sigue el principio de **separar servicios por una razón y no por moda**. Los límites iniciales de los microservicios se derivan de capacidades de negocio significativas, bounded contexts, propiedad independiente de los datos, contratos explícitos y necesidades justificadas de escalabilidad o despliegue.
