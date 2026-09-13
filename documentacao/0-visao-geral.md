# Visão geral do trabalho

## Descrição Geral e Objetivos

Como desenvolvedor(a), sua missão será criar uma solução completa que permita cadastrar, consultar, atualizar e excluir usuários, aplicando os conceitos estudados ao longo da disciplina sobre:

- APIs REST;
- Web Services;
- Autenticação;
- Autorização;
- Controle de acesso;
- Segurança de aplicações web.

Objetivo da atividade é desenvolver uma aplicação web funcional composta por:

- Back-end: Uma API REST para gerenciamento de usuários.
- Front-end: Uma interface web simples que permita utilizar e demonstrar os recursos da API.

O projeto deverá evidenciar a aplicação prática dos conceitos de autenticação, autorização e segurança estudados na disciplina.
Sua aplicação deverá permitir que diferentes usuários realizem operações de acordo com seu perfil de acesso.

## Requisitos

### Requisitos Funcionais

A solução deverá conter as funcionalidades obrigatórias:

1. Cadastro de usuários: Permitir o registro de novos usuários contendo informações como: Nome; E-mail; Senha; Perfil de acesso.
2. Consulta de usuários: Permitir listar os usuários cadastrados. Por exemplo: Listar todos os usuários; consultar um usuário específico.
3. Atualização de usuários: Permitir alterar informações de um usuário já cadastrado. Por exemplo: Nome; E-mail; Perfil.
4. Exclusão de usuários: Permitir remover usuários do sistema.
5. Login: Criar um endpoint responsável pela autenticação dos usuários. O login deverá validar e-mail; senha. Quando as credenciais forem válidas, a API deverá gerar um token JWT.

### Requisitos Interface frontend

Além da API, desenvolva uma interface web simples para demonstrar o funcionamento da solução;O objetivo do front-end é permitir que o professor visualize e teste as funcionalidades desenvolvidas. A interface pode conter:

1. Tela de login;
2. Listagem de usuários;
3. Cadastro de usuários;
4. Edição de usuários;
5. Exclusão de usuários;
6. Exibição das respostas da API.

### Documentação

#### Parte 1 – Modelagem da API

No relatório ou documentação do projeto, apresente pelo menos quatro endpoints REST utilizados na solução.
Para cada endpoint informe: Método HTTP; URL; Finalidade; Código de resposta esperado.

Exemplo:

| Método | Endpoint       | Finalidade        | Resposta       |
| ------ | -------------- | ----------------- | -------------- |
| GET    | /usuarios      | Listar usuários   | 200 OK         |
| POST   | /usuarios      | Criar usuário     | 201 Created    |
| PUT    | /usuarios/{id} | Atualizar usuário | 200 OK         |
| DELETE | /usuarios/{id} | Excluir usuário   | 204 No Content |

#### Parte 2 – Segurança com JWT

Sua API deverá utilizar autenticação baseada em JWT (JSON Web Token). Na documentação explique:

1. Processo de login: Como o usuário envia suas credenciais.
2. Geração do token: Como o sistema cria o JWT após autenticação.
3. Informações armazenadas no token: Por exemplo ID do usuário; Nome; Perfil; Data de emissão.
4. Política de expiração: Informe o tempo de validade do token. Por exemplo 30 minutos; 1 hora; 24 horas.
5. Justifique sua escolha.

#### Parte 3 – Controle de acesso (RBAC)

A aplicação deverá possuir três perfis de usuário:

1. Administrador: Possui acesso total ao sistema. Por exemplo: Criar usuários; editar usuários; excluir usuários; consultar usuários.
2. Operador: Possui acesso intermediário. Por exemplo: Consultar usuários; atualizar informações.
3. Cliente: Possui acesso restrito. Por exemplo: Visualizar apenas seus próprios dados.

Importante! Os endpoints devem ser protegidos por regras de autorização. Isso significa que nem todos os usuários poderão acessar todas as funcionalidades.

#### Parte 4 – OAuth 2.0

Na documentação, explique como uma aplicação parceira poderia acessar sua API utilizando OAuth 2.0; Sua explicação deve contemplar:

1. Concessão de acesso: Como o usuário autoriza o acesso.
2. Utilização de tokens: Como o token é utilizado para acessar recursos protegidos.
3. Benefícios: Explique vantagens como maior segurança; delegação de permissões; não compartilhamento de senhas.

Observação: não é necessário implementar OAuth 2.0. Basta explicar seu funcionamento no contexto da solução proposta.

#### Parte 5 – Análise de Segurança

Identifique pelo menos três riscos de segurança relacionados à API. Para cada risco apresente uma medida de mitigação.

Exemplo:

| Risco                            | Solução                          |
| -------------------------------- | -------------------------------- |
| Roubo de token JWT               | Utilizar HTTPS e expiração curta |
| Senhas armazenadas em texto puro | Utilizar hash com bcrypt         |
| Acesso indevido a endpoints      | Implementar RBAC                 |

## Entrega - O que você deve entregar

O envio deverá conter:

- Código-fonte do projeto: Arquivos do back-end e do front-end.
- Documentação da API incluindo: Endpoints; Métodos HTTP; Códigos de resposta; Perfis de acesso; JWT; OAuth 2.0; Análise de segurança.
- Evidências de funcionamento apresentadas por meio de: Capturas de tela; Vídeo de demonstração; Link do projeto (quando disponível).
- Arquivo README, que deve explicar: Objetivo do projeto; Tecnologias utilizadas; Como instalar; Como executar; Como testar a aplicação.

Não é necessário publicar a aplicação na internet. Ela pode ser executada localmente e demonstrada por meio de capturas de tela, vídeo ou documentação.

# Critérios de avaliação

Utilize esse checklist para garantir que você contemplou essas solicitações:

- Correta aplicação dos princípios REST;
- Implementação dos endpoints solicitados;
- Funcionamento da autenticação com JWT;
- Controle de acesso por perfis (RBAC);
- Aplicação de boas práticas de segurança;
- Clareza da documentação;
- Funcionamento geral da solução;
- Capacidade de justificar as decisões adotadas.

Ao final da atividade, espera-se que você desenvolva e apresente uma aplicação funcional que demonstre a implementação de operações CRUD (criação, consulta, atualização e exclusão) de usuários, a autenticação por meio de JWT (JSON Web Token), o controle de acesso baseado em perfis de usuário (Administrador, Operador e Cliente) e a proteção adequada dos endpoints da aplicação. Além disso, a solução deve evidenciar a aplicação de boas práticas de segurança no desenvolvimento de APIs, bem como a compreensão dos conceitos relacionados ao protocolo OAuth 2.0 e à análise de riscos envolvidos nos processos de autenticação e autorização.
