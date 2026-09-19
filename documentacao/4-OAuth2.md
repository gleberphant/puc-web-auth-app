# OAuth 2.0: acesso de uma aplicação parceira

Uma aplicação parceira poderia acessar a API por meio do fluxo Authorization Code do OAuth 2.0. Nesse fluxo, a aplicação não recebe nem precisa armazenar a senha do usuário.

## Concessão de acesso

1. A aplicação parceira redireciona o usuário para o servidor de autorização.
2. O usuário faz login e visualiza quais dados e operações serão liberados.
3. O usuário autoriza ou recusa o acesso solicitado.
4. Após a autorização, o servidor redireciona o usuário de volta para a aplicação parceira com um código temporário.
5. A aplicação parceira troca esse código por um token de acesso usando uma comunicação segura com o servidor de autorização.

O acesso deve ser concedido com escopos, limitando as permissões ao mínimo necessário. Por exemplo, uma aplicação pode receber permissão apenas para consultar os dados do usuário, sem poder editar ou excluir usuários.

## Utilização de tokens

Para acessar um endpoint protegido, a aplicação parceira envia o token no cabeçalho HTTP `Authorization`:

```http
Authorization: <token_de_acesso>
```

A API valida a assinatura, a validade e os escopos do token antes de liberar o recurso. Se o token estiver expirado, inválido ou não possuir a permissão necessária, a API deve negar a solicitação.

## Benefícios

- **Maior segurança:** a senha do usuário não é compartilhada com a aplicação parceira.
- **Delegação de permissões:** o usuário decide quais recursos serão acessados e a aplicação recebe somente os escopos autorizados.
- **Revogação de acesso:** o usuário ou o servidor pode invalidar a autorização sem precisar alterar a senha principal.
- **Validade controlada:** tokens de curta duração reduzem o impacto de um eventual vazamento.
- **Integração padronizada:** aplicações diferentes podem integrar-se usando um protocolo conhecido e amplamente adotado.

> O OAuth 2.0 é apresentado neste projeto como conceito de integração. Sua implementação não faz parte da aplicação atual.
