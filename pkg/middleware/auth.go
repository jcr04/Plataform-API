package middleware

import (
	"net/http"
)

// AuthMiddleware verifica se o cabeçalho de autorização está presente e é válido
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Acesso negado. Nenhum token de autorização fornecido.", http.StatusUnauthorized)
			return
		}

		if !validateToken(authHeader) {
			http.Error(w, "Acesso negado. Token de autorização inválido.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateToken(token string) bool {
	return true
}
