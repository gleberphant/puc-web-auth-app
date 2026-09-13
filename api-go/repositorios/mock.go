package repositorios

import (
	"github.com/gleberphant/puc-react-app/api-go/modelos"
	"golang.org/x/crypto/bcrypt"
)

// var mapaPermissões2 = MapaPermissoes{
// 	"admin": {
// 		"/usuarios": {"GET": true, "POST": true, "PUT": true, "DELETE": true},
// 		"/login":    {"GET": true, "POST": true},
// 	},
// 	"usuario": {
// 		"/usuarios": {"GET": true, "POST": true, "PUT": true, "DELETE": true},
// 		"/login":    {"GET": true, "POST": true},
// 	},
// 	"cliente": {
// 		"/usuarios": {"GET": true, "POST": true, "PUT": true, "DELETE": true},
// 		"/login":    {"GET": true, "POST": true},
// 	},
// }

type MapaPermissoesType map[string]map[string]map[string]bool

var permissoesFront = map[string][]string{
	"admin":   {"home", "listar usuarios", "exibir usuarios", "editar usuarios", "novo usuario", "sobre"},
	"usuario": {"home", "listar usuarios", "exibir usuarios", "sobre"},
	"cliente": {"home", "sobre"},
}

func MapaPermissoesFrontMock() *map[string][]string {
	return &permissoesFront
}

var permissoes = MapaPermissoesType{
	"/usuarios": {
		"GET":    {"admin": true, "usuario": true, "cliente": false},
		"POST":   {"admin": true, "usuario": false, "cliente": false},
		"PUT":    {"admin": true, "usuario": false, "cliente": false},
		"DELETE": {"admin": true, "usuario": false, "cliente": false},
	},

	"/login": {
		"GET":    {"admin": true, "usuario": true, "cliente": true},
		"POST":   {"admin": true, "usuario": true, "cliente": true},
		"PUT":    {"admin": true, "usuario": true, "cliente": true},
		"DELETE": {"admin": true, "usuario": true, "cliente": true},
	},
}

func MapaPermissoesMock() *MapaPermissoesType {
	return &permissoes
}

func encriptarSenha(senha string) string {
	senhaEncriptada, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		panic("Falha na criptografia de senha")
	}
	return string(senhaEncriptada)
}

var repoUsuario = []modelos.Usuario{
	{
		Uid:    "00000000-0000-0000-0000-000000000000",
		Login:  "admin",
		Senha:  encriptarSenha("admin"),
		Nome:   "Adminsitrador",
		Email:  "admin@admin",
		Perfil: "admin",
	},
	{
		Uid:    "c6f23200-df9d-45a8-996e-2b92afd6a215",
		Login:  "usuario1",
		Senha:  encriptarSenha("usuario"),
		Nome:   "Usuario1 Nome completo ",
		Email:  "usuario@usuario",
		Perfil: "usuario",
	},
	{
		Uid:    "7746da64-fc2e-429b-aa17-c1c4b4c76962",
		Login:  "usuario2",
		Senha:  encriptarSenha("123456"),
		Nome:   "Usuario2 Nome completo",
		Email:  "usuario2@usuario2",
		Perfil: "usuario",
	},

	{
		Uid:    "12345678-1234-1234-1234-123456789000",
		Login:  "cliente1",
		Senha:  encriptarSenha("cliente"),
		Nome:   "Cliente nome completo",
		Email:  "usuario2@usuario2",
		Perfil: "cliente",
	},
}

func RepositorioUsuariosMock() []modelos.Usuario {
	return repoUsuario
}
