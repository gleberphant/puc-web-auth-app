package servicos

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gleberphant/puc-react-app/api-go/repositorios"
	"golang.org/x/crypto/bcrypt"
)

const (
	janelaTentativas = 15 * time.Minute
	tempoBloqueio    = 15 * time.Minute
	maxFalhasIP      = 10
	maxFalhasConta   = 5
)

var (
	ErrLoginBloqueado   = errors.New("login temporariamente bloqueado")
	ErrMuitasTentativas = errors.New("muitas tentativas de login")
	controleLoginMutex  sync.Mutex
	tentativasIP        = map[string]*controleTentativas{}
	tentativasConta     = map[string]*controleTentativas{}
)

type controleTentativas struct {
	falhas       int
	inicioJanela time.Time
	bloqueadoAte time.Time
}

func VerificaLoginSenha(login string, senhaTexto string, ip string) (map[string]string, error) {
	login = strings.TrimSpace(login)
	conta := strings.ToLower(login)
	err := aguardarTentativa(ip, conta)
	if err != nil {
		return nil, err
	}

	for _, usuario := range repositorios.RepositorioUsuariosMock() {
		if usuario.Login == login &&
			(bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senhaTexto)) == nil) {
			registrarSucesso(ip, conta)
			return map[string]string{
				"uid":    usuario.Uid,
				"nome":   usuario.Nome,
				"login":  usuario.Login,
				"perfil": usuario.Perfil,
			}, nil
		}
	}

	registrarFalha(ip, conta)
	return nil, errors.New("login ou senha invalidos")
}

func aguardarTentativa(ip string, conta string) error {
	agora := time.Now()
	controleLoginMutex.Lock()
	porIP := obterControle(tentativasIP, ip, agora)
	porConta := obterControle(tentativasConta, conta, agora)
	if porConta.bloqueadoAte.After(agora) {
		controleLoginMutex.Unlock()
		return ErrLoginBloqueado
	}
	if porIP.falhas >= maxFalhasIP {
		controleLoginMutex.Unlock()
		return ErrMuitasTentativas
	}
	falhas := porIP.falhas
	if porConta.falhas > falhas {
		falhas = porConta.falhas
	}
	controleLoginMutex.Unlock()

	if falhas == 0 {
		return nil
	}
	atraso := time.Duration(1<<min(falhas-1, 3)) * time.Second
	time.Sleep(atraso)
	return nil
}

func obterControle(controles map[string]*controleTentativas, chave string, agora time.Time) *controleTentativas {
	controle, ok := controles[chave]
	if !ok || agora.Sub(controle.inicioJanela) >= janelaTentativas {
		controle = &controleTentativas{inicioJanela: agora}
		controles[chave] = controle
	}
	return controle
}

func registrarFalha(ip string, conta string) {
	agora := time.Now()
	controleLoginMutex.Lock()
	porIP := obterControle(tentativasIP, ip, agora)
	porConta := obterControle(tentativasConta, conta, agora)
	porIP.falhas++
	porConta.falhas++
	if porConta.falhas >= maxFalhasConta {
		porConta.bloqueadoAte = agora.Add(tempoBloqueio)
	}
	falhasConta := porConta.falhas
	bloqueadoAte := porConta.bloqueadoAte
	controleLoginMutex.Unlock()

	if falhasConta >= 3 {
		log.Printf("tentativa suspeita de login: conta=%q ip=%q falhas=%d bloqueada_ate=%s", conta, ip, falhasConta, bloqueadoAte.Format(time.RFC3339))
	}
}

func registrarSucesso(ip string, conta string) {
	controleLoginMutex.Lock()
	delete(tentativasIP, ip)
	delete(tentativasConta, conta)
	controleLoginMutex.Unlock()
}
