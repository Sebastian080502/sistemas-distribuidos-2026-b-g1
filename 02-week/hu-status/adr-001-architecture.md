# ADR-001: Arquitectura basada en microservicios

- Estado: Aceptado
- Fecha: 2026-08-17
- Decisor: Juan Sebastian Osorio Fierro
- Ámbito: Arquitectura general de la Plataforma Distribuida de Reservas
- Historia: HU-ADR-001 — Formalizar la decisión arquitectónica

## 1. Contexto

La Plataforma Distribuida de Reservas tiene como objetivo gestionar espacios reservables, disponibilidad, reservas, pagos y notificaciones.

El proyecto se desarrolla como parte de la asignatura de Sistemas Distribuidos y debe demostrar conceptos relacionados con sistemas distribuidos, comunicación entre servicios, independencia de componentes, tolerancia a fallos, consistencia, escalabilidad y evolución arquitectónica.

El proyecto será desarrollado inicialmente por un único estudiante, por lo que la arquitectura debe mantener un alcance controlado y evitar una complejidad que no pueda ser implementada, probada y documentada durante el periodo académico disponible.

Al mismo tiempo, la arquitectura debe permitir que el sistema pueda crecer incorporando nuevas capacidades de negocio sin tener que reconstruir completamente la solución.

Durante el análisis inicial del dominio se identificaron las siguientes capacidades principales:

- Identidad y acceso.
- Gestión de espacios.
- Gestión de reservas.
- Gestión de pagos.
- Gestión de notificaciones.

Estas capacidades presentan responsabilidades diferentes y pueden evolucionar de manera independiente.

## 2. Problema

Se debe seleccionar una arquitectura que permita demostrar los principios de sistemas distribuidos sin introducir una cantidad innecesaria de complejidad.

La arquitectura debe permitir:

- Separar responsabilidades de negocio.
- Mantener límites claros entre capacidades.
- Evitar una base de datos compartida entre servicios.
- Permitir comunicación síncrona y asíncrona.
- Aislar fallos entre componentes.
- Permitir escalabilidad independiente cuando sea necesario.
- Facilitar la incorporación futura de nuevas capacidades.
- Mantener el proyecto implementable por un único desarrollador.
- Mantener una documentación arquitectónica clara.

## 3. Decisión

Se utilizará una arquitectura basada en microservicios.

Cada microservicio será responsable de una capacidad de negocio definida y tendrá un límite de responsabilidad explícito.

Los microservicios iniciales serán:

1. Servicio de Identidad.
2. Servicio de Espacios.
3. Servicio de Reservas.
4. Servicio de Pagos.
5. Servicio de Notificaciones.

Estos servicios corresponden a los bounded contexts identificados en el Context Map inicial.

La arquitectura no considera que la cantidad de microservicios sea un objetivo en sí mismo.

La incorporación de un nuevo microservicio deberá justificarse mediante una responsabilidad de negocio clara, propiedad independiente de datos y una necesidad real de evolución, escalabilidad o despliegue independiente.

## 4. Justificación de la decisión

La arquitectura de microservicios se selecciona porque el proyecto necesita demostrar características propias de los sistemas distribuidos.

Entre ellas:

- Comunicación entre procesos independientes.
- Comunicación síncrona mediante APIs.
- Comunicación asíncrona mediante eventos.
- Propiedad independiente de datos.
- Consistencia distribuida.
- Idempotencia.
- Aislamiento de fallos.
- Escalabilidad independiente.
- Despliegue independiente.
- Evolución independiente de capacidades.

La arquitectura también permite que el proyecto pueda crecer progresivamente sin convertir el sistema en un único componente difícil de mantener.

## 5. Alternativas consideradas

### 5.1 Monolito tradicional

Una primera alternativa sería implementar todo el sistema como una única aplicación.

#### Ventajas

- Menor complejidad inicial.
- Desarrollo más rápido.
- Un único proceso de ejecución.
- Una única base de datos.
- Menor complejidad de comunicación.

#### Desventajas

- No permite demostrar adecuadamente comunicación entre microservicios.
- No existe aislamiento real entre capacidades.
- Las fallas pueden afectar a toda la aplicación.
- El escalamiento debe realizarse sobre todo el sistema.
- No existe despliegue independiente por capacidad.
- No permite demostrar adecuadamente consistencia distribuida y comunicación asíncrona.

#### Decisión

Descartado.

Aunque sería más sencillo para un único desarrollador, no cumple adecuadamente el objetivo académico principal del proyecto.

### 5.2 Monolito modular

Otra alternativa sería implementar un único despliegue con módulos internos claramente separados.

#### Ventajas

- Mantiene límites de dominio.
- Menor complejidad operacional.
- Facilita las pruebas.
- Permite una evolución inicial ordenada.
- Puede ser adecuado para sistemas pequeños.

#### Desventajas

- Los módulos siguen compartiendo el mismo proceso.
- No existe independencia real de despliegue.
- No permite demostrar completamente la comunicación distribuida.
- Las fallas de infraestructura pueden afectar al proceso completo.
- La escalabilidad independiente de cada capacidad es limitada.

#### Decisión

Descartado como arquitectura final.

Algunos principios del monolito modular, como los límites claros y la separación de responsabilidades, sí serán utilizados dentro de cada microservicio.

### 5.3 Microservicios

La tercera alternativa consiste en separar las principales capacidades de negocio en servicios independientes.

#### Ventajas

- Permite independencia de despliegue.
- Permite escalabilidad independiente.
- Facilita el aislamiento de fallos.
- Permite utilizar comunicación síncrona y asíncrona.
- Cada servicio puede ser propietario de sus datos.
- Permite demostrar conceptos propios de sistemas distribuidos.
- Facilita la incorporación de nuevas capacidades.

#### Desventajas

- Mayor complejidad operacional.
- Comunicación distribuida.
- Posibilidad de fallos de red.
- Mayor dificultad de pruebas de integración.
- Necesidad de observabilidad.
- Necesidad de manejar consistencia distribuida.
- Mayor cantidad de componentes que mantener.

#### Decisión

Seleccionado.

La complejidad adicional está justificada por los objetivos académicos y arquitectónicos del proyecto.

## 6. Límites de los microservicios

Los límites iniciales se derivan de los bounded contexts.

### Servicio de Identidad

Responsabilidad:

- Usuarios.
- Autenticación.
- Autorización.
- Roles y permisos.

No será responsable de reservas, pagos o espacios.

### Servicio de Espacios

Responsabilidad:

- Espacios reservables.
- Disponibilidad.
- Periodos bloqueados.

No será responsable de crear o administrar reservas.

### Servicio de Reservas

Responsabilidad:

- Creación de reservas.
- Estados.
- Confirmación.
- Cancelación.
- Ciclo de vida de la reserva.

Será el principal coordinador del flujo de negocio.

### Servicio de Pagos

Responsabilidad:

- Operaciones de pago.
- Estados de pago.
- Identificadores de transacción.
- Eventos relacionados con pagos.

No será responsable del ciclo de vida completo de la reserva.

### Servicio de Notificaciones

Responsabilidad:

- Recepción de eventos.
- Creación de notificaciones.
- Entrega.
- Estado de entrega.

No será responsable de modificar directamente reservas o pagos.

## 7. Propiedad de los datos

Cada microservicio será propietario de sus datos.

La distribución inicial será:

| Servicio       | Datos principales                        |
| -------------- | ---------------------------------------- |
| Identidad      | Usuarios, credenciales, roles y permisos |
| Espacios       | Espacios, disponibilidad y bloqueos      |
| Reservas       | Reservas y estados                       |
| Pagos          | Pagos y estados                          |
| Notificaciones | Notificaciones y entregas                |

No se permitirá que un microservicio consulte directamente las tablas de otro microservicio.

La comunicación de información entre servicios se realizará mediante:

- APIs.
- Eventos.
- Contratos explícitos.

## 8. Estrategia de comunicación

Se utilizarán dos mecanismos principales.

### Comunicación síncrona

Se utilizará cuando el servicio consumidor necesite una respuesta inmediata.

Ejemplo:

Servicio de Reservas
|
| Consulta disponibilidad
v
Servicio de Espacios

Otro ejemplo:

Servicio de Reservas
|
| Solicitud de pago
v
Servicio de Pagos

La comunicación síncrona inicial se implementará mediante APIs REST.

### Comunicación asíncrona

Se utilizará cuando una operación pueda continuar de manera desacoplada.

Ejemplo:

Servicio de Pagos
|
| PaymentConfirmed
v
Servicio de Reservas

Otro ejemplo:

Servicio de Reservas
|
| ReservationConfirmed
v
Servicio de Notificaciones

El mecanismo de mensajería será definido en una decisión arquitectónica posterior cuando se establezcan los requisitos técnicos de infraestructura.

## 9. Consistencia

La arquitectura utilizará diferentes estrategias de consistencia.

### Operaciones críticas

Las operaciones relacionadas con:

- Disponibilidad.
- Creación de reservas.
- Confirmación de reservas.
- Estados críticos de pago.

deben garantizar que no se produzcan inconsistencias graves como:

- Doble reserva.
- Pago duplicado.
- Estados imposibles.

### Operaciones no críticas

Las notificaciones, reportes y capacidades analíticas podrán utilizar consistencia eventual.

Por ejemplo:

Una reserva puede estar confirmada aunque la notificación todavía se encuentre pendiente de entrega.

## 10. Idempotencia

Las operaciones distribuidas deberán considerar la posibilidad de recibir solicitudes o eventos duplicados.

El procesamiento de eventos importantes deberá ser idempotente cuando corresponda.

Por ejemplo:

Si el evento PaymentConfirmed es recibido dos veces, el sistema no debe generar dos confirmaciones financieras ni producir efectos duplicados.

La implementación concreta de la idempotencia será definida durante el desarrollo del servicio correspondiente.

## 11. Aislamiento de fallos

Los servicios deben minimizar el impacto de fallos entre ellos.

Por ejemplo:

- Un fallo del Servicio de Notificaciones no debe cancelar una reserva confirmada.
- Un evento de pago fallido debe poder ser procesado sin afectar la disponibilidad del Servicio de Espacios.
- Un servicio temporalmente indisponible no debe provocar acceso directo a su base de datos desde otro servicio.
- Los consumidores de eventos deben manejar eventos duplicados cuando sea necesario.

El comportamiento exacto ante fallos será documentado y probado durante la implementación.

## 12. Escalabilidad

La arquitectura permitirá escalar servicios de forma independiente.

Por ejemplo:

Si la consulta de disponibilidad recibe una carga superior a la creación de reservas, el Servicio de Espacios podrá escalarse independientemente.

De igual forma, si el procesamiento de notificaciones aumenta significativamente, el Servicio de Notificaciones podrá aumentar sus instancias sin escalar todos los demás servicios.

La escalabilidad no implica agregar microservicios innecesarios.

Cada nuevo servicio debe estar justificado por una capacidad de negocio o una necesidad técnica real.

## 13. Extensibilidad

La arquitectura permitirá agregar nuevas capacidades mediante nuevos servicios cuando sea necesario.

Ejemplos potenciales:

- Servicio de Reseñas.
- Servicio de Promociones.
- Servicio de Analítica.
- Servicio de Búsqueda.
- Servicio de Fidelización.
- Servicio de Auditoría.

Estos servicios no forman parte del MVP inicial.

La arquitectura actual deberá permitir incorporarlos sin modificar directamente las bases de datos de los servicios existentes.

## 14. Consecuencias positivas

La decisión proporciona los siguientes beneficios:

- Cumplimiento del enfoque de Sistemas Distribuidos.
- Límites de negocio claramente definidos.
- Independencia de despliegue.
- Escalabilidad independiente.
- Aislamiento de fallos.
- Comunicación síncrona y asíncrona.
- Propiedad independiente de datos.
- Mayor facilidad para incorporar nuevas capacidades.
- Posibilidad de demostrar patrones y problemas reales de sistemas distribuidos.
- Arquitectura preparada para evolución futura.

## 15. Consecuencias negativas

La decisión también introduce costos:

- Mayor complejidad de desarrollo.
- Mayor complejidad de despliegue.
- Necesidad de administrar varios servicios.
- Mayor complejidad de pruebas.
- Problemas de latencia de red.
- Fallos parciales.
- Necesidad de observabilidad.
- Necesidad de manejar consistencia distribuida.
- Mayor esfuerzo de documentación.

Debido a que el proyecto es desarrollado individualmente, se limitará el MVP para mantener la complejidad bajo control.

## 16. Restricciones de alcance

Para evitar que la arquitectura se vuelva inmanejable, se establecen las siguientes restricciones:

- El MVP utilizará inicialmente cinco microservicios candidatos.
- No se crearán servicios solamente para aumentar el número de microservicios.
- Las capacidades futuras permanecerán fuera del MVP.
- Cada servicio deberá tener pruebas automatizadas.
- Cada servicio deberá tener documentación mínima.
- Cada servicio deberá tener una responsabilidad claramente delimitada.
- No se utilizarán bases de datos compartidas.
- Las decisiones arquitectónicas importantes deberán documentarse mediante ADR.

## 17. Criterios para agregar un nuevo microservicio

Un nuevo microservicio solamente podrá incorporarse cuando cumpla una o más condiciones justificadas:

1. Representa una capacidad de negocio independiente.
2. Posee un límite de dominio claro.
3. Requiere escalabilidad independiente.
4. Requiere despliegue independiente.
5. Tiene datos que deben ser administrados por separado.
6. Reduce significativamente el acoplamiento del sistema.
7. Permite aislar una responsabilidad con un ciclo de evolución diferente.

La cantidad de microservicios no será utilizada como métrica de éxito.

## 18. Estado de la decisión

Aceptado.

La arquitectura basada en microservicios será utilizada como dirección arquitectónica para la Plataforma Distribuida de Reservas.

Los límites definidos en este ADR corresponden al diseño inicial y podrán ser refinados mediante futuros ADR cuando aparezca nueva información relevante.

## 19. Regla de inmutabilidad

Este ADR, una vez aceptado, **no debe modificarse**.

Cualquier cambio posterior a esta decisión arquitectónica deberá documentarse en un nuevo archivo (`adr-002-*.md`) que:

- Referencie explícitamente ADR-001 como el registro que se sustituye.
- Indique qué cambió y por qué.
- Documente el contexto, la decisión, las alternativas y las consecuencias de la nueva decisión.

## 20. Relación con otros documentos

Esta decisión está relacionada con:

- [`context-map.md`](./context-map.md)
- [`user-stories.md`](./user-stories.md)
- [`../../01-week/hu-status/prd.md`](../../01-week/hu-status/prd.md)
- [`README.md`](./README.md)
- [`README.es.md`](./README.es.md)

El Context Map define las relaciones entre los bounded contexts.

Este ADR documenta la decisión de utilizar dichos límites como base para la arquitectura de microservicios.
