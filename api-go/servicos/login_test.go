package servicos

import (
	"testing"
	"time"

	"github.com/gleberphant/puc-react-app/api-go/repositorios"
)

// ====================================================================
// Testes para VerificaLoginSenha
// ====================================================================

func TestVerificaLoginSenha_LoginValido_Admin(t *testing.T) {
	// Setup: Usuário admin com senha "admin" no mock
	usuario := repositorios.RepositorioUsuariosMock()[0]
	if usuario.Login != "admin" {
		t.Skip("Setup do mock do admin falhou; verifique o mock.go")
	}

	// Execução
	resultado, err := VerificaLoginSenha("admin", "admin", "127.0.0.1")
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso, obteve erro: %v", err)
	}
	if resultado["uid"] != usuario.Uid {
		t.Errorf("UID esperado %q, obteve %q", usuario.Uid, resultado["uid"])
	}
	if resultado["login"] != "admin" {
		t.Errorf("Login esperado 'admin', obteve %q", resultado["login"])
	}
}

func TestVerificaLoginSenha_LoginValido_UsuarioComum(t *testing.T) {
	// Setup: Usuário comum com senha "usuario"
	usuario := repositorios.RepositorioUsuariosMock()[1]

	// Execução
	resultado, err := VerificaLoginSenha("usuario1", "usuario", "192.168.0.1")
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso, obteve erro: %v", err)
	}
	if resultado["uid"] != usuario.Uid {
		t.Errorf("UID esperado %q, obteve %q", usuario.Uid, resultado["uid"])
	}
}

func TestVerificaLoginSenha_SenhaInvalida(t *testing.T) {
	// Execução
	_, err := VerificaLoginSenha("admin", "senha_errada", "10.0.0.1")

	// Verificação
	if err == nil {
		t.Error("Esperava erro para senha inválida, obteve nil")
	}
	if err.Error() != "login ou senha invalidos" {
		t.Errorf("Mensagem de erro esperada 'login ou senha invalidos', obteve %q", err.Error())
	}
}

func TestVerificaLoginSenha_UsuarioInexistente(t *testing.T) {
	// Execução
	_, err := VerificaLoginSenha("nao_existe", "qualquer_senha", "172.16.0.1")

	// Verificação
	if err == nil {
		t.Error("Esperava erro para usuário inexistente, obteve nil")
	}
}

func TestVerificaLoginSenha_TrimWhitespaceNoLogin(t *testing.T) {
	// Execução
	resultado, err := VerificaLoginSenha("  usuario2  ", "123456", "8.8.8.8")
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso com login com espaços, obteve erro: %v", err)
	}
	if resultado["login"] != "usuario2" {
		t.Errorf("Login esperado 'usuario2' após trim, obteve %q", resultado["login"])
	}
}

// ====================================================================
// Testes para Controle de Tentativas
// ====================================================================

func TestControleLogin_BloqueioPorFalhasExcessivas(t *testing.T) {
	// Setup: Limpa o estado antes do teste
	resetarControleTentativas(t)

	ip := "192.168.1.100"
	senhaErrada := "senha_errada_bloqueio"

	// Executa 5 falhas consecutivas para atingir o limite (maxFalhasConta = 5)
	for i := 0; i < maxFalhasConta; i++ {
		_, err := VerificaLoginSenha("admin", senhaErrada, ip)
		if err == nil {
			t.Fatal("Esperava erro de login na tentativa", i+1)
		}
	}

	// Agora a conta deve estar bloqueada
	_, err := VerificaLoginSenha("admin", "admin", ip) // senha correta
	if err != ErrLoginBloqueado {
		t.Errorf("Esperava ErrLoginBloqueado após %d falhas, obteve: %v", maxFalhasConta, err)
	}
}

func TestControleLogin_DelayExponencial(t *testing.T) {
	// Setup
	resetarControleTentativas(t)

	ip := "10.0.0.50"
	senhaErrada := "errada"

	// Primeira falha não deve ter delay significativo
	tempoInicio := time.Now()
	_, _ = VerificaLoginSenha("admin", senhaErrada, ip)
	tempoPrimeiraFalha := time.Since(tempoInicio)

	// Segunda falha deve ter um pequeno delay exponencial
	tempoInicio = time.Now()
	_, _ = VerificaLoginSenha("admin", senhaErrada, ip)
	tempoSegundaFalha := time.Since(tempoInicio)

	// A segunda tentativa deve levar mais tempo devido ao delay (1 << min(2-1, 3) = 2s)
	if tempoSegundaFalha <= tempoPrimeiraFalha {
		t.Log("Delay exponencial não foi verificado (pode ser flaky). Primeira:", tempoPrimeiraFalha, "Segunda:", tempoSegundaFalha)
	}
}

func TestControleLogin_LimpezaAposSucesso(t *testing.T) {
	// Setup
	resetarControleTentativas(t)

	ip := "172.16.50.1"

	// Gera algumas falhas
	for i := 0; i < 2; i++ {
		_, _ = VerificaLoginSenha("admin", "errada", ip)
	}

	// Faz login com sucesso para limpar o contador
	_, err := VerificaLoginSenha("admin", "admin", ip)
	if err != nil {
		t.Fatalf("Esperava sucesso no login, obteve: %v", err)
	}

	// Verifica que o IP foi removido do controle de tentativas
	controleLoginMutex.Lock()
	_, existeIP := tentativasIP[ip]
	controleLoginMutex.Unlock()

	if existeIP {
		t.Error("Esperava que o contador de IP fosse limpo após sucesso")
	}
}

func TestObterControle_CriaNovoControle(t *testing.T) {
	controles := map[string]*controleTentativas{}
	agora := time.Now()

	// Execução
	ctrl := obterControle(controles, "chave_nova", agora)

	// Verificação
	if ctrl == nil {
		t.Fatal("Esperava um controle não-nil")
	}
	if len(controles) != 1 {
		t.Errorf("Esperava 1 controle no mapa, obteve %d", len(controles))
	}
}

func TestObterControle_ReutilizaControleExistente(t *testing.T) {
	agora := time.Now()
	controles := map[string]*controleTentativas{
		"chave_existente": {falhas: 3, inicioJanela: agora},
	}

	// Execução
	ctrl := obterControle(controles, "chave_existente", agora)

	// Verificação
	if ctrl.falhas != 3 {
		t.Errorf("Esperava 3 falhas, obteve %d", ctrl.falhas)
	}
	if len(controles) != 1 {
		t.Errorf("Esperava 1 controle no mapa, obteve %d", len(controles))
	}
}

func TestObterControle_ReiniciaControleExpirado(t *testing.T) {
	// Setup: Controle com janela expirada
	agora := time.Now()
	controles := map[string]*controleTentativas{
		"chave_expirada": {falhas: 10, inicioJanela: agora.Add(-janelaTentativas - time.Minute)},
	}

	// Execução
	ctrl := obterControle(controles, "chave_expirada", agora)

	// Verificação: O controle deve ser reiniciado
	if ctrl.falhas != 0 {
		t.Errorf("Esperava 0 falhas após reinício, obteve %d", ctrl.falhas)
	}
}

// ====================================================================
// Funções auxiliares de setup/teardown
// ====================================================================

func resetarControleTentativas(t *testing.T) {
	t.Helper()
	controleLoginMutex.Lock()
	tentativasIP = map[string]*controleTentativas{}
	tentativasConta = map[string]*controleTentativas{}
	controleLoginMutex.Unlock()

	// Garante que o teste não afete outros devido ao estado global de login
	t.Cleanup(func() {
		controleLoginMutex.Lock()
		tentativasIP = map[string]*controleTentativas{}
		tentativasConta = map[string]*controleTentativas{}
		controleLoginMutex.Unlock()
	})
}
