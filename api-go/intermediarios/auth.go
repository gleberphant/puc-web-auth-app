package intermediarios

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gleberphant/puc-react-app/api-go/configs"
	"github.com/gleberphant/puc-react-app/api-go/repositorios"
	"github.com/golang-jwt/jwt/v5"
)

// intermediário de autorização
func AuthMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/login" || req.Method == http.MethodOptions {
			log.Printf("Rota pública sem login")
			next.ServeHTTP(res, req)
			return
		}

		tokenString := req.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(res, "Token ausente", http.StatusUnauthorized)
			return
		}

		// verifica assinatura do token
		token, err := jwt.ParseWithClaims(
			tokenString,
			jwt.MapClaims{},
			func(token *jwt.Token) (interface{}, error) { // Impede aceitar outro algoritmo de assinatura
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(configs.JWTSECRET), nil
			},
		)

		if err != nil || !token.Valid {
			log.Printf("Token inválido ou expirado")
			http.Error(res, "Token inválido ou expirado", http.StatusUnauthorized)
			return
		}

		// extrai claims e perfil

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("Erro ao ler as claims do token")
			http.Error(res, "Erro interno ao processar token", http.StatusInternalServerError)
			return
		}

		perfil, ok := claims["perfil"].(string)

		if !ok {
			log.Printf("Aviso: Campo 'perfil' não encontrado no token")
			http.Error(res, "Perfil não identificado no token", http.StatusBadRequest)
			return
		}

		rotasPermitidas := *repositorios.MapaPermissoesMock()

		var rota string

		if strings.Count(req.URL.Path, "/") > 1 {
			rota = path.Dir(req.URL.Path)
		} else {
			rota = req.URL.Path
		}

		// for key := range rotasPermitidas {
		// 	if strings.HasPrefix(req.URL.Path, key) {
		// 		rotaBase = key
		// 		break
		// 	}
		// }

		metodosPermitidos, ok := rotasPermitidas[rota]
		if !ok {
			log.Printf("Rota não configurada no mapa de permissões: %s", req.URL.Path)
			http.Error(res, "Rota não encontrada", http.StatusNotFound)
			return
		}

		perfisPermitidos, ok := metodosPermitidos[req.Method]
		if !ok {
			log.Printf("Método %s não configurado para a rota %s", req.Method, req.URL.Path)
			http.Error(res, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		autorizado, ok := perfisPermitidos[perfil]

		if !autorizado || !ok {
			log.Printf("Perfil %s não autorizado para %s [%s]", perfil, req.Method, req.URL.Path)
			res.WriteHeader(http.StatusForbidden)
			json.NewEncoder(res).Encode(map[string]string{"error": "Perfil não autorizado"})
			return

		}

		next.ServeHTTP(res, req)
	})
}
