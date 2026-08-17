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

| HU ID      | Título                                                                              | Estado (todo/doing/done) | Evidencia (URL de PR o commit)                         |
| ---------- | ----------------------------------------------------------------------------------- | ------------------------ | ------------------------------------------------------ |
| HU-ADR-001 | Formalizar la decisión arquitectónica de la Plataforma Distribuida de Reservas      | done                     | [`adr-001-architecture.md`](./adr-001-architecture.md) |
| HU-ARC-001 | Identificar los bounded contexts de la plataforma de reservas                       | done                     | [`context-map.md`](./context-map.md)                   |
| HU-ARC-002 | Definir el Context Map inicial y las relaciones entre bounded contexts              | done                     | [`context-map.md`](./context-map.md)                   |
| HU-ARC-003 | Documentar la decisión arquitectónica inicial de microservicios                     | done                     | [`adr-001-architecture.md`](./adr-001-architecture.md) |
| HU-ARC-004 | Definir criterios de aceptación verificables para el backlog arquitectónico inicial | done                     | [`user-stories.md`](./user-stories.md)                 |

## 2. Mi contribución individual

- Completé el ejercicio de ramas de la Sesión 2 de la Semana 02: `main` → `qa` → `develop`, rama de HU, Pull Request y cherry-pick.
- Analicé el dominio de la plataforma de reservas con Domain-Driven Design.
- Identifiqué Identidad y Acceso, Gestión de Espacios, Gestión de Reservas, Gestión de Pagos y Gestión de Notificaciones como bounded contexts.
- Establecí que cada futuro microservicio es propietario de sus datos y no debe leer la base de datos de otro servicio.
- Documenté llamadas REST síncronas para consultas inmediatas y eventos asíncronos para pagos, cambios de reserva y notificaciones.
- Formalicé la decisión de microservicios en el ADR-001, incluyendo alternativas, consecuencias y la regla de inmutabilidad.
- Convertí la decisión arquitectónica en historias de usuario verificables con criterios Given/When/Then.

## 3. Bloqueantes y riesgos

- La Semana 02 se está recuperando después de que la definición del producto se completara más tarde de lo previsto.
- El message broker todavía no se ha seleccionado; esa decisión corresponde a un ADR posterior.
- El proyecto lo desarrolla una sola persona, por lo que la cantidad de servicios debe mantenerse controlada.
- GitHub Projects todavía no está publicado porque GitHub CLI no está disponible en este entorno.
- El primer Sprint formal de desarrollo todavía no ha comenzado.

## 4. Plan para la próxima semana

- Crear el GitHub Project **Distributed Reservation Platform - Sprint** con columnas Backlog, Ready, In progress, In review y Done.
- Definir límites WIP: Ready = 5, In progress = 2, In review = 3.
- Definir el Sprint Goal y un Sprint Backlog reducido, comenzando por Gestión de Espacios.
- Crear la primera rama de HU (`hu-xxx-dev`) y abrir un Pull Request hacia `develop`.
- Iniciar el primer microservicio cuando los límites arquitectónicos se mantengan estables.

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

- GitHub Projects permanece pendiente: el nombre previsto es **Distributed Reservation Platform - Sprint**. Todavía no hay URL del tablero.
- El primer Sprint comienza después de crear el tablero y de adjuntar las capturas del ejercicio Git.

## 6. Enlaces de evidencia

- Resúmenes de sesión: [`resumen-sesiones/resumen_sistemas_distribuidos_semana_2.svg`](./resumen-sesiones/resumen_sistemas_distribuidos_semana_2.svg)
- Ejercicio de ramas Git: [`ejercicio-ramas-git/README.md`](./ejercicio-ramas-git/README.md)
- Brief del producto: [`../../01-week/hu-status/prd.md`](../../01-week/hu-status/prd.md)
- Context Map: [`context-map.md`](./context-map.md)
- ADR-001: [`adr-001-architecture.md`](./adr-001-architecture.md)
- Historias de usuario y criterios de aceptación: [`user-stories.md`](./user-stories.md)
- GitHub Project: pendiente — el tablero todavía no fue creado
- Material de aprendizaje del curso (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repositorio: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1

El proyecto sigue el principio de **separar servicios por una razón y no por moda**. Los límites iniciales de los microservicios se derivan de capacidades de negocio, bounded contexts, propiedad independiente de los datos, contratos explícitos y necesidades justificadas de escalabilidad o despliegue.
