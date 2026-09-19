package servicos

import (
	"errors"

	"github.com/gleberphant/puc-react-app/api-go/modelos"
	"github.com/gleberphant/puc-react-app/api-go/repositorios"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var repoUsuario []modelos.Usuario = repositorios.RepositorioUsuariosMock()

func CriarUsuarios(novoUsuario modelos.Usuario) error {
	if novoUsuario.Uid == "" {
		novoUsuario.Uid = uuid.New().String()
	}

	for _, usuario := range repoUsuario {
		if usuario.Uid == novoUsuario.Uid {
			return errors.New("usuario ja existe")
		}
	}

	senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	novoUsuario.Senha = string(senhaCriptografada)

	repoUsuario = append(repoUsuario, novoUsuario)
	return nil
}

func ListarUsuarios() ([]modelos.Usuario, error) {
	listaUsuarios := make([]modelos.Usuario, len(repoUsuario))
	copy(listaUsuarios, repoUsuario)
	return listaUsuarios, nil
}

func ExibirUsuario(uid string) (*modelos.Usuario, error) {
	for _, u := range repoUsuario {
		if u.Uid == uid {
			return &u, nil
		}
	}

	return nil, errors.New("Usuario não encontrado")
}

func EditarUsuarios(novoUsuario modelos.Usuario) error {
	senhaCriptografada, err := bcrypt.GenerateFromPassword([]byte(novoUsuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	novoUsuario.Senha = string(senhaCriptografada)

	for i := range repoUsuario {
		if repoUsuario[i].Uid == novoUsuario.Uid {
			repoUsuario[i] = novoUsuario
			return nil
		}
	}

	return errors.New("Usuario não encontrato")
}

func DeletarUsuarios(uid string) error {
	for i := range repoUsuario {
		if repoUsuario[i].Uid == uid {
			repoUsuario = append(repoUsuario[:i], repoUsuario[i+1:]...)
			return nil
		}
	}
	return errors.New("Usuario não encontrato")
}
