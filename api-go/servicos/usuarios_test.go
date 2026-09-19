package servicos

import (
	"testing"

	"github.com/gleberphant/puc-react-app/api-go/modelos"
	"github.com/gleberphant/puc-react-app/api-go/repositorios"
)

// ====================================================================
// Testes para CriarUsuarios
// ====================================================================

func TestCriarUsuarios_Sucesso(t *testing.T) {
	// Setup: Salva estado original e limpa
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	novoUsuario := modelos.Usuario{
		Login:  "novo_usuario",
		Senha:  "senha_secreta",
		Email:  "novo@teste.com",
		Nome:   "Novo Usuario",
		Perfil: "usuario",
	}

	// Execução
	err := CriarUsuarios(novoUsuario)
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso na criação, obteve erro: %v", err)
	}

	// Verifica se o usuário foi adicionado
	if len(repoUsuario) != len(repositorios.RepositorioUsuariosMock())+1 {
		t.Errorf("Esperava %d usuários, obteve %d", len(repositorios.RepositorioUsuariosMock())+1, len(repoUsuario))
	}

	// Verifica se a senha foi criptografada
	for _, u := range repoUsuario {
		if u.Login == "novo_usuario" {
			if u.Senha == "senha_secreta" {
				t.Error("Esperava senha criptografada, obteve senha em texto plano")
			}
		}
	}
}

func TestCriarUsuarios_Duplicado(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	// Usa UID de um usuário existente no mock
	usuarioExistente := repositorios.RepositorioUsuariosMock()[0]
	novoUsuario := modelos.Usuario{
		Uid:   usuarioExistente.Uid, // UID duplicado
		Login: "outro_nome",
		Senha: "senha",
	}

	// Execução
	err := CriarUsuarios(novoUsuario)

	// Verificação
	if err == nil {
		t.Error("Esperava erro de usuário duplicado, obteve nil")
	}
	if err.Error() != "usuario ja existe" {
		t.Errorf("Mensagem esperada 'usuario ja existe', obteve %q", err.Error())
	}
}

// ====================================================================
// Testes para ListarUsuarios
// ====================================================================

func TestListarUsuarios(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	mockUsuarios := repositorios.RepositorioUsuariosMock()
	repoUsuario = make([]modelos.Usuario, len(mockUsuarios))
	copy(repoUsuario, mockUsuarios)
	t.Cleanup(func() { repoUsuario = repoOriginal })

	// Execução
	resultado, err := ListarUsuarios()
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso, obteve erro: %v", err)
	}
	if len(resultado) != len(repoUsuario) {
		t.Errorf("Esperava %d usuários, obteve %d", len(repoUsuario), len(resultado))
	}
}

func TestListarUsuarios_RetornaCopia(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	// Execução
	resultado, _ := ListarUsuarios()

	// Modifica o resultado
	if len(resultado) > 0 {
		resultado[0].Login = "MODIFICADO"
	}

	// Verifica se o repositório original foi preservado
	if repoUsuario[0].Login == "MODIFICADO" {
		t.Error("Esperava cópia da lista, mas o repositório original foi modificado")
	}
}

// ====================================================================
// Testes para ExibirUsuario
// ====================================================================

func TestExibirUsuario_Sucesso(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	uidAlvo := "00000000-0000-0000-0000-000000000000"

	// Execução
	resultado, err := ExibirUsuario(uidAlvo)
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso, obteve erro: %v", err)
	}
	if resultado == nil {
		t.Fatal("Esperava usuário não-nill")
	}
	if resultado.Uid != uidAlvo {
		t.Errorf("UID esperado %q, obteve %q", uidAlvo, resultado.Uid)
	}
	if resultado.Login != "admin" {
		t.Errorf("Login esperado 'admin', obteve %q", resultado.Login)
	}
}

func TestExibirUsuario_NaoEncontrado(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	uidInexistente := "uid-que-nao-existe"

	// Execução
	_, err := ExibirUsuario(uidInexistente)

	// Verificação
	if err == nil {
		t.Error("Esperava erro para usuário não encontrado, obteve nil")
	}
	if err.Error() != "Usuario não encontrado" {
		t.Errorf("Mensagem esperada 'Usuario não encontrado', obteve %q", err.Error())
	}
}

// ====================================================================
// Testes para EditarUsuarios
// ====================================================================

func TestEditarUsuarios_Sucesso(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	uidAlvo := "7746da64-fc2e-429b-aa17-c1c4b4c76962"
	usuarioAtualizado := modelos.Usuario{
		Uid:    uidAlvo,
		Login:  "usuario2_editado",
		Senha:  "nova_senha123",
		Nome:   "Usuario2 Editado",
		Email:  "novo_email@usuario2.com",
		Perfil: "usuario",
	}

	// Execução
	err := EditarUsuarios(usuarioAtualizado)
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso, obteve erro: %v", err)
	}

	// Verifica se os dados foram atualizados
	usuarioEditado, err := ExibirUsuario(uidAlvo)
	if err != nil {
		t.Fatalf("Erro ao buscar usuário editado: %v", err)
	}
	if usuarioEditado.Login != "usuario2_editado" {
		t.Errorf("Login esperado 'usuario2_editado', obteve %q", usuarioEditado.Login)
	}
	if usuarioEditado.Email != "novo_email@usuario2.com" {
		t.Errorf("Email esperado 'novo_email@usuario2.com', obteve %q", usuarioEditado.Email)
	}
}

func TestEditarUsuarios_NaoEncontrado(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	uidInexistente := "uid-que-nao-existe"
	usuarioAtualizado := modelos.Usuario{
		Uid:   uidInexistente,
		Login: "nao_importa",
		Senha: "senha",
	}

	// Execução
	err := EditarUsuarios(usuarioAtualizado)

	// Verificação
	if err == nil {
		t.Error("Esperava erro para usuário não encontrado, obteve nil")
	}
	if err.Error() != "Usuario não encontrato" {
		t.Errorf("Mensagem esperada 'Usuario não encontrato', obteve %q", err.Error())
	}
}

// ====================================================================
// Testes para DeletarUsuarios
// ====================================================================

func TestDeletarUsuarios_Sucesso(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	uidDeletar := repoUsuario[2].Uid // Terceriro usuário do mock
	tamanhoAntes := len(repoUsuario)
	t.Cleanup(func() { repoUsuario = repoOriginal })

	// Execução
	err := DeletarUsuarios(uidDeletar)
	// Verificação
	if err != nil {
		t.Errorf("Esperava sucesso, obteve erro: %v", err)
	}
	if len(repoUsuario) != tamanhoAntes-1 {
		t.Errorf("Esperava %d usuários após deleção, obteve %d", tamanhoAntes-1, len(repoUsuario))
	}

	// Verifica se o usuário foi realmente removido
	_, err = ExibirUsuario(uidDeletar)
	if err == nil {
		t.Error("Esperava erro ao buscar usuário deletado")
	}
}

func TestDeletarUsuarios_NaoEncontrado(t *testing.T) {
	// Setup
	repoOriginal := repoUsuario
	repoUsuario = repositorios.RepositorioUsuariosMock()
	t.Cleanup(func() { repoUsuario = repoOriginal })

	uidInexistente := "uid-que-nao-existe"

	// Execução
	err := DeletarUsuarios(uidInexistente)

	// Verificação
	if err == nil {
		t.Error("Esperava erro para usuário não encontrado, obteve nil")
	}
	if err.Error() != "Usuario não encontrato" {
		t.Errorf("Mensagem esperada 'Usuario não encontrato', obteve %q", err.Error())
	}
}
