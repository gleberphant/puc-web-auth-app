# Rotas

## Rotas - DE USUÁRIOS

| Método | Endpoint        | Finalidade      | Resposta    |
| ------ | --------------- | --------------- | ----------- |
| GET    | /usuarios       | Listar usuários | 200 OK      |
| GET    | /usuarios/{uid} | Exibir usuário  | 200 OK      |
| POST   | /usuarios       | Criar usuário   | 201 Created |
| PUT    | /usuarios/{uid} | Editar usuário  | 200 OK      |
| DELETE | /usuarios/{uid} | Excluir usuário | 200 OK      |

## Rotas - DE AUTENTICAÇÃO

| Método | Endpoint | Finalidade           | Resposta |
| ------ | -------- | -------------------- | -------- |
| POST   | /login   | Fazer login          | 200 OK   |
| GET    | /login   | Exibi usuario logado | 200 OK   |

## Rotas - DE INFORMAÇÃO

| Método | Endpoint | Finalidade             | Resposta |
| ------ | -------- | ---------------------- | -------- |
| POST   | /        | Pagina sobre aplicação | 200 OK   |
| POST   | /sobre   | Pagina Sobre o autor   | 200 OK   |
