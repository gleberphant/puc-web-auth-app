package manipuladores

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gleberphant/puc-react-app/api-go/configs"
	"github.com/gleberphant/puc-react-app/api-go/servicos"
	"github.com/golang-jwt/jwt/v5"
)

func InjetarRotasLogin(roteador *http.ServeMux) {
	roteador.HandleFunc("POST /login", LoginPost)
}

// retorna o proprio usuario logado
func LoginGet(res http.ResponseWriter, req *http.Request) {
	// define struct que vai receber oo body extraido do request
	var requestBody struct {
		Uid string `json:"uid"`
	}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Error Decoder: %s", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}

	// chama o service
	usuario, err := servicos.ExibirUsuario(requestBody.Uid)
	// confirmação do service
	if err != nil {
		log.Printf("Error: %s", err.Error())
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(map[string]string{"error": "Usuario não autorizado"})
		return
	}

	// responde ao cliente

	responseBody := map[string]any{"usuario": usuario}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(responseBody)
}

// requisição da autenticação de um usuario com username e senha
func LoginPost(res http.ResponseWriter, req *http.Request) {
	// define struct que vai receber oo body extraido do request
	var requestBody struct {
		Login string `json:"login"`
		Senha string `json:"senha"`
	}

	// extrai login e senha do body
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Error Decoder: %s", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}

	// chama o service
	ip := req.RemoteAddr
	if host, _, erroEndereco := net.SplitHostPort(req.RemoteAddr); erroEndereco == nil {
		ip = host
	}
	usuarioLogado, err := servicos.VerificaLoginSenha(requestBody.Login, requestBody.Senha, ip)
	// confirmação do service
	if err != nil {
		log.Printf("Error: %s", err.Error())
		status := http.StatusUnauthorized
		res.WriteHeader(status)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}

	// cria token jwt
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uid":    usuarioLogado["uid"],
			"login":  usuarioLogado["login"],
			"perfil": usuarioLogado["perfil"],
			"exp":    time.Now().Add(time.Hour).Unix(),
			"iat":    time.Now().Unix(),
		})

	// transforma em strings
	tokenString, err := tokenJwt.SignedString([]byte(configs.JWTSECRET))
	if err != nil {
		log.Printf("Error: %s", err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{"error": "Erro ao gerar Token "})
		return
	}

	// body da resposta ao cliente
	responseBody := map[string]any{"usuario": usuarioLogado, "token": tokenString, "permissoes": map[string][]string{}}

	// responde ao cliente
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(responseBody)
}
