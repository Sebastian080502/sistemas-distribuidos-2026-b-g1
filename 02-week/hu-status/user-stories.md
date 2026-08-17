# Backlog de Historias de Usuario

## Plataforma Distribuida de Reservas

Este documento contiene el backlog inicial de historias de usuario de la Plataforma Distribuida de Reservas.

Las historias se derivan de las necesidades identificadas en el PRD y están relacionadas con los bounded contexts definidos en el Context Map.

Cada historia debe contar con criterios de aceptación verificables antes de considerarse lista para implementación.

---

## HU-ADR-001 — Formalizar la decisión arquitectónica

### Historia

Como estudiante del curso de Sistemas Distribuidos, quiero documentar la decisión arquitectónica inicial de la Plataforma Distribuida de Reservas, para dejar justificados los bounded contexts, los microservicios candidatos y las alternativas descartadas.

### Bounded Context

Arquitectura / transversal.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Bounded contexts identificados

Dado el análisis inicial del dominio,
cuando se documente la arquitectura,
entonces deben identificarse los bounded contexts principales de la plataforma.

#### CA-002 — Justificación de microservicios

Dado que el proyecto corresponde a Sistemas Distribuidos,
cuando se seleccione la arquitectura,
entonces debe justificarse el uso de microservicios.

#### CA-003 — Alternativas consideradas

Cuando se documente la decisión arquitectónica,
entonces deben registrarse las alternativas consideradas y las razones para descartarlas.

#### CA-004 — Responsabilidad de cada servicio

Cuando se defina la arquitectura inicial,
entonces cada microservicio debe tener una responsabilidad de negocio claramente delimitada.

#### CA-005 — Comunicación síncrona y asíncrona

Cuando se documente la comunicación entre servicios,
entonces deben identificarse las interacciones síncronas y asíncronas previstas.

#### CA-006 — Propiedad de datos

Cuando se documente la propiedad de datos,
entonces debe establecerse que cada microservicio es responsable de sus propios datos.

#### CA-007 — Inmutabilidad del ADR

Cuando el ADR sea aceptado,
entonces cualquier modificación posterior de la decisión arquitectónica deberá registrarse mediante un nuevo ADR.

### Evidencia

- [`adr-001-architecture.md`](./adr-001-architecture.md)
- [`context-map.md`](./context-map.md)

---

## HU-ARC-001 — Identificar bounded contexts

### Historia

Como equipo, quiero definir los bounded contexts de la plataforma para establecer límites de negocio claros.

### Prioridad

Alta.

### Criterios de aceptación

- Identity & Access está definido.
- Space Management está definido.
- Reservation Management está definido.
- Payment Management está definido.
- Notification Management está definido.
- Cada contexto tiene una responsabilidad claramente documentada.

### Evidencia

- [`context-map.md`](./context-map.md)

---

## HU-ARC-002 — Definir el Context Map inicial

### Historia

Como equipo, quiero documentar el Context Map para establecer las relaciones entre los bounded contexts.

### Prioridad

Alta.

### Criterios de aceptación

- Las relaciones entre contextos están documentadas.
- El propietario de cada dato está identificado.
- No existe acceso directo entre bases de datos.
- Los contratos de comunicación síncrona y asíncrona están identificados.

### Evidencia

- [`context-map.md`](./context-map.md)

---

## HU-ARC-003 — Documentar la decisión arquitectónica mediante ADR-001

### Historia

Como equipo, quiero documentar la decisión arquitectónica mediante ADR-001.

### Prioridad

Alta.

### Criterios de aceptación

- Contexto documentado.
- Decisión documentada.
- Alternativas consideradas.
- Alternativas rechazadas justificadamente.
- Consecuencias positivas y negativas documentadas.

### Evidencia

- [`adr-001-architecture.md`](./adr-001-architecture.md)

---

## HU-RES-001 — Consultar espacios disponibles

### Historia

Como usuario, quiero consultar los espacios disponibles para una fecha y rango horario, para poder seleccionar un espacio que pueda reservar.

### Bounded Context

Gestión de Espacios.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Consulta válida

Dado un rango de fecha y hora válido,
cuando el usuario consulte la disponibilidad,
entonces el sistema debe mostrar los espacios disponibles para ese periodo.

#### CA-002 — Espacio ya reservado

Dado un espacio que ya tiene una reserva confirmada durante el periodo solicitado,
cuando se consulte su disponibilidad,
entonces dicho espacio no debe aparecer como disponible.

#### CA-003 — Rango inválido

Dado un rango de fecha y hora inválido,
cuando se realice la consulta,
entonces el sistema debe rechazar la solicitud indicando el error.

#### CA-004 — Consulta sin efectos colaterales

Cuando se consulte la disponibilidad,
entonces la consulta no debe modificar ninguna reserva existente.

---

## HU-RES-002 — Consultar la disponibilidad de un espacio seleccionado

### Historia

Como usuario, quiero consultar la disponibilidad de un espacio específico, para saber si puedo reservarlo en una fecha y horario determinados.

### Bounded Context

Gestión de Espacios.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Espacio disponible

Dado un espacio existente y un periodo válido sin reservas confirmadas superpuestas,
cuando el usuario consulte ese espacio,
entonces el sistema debe indicar que el espacio está disponible.

#### CA-002 — Espacio no disponible

Dado un espacio con una reserva confirmada que se superpone con el periodo solicitado,
cuando el usuario consulte ese espacio,
entonces el sistema debe indicar que no está disponible.

#### CA-003 — Espacio inexistente

Dado un identificador de espacio que no existe,
cuando se consulte su disponibilidad,
entonces el sistema debe rechazar la solicitud indicando que el espacio no fue encontrado.

#### CA-004 — Periodo bloqueado

Dado un espacio con un periodo bloqueado por el administrador,
cuando se consulte ese periodo,
entonces el sistema debe indicar que el espacio no está disponible.

---

## HU-RES-003 — Crear una reserva para un espacio disponible

### Historia

Como usuario, quiero crear una reserva de un espacio disponible, para asegurar el uso del espacio en una fecha y horario específicos.

### Bounded Context

Gestión de Reservas.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Reserva válida

Dado un usuario autenticado, un espacio existente y un periodo disponible,
cuando el usuario cree una reserva,
entonces el sistema debe crear la reserva en un estado inicial válido y devolver su identificador.

#### CA-002 — Conflicto de disponibilidad

Dado un espacio que no está disponible en el periodo solicitado,
cuando el usuario intente crear la reserva,
entonces el sistema debe rechazar la operación y no crear la reserva.

#### CA-003 — Superposición de reservas confirmadas

Dado que ya existe una reserva confirmada para el mismo espacio y un periodo superpuesto,
cuando otro usuario intente crear una reserva,
entonces el sistema debe rechazarla para cumplir BR-001.

#### CA-004 — Datos incompletos

Dado que faltan datos obligatorios de la reserva,
cuando se envíe la solicitud,
entonces el sistema debe rechazarla indicando los campos inválidos.

---

## HU-RES-004 — Consultar el estado de una reserva

### Historia

Como usuario, quiero consultar el estado de una reserva, para conocer en qué punto del ciclo de vida se encuentra.

### Bounded Context

Gestión de Reservas.

### Prioridad

Alta.

### Dependencia

HU-RES-004 puede implementarse después de HU-RES-003.

### Criterios de aceptación

#### CA-001 — Consulta de reserva existente

Dada una reserva existente,
cuando el usuario autorizado consulte su estado,
entonces el sistema debe devolver el identificador, el espacio, el periodo y el estado actual.

#### CA-002 — Reserva inexistente

Dado un identificador de reserva que no existe,
cuando se consulte,
entonces el sistema debe indicar que la reserva no fue encontrada.

#### CA-003 — Acceso no autorizado

Dado un usuario que no es el dueño de la reserva ni un administrador autorizado,
cuando intente consultar la reserva,
entonces el sistema debe rechazar el acceso.

#### CA-004 — Estados visibles

Cuando se consulte una reserva,
entonces el estado devuelto debe corresponder a uno de los estados definidos: PENDIENTE, PAGO_PENDIENTE, CONFIRMADA o CANCELADA.

---

## HU-RES-005 — Procesar el pago asociado a una reserva

### Historia

Como usuario, quiero procesar el pago asociado a una reserva, para poder confirmar el uso del espacio.

### Bounded Context

Gestión de Pagos.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Pago iniciado

Dada una reserva en estado que permite pago,
cuando el usuario solicite el pago,
entonces el sistema debe crear una operación de pago y asociarla a la reserva.

#### CA-002 — Idempotencia

Dada una misma operación de pago enviada más de una vez,
cuando se procese,
entonces el sistema no debe generar múltiples efectos financieros.

#### CA-003 — Resultado de pago

Cuando el procesamiento del pago finalice,
entonces el Servicio de Pagos debe publicar un evento de resultado, por ejemplo PaymentConfirmed o PaymentFailed.

#### CA-004 — Reserva inexistente o no pagable

Dada una reserva inexistente o que no admite pago,
cuando se solicite el pago,
entonces el sistema debe rechazar la operación.

---

## HU-RES-006 — Confirmar una reserva después de un pago exitoso

### Historia

Como usuario, quiero que mi reserva se confirme después de un pago exitoso, para quedar con el espacio asignado de forma definitiva.

### Bounded Context

Gestión de Reservas.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Confirmación por pago exitoso

Dada una reserva pendiente de pago,
cuando se reciba el evento PaymentConfirmed,
entonces la reserva debe pasar al estado CONFIRMADA.

#### CA-002 — Pago fallido

Dada una reserva pendiente de pago,
cuando se reciba el evento PaymentFailed,
entonces la reserva no debe pasar a CONFIRMADA.

#### CA-003 — Evento duplicado

Dado que PaymentConfirmed se recibe dos veces,
cuando se procese el segundo evento,
entonces el sistema no debe generar una segunda confirmación ni efectos duplicados.

#### CA-004 — Transición inválida

Dada una reserva CANCELADA,
cuando se reciba un evento de pago,
entonces el sistema no debe confirmarla.

---

## HU-RES-007 — Recibir una notificación cuando una reserva cambie de estado

### Historia

Como usuario, quiero recibir una notificación cuando mi reserva cambie de estado, para mantenerme informado sin consultar permanentemente el sistema.

### Bounded Context

Gestión de Notificaciones.

### Prioridad

Media.

### Criterios de aceptación

#### CA-001 — Notificación por confirmación

Dada una reserva que pasa a CONFIRMADA,
cuando se publique ReservationConfirmed,
entonces el Servicio de Notificaciones debe crear y entregar una notificación al usuario.

#### CA-002 — Notificación por cancelación

Dada una reserva cancelada,
cuando se publique ReservationCancelled,
entonces el usuario debe recibir una notificación de cancelación.

#### CA-003 — Fallo de notificación

Dado un fallo en la entrega de la notificación,
cuando la reserva ya esté confirmada,
entonces la reserva no debe invalidarse.

#### CA-004 — Eventos duplicados

Dado un evento de notificación duplicado,
cuando se procese,
entonces el sistema no debe generar entregas duplicadas incontroladas.

---

## HU-RES-008 — Gestionar espacios reservables y sus reglas de disponibilidad

### Historia

Como administrador de espacios, quiero registrar y actualizar espacios y sus reglas de disponibilidad, para controlar qué recursos pueden reservarse y en qué periodos.

### Bounded Context

Gestión de Espacios.

### Prioridad

Alta.

### Criterios de aceptación

#### CA-001 — Registrar espacio

Dado un administrador autenticado y datos válidos del espacio,
cuando registre un espacio,
entonces el sistema debe persistirlo y dejarlo consultable.

#### CA-002 — Actualizar espacio

Dado un espacio existente,
cuando el administrador actualice su información,
entonces los cambios deben quedar persistidos sin afectar reservas confirmadas de forma inconsistente.

#### CA-003 — Bloquear periodo

Dado un espacio existente,
cuando el administrador bloquee un periodo,
entonces ese periodo no debe aparecer como disponible para nuevas reservas.

#### CA-004 — Datos inválidos

Dado un registro o actualización con datos inválidos,
cuando se envíe la solicitud,
entonces el sistema debe rechazarla indicando el error.

---

## Definition of Ready

Una historia podrá entrar al Sprint cuando:

- El objetivo de negocio esté claramente definido.
- El bounded context responsable esté identificado.
- Los criterios de aceptación sean verificables.
- Las dependencias conocidas estén identificadas.
- El alcance sea suficientemente pequeño para el Sprint.
- No existan preguntas críticas sin resolver que impidan su implementación.

## Definition of Done

Una historia se considerará terminada cuando:

- La funcionalidad cumpla sus criterios de aceptación.
- El código esté implementado.
- Las pruebas correspondientes estén agregadas.
- No existan errores críticos conocidos.
- El cambio haya sido revisado mediante Pull Request.
- La documentación necesaria haya sido actualizada.
- El cambio esté integrado en la rama correspondiente.
- La evidencia de la historia esté registrada en el estado semanal.

Las historias arquitectónicas HU-ARC-001 a HU-ARC-003 y HU-ADR-001 acompañan la definición de la arquitectura y no se desarrollan como funcionalidades de producto independientes.
