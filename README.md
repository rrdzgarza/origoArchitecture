# Origo: Definiciones de Dominios y Puertos

Este repositorio es la **única fuente de verdad** para las definiciones de la arquitectura del ecosistema de aplicaciones de Origo. Su propósito es centralizar los **modelos de dominio** y los **puertos** (interfaces) que forman el núcleo de la lógica de negocio, siguiendo los principios de una **arquitectura hexagonal**.

**Este repositorio contiene únicamente definiciones (`structs` e `interfaces` de Go), no implementaciones.**

## Propósito y Alcance

El objetivo de este repositorio es proporcionar una base arquitectónica coherente y desacoplada para todos los proyectos consumidores (APIs, ERPs, etc.). Al centralizar las definiciones, se garantiza que todos los componentes del ecosistema hablen el mismo lenguaje y se adhieran a los mismos contratos.

El alcance de este repositorio se limita estrictamente a:

-   **Definición de Dominios**: `Structs` de Go que modelan las entidades de negocio principales (ej. `User`, `Tenant`, `Lessor`). Se encuentran en el directorio `domains/`.
-   **Definición de Puertos**: `Interfaces` de Go que definen los contratos para los servicios y repositorios (ej. `AuthenticationService`, `UserRepository`). Se encuentran en el directorio `ports/`.
-   **Definición de Errores Compartidos**: Un conjunto de errores estándar (ej. `ErrNotFound`) para un manejo consistente a través de las aplicaciones. Se encuentra en `domains/shared/errors`.

Cualquier lógica de implementación, como constructores de entidades, lógica de negocio compleja o pruebas unitarias, **debe residir en los repositorios consumidores** que importan este paquete.

## Estructura del Repositorio

-   `domains/`: Contiene los paquetes de dominio, cada uno representando un área de negocio.
    -   `authentication/`, `lessors/`, `tenants/`, etc.
    -   `shared/`: Contiene tipos de datos y errores transversales a todos los dominios.
-   `ports/`: Contiene las interfaces que definen los contratos de la aplicación.

## Uso en Otro Proyecto

Para utilizar estas definiciones en otro proyecto de Go, impórtalas utilizando la ruta del módulo:

```go
import (
    "github.com/rrdzgarza/origoDomains/domains/authentication"
    "github.com/rrdzgarza/origoDomains/ports"
    "github.com/rrdzgarza/origoDomains/domains/shared/errors"
)

// Ejemplo de una implementación de un puerto en un proyecto consumidor.
type PostgresUserRepository struct {
    // ... dependencias de base de datos
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id shared.UUID) (*authentication.User, error) {
    // ... lógica para buscar un usuario en PostgreSQL.
    // Si no se encuentra, se debe devolver el error estándar:
    return nil, errors.ErrNotFound
}
```

---
*Para una guía de desarrollo detallada, consulta el archivo `AGENTS.md`.*