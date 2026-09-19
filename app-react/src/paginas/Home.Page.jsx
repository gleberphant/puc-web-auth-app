import { Button, Card } from "react-bootstrap";

export default function HomePage({ logout, usuarioLogado }) {
  console.log("ususario no home", usuarioLogado);
  return (
    <>
      <Card>
        <Card.Header>
          <h1>
            Bem Vindo <b>{usuarioLogado?.login ?? "*indefinido*"}</b>
          </h1>
        </Card.Header>
        <Card.Body>
          <p>Uid: {usuarioLogado?.uid ?? "*indefinido*"}</p>
          <p>Login: {usuarioLogado?.login ?? "*indefinido*"}</p>
          <p>Email: {usuarioLogado?.email ?? "*indefinido*"}</p>
          <p>Nome Completo: {usuarioLogado?.nome ?? "*indefinido*"}</p>
          <p>Perfil: {usuarioLogado?.perfil ?? "*indefinido*"}</p>
        </Card.Body>
        <Card.Footer>
          <Button variant="outline-danger" onClick={logout}>
            Clique para Sair
          </Button>
        </Card.Footer>
      </Card>
      <p></p>
    </>
  );
}
