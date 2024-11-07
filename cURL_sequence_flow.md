# cURL Sequence flow


Para clarificar, aquí tienes el flujo de trabajo completo con curl:

1. Inicio de Sesión:

    curl -X POST http://localhost:9292/login -d "username=user&password=password"

Recibirás accessToken y refreshToken.



2. Acceso a un Recurso Protegido:

    curl -X GET http://localhost:9292/protected -H "Authorization: Bearer {accessToken}"


3. Renovación del accessToken cuando expira:

    curl -X POST http://localhost:9292/refresh -d "refreshToken={refreshToken}"

Esto te devolverá un nuevo accessToken.


**Este flujo asegura que el usuario tenga una experiencia de sesión continua, siempre y cuando el refreshToken siga siendo válido.**


