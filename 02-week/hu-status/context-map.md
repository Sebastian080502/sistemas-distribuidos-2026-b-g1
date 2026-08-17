# Mapa de Contextos — Plataforma Distribuida de Reservas

## 1. Propósito

Este documento define los bounded contexts iniciales y las relaciones entre ellos para la Plataforma Distribuida de Reservas.

El Context Map se utiliza para establecer límites claros del dominio antes de implementar los microservicios.

La arquitectura inicial se basa en capacidades de negocio y no en capas técnicas.

El objetivo es evitar una fragmentación innecesaria de servicios y garantizar que cada posible microservicio tenga una responsabilidad clara y sea propietario de sus propios datos.

## 2. Bounded Contexts

El análisis inicial del dominio identifica los siguientes bounded contexts:

| Bounded Context           | Responsabilidad principal                                                  | Propiedad de datos                                       |
| ------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------- |
| Identidad y Acceso        | Gestionar la identidad, autenticación y autorización de los usuarios.      | Usuarios, credenciales, roles y permisos                 |
| Gestión de Espacios       | Gestionar los espacios disponibles para reserva y su disponibilidad.       | Espacios, reglas de disponibilidad y periodos bloqueados |
| Gestión de Reservas       | Gestionar el ciclo de vida de las reservas.                                | Reservas y estados de las reservas                       |
| Gestión de Pagos          | Gestionar los pagos asociados a las reservas.                              | Pagos, estados de pago y transacciones                   |
| Gestión de Notificaciones | Gestionar las notificaciones generadas por eventos relevantes del negocio. | Notificaciones, estado de entrega y preferencias         |

## 3. Contexto de Identidad y Acceso

### Responsabilidad

El contexto de Identidad y Acceso es responsable de gestionar la identidad y las reglas de acceso de los usuarios que interactúan con la plataforma.

Sus principales responsabilidades son:

- Registrar usuarios.
- Autenticar usuarios.
- Autorizar operaciones.
- Gestionar roles.
- Gestionar permisos.
- Validar el acceso a los recursos.

### Propiedad de datos

El contexto es propietario de:

- Usuarios.
- Credenciales.
- Roles.
- Permisos.

Los demás servicios no deben acceder directamente a estos datos.

### Interacciones principales

Identidad y Acceso proporciona información de identidad y validaciones de autorización a los demás servicios cuando sea necesario.

## 4. Contexto de Gestión de Espacios

### Responsabilidad

El contexto de Gestión de Espacios es responsable de administrar los espacios físicos que pueden ser reservados y las reglas relacionadas con su disponibilidad.

Sus responsabilidades son:

- Registrar espacios.
- Actualizar información de los espacios.
- Gestionar disponibilidad.
- Definir periodos no disponibles.
- Validar si un espacio puede ser solicitado durante un periodo determinado.

### Propiedad de datos

El contexto es propietario de:

- Espacios.
- Reglas de disponibilidad.
- Periodos bloqueados.

### Interacciones principales

Gestión de Reservas consulta a Gestión de Espacios cuando necesita validar la disponibilidad de un espacio.

## 5. Contexto de Gestión de Reservas

### Responsabilidad

El contexto de Gestión de Reservas es responsable del ciclo de vida de las reservas.

Sus responsabilidades son:

- Crear reservas.
- Validar las condiciones de una reserva.
- Mantener los estados de las reservas.
- Confirmar reservas.
- Cancelar reservas.
- Publicar eventos relacionados con las reservas.

### Propiedad de datos

El contexto es propietario de:

- Reservas.
- Estados de las reservas.
- Información relacionada con el ciclo de vida de las reservas.

### Interacciones principales

Gestión de Reservas se comunica con:

- Identidad y Acceso para identificar y autorizar usuarios.
- Gestión de Espacios para validar disponibilidad.
- Gestión de Pagos para coordinar el pago.
- Gestión de Notificaciones mediante eventos de reserva.

Gestión de Reservas representa el contexto de negocio central de la plataforma inicial porque coordina el flujo principal de las reservas.

## 6. Contexto de Gestión de Pagos

### Responsabilidad

El contexto de Gestión de Pagos es responsable de administrar los pagos asociados a las reservas.

Sus responsabilidades son:

- Crear operaciones de pago.
- Consultar el estado de un pago.
- Procesar resultados de pago.
- Evitar efectos financieros duplicados.
- Publicar eventos relacionados con pagos.

### Propiedad de datos

El contexto es propietario de:

- Pagos.
- Estados de pago.
- Identificadores de transacciones.
- Información relacionada con el procesamiento de pagos.

### Interacciones principales

Gestión de Pagos recibe solicitudes de pago asociadas a las reservas.

Después del procesamiento puede publicar eventos como:

- PaymentConfirmed.
- PaymentFailed.
- PaymentCancelled.

Gestión de Reservas consume estos eventos para actualizar el ciclo de vida de la reserva.

## 7. Contexto de Gestión de Notificaciones

### Responsabilidad

El contexto de Gestión de Notificaciones es responsable de entregar las notificaciones generadas por eventos relevantes del negocio.

Sus responsabilidades son:

- Recibir eventos de notificación.
- Crear registros de notificación.
- Entregar notificaciones.
- Consultar el estado de entrega.
- Gestionar fallos de entrega.

### Propiedad de datos

El contexto es propietario de:

- Notificaciones.
- Estado de entrega.
- Preferencias de notificación.

### Interacciones principales

Gestión de Notificaciones consume eventos como:

- ReservationCreated.
- ReservationConfirmed.
- ReservationCancelled.
- PaymentFailed.

La entrega de notificaciones no forma parte de la transacción principal de la reserva.

Por lo tanto, un fallo en las notificaciones no debe invalidar una reserva que ya fue confirmada.

## 8. Relaciones entre Contextos

Las relaciones iniciales entre los bounded contexts son:

| Contexto origen     | Contexto destino          | Relación                                            | Comunicación |
| ------------------- | ------------------------- | --------------------------------------------------- | ------------ |
| Identidad y Acceso  | Gestión de Reservas       | Proporciona información de identidad y autorización | Síncrona     |
| Gestión de Espacios | Gestión de Reservas       | Proporciona información de disponibilidad           | Síncrona     |
| Gestión de Reservas | Gestión de Pagos          | Solicita el procesamiento del pago de una reserva   | Síncrona     |
| Gestión de Pagos    | Gestión de Reservas       | Publica el resultado del pago                       | Asíncrona    |
| Gestión de Reservas | Gestión de Notificaciones | Publica eventos relacionados con las reservas       | Asíncrona    |
| Gestión de Pagos    | Gestión de Notificaciones | Publica eventos relacionados con los pagos          | Asíncrona    |

## 9. Context Map inicial

```text
                         +-------------------------+
                         |  Identidad y Acceso     |
                         +------------+------------+
                                      |
                                      | Identidad /
                                      | Autorización
                                      | Síncrona
                                      v
+----------------------+      +-------------------------+
| Gestión de Espacios  |----->| Gestión de Reservas     |
+----------------------+      +-----------+-------------+
        |                                 |
        | Disponibilidad                  | Solicitud de pago
        | Síncrona                        | Síncrona
        |                                 v
        |                       +-------------------------+
        +---------------------->| Gestión de Pagos        |
                                +-----------+-------------+
                                            |
                                            | Eventos de pago
                                            | Asíncrona
                                            v
                                +-------------------------+
                                | Gestión de              |
                                | Notificaciones          |
                                +-------------------------+

Gestión de Reservas
        |
        | Eventos de reserva
        | Asíncrona
        v
Gestión de Notificaciones
```

## 10. Estrategia de comunicación

La arquitectura utilizará comunicación síncrona y asíncrona.

### Comunicación síncrona

La comunicación síncrona se utilizará cuando el contexto que realiza la solicitud necesite una respuesta inmediata.

Ejemplos iniciales:

- Gestión de Reservas consulta la disponibilidad de un espacio.
- Gestión de Reservas solicita la creación de un pago.
- Los servicios validan información de identidad o autorización cuando sea necesario.

La comunicación síncrona inicial utilizará contratos basados en APIs REST.

### Comunicación asíncrona

La comunicación asíncrona se utilizará cuando el emisor no necesite una respuesta inmediata y el proceso pueda desacoplarse.

Los eventos iniciales incluyen:

- PaymentConfirmed.
- PaymentFailed.
- ReservationCreated.
- ReservationConfirmed.
- ReservationCancelled.

El mecanismo específico de mensajería será seleccionado durante la etapa de implementación y diseño detallado de la arquitectura.

## 11. Propiedad de los datos

Cada bounded context es propietario de sus datos de negocio.

La propiedad inicial es:

| Bounded Context           | Datos propios                                            |
| ------------------------- | -------------------------------------------------------- |
| Identidad y Acceso        | Usuarios, credenciales, roles y permisos                 |
| Gestión de Espacios       | Espacios, reglas de disponibilidad y periodos bloqueados |
| Gestión de Reservas       | Reservas y estados de las reservas                       |
| Gestión de Pagos          | Pagos y estados de pago                                  |
| Gestión de Notificaciones | Notificaciones y estado de entrega                       |

Ningún bounded context puede leer o modificar directamente la base de datos de otro bounded context.

La información entre contextos debe obtenerse mediante APIs explícitas o eventos asíncronos.

## 12. Estrategia de consistencia

La plataforma utilizará diferentes modelos de consistencia dependiendo de la operación de negocio.

### Consistencia fuerte

Se requiere consistencia fuerte o un mecanismo equivalente de coordinación para:

- Crear una reserva.
- Validar disponibilidad.
- Confirmar una reserva.
- Aplicar cambios críticos en el estado de un pago.

El objetivo es evitar:

- Reservas duplicadas.
- Efectos financieros duplicados.
- Estados de reserva inválidos.

### Consistencia eventual

La consistencia eventual es aceptable para:

- Notificaciones.
- Analítica.
- Reportes.
- Modelos de lectura no críticos.

Por ejemplo, una reserva puede quedar confirmada antes de que el servicio de notificaciones termine de entregar la notificación de confirmación.

## 13. Aislamiento de fallos

Los bounded contexts están diseñados para reducir el impacto de las fallas entre servicios.

Por ejemplo:

- Un fallo en las notificaciones no debe cancelar una reserva confirmada.
- Una interrupción temporal del servicio de notificaciones no debe impedir que Gestión de Reservas mantenga el estado de la reserva.
- Los fallos de pago deben representarse explícitamente mediante eventos de pago.
- Los eventos duplicados deben procesarse de manera idempotente cuando sea necesario.

Esta separación permite que las capacidades individuales evolucionen y fallen de manera independiente.

## 14. Candidatos iniciales a microservicios

Los bounded contexts iniciales son candidatos para convertirse en microservicios independientes:

1. Servicio de Identidad.
2. Servicio de Espacios.
3. Servicio de Reservas.
4. Servicio de Pagos.
5. Servicio de Notificaciones.

Los límites definitivos de implementación serán confirmados después de revisar el modelo de dominio y las decisiones arquitectónicas.

El proyecto no incorporará microservicios adicionales únicamente para aumentar la cantidad de servicios.

Un nuevo servicio deberá tener:

- Una responsabilidad de negocio significativa.
- Un límite de dominio claro.
- Propiedad independiente de datos.
- Un contrato de comunicación definido.
- Una necesidad justificada de escalabilidad o despliegue independiente.

## 15. Extensibilidad futura

La arquitectura permitirá incorporar nuevos bounded contexts cuando exista una justificación.

Algunos posibles contextos futuros son:

- Reseñas.
- Promociones.
- Búsqueda.
- Fidelización.
- Analítica.
- Auditoría.

Estas capacidades se encuentran fuera del MVP inicial.

Su incorporación futura no deberá requerir acceso directo a las bases de datos de los servicios existentes.

## 16. Principios arquitectónicos

El Context Map sigue los siguientes principios:

1. Las capacidades de negocio definen los límites de los servicios.
2. Cada servicio es propietario de sus datos.
3. Los servicios se comunican mediante contratos explícitos.
4. Las bases de datos no se comparten entre servicios.
5. La comunicación síncrona se utiliza cuando se necesita una respuesta inmediata.
6. La comunicación asíncrona se utiliza cuando los procesos pueden desacoplarse.
7. Las operaciones críticas deben considerar la idempotencia.
8. Los fallos de servicios no críticos no deben invalidar transacciones de negocio ya completadas.
9. Los nuevos servicios deben tener un límite justificado.
10. La cantidad de servicios debe mantenerse proporcional a las necesidades reales del proyecto.

## 17. Conclusión inicial

El análisis inicial del dominio identifica cinco bounded contexts que son candidatos adecuados para convertirse en microservicios independientes.

El contexto de Gestión de Reservas representa la capacidad central del negocio de la plataforma inicial, coordinando disponibilidad, reservas, pagos y eventos relacionados con las reservas.

El Context Map establece los límites iniciales y las relaciones de comunicación que serán utilizados como entrada para la decisión arquitectónica documentada en [`adr-001-architecture.md`](./adr-001-architecture.md).

La arquitectura podrá ser refinada durante la implementación a medida que se obtenga nuevo conocimiento sobre el dominio.
