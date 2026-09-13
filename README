<div align="center">

![Logo da PUC-PR](https://www.pucpr.br/wp-content/themes/pucpr/_assets/images/logo-pucpr-vermelha.svg)

# PUC Web Auth App

Aplicacao web para gerenciamento de usuarios, autenticacao e autorizacao baseada em perfis.

</div>

## Sumario

- [Objetivo](#objetivo)
- [Arquitetura](#arquitetura)
- [Tecnologias utilizadas](#tecnologias-utilizadas)
- [Pre-requisitos](#pre-requisitos)
- [Instalacao](#instalacao)
- [Como executar](#como-executar)
- [Como testar](#como-testar)
- [Documentacao](#documentacao)
- [Seguranca](#seguranca)
- [Autor](#autor)

## Objetivo

Este projeto foi desenvolvido como atividade pratica da disciplina de DevOps e Desenvolvimento Web do curso de Analise e Desenvolvimento de Sistemas da PUC-PR. Seu objetivo e demonstrar uma aplicacao web completa com:

- API REST para cadastro, consulta, edicao e exclusao de usuarios;
- interface web para utilizar a API;
- autenticacao de usuarios com login e senha;
- autenticacao baseada em JWT;
- autorizacao por perfis usando RBAC;
- praticas de seguranca, como hash de senhas, expiracao de tokens e controle de tentativas de login.

## Arquitetura

O projeto possui dois componentes principais:

- `api-go/`: backend em Go, responsavel pela API REST, autenticacao, autorizacao e regras de negocio;
- `app-react/`: frontend em React e Vite, responsavel pelas telas de login e gerenciamento de usuarios.

Por padrao, o backend executa em `http://localhost:4000` e o frontend de desenvolvimento em `http://localhost:5173`.

O repositorio de usuarios atual e um mock em memoria. Portanto, os dados cadastrados sao perdidos quando o backend e reiniciado.

## Tecnologias utilizadas

### Backend

- Go `1.26.6`;
- `net/http` para o servidor HTTP;
- JWT com `github.com/golang-jwt/jwt/v5`;
- bcrypt com `golang.org/x/crypto/bcrypt`;
- UUID com `github.com/google/uuid`;
- CORS e middlewares HTTP.

### Frontend

- React `19`;
- Vite `8`;
- React Router;
- React Bootstrap e Bootstrap Icons;
- ESLint.

### Desenvolvimento e testes

- Docker;
- Bruno ou outra ferramenta HTTP compativel com os arquivos de `test-api/`;
- scripts Bash e arquivos de configuracao do VS Code.

## Pre-requisitos

Instale os seguintes programas:

- Go `1.26.6` ou compativel;
- Node.js compativel com o frontend e npm;
- Docker, caso deseje executar em container;
- Bruno, opcionalmente, para executar a colecao de requisicoes HTTP.

Confira as instalacoes:

```bash
go version
node --version
npm --version
docker --version
```

## Instalacao

### Backend Go

```bash
cd api-go
go mod download
```

### Frontend React

Em outro terminal, a partir da raiz do projeto:

```bash
cd app-react
npm install
```

## Como executar

### Executar o backend

No diretorio `api-go`:

```bash
go run .
```

O backend ficara disponivel em `http://localhost:4000`.

### Executar o frontend

Em outro terminal, no diretorio `app-react`:

```bash
npm run dev
```

Abra no navegador o endereco informado pelo Vite, normalmente `http://localhost:5173`.

### Executar com Docker

Para construir e executar a imagem da API:

```bash
cd api-go
docker build -t puc-react-api:dev .
docker run --rm -p 4000:4000 puc-react-api:dev
```

O arquivo `executar.sh` tambem documenta a inicializacao do backend em Docker.

## Como testar

### Validar o backend

No diretorio `api-go`:

```bash
go test ./...
go vet ./...
```

Atualmente o projeto nao possui testes automatizados dedicados, mas `go test ./...` verifica a compilacao de todos os pacotes.

### Validar o frontend

No diretorio `app-react`:

```bash
npm run lint
npm run build
```

`npm run lint` verifica os padroes do codigo e `npm run build` confirma que o frontend pode ser compilado para producao.

### Testar a API manualmente

Os arquivos de requisicao HTTP estao em `test-api/` e podem ser executados pelo Bruno. Eles cobrem login e operacoes de usuarios, incluindo criacao, consulta, edicao e exclusao.

Tambem e possivel verificar a API diretamente:

```bash
curl http://localhost:4000
```

## Documentacao

A pasta [`documentacao/`](documentacao/) contem o detalhamento da aplicacao:

- [Visao geral](documentacao/0-visao-geral.md): objetivos, requisitos funcionais e criterios da atividade;
- [Modelo da API](documentacao/1-modelo-api.md): rotas, metodos HTTP e respostas esperadas;
- [Seguranca com JWT](documentacao/2-seguranca-jwt.md): login, claims, assinatura e expiracao de tokens;
- [Controle RBAC](documentacao/3-controle-rbac.MD): perfis e permissoes por rota e metodo;
- [OAuth 2.0](documentacao/4-OAuth2.md): proposta de integracao com aplicacoes parceiras;
- [Analise de seguranca](documentacao/5-analise-seguranca.md): riscos e medidas de mitigacao.

O OAuth 2.0 esta documentado apenas como proposta conceitual e nao e implementado na aplicacao atual.

## Seguranca

O projeto utiliza:

- bcrypt para armazenar senhas como hash;
- JWT assinado, com `uid`, `login`, `perfil`, `iat` e `exp`;
- expiracao do token em uma hora;
- RBAC para autorizar operacoes conforme o perfil;
- limite de tentativas por IP e por conta, atraso progressivo, logs de tentativas suspeitas e bloqueio temporario;
- respostas genericas para falhas de login, sem registrar senhas.

Em producao, a chave `JWTSECRET` deve ser configurada por variavel de ambiente. O controle de tentativas tambem deve usar um armazenamento compartilhado, como Redis, quando houver mais de uma instancia da API.

## Autor

HANDERSON GLEBER DE LIMA CAVALCANTI (GRAVATINHA)

E-mail: handerson.gleber@gmail.com

[Instagram](https://www.instagram.com/handersongleber/)

## Licenca

Este projeto esta sob licenca livre.
