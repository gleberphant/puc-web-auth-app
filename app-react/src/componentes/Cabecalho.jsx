import {
  OverlayTrigger,
  Popover,
  Navbar,
  Nav,
  Container,
} from "react-bootstrap";
import { NavLink } from "react-router-dom";
import "../estilos/Layout.Page.css";
import PUCBRASAO from "../assets/images/pucpr-brasao-redondo.png";

export function Cabecalho({ logout, usuarioLogado }) {
  return (
    <Navbar className="app-navbar" expand="md" data-bs-theme="dark">
      <Container>
        <Navbar.Brand href="#home">
          <img height="50" src={PUCBRASAO}></img>
        </Navbar.Brand>

        <Navbar.Toggle aria-controls="basic-navbar-nav" />

        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto" justify="true">
            <Nav.Link as={NavLink} to="/">
              Home
            </Nav.Link>
            <Nav.Link as={NavLink} to="/usuarios">
              Listar Usuarios
            </Nav.Link>
            <Nav.Link as={NavLink} to="/cadastro">
              Novo Usuario
            </Nav.Link>

            <Nav.Link as={NavLink} to="/sobre">
              Sobre
            </Nav.Link>
          </Nav>
          <Nav className="justify-content-end">
            <OverlayTrigger
              trigger="click"
              placement="bottom"
              overlay={
                <Popover id="popover-positioned-bottom">
                  <Popover.Header as="h3">
                    {usuarioLogado?.login ?? "*indefinido*"}
                  </Popover.Header>
                  <Popover.Body>
                    <p>Uid: {usuarioLogado?.uid ?? "*indefinido*"}</p>
                    <p>Login: {usuarioLogado?.login ?? "*indefinido*"}</p>
                    <p>Email: {usuarioLogado?.email ?? "*indefinido*"}</p>
                    <p>
                      Nome Completo: {usuarioLogado?.nome ?? "*indefinido*"}
                    </p>
                    <p>Perfil: {usuarioLogado?.perfil ?? "*indefinido*"}</p>
                  </Popover.Body>
                </Popover>
              }
            >
              <Nav.Link as="button">Perfil</Nav.Link>
            </OverlayTrigger>
            <Nav.Link as="button" onClick={logout}>
              Sair
            </Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
