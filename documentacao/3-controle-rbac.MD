# PERFIS DOS USUÁRIOS

(1) Administrador: Possui acesso total ao sistema. Por exemplo: Criar usuários; editar usuários; excluir usuários; consultar usuários.

(2) Operador: Possui acesso intermediário. Por exemplo: Consultar usuários; atualizar informações.

(3) Cliente: Possui acesso restrito. Por exemplo: Visualizar apenas seus próprios dados.

# Mapa de permissões

| rota        | método | perfis                        |
| :---------- | :----- | :---------------------------- |
| `/usuarios` | GET    | `admin`, `usuario`            |
| `/usuarios` | POST   | `admin`                       |
| `/usuarios` | PUT    | `admin`                       |
| `/usuarios` | DELETE | `admin`                       |
| `/login`    | GET    | `admin`, `usuario`, `cliente` |
| `/login`    | POST   | `admin`, `usuario`, `cliente` |
| `/login`    | PUT    | `admin`, `usuario`, `cliente` |
| `/login`    | DELETE | `admin`, `usuario`, `cliente` |
