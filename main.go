package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_jwt_secret_key")
var refreshTokens = map[string]string{} // Almacenar refresh tokens por usuario

type Claims struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

// Genera un token de acceso con una duración corta
func generateAccessToken(username string, roles []string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Genera un refresh token con una duración más larga
func generateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Endpoint para iniciar sesión y obtener tokens
func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validación ficticia de usuario y contraseña
	if username == "user" && password == "password" {
		roles := []string{"user"}

		accessToken, err := generateAccessToken(username, roles)
		if err != nil {
			http.Error(w, "Error generando el token de acceso", http.StatusInternalServerError)
			return
		}

		refreshToken, err := generateRefreshToken(username)
		if err != nil {
			http.Error(w, "Error generando el refresh token", http.StatusInternalServerError)
			return
		}

		// Guardar el refresh token en un almacenamiento temporal
		refreshTokens[username] = refreshToken

		// Retornar ambos tokens
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"accessToken": "%s", "refreshToken": "%s"}`, accessToken, refreshToken)))
		return
	}
	http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
}

// Middleware para proteger rutas
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "No autorizado", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Endpoint protegido que requiere autenticación
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Acceso autorizado a datos protegidos"))
}

// Endpoint para refrescar el token de acceso
func refreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.FormValue("refreshToken")

	// Verifica si el refresh token es válido y existe
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Refresh token no válido", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		http.Error(w, "Refresh token inválido", http.StatusUnauthorized)
		return
	}

	username := claims["sub"].(string)
	storedToken, exists := refreshTokens[username]
	if !exists || storedToken != refreshToken {
		http.Error(w, "Refresh token no autorizado", http.StatusUnauthorized)
		return
	}

	// Genera un nuevo token de acceso
	roles := []string{"user"} // Aquí podrías recuperar roles de una base de datos
	newAccessToken, err := generateAccessToken(username, roles)
	if err != nil {
		http.Error(w, "Error generando el nuevo token de acceso", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"accessToken": "%s"}`, newAccessToken)))
}

func main() {
	r := chi.NewRouter()

	// Middleware de logging y recuperación de pánico
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rutas
	r.Post("/login", loginHandler)
	r.With(authMiddleware).Get("/protected", protectedHandler)
	r.Post("/refresh", refreshHandler)

	// Iniciar el servidor
	log.Println("Servidor escuchando en :9292")
	log.Fatal(http.ListenAndServe(":9292", r))
}
