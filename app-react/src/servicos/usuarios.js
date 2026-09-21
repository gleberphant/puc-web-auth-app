import { CheckToken } from "./autenticacao";

//const HOST_URL = "http://localhost:4000";
const HOST_URL = "https://api-go-363261584146.us-central1.run.app/";

const ENDPOINT = {
  CRIAR: { URL: () => `${HOST_URL}/usuarios`, METODO: "POST" },
  LISTAR: { URL: () => `${HOST_URL}/usuarios`, METODO: "GET" },
  EXIBIR: { URL: (uid) => `${HOST_URL}/usuarios/${uid}`, METODO: "GET" },
  EDITAR: { URL: (uid) => `${HOST_URL}/usuarios/${uid}`, METODO: "PUT" },
  EXCLUIR: { URL: (uid) => `${HOST_URL}/usuarios/${uid}`, METODO: "DELETE" },
};

// criar usuario
export async function CriarUsuario(novoUsuario) {
  try {
    console.info("Criando  novo usuario ", novoUsuario);

    const token = CheckToken();

    if (!token) {
      throw new Error("Token inválido");
    }

    const resposta = await fetch(ENDPOINT.CRIAR.URL(), {
      method: ENDPOINT.CRIAR.METODO,
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify(novoUsuario),
    });

    const responseBody = await resposta.json();

    if (!resposta.ok) {
      throw new Error(
        `${resposta.status} ${resposta.statusText}: ${
          responseBody.error ?? "Erro na requisição"
        }`,
      );
    }

    console.info("Response body: ", responseBody);

    return [resposta.status, null];
  } catch (err) {
    console.error("catch", err);
    return [null, err];
  }
}
// listar usuario
export async function ListarUsuarios() {
  try {
    console.info("Listanto Usuarios");

    const token = CheckToken();

    if (!token) {
      throw new Error("Token inválido");
    }

    const resposta = await fetch(ENDPOINT.LISTAR.URL(), {
      method: ENDPOINT.LISTAR.METODO,
      headers: {
        Authorization: token,
      },
    });

    const responseBody = await resposta.json();

    if (!resposta.ok) {
      throw new Error(
        `${resposta.status} ${resposta.statusText}: ${
          responseBody.error ?? "Erro na requisição"
        }`,
      );
    }

    console.info("Response body: ", responseBody);

    return [[...responseBody.usuarios], null];
  } catch (err) {
    console.error("Error na requisição ", err);

    return [null, err];
  }
}

// editar usuario
export async function EditarUsuario(novoUsuario) {
  try {
    console.info("Editando o usuario usuario ", novoUsuario);

    const token = CheckToken();

    if (!token) {
      throw new Error("Token inválido");
    }

    const resposta = await fetch(ENDPOINT.EDITAR.URL(novoUsuario.Uid), {
      method: ENDPOINT.EDITAR.METODO,
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
      body: JSON.stringify(novoUsuario),
    });

    const responseBody = await resposta.json();

    if (!resposta.ok) {
      throw new Error(
        `${resposta.status} ${resposta.statusText}: ${
          responseBody.error ?? "Erro na requisição"
        }`,
      );
    }

    console.info("Response body: ", responseBody);

    return [resposta.status, null];
  } catch (err) {
    console.error("catch", err);
    return [null, err];
  }
}

// deletar usuario
export async function ExcluirUsuario(uid) {
  try {
    const token = CheckToken();

    if (!token) {
      throw new Error("Token inválido");
    }

    const resposta = await fetch(ENDPOINT.EXCLUIR.URL(uid), {
      method: ENDPOINT.EXCLUIR.METODO,
      headers: {
        Authorization: token,
      },
    });

    const responseBody = await resposta.json();

    if (!resposta.ok) {
      throw new Error(
        `${resposta.status} ${resposta.statusText}: ${
          responseBody.error ?? "Erro na requisição"
        }`,
      );
    }

    return [resposta.status, null];
  } catch (erro) {
    console.error("Erro ao excluir usuário:", erro);
    return [null, erro];
  }
}
