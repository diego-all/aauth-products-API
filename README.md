# aauth-products-api


> Nota: Remover los {}

## Endpoint 1: /login

Este endpoint simula el proceso de inicio de sesión. Envía las credenciales username y password en una solicitud POST para obtener el accessToken y el refreshToken.

    curl -X POST http://localhost:9292/login -d "username=user&password=password"

Respuesta esperada
Si las credenciales son correctas, recibirás un JSON con ambos tokens:

    {
    "accessToken": "jwt_token_generado_para_acceso",
    "refreshToken": "jwt_token_generado_para_refresco"
    }


## Endpoint 2: /protected

Este endpoint está protegido y solo es accesible si incluyes el accessToken en la solicitud. Para acceder, envía el token en la cabecera Authorization en una solicitud GET

Solicitud de Acceso Protegido


    curl -X GET http://localhost:9292/protected -H "Authorization: Bearer {accessToken}"



## Endpoint 3: /refresh

Para usar el endpoint /refresh y obtener un nuevo accessToken cuando el actual ha expirado.

    curl -X POST http://localhost:9292/refresh -d "refreshToken={refreshToken}"



**Respuesta con un Nuevo accessToken:**

Si el refreshToken es válido y no ha expirado, el servidor te devuelve un nuevo accessToken, que podrás usar para futuras solicitudes a los recursos protegidos.
Si el refreshToken ha expirado o es inválido, el usuario tendrá que autenticarse de nuevo (es decir, volver a hacer login).


### Concepto de sesión

En este programa, la "sesión" se maneja de forma ligera, utilizando tokens en lugar de almacenar información de la sesión en el servidor. Esto implica:

Sin Estado en el Servidor: El servidor no almacena el estado de la sesión del usuario. En su lugar, depende de los tokens JWT para verificar la autenticación. Cada solicitud incluye un token que es auto-suficiente, conteniendo toda la información necesaria (como el nombre de usuario y roles) para validar la solicitud.

Control de Tiempo de Vida: La sesión tiene dos tiempos de vida distintos:

Token de Acceso: Este define la duración de la sesión activa en la aplicación (por ejemplo, 15 minutos). Si el accessToken expira, el usuario debe solicitar un nuevo token usando el refreshToken.
Token de Refresh: Controla la duración general de la sesión. Mientras el refreshToken sea válido, el usuario puede continuar renovando el accessToken. Una vez que el refreshToken expira, el usuario debe autenticarse nuevamente con sus credenciales.



### Resumen del Proceso Completo

1. Inicio de Sesión (/login): El usuario envía sus credenciales y recibe accessToken y refreshToken.

2. Acceso a Recursos Protegidos (/protected): El usuario utiliza accessToken en el encabezado Authorization para acceder a los recursos protegidos.

3. Renovación de Token (/refresh): Cuando el accessToken expira, el usuario usa refreshToken para obtener un nuevo accessToken.

4. Reinicio de Sesión: Si el refreshToken también expira, el usuario debe autenticarse de nuevo.

