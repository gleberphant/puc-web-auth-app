# Tarefas

# Requisitos funcionais - backend

[x] Cadastro de usuários: Permitir o registro de novos usuários contendo informações como: Nome; E-mail; Senha; Perfil de acesso.
[x] Consulta de usuários: Permitir listar os usuários cadastrados. Por exemplo: Listar todos os usuários; consultar um usuário específico.
[x] Atualização de usuários: Permitir alterar informações de um usuário já cadastrado. Por exemplo: Nome; E-mail; Perfil.
[x] Exclusão de usuários: Permitir remover usuários do sistema.
[x] Login: Criar um endpoint responsável pela autenticação dos usuários. O login deverá validar e-mail; senha. Quando as credenciais forem válidas, a API deverá gerar um token JWT.

# Requisitos funcionais - frontend

Além da API, desenvolva uma interface web simples para demonstrar o funcionamento da solução;O objetivo do front-end é permitir que o professor visualize e teste as funcionalidades desenvolvidas. A interface pode conter:

[x]Tela de login;
[x]Listagem de usuários;
[x]Cadastro de usuários;
[x]Edição de usuários;
[x]Exclusão de usuários;
[x]Exibição das respostas da API.

# requisitos da documentação

[x]Parte 1 – Modelagem da API
[x]Parte 2 – Segurança com JWT
[x]Parte 3 – Controle de acesso (RBAC)
[x]Parte 4 – OAuth 2.0.;Observação: não é necessário implementar OAuth 2.0. Basta explicar seu funcionamento no contexto da solução proposta;
[x]Parte 5 – Análise de Segurança

# requisitos gerais

[x] Funcionamento da autenticação com JWT;
[_] Correta aplicação dos princípios REST;
[_] Implementação dos endpoints solicitados;
[_] Controle de acesso por perfis (RBAC);
[_] Aplicação de boas práticas de segurança;
[_] Clareza da documentação;
[_] Funcionamento geral da solução;
[_] Capacidade de justificar as decisões adotadas.

## tarefas documentação

### Parte 1 – Modelagem da API

[x] No relatório ou documentação do projeto, apresente pelo menos quatro endpoints REST utilizados na solução.

### Parte 2 – Segurança com JWT

Sua API deverá utilizar autenticação baseada em JWT (JSON Web Token). Na documentação explique:

[x]Processo de login: Como o usuário envia suas credenciais.
[x]Geração do token: Como o sistema cria o JWT após autenticação.
[x]Informações armazenadas no token: Por exemplo ID do usuário; Nome; Perfil; Data de emissão.
[x]Política de expiração: Informe o tempo de validade do token. Por exemplo 30 minutos; 1 hora; 24 horas.

### Parte 3 – Controle de acesso (RBAC)

Os endpoints devem ser protegidos por regras de autorização. Isso significa que nem todos os usuários poderão acessar todas as funcionalidad; A aplicação deverá possuir três perfis de usuário:

[x] Administrador: Possui acesso total ao sistema. Por exemplo: Criar usuários; editar usuários; excluir usuários; consultar usuários.
[x] Operador: Possui acesso intermediário. Por exemplo: Consultar usuários; atualizar informações.
[x] Cliente: Possui acesso restrito. Por exemplo: Visualizar apenas seus próprios dados.

### Parte 4 – OAuth 2.0

Na documentação, explique como uma aplicação parceira poderia acessar sua API utilizando OAuth 2.0; Observação: não é necessário implementar OAuth 2.0. Basta explicar seu funcionamento no contexto da solução proposta;Sua explicação deve contemplar:

[] Concessão de acesso: Como o usuário autoriza o acesso.
[] Utilização de tokens: Como o token é utilizado para acessar recursos protegidos.
[] Benefícios: Explique vantagens como maior segurança; delegação de permissões; não compartilhamento de senhas.

### Parte 5 – Análise de Segurança

[] Identifique pelo menos três riscos de segurança relacionados à API. Para cada risco apresente uma medida de mitigação.

## tarefas do sistema

### backend

[x] Cadastro de usuários: Permitir o registro de novos usuários contendo informações como: Nome; E-mail; Senha; Perfil de acesso.
[x] Consulta de usuários: Permitir listar os usuários cadastrados. Por exemplo: Listar todos os usuários; consultar um usuário específico.
[x] Atualização de usuários: Permitir alterar informações de um usuário já cadastrado. Por exemplo: Nome; E-mail; Perfil.
[x] Exclusão de usuários: Permitir remover usuários do sistema.
[x] Login: Criar um endpoint responsável pela autenticação dos usuários. O login deverá validar e-mail; senha. Quando as credenciais forem válidas, a API deverá gerar um token JWT.

### frontend

Além da API, desenvolva uma interface web simples para demonstrar o funcionamento da solução;O objetivo do front-end é permitir que o professor visualize e teste as funcionalidades desenvolvidas. A interface pode conter:

[x]Tela de login;
[x]Listagem de usuários;
[x]Cadastro de usuários;
[x]Edição de usuários;
[x]Exclusão de usuários;
[x]Exibição das respostas da API.
[x] Cadastro de usuários
[x] Listar todos os usuários;
[x] alterar informações de um usuário já cadastrado
[x] Exclusão de usuários
[x] Login: endpoint responsável pela autenticação dos usuários. O login deverá validar e-mail; senha. Quando as credenciais forem válidas, a API deverá gerar um token JWT.
[x] autenticação com jwt no backend
[x] front end em react;
[x] formulário de login html
[x] telas de listagem , criação e exclusao de usuarios
[x] edição do usuario
[x] Funcionamento da autenticação com JWT;
[x] Implementação dos endpoints solicitados;

## PERFIS DE SISTEMA

[x] ADM: acesos total ao sistema
[x] OPERADOR: acesso intermediarios; apenas listar os usuarios; nao poderá criar OU editar outtros usuarios
[x] CLIENTE: acesso restrito, apenas consultar dados do seu usuario; não podera listar, criar ou editar outros usuarios

## tarefas requisitos atividade

[x] Controle de acesso por perfis (RBAC);
[_] consultar um usuário específico; filtro na tela de listagem de usuario
[_] Justificar as decisões adotadas.
[_] Documentação clara e completa;
[_] Funcionamento geral da solução;
[ ] integração OAtuh2
[_] verificar a aplicação dos princípios REST;
[_] Aplicação de boas práticas de segurança;
[ ] configura git actions e conteiners

# O envio deverá conter:

[] Código-fonte do projeto: Arquivos do back-end e do front-end.
[] Documentação da API incluindo: Endpoints; Métodos HTTP; Códigos de resposta; Perfis de acesso; JWT; OAuth 2.0; Análise de segurança.
[] Evidências de funcionamento apresentadas por meio de: Capturas de tela; Vídeo de demonstração; Link do projeto (quando disponível).
[] Arquivo README, que deve explicar: Objetivo do projeto; Tecnologias utilizadas; Como instalar; Como executar; Como testar a aplicação.

Não é necessário publicar a aplicação na internet. Ela pode ser executada localmente e demonstrada por meio de capturas de tela, vídeo ou documentação.

---

# relatorio de desenvolvimento

o crud esta concluido; falta implementar consultar um usuario específico por meiod e filtros na tela de listagem
