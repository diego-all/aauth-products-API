


```mermaid
sequenceDiagram
    participant Cliente
    participant Servidor

    Cliente->>Servidor: POST /login (credenciales)
    Servidor-->>Cliente: accessToken, refreshToken

    Cliente->>Servidor: GET /protected (Authorization: Bearer accessToken)
    alt Token válido
        Servidor-->>Cliente: Datos protegidos
    else Token no válido o expirado
        Servidor-->>Cliente: No autorizado
    end

    Cliente->>Servidor: POST /refresh (refreshToken)
    alt Refresh token válido y autorizado
        Servidor-->>Cliente: Nuevo accessToken
        Cliente->>Servidor: GET /protected (Authorization: Bearer nuevo accessToken)
        Servidor-->>Cliente: Datos protegidos
    else Refresh token no válido o expirado
        Servidor-->>Cliente: No autorizado
    end
