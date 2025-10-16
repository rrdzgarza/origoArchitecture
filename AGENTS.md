# AGENTS.md: Guía para el Desarrollo del Dominio en Go

## 1. Descripción General del Proyecto

Este repositorio centraliza los modelos de dominio y los puertos para un ecosistema de aplicaciones desacopladas (API, ERP, LMS, etc.) basadas en una **arquitectura hexagonal**. El objetivo es tener una única fuente de verdad para la lógica de negocio principal y sus interfaces. Este proyecto está escrito en **Go**.

Los paquetes en el directorio `domains` (`authentication`, `lessors`, `tenants`, etc.) representan los dominios de negocio principales. El paquete `ports` contiene las interfaces para interactuar con estos dominios.

## 2. Estructura y Arquitectura Hexagonal en Go

Este proyecto sigue una estructura de paquetes donde cada paquete en el directorio `domains` es un dominio de negocio importable.

-   **Paquetes de Dominio (`domains/`)**: Aquí reside el **núcleo del dominio (el hexágono)**. Cada paquete (ej. `domains/authentication`) contiene la lógica de negocio pura, implementada con `structs` e `interfaces` de Go.
-   **Paquete `domains/shared/`**: Este paquete contiene tipos de datos transversales utilizados en múltiples dominios, como el tipo `shared.UUID` para identificadores únicos.
-   **Puertos (Interfaces)**: Los puertos se definen como `interfaces` de Go dentro del paquete `ports`. Estos puertos definen cómo el núcleo de la aplicación se comunica con el mundo exterior (e.g., bases de datos, APIs web, colas de mensajes).
-   **Adaptadores**: La implementación de los puertos debe residir en los proyectos consumidores (ej. la API), típicamente en un directorio como `internal/infrastructure/postgres`.

## 3. Modelo de Datos Multi-Tenant

La arquitectura es **multi-lessor y multi-tenant**.

-   **`LessorID`**: Identifica al proveedor del servicio.
-   **`TenantID`**: Identifica a la cuenta del cliente final.

Casi todas las entidades principales (como `User`, `Log`, `Email`) contienen `LessorID` y `TenantID`. **Es fundamental que cualquier nueva entidad que almacene datos de cliente incluya estos identificadores**.

## 4. Instrucciones para el Desarrollo

### Interfaces de Repositorio

Las interfaces de repositorio se encuentran en el archivo `ports/repositories.go`. Estas interfaces definen los métodos para la persistencia de datos (CRUD) para cada dominio.

### Manejo de Errores

El paquete `domains/shared/errors` define un conjunto de errores personalizados que deben ser utilizados por las implementaciones de los repositorios y servicios. Esto permite un manejo de errores más robusto y consistente en toda la aplicación.

### Cómo Añadir una Nueva Entidad a un Dominio

1.  **Identifica el Paquete**: Navega al paquete correspondiente en la raíz del proyecto.
2.  **Crea el Archivo**: Crea un nuevo archivo `.go` (ej. `session.go`).
3.  **Define el `struct`**: Implementa el `struct`. Si la entidad pertenece a un tenant, **añade los campos `LessorID shared.UUID` y `TenantID shared.UUID`**.
4.  **Declara el Paquete**: Asegúrate de que el archivo comience con `package <nombre_del_paquete>`.

### Cómo Modificar un Puerto

Cuando necesites añadir o cambiar una funcionalidad, probablemente necesitarás modificar la interfaz del puerto correspondiente en el directorio `ports`. Asegúrate de que tus cambios sean abstractos y no contengan detalles de implementación.

## 5. Sugerencias de Mejora y Próximos Pasos

-   **Añadir Pruebas Unitarias**: Crear archivos `_test.go` para validar la lógica de negocio.
-   **Implementar Constructores**: Añadir funciones constructoras (ej. `NewUser(...)`) para asegurar que las entidades se creen en un estado válido.
-   **Usar Linters**: Configurar `golangci-lint` para asegurar la calidad y consistencia del código.
-   **Propagación de Contexto**: Todos los métodos de servicio deben aceptar un `context.Context` como primer argumento. Este patrón debe mantenerse para todos los métodos nuevos para manejar plazos, cancelaciones y valores con alcance de solicitud.
-   **Puerto de Configuración**: A medida que el sistema crezca, es posible que necesites un puerto dedicado para la configuración de la aplicación.
-   **Puerto de Feature Flags**: Para habilitar o deshabilitar funcionalidades dinámicamente, un puerto de feature flags sería beneficioso.
-   **Idempotencia**: Para operaciones que no deben repetirse, considera añadir soporte para claves de idempotencia en las firmas de los métodos de los puertos relevantes.
-   **Versionado**: A medida que la API evolucione, considera una estrategia para versionar las interfaces, quizás creando nuevos paquetes (e.g., `ports/v2`).

Recuerda, tu objetivo es mantener el núcleo limpio, abstracto e independiente de cualquier tecnología o implementación específica. ¡Buena suerte!