//estilos
import "../estilos/Login.Page.css";
import PUCBRASAO from "../assets/images/pucpr-brasao-redondo.png";

//dependencias
import { Container, Card, Button, Form } from "react-bootstrap";
import { useState } from "react";

//meus componentes
import Carregando from "../componentes/Carregando";
import { fazerLogin } from "../servicos/autenticacao";

export default function LoginPage({ loginCallback }) {
  const [carregando, setCarregando] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    setCarregando(true);

    const f = new FormData(e.target);

    const [usuario, err] = await fazerLogin(f.get("login"), f.get("senha"));

    if (err != null || usuario == null) {
      alert(err);
      setCarregando(false);
      return;
    }

    console.log("Usuario recebido ", usuario);
    loginCallback(usuario);
    setCarregando(false);
  };

  if (carregando) return <Carregando></Carregando>;
  return (
    <Container fluid className="login-page">
      <Card className="login-card">
        <img className="login-logo" src={PUCBRASAO} alt="Brasão da PUCPR" />

        <h4 className="text-center mb-4">Sistemas Web Seguros</h4>

        <Form onSubmit={handleSubmit}>
          <Form.Group className="mb-3" controlId="formGroupEmail">
            <Form.Label>Email</Form.Label>
            <Form.Control
              name="login"
              type="text"
              placeholder="Digite seu Login"
              required
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formGroupSenha">
            <Form.Label>Senha</Form.Label>
            <Form.Control
              name="senha"
              type="password"
              placeholder="Password"
              required
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formGroupCheckbox">
            <Form.Check type="checkbox" label="Não sou robô" required />
          </Form.Group>

          <Button
            className="d-block mx-auto px-5"
            variant="crimson"
            type="submit"
          >
            Entrar
          </Button>
        </Form>
        <p
          style={{
            textAlign: "center",
            fontSize: "12px",
            paddingTop: "24px",
          }}
        >
          Desenvolvido por Handerson Gleber de Lima Cavalcanti
        </p>
      </Card>
    </Container>
  );
}
