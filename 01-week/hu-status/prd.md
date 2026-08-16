# Brief del Producto — Plataforma Distribuida de Reservas

## 1. Contexto

La Plataforma Distribuida de Reservas es un proyecto para la asignatura de Sistemas Distribuidos.

El sistema permitirá gestionar y reservar espacios físicos durante fechas y periodos de tiempo específicos.

El proyecto será desarrollado como un sistema distribuido basado en microservicios. La arquitectura estará orientada a soportar límites de servicio independientes, propiedad independiente de los datos, comunicación distribuida, escalabilidad y evolución futura.

## 2. Problema

Las organizaciones que administran espacios e instalaciones necesitan coordinar la disponibilidad, las reservas, los pagos y las notificaciones.

Cuando estos procesos no cuentan con una gestión centralizada y confiable, pueden presentarse los siguientes problemas:

- Dos usuarios pueden intentar reservar el mismo espacio durante periodos que se superponen.
- Los usuarios pueden no disponer de información confiable sobre la disponibilidad de los espacios.
- Los procesos de reserva y pago pueden quedar inconsistentes.
- Las solicitudes o mensajes duplicados pueden producir operaciones financieras duplicadas.
- Los usuarios pueden no recibir información oportuna sobre los cambios de sus reservas.
- Las fallas en un proceso pueden afectar otras partes del flujo de reserva.
- Los estados de las reservas pueden no contar con suficiente trazabilidad.

El proyecto busca solucionar estos problemas mediante una arquitectura distribuida basada en capacidades de negocio independientes.

## 3. Necesidades y problemas

El sistema inicial debe atender las siguientes necesidades:

- Consultar espacios disponibles.
- Consultar la disponibilidad de un espacio específico.
- Crear reservas.
- Consultar el estado de una reserva.
- Procesar pagos asociados a las reservas.
- Confirmar reservas después de un pago exitoso.
- Notificar a los usuarios sobre eventos relacionados con sus reservas.
- Administrar espacios y sus reglas de disponibilidad.
- Evitar reservas conflictivas.
- Evitar el procesamiento duplicado de operaciones críticas.

## 4. Objetivo del producto

El objetivo principal es proporcionar una plataforma confiable y extensible para gestionar y reservar espacios.

El sistema debe mantener la consistencia en las operaciones críticas y permitir la evolución y despliegue independiente de sus capacidades de negocio.

La arquitectura también debe permitir incorporar nuevas capacidades de negocio sin requerir un rediseño completo del sistema existente.

## 5. Usuarios objetivo

### Usuario

Persona que desea encontrar y reservar un espacio.

El usuario podrá:

- Consultar espacios disponibles.
- Consultar disponibilidad.
- Crear una reserva.
- Consultar el estado de una reserva.
- Realizar el pago asociado.
- Recibir notificaciones sobre cambios en sus reservas.

### Administrador de espacios

Persona responsable de administrar los espacios disponibles para reserva.

El administrador podrá:

- Registrar espacios.
- Actualizar información de los espacios.
- Configurar disponibilidad.
- Bloquear periodos no disponibles.
- Consultar las reservas asociadas a los espacios.

## 6. Proceso principal

El flujo principal del negocio es:

Usuario
|
v
Consultar espacios
|
v
Consultar disponibilidad
|
v
Crear reserva
|
v
Procesar pago
|
v
Confirmar reserva
|
v
Notificar al usuario

El ciclo de vida de la reserva será refinado durante el análisis del dominio.

Estados iniciales de la reserva:

- PENDIENTE
- PAGO_PENDIENTE
- CONFIRMADA
- CANCELADA

## 7. Alcance del MVP

El primer MVP se enfocará en el proceso principal de reservas.

### Incluido

- Identificación y acceso de usuarios.
- Gestión de espacios.
- Consulta de disponibilidad de espacios.
- Creación de reservas.
- Consulta del estado de las reservas.
- Abstracción del procesamiento de pagos.
- Confirmación de reservas.
- Cancelación de reservas.
- Notificaciones a usuarios.
- Propiedad independiente de los datos.
- Comunicación REST.
- Mensajería asíncrona para determinados flujos.
- Idempotencia para operaciones críticas.
- Pruebas automatizadas.
- Ejecución local mediante contenedores.
- Documentación técnica.

### Extensiones futuras

La arquitectura podrá soportar posteriormente:

- Reseñas y calificaciones.
- Promociones y descuentos.
- Programa de fidelización.
- Búsqueda avanzada.
- Analítica.
- Reportes.
- Auditoría.
- Canales adicionales de notificación.
- Proveedores de pago adicionales.
- Caché.
- Observabilidad avanzada.

Estas capacidades no forman parte del MVP inicial.

## 8. Fuera de alcance

Las siguientes capacidades están fuera del alcance inicial:

- Seguimiento GPS en tiempo real.
- Geolocalización compleja.
- Funcionalidad de red social.
- Chat en tiempo real.
- Funcionalidad de marketplace.
- Precios dinámicos.
- Recomendaciones mediante inteligencia artificial.
- Funcionalidad completa de ERP.
- Contabilidad compleja.
- Integración bancaria completa.
- Múltiples aplicaciones móviles.
- Kubernetes como requisito inicial.
- Una cantidad artificialmente grande de microservicios.

El proyecto priorizará un MVP controlado que pueda ser desarrollado durante el periodo académico disponible.

## 9. Reglas de negocio principales

### BR-001 — Disponibilidad del espacio

Un espacio no puede tener dos reservas confirmadas que se superpongan en fecha y hora.

### BR-002 — Creación de reserva

Una reserva solamente puede crearse cuando el espacio y el periodo solicitado se encuentran disponibles.

### BR-003 — Confirmación del pago

Una reserva no puede pasar al estado confirmado hasta que el pago asociado alcance el estado exitoso requerido.

### BR-004 — Idempotencia del pago

Procesar varias veces la misma operación de pago no debe generar múltiples efectos financieros.

### BR-005 — Ciclo de vida de la reserva

Los cambios de estado de una reserva deben respetar las transiciones válidas definidas por el dominio.

### BR-006 — Independencia de las notificaciones

Una falla en el envío de una notificación no debe invalidar una reserva que ya fue confirmada.

### BR-007 — Propiedad de los datos

Cada servicio debe ser propietario de los datos asociados a su capacidad de negocio.

### BR-008 — Independencia de los servicios

Un servicio no debe acceder directamente a la base de datos de otro servicio.

## 10. Requisitos de consistencia

El sistema no utilizará el mismo modelo de consistencia para todas las operaciones.

### Consistencia fuerte

Las siguientes operaciones requieren consistencia fuerte o un mecanismo equivalente de coordinación:

- Creación de una reserva.
- Validación de disponibilidad durante la reserva.
- Confirmación de una reserva.
- Aplicación del resultado de un pago.

El objetivo principal es evitar reservas conflictivas y efectos financieros duplicados.

### Consistencia eventual

Las siguientes operaciones pueden utilizar consistencia eventual:

- Notificaciones.
- Analítica.
- Reportes futuros.
- Métricas.
- Modelos de lectura no críticos.

## 11. Semántica de entrega

### Consultas de usuario

Las consultas dirigidas al usuario utilizarán normalmente comunicación síncrona mediante solicitud y respuesta.

### Operaciones críticas

Los comandos y eventos distribuidos críticos podrán utilizar entrega at-least-once.

Debido a que la entrega at-least-once puede producir mensajes duplicados, los consumidores deberán implementar procesamiento idempotente cuando sea necesario.

Ejemplos:

- PaymentConfirmed
- ReservationConfirmed
- ReservationCancelled

### Notificaciones

Los eventos de notificación podrán utilizar entrega at-least-once, ya que los mensajes duplicados podrán manejarse mediante idempotencia o mecanismos de deduplicación.

## 12. Dirección arquitectónica inicial

El proyecto utilizará microservicios como dirección arquitectónica desde el inicio.

Sin embargo, la cantidad de servicios no se definirá arbitrariamente.

Los límites de los servicios se derivarán de:

- Capacidades de negocio.
- Bounded Contexts.
- Reglas de negocio independientes.
- Propiedad de los datos.
- Contratos de comunicación.
- Necesidades justificadas de despliegue o escalabilidad independiente.

Las capacidades candidatas iniciales son:

- Identidad y acceso.
- Gestión de espacios.
- Gestión de reservas.
- Gestión de pagos.
- Gestión de notificaciones.

Estas capacidades serán validadas durante el análisis de Domain-Driven Design de la Semana 02.

## 13. Escalabilidad

La escalabilidad es un requisito arquitectónico importante.

El proyecto debe poder incorporar nuevos servicios cuando una nueva capacidad de negocio justifique un límite independiente.

Algunos posibles servicios futuros son:

- Servicio de reseñas.
- Servicio de analítica.
- Servicio de promociones.
- Servicio de búsqueda.
- Servicio de fidelización.
- Servicio de auditoría.

Los nuevos servicios solamente se incorporarán cuando exista una razón de negocio o arquitectónica que lo justifique.

El objetivo no es maximizar la cantidad de microservicios, sino mantener límites claros y permitir la evolución independiente del sistema.

## 14. Dirección tecnológica

La implementación inicial utilizará un entorno de desarrollo basado en contenedores.

La arquitectura contemplará:

- Docker.
- APIs REST.
- Message Broker.
- Bases de datos independientes.
- Pruebas automatizadas.

Las tecnologías y frameworks específicos serán seleccionados durante las etapas de diseño arquitectónico e implementación.

## 15. Documentación

La documentación se desarrollará junto con el sistema.

El proyecto incluirá progresivamente:

- Brief del producto.
- Modelo de dominio.
- Bounded Contexts.
- Context Map.
- Diagramas C4.
- Architecture Decision Records.
- Contratos de API.
- Contratos de eventos.
- Documentación de pruebas.
- Documentación de seguridad.
- Documentación de despliegue.

La documentación se mantendrá actualizada a medida que evolucione la arquitectura.

## 16. Gestión del proyecto

El proyecto utilizará un enfoque Scrumban.

### Elementos de Scrum

- Sprints semanales.
- Sprint Goal.
- Sprint Planning.
- Sprint Review.
- Sprint Retrospective.

### Elementos de Kanban

El trabajo se gestionará mediante GitHub Projects utilizando el siguiente flujo:

BACKLOG
READY
IN PROGRESS
REVIEW
TESTING
DONE

Las actividades actuales de la Semana 01 y Semana 02 corresponden a la etapa de preparación y recuperación del proyecto.

El primer Sprint formal de desarrollo comenzará después de completar estas actividades.

## 17. Criterios de éxito

El proyecto se considerará exitoso cuando:

- El sistema esté implementado utilizando microservicios reales.
- Cada servicio tenga una responsabilidad de negocio claramente definida.
- Cada servicio sea propietario de sus datos.
- Los servicios se comuniquen mediante contratos explícitos.
- Se demuestre comunicación síncrona y asíncrona cuando sea apropiado.
- Las operaciones críticas manejen correctamente el procesamiento duplicado.
- Las fallas distribuidas sean manejadas adecuadamente.
- El sistema pueda ejecutarse localmente mediante contenedores.
- Las pruebas automatizadas cubran el comportamiento crítico del negocio.
- La arquitectura pueda evolucionar incorporando nuevas capacidades de negocio.
- La documentación permanezca sincronizada con la implementación.
- El proyecto pueda ser demostrado y defendido técnicamente.

## 18. Preguntas abiertas

Las siguientes preguntas serán resueltas durante el análisis del dominio y la arquitectura:

- ¿Qué categorías específicas de espacios serán soportadas?
- ¿Las reservas se pagan inmediatamente o pueden permanecer pendientes?
- ¿Qué estados de pago serán necesarios?
- ¿Qué reglas de cancelación aplican?
- ¿Qué sucede cuando un pago falla?
- ¿Cuánto tiempo puede permanecer pendiente una reserva?
- ¿Qué canales de notificación serán soportados?
- ¿Un espacio será administrado por un único administrador o por varios?
- ¿Qué mecanismo de autenticación se utilizará?
- ¿Qué requisitos adicionales de infraestructura serán establecidos durante el proyecto?

Cada pregunta resuelta se convertirá, cuando corresponda, en una regla de negocio, criterio de aceptación o decisión arquitectónica.

## 19. Glosario de negocio

| Término               | Definición                                                                                               |
| --------------------- | -------------------------------------------------------------------------------------------------------- |
| Espacio               | Instalación o recurso físico que puede ser reservado durante un periodo determinado.                     |
| Disponibilidad        | Condición que determina si un espacio puede ser reservado para una fecha y hora determinadas.            |
| Reserva               | Solicitud o asignación confirmada de un espacio durante un periodo específico.                           |
| Estado de reserva     | Estado actual de una reserva dentro de su ciclo de vida.                                                 |
| Pago                  | Operación financiera asociada a una reserva.                                                             |
| Notificación          | Mensaje que informa al usuario sobre un evento o cambio relevante de una reserva.                        |
| Administrador         | Usuario responsable de gestionar espacios y su disponibilidad.                                           |
| Usuario               | Persona que utiliza la plataforma para consultar y reservar espacios.                                    |
| Bounded Context       | Límite del dominio dentro del cual un modelo y lenguaje de negocio específico son válidos.               |
| Microservicio         | Servicio desplegable de forma independiente que es responsable de una capacidad de negocio definida.     |
| Evento de dominio     | Registro de un evento importante del negocio que puede ser consumido por otros servicios.                |
| Idempotencia          | Propiedad que permite procesar una operación varias veces sin producir efectos de negocio duplicados.    |
| Consistencia eventual | Modelo de consistencia en el que los datos distribuidos pueden diferir temporalmente antes de converger. |

## 20. Principio arquitectónico inicial

El proyecto seguirá el principio:

> Separar servicios por una razón, no por moda.

La arquitectura priorizará límites de negocio significativos, propiedad independiente de los datos, contratos explícitos, escalabilidad justificada y despliegue independiente.

El objetivo no es maximizar la cantidad de microservicios, sino construir una arquitectura distribuida comprensible, comprobable, escalable y defendible.
