# PERFIS DOS USUÁRIOS

(1) Administrador: Possui acesso total ao sistema. P
(2) Operador: Possui acesso intermediário.
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
