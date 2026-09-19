import "../estilos/ListaUsuarios.Page.css";

import { useEffect, useState } from "react";
import { Button, Spinner, Table, Alert } from "react-bootstrap";

import {
  EditarUsuario,
  ExcluirUsuario,
  ListarUsuarios,
} from "../servicos/usuarios";

import Carregando from "../componentes/Carregando";
import EditarUsuarioModal from "./EditarUsuario.Modal";
import ExibirUsuarioModal from "./ExibirUsuario.Modal";

export default function ListaUsuariosPage() {
  const [mensagem, setMensagem] = useState("");
  const [listaUsuarios, setListaUsuarios] = useState([]);
  const [carregando, setCarregando] = useState(true);
  const [excluindo, setExcluindo] = useState(false);
  const [usuarioSelecionado, setUsuarioSelecionado] = useState({});
  const [modalAberta, setModalAberta] = useState(null);

  const carregarLista = async () => {
    setCarregando(true);

    const [lista, erro] = await ListarUsuarios();

    if (erro) {
      setListaUsuarios([]);
      alert(`Falha no carregamento da lista: ${erro.message}`);
    } else {
      setListaUsuarios(lista);
    }

    setCarregando(false);
  };

  useEffect(() => {
    const carregarListaInicial = async () => {
      setCarregando(true);

      const [lista, erro] = await ListarUsuarios();

      if (erro) {
        setListaUsuarios([]);
        //alert(`Falha no carregamento da lista: ${erro.message}`);

        setMensagem(`${erro.message}`);
      } else {
        setListaUsuarios(lista);
      }

      setCarregando(false);
    };

    carregarListaInicial();
  }, []);

  const abrirDetalhes = (usuario) => {
    setUsuarioSelecionado(usuario);
    setModalAberta("detalhes");
  };

  const abrirEdicao = (usuario) => {
    setUsuarioSelecionado(usuario);
    setModalAberta("edicao");
  };

  const fecharModal = () => {
    setModalAberta(null);
  };

  const salvarEdicao = async (usuario) => {
    const [, erro] = await EditarUsuario(usuario);

    if (erro) {
      alert(`Falha ao editar usuário: ${erro.message}`);
      return false;
    }

    await carregarLista();
    return true;
  };

  const excluirUsuario = async (usuario) => {
    if (!window.confirm(`Deseja realmente remover ${usuario.Nome}?`)) {
      return;
    }

    setExcluindo(true);

    const [, erro] = await ExcluirUsuario(usuario.Uid);

    setExcluindo(false);

    if (erro) {
      alert(`Falha ao excluir usuário: ${erro.message}`);
      return;
    }

    await carregarLista();
  };

  if (carregando) {
    return <Carregando />;
  }

  return (
    <div className="lista-usuarios">
      {mensagem && (
        <Alert variant="danger" dismissible onClose={() => setMensagem("")}>
          {mensagem}
        </Alert>
      )}

      <ExibirUsuarioModal
        usuarioSelecionado={usuarioSelecionado}
        show={modalAberta === "detalhes"}
        fechar={fecharModal}
      />

      <EditarUsuarioModal
        key={usuarioSelecionado?.Uid ?? "edicao"}
        usuarioSelecionado={usuarioSelecionado}
        show={modalAberta === "edicao"}
        fechar={fecharModal}
        onSubmit={salvarEdicao}
      />

      <Table striped hover size="sm">
        <thead>
          <tr>
            <th>Login</th>
            <th>Nome completo</th>
            <th>E-mail</th>
            <th>Perfil</th>
            <th>Ações</th>
          </tr>
        </thead>

        <tbody>
          {listaUsuarios.map((usuario) => (
            <tr key={usuario.Uid}>
              <td>{usuario.Login}</td>
              <td>{usuario.Nome}</td>
              <td>{usuario.Email}</td>
              <td>{usuario.Perfil}</td>
              <td>
                <Button
                  variant="link"
                  size="sm"
                  onClick={() => abrirDetalhes(usuario)}
                  aria-label={`Ver detalhes de ${usuario.Nome}`}
                  className="bi bi-eye-fill"
                  style={{ fontSize: "24px" }}
                />

                <Button
                  variant="link"
                  size="sm"
                  onClick={() => abrirEdicao(usuario)}
                  aria-label={`Editar ${usuario.Nome}`}
                  className="bi bi-pen-fill"
                  style={{ fontSize: "22px", color: "orange" }}
                />

                {!excluindo ? (
                  <Button
                    variant="link"
                    size="sm"
                    onClick={() => excluirUsuario(usuario)}
                    aria-label={`Excluir ${usuario.Nome}`}
                    className="bi bi-person-x-fill"
                    style={{ fontSize: "24px", color: "crimson" }}
                  />
                ) : (
                  <Spinner
                    animation="border"
                    size="sm"
                    role="status"
                    aria-label="Excluindo usuário"
                  />
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}
