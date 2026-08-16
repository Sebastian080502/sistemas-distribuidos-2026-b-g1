<!-- PLANTILLA HU-STATUS - no elimine los marcadores <!-- ... -->.

     La calificación semanal se lee AUTOMÁTICAMENTE desde este archivo:
       01-week/hu-status/README.md  (dentro de SU fork). Inglés. -->

# Estado Semanal - Semana 01

<!-- CONFIG-START - debe coincidir con el CONFIG de su repositorio de perfil (username/username) -->

- FULL_NAME: Juan Sebastian Osorio Fierro
- GITHUB_USER: Sebastian080502
- TEAM: Grupo 1 - Distributed Reservation Platform
- SPRINT_GOAL: Preparar el producto, el dominio, el backlog, la estrategia de consistencia y la organización del proyecto necesarias para iniciar el primer Sprint formal de desarrollo.
<!-- CONFIG-END -->

## 1. Historias de usuario trabajadas esta semana

| HU ID      | Título                                                        | Estado (todo/doing/done) | Evidencia (URL de PR o commit)                               |
| ---------- | ------------------------------------------------------------- | ------------------------ | ------------------------------------------------------------ |
| HU-RES-001 | Consultar espacios disponibles por fecha y hora               | doing                    | Pendiente - commit de recuperación de Semana 01              |
| HU-RES-002 | Consultar la disponibilidad de un espacio seleccionado        | todo                     | Pendiente - implementación planificada para un Sprint futuro |
| HU-RES-003 | Crear una reserva para un espacio disponible                  | todo                     | Pendiente - implementación planificada para un Sprint futuro |
| HU-RES-004 | Consultar el estado de una reserva                            | todo                     | Pendiente - implementación planificada para un Sprint futuro |
| HU-RES-005 | Procesar el pago asociado a una reserva                       | todo                     | Pendiente - implementación planificada para un Sprint futuro |
| HU-RES-006 | Confirmar una reserva después de un pago exitoso              | todo                     | Pendiente - flujo distribuido por implementar                |
| HU-RES-007 | Recibir una notificación cuando una reserva cambie de estado  | todo                     | Pendiente - integración asíncrona futura                     |
| HU-RES-008 | Gestionar espacios reservables y sus reglas de disponibilidad | doing                    | Pendiente - definición del dominio en progreso               |

## 2. Mi contribución individual

- Definí el dominio del proyecto como una **Plataforma Distribuida de Gestión y Reservas de Espacios**, enfocada en la gestión de espacios reservables, disponibilidad, reservas, pagos y notificaciones.
- Identifiqué el problema principal a resolver: evitar reservas conflictivas, proporcionar información confiable sobre la disponibilidad, coordinar los procesos de reserva y pago, y mantener trazabilidad del ciclo de vida de una reserva.
- Definí el alcance inicial del producto y establecí límites para evitar un crecimiento descontrolado durante el periodo de desarrollo disponible.
- Identifiqué los principales actores y operaciones de negocio del MVP inicial: usuarios, administradores de espacios, consultas de disponibilidad, creación de reservas, procesamiento de pagos, confirmación de reservas y notificaciones.
- Definí el backlog inicial mediante las HU-RES-001 a HU-RES-008, derivando las historias de las capacidades de negocio y no de detalles técnicos de implementación.
- Analicé los requisitos de consistencia de las operaciones principales. La creación de reservas y la confirmación de pagos requieren consistencia fuerte y procesamiento idempotente, mientras que las notificaciones y futuras capacidades de reportes pueden utilizar consistencia eventual.
- Definí la estrategia inicial de semántica de entrega. Los comandos y eventos críticos utilizarán at-least-once cuando sea necesario, junto con consumidores idempotentes cuando exista posibilidad de procesamiento duplicado. Las consultas dirigidas al usuario utilizarán comunicación síncrona de solicitud y respuesta.
- Establecí que el proyecto utilizará **microservicios desde el inicio**, con límites de servicio derivados de capacidades de negocio y bounded contexts, en lugar de capas técnicas.
- Establecí que cada microservicio deberá ser propietario de sus datos y que se evitará compartir una base de datos entre servicios para prevenir la construcción de un monolito distribuido.
- Definí la escalabilidad como un requisito del proyecto. El MVP inicial contendrá únicamente los servicios justificados por el dominio, mientras que la arquitectura permitirá incorporar nuevos bounded contexts cuando exista una razón de negocio o arquitectónica.
- Establecí la documentación como un entregable fundamental del proyecto, incluyendo el Brief del Producto, documentación del dominio, diagramas C4, ADRs, contratos de API, contratos de eventos, estrategia de pruebas, seguridad y despliegue.
- Definí el enfoque de gestión del proyecto como **Scrumban**, utilizando Sprints semanales para el desarrollo y Kanban mediante GitHub Projects para visualizar continuamente el trabajo.
- Preparé la organización inicial del proyecto necesaria para completar las actividades de recuperación e iniciar el primer Sprint formal de desarrollo.

## 3. Bloqueantes y riesgos

- La definición del proyecto se realizó después del inicio de las primeras semanas del curso, por lo que las actividades de las Semanas 01 y 02 se están recuperando antes de iniciar el primer Sprint formal de desarrollo.
- El proyecto tuvo que seleccionarse cuidadosamente para cumplir los requisitos de Sistemas Distribuidos y mantener al mismo tiempo un alcance controlado para un único desarrollador.
- Los límites definitivos de los microservicios todavía no se han fijado arbitrariamente. Serán validados mediante el análisis del dominio y los bounded contexts para evitar una fragmentación innecesaria.
- El proyecto está siendo desarrollado individualmente, lo que representa un riesgo de alcance y tiempo. Por esta razón, el MVP inicial debe mantenerse controlado.
- El número de microservicios no se incrementará únicamente para aumentar la complejidad aparente del proyecto. Cada nuevo servicio deberá representar una capacidad de negocio justificada.
- El primer Sprint formal de desarrollo todavía no ha comenzado. El trabajo actual corresponde a la preparación y recuperación de las primeras actividades académicas.
- Las decisiones técnicas deberán mantenerse abiertas hasta que se complete el análisis arquitectónico correspondiente.

## 4. Plan para la próxima semana

- Completar el análisis de Domain-Driven Design de la plataforma de reservas.
- Identificar y validar los bounded contexts y sus responsabilidades.
- Crear el Context Map inicial y definir las relaciones entre los bounded contexts.
- Definir los límites iniciales de los microservicios a partir de las capacidades de negocio.
- Documentar la decisión arquitectónica mediante ADR-001.
- Definir la estrategia inicial de comunicación entre servicios, incluyendo comunicación REST síncrona y mensajería asíncrona cuando sea apropiado.
- Definir la propiedad de los datos y la estrategia de persistencia de cada servicio.
- Refinar el backlog inicial y agregar criterios de aceptación verificables a las principales historias de usuario.
- Configurar el tablero de GitHub Projects y preparar el primer Sprint formal.
- Iniciar la implementación del primer microservicio después de validar los límites del dominio y la arquitectura.

## 5. Autoverificación de cumplimiento

- [x] Problema y alcance inicial del proyecto definidos
- [x] Backlog inicial creado a partir de necesidades de negocio
- [x] Estrategia de consistencia definida para operaciones críticas
- [x] Semántica de entrega considerada para operaciones distribuidas
- [x] Microservicios establecidos como dirección arquitectónica
- [x] Estrategia inicial de escalabilidad definida
- [x] No se han incluido secretos ni credenciales
- [ ] Bounded Contexts y Context Map completados
- [ ] ADR-001 publicado
- [ ] Pruebas automatizadas agregadas
- [ ] Rama HU por entorno + PR completada
- [ ] Primer Sprint formal iniciado

Notas sobre los elementos pendientes:

- Los Bounded Contexts, el Context Map y el ADR-001 corresponden al trabajo arquitectónico de recuperación de la Semana 02.
- Todavía no se ha implementado ningún microservicio porque el proyecto se encuentra en la etapa de preparación y definición arquitectónica.
- El primer Sprint formal comenzará después de completar las actividades de recuperación y configurar GitHub Projects.
- La evidencia de ramas y PRs se agregará cuando la primera HU de implementación entre en desarrollo siguiendo el flujo definido para el repositorio.

## 6. Enlaces de evidencia

- Brief del producto: [`prd.md`](./prd.md) - Plataforma Distribuida de Gestión y Reservas de Espacios.
- Material de aprendizaje del curso (OVAs): https://code-corhuila.github.io/ova-web/2026-B/distribuidos/
- Repositorio: https://github.com/Sebastian080502/sistemas-distribuidos-2026-b-g1
- GitHub Project: Pendiente - tablero por crear durante la etapa de preparación del proyecto.

El proyecto seguirá el principio de **separar servicios por una razón y no por moda**. Los microservicios se derivarán de límites de negocio significativos, contratos explícitos, propiedad independiente de los datos y necesidades justificadas de escalabilidad o despliegue.
