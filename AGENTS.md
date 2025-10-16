# AGENTS.md: Guía para Definiciones de Dominio y Puertos

## Propósito de este Repositorio

Tu única función en este repositorio es **definir la arquitectura central** del ecosistema Origo. Este repositorio contiene exclusivamente las **definiciones** de los modelos de dominio (`structs`) y los puertos (`interfaces`), no su implementación.

**Regla fundamental: No escribas lógica de implementación aquí.** Esto incluye constructores, métodos con lógica de negocio, pruebas unitarias o cualquier otro código que no sea una `struct` o una `interface` de Go.

## Tu Tarea como Agente

Tu trabajo se centra en las siguientes tareas:

### 1. Modificar o Añadir `structs` de Dominio

-   **Ubicación**: `domains/<nombre-del-dominio>/`
-   **Acción**: Crea o modifica archivos `.go` para definir las `structs` que representan las entidades del negocio.
-   **Ejemplo**: Si necesitas añadir un campo `Status` a la `struct` `User`, modificarías `domains/authentication/user.go`.

```go
// domains/authentication/user.go

package authentication

import "github.com/rrdzgarza/origoArchitecture/domains/shared"

// CORRECTO: Solo campos de datos.
type User struct {
    ID        shared.UUID
    Email     string
    Status    string // Nuevo campo añadido.
}
```

### 2. Modificar o Añadir `interfaces` (Puertos)

-   **Ubicación**: `ports/`
-   **Acción**: Crea o modifica `interfaces` que definen los contratos para los servicios o repositorios.
-   **Ejemplo**: Si necesitas un método para encontrar usuarios por `Status`, lo añadirías a la `UserRepository`.

```go
// ports/repositories.go

package ports

// ... imports

type UserRepository interface {
    CreateUser(ctx context.Context, user *authentication.User) error
    GetUserByID(ctx context.Context, id shared.UUID) (*authentication.User, error)
    // CORRECTO: Nuevo método añadido a la interfaz.
    GetUsersByStatus(ctx context.Context, status string) ([]*authentication.User, error)
}
```

### 3. Modificar o Añadir Errores Compartidos

-   **Ubicación**: `domains/shared/errors/`
-   **Acción**: Define nuevos errores estándar que serán utilizados en todo el ecosistema.

```go
// domains/shared/errors/errors.go

package errors

import "errors"

var (
    ErrNotFound = errors.New("not found")
    // CORRECTO: Nuevo error estándar añadido.
    ErrDuplicateRecord = errors.New("duplicate record")
)
```

## Lo que NO Debes Hacer

-   **NO** añadir funciones constructoras (ej. `NewUser()`).
-   **NO** añadir métodos con lógica a las `structs`.
-   **NO** crear archivos `_test.go`.
-   **NO** implementar las interfaces definidas en `ports/`.

Recuerda: este es el **plano arquitectónico**, no el edificio. Las implementaciones se realizan en otros repositorios.