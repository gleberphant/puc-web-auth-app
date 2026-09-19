// estilos
import "./estilos/App.css";
import "bootstrap-icons/font/bootstrap-icons.css";

// dependencias
import { useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { fazerLogout } from "./servicos/autenticacao.js";

// minhas paginas
import HomePage from "./paginas/Home.Page";
import LoginPage from "./paginas/Login.Page";
import LayoutPage from "./paginas/Layout.Page.jsx";
import SobrePage from "./paginas/Sobre.Page.jsx";
import CadastroUsuarioPage from "./paginas/CriarUsuario.Page.jsx";
import ListarUsuariosPage from "./paginas/ListaUsuarios.Page.jsx";

export default function App() {
  const [logado, setLogado] = useState(false);
  const [usuarioLogado, setUsuarioLogado] = useState(null);

  /** TODO:
   * a autenticação no front precisa ser também um midleware;
   * caso usuario não seja autorizado (error 401 ou 403)  ele precisa ser redirecionado para login
   * criar um midleware para o roteador do front consultar o back se a rota é autorizada ao perfil
   *  posso carregar o mapa de permissoes no FRONT ou então fazer uma consulta em cada rota
   * */

  const logout = () => {
    fazerLogout();
    setLogado(false);
  };

  const login = (usuario) => {
    console.log("App recebeu usuario, ", usuario);
    setUsuarioLogado(usuario);
    setLogado(true);
  };

  if (logado && usuarioLogado != null)
    return (
      <BrowserRouter>
        <Routes>
          <Route
            element={
              <LayoutPage logout={logout} usuarioLogado={usuarioLogado} />
            }
          >
            <Route
              path="/"
              element={
                <HomePage logout={logout} usuarioLogado={usuarioLogado} />
              }
            />
            <Route path="/sobre" element={<SobrePage />} />
            <Route path="/usuarios" element={<ListarUsuariosPage />} />
            <Route path="/cadastro" element={<CadastroUsuarioPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    );
  else
    return (
      <>
        <LoginPage loginCallback={login} />
      </>
    );
}
