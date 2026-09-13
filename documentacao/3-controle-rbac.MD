# PERFIS DOS USUÁRIOS

(1) Administrador: Possui acesso total ao sistema.
(2) Usuário: Possui acesso intermediário.
(3) Cliente: Possui acesso restrito.

# Mapa de permissões

| rota        | método | perfis permitidos             |
| :---------- | :----- | :---------------------------- |
| `/usuarios` | GET    | `admin`, `usuario`            |
| `/usuarios` | POST   | `admin`                       |
| `/usuarios` | PUT    | `admin`                       |
| `/usuarios` | DELETE | `admin`                       |
| `/login`    | GET    | `admin`, `usuario`, `cliente` |
| `/login`    | POST   | `admin`, `usuario`, `cliente` |
| `/login`    | PUT    | `admin`, `usuario`, `cliente` |
| `/login`    | DELETE | `admin`, `usuario`, `cliente` |

# Como funciona o controle RBAC

O projeto utiliza RBAC (Role-Based Access Control), ou controle de acesso baseado em perfis. Após o login, o perfil do usuário (`admin`, `usuario` ou `cliente`) é gravado no token JWT. Em cada requisição protegida, o middleware de autenticação valida a assinatura do token, extrai o perfil e consulta o mapa de permissões usando a rota e o método HTTP solicitados.

Se o token estiver ausente, inválido ou expirado, a API retorna `401 Unauthorized`. Quando o token é válido, mas o perfil não possui permissão para aquela operação, a API retorna `403 Forbidden`. A rota `/login` e as requisições `OPTIONS` são públicas; as demais passam pela verificação de perfil antes de serem encaminhadas ao manipulador correspondente.
