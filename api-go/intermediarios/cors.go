package intermediarios

import (
	"net/http"

	"github.com/gleberphant/puc-react-app/api-go/configs"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")
		allowedOrigin := configs.CORS_ORIGIN

		if allowedOrigin == "*" {
			res.Header().Set("Access-Control-Allow-Origin", "*")
		} else if origin != "" && (origin == allowedOrigin || origin == "http://localhost:5173") {
			res.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			res.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		}

		// Métodos HTTP permitidos
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Cabeçalhos que o cliente tem permissão de enviar
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Permite envio de cookies/autenticação se necessário
		res.Header().Set("Access-Control-Allow-Credentials", "true")

		// Trata a requisição de Preflight (feita automaticamente pelo navegador via OPTIONS)
		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}

		// Prossegue para o próximo handler/mux
		next.ServeHTTP(res, req)
	})
}
