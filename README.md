# aauth-products-api



## Endpoint 1: /login

Este endpoint simula el proceso de inicio de sesión. Envía las credenciales username y password en una solicitud POST para obtener el accessToken y el refreshToken.

    curl -X POST http://localhost:8080/login -d "username=user&password=password"

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



        curl -X GET http://localhost:9292/protected -H "Authorization: Bearer {eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIiLCJyb2xlcyI6WyJ1c2VyIl0sImV4cCI6MTczMDg3OTY5N30.j5p6cS9fcc0sR3d0jMOMqiBlc5PyatxtkugKX1q4aQY}"




3. Revisar el Middleware de Autorización

curl -X GET http://localhost:9292/protected -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIiLCJyb2xlcyI6WyJ1c2VyIl0sImV4cCI6MTczMDg4MDMwNn0.FSBiRuDxAslZGHm-DfHnJxsT0no9DmybjJ2hQFhtoI4"


Revisa el código de tu middleware para asegurarte de que el token esté siendo validado correctamente. A continuación, algunos puntos importantes:

Decodificación del Token: El middleware debe decodificar el token y verificar su validez, incluyendo la expiración y los roles.

Revisión del Algoritmo de Firma: Confirma que el algoritmo (HS256 en tu caso) y la clave secreta utilizados en el middleware coinciden con los del token generado en /login.

4. Mensaje de Error Específico
Revisa los mensajes de error en la respuesta o en los registros de tu servidor. Esto puede darte información adicional sobre por qué la solicitud es rechazada. Algunos motivos comunes incluyen:

Token inválido
Token expirado
Problemas de roles o permisos (si el middleware revisa roles)
5. Habilita Mensajes de Depuración
Agrega mensajes de depuración en el middleware de autorización, especialmente en las partes donde:

Se decodifica el token.
Se verifica la validez y la expiración.
Se revisan los permisos o roles.
Con estos mensajes, podrás identificar en qué punto exacto se está rechazando el token.

Si sigues teniendo problemas, revisa y comparte el código del middleware o la función que verifica el accessToken para que podamos ver en detalle cómo se está manejando la autorización.