# Processo de login: Como o usuário envia suas credenciais?

- Ao acessar a aplicação, o frontend verifica se existe um token JWT salvo no localStorage;
- se o token não existir, a aplicação renderiza a tela de login;
- o usuário informa seu login e senha no formulário;
- o frontend envia essas credenciais para a rota /login do backend em uma requisição HTTP do tipo POST;
- os dados são enviados em formato JSON, normalmente com os campos login e senha.

Esse processo garante que as credenciais não sejam enviadas em query string nem em parâmetros visíveis na URL, reduzindo risco de exposição em logs ou em histórico de navegação.

# Geração do token: Como o sistema cria o JWT após autenticação?

- o backend recebe o JSON com as credenciais;
- verifica se o usuário existe e se a senha informada confere com a registrada no sistema;
- se a autenticação for válida, o backend cria um objeto JWT com claims relevantes para a sessão, como uid, login e perfil;
- o token é assinado com uma chave secreta, garantindo que ele não possa ser alterado sem que a assinatura seja invalidada;
- o servidor retorna o token ao cliente junto com os dados do usuário autenticado.

No projeto atual, essa etapa acontece na rota de login do Go, e a validação de autenticação em seguida é feita pelo middleware de autorização, que verifica a assinatura e o conteúdo do token antes de liberar o acesso às rotas protegidas.

# Quais Informações armazenadas no token?

- uid;
- login;
- perfil;
- data de emissão (iat);
- data de expiração (exp).

Essas informações permitem que o servidor identifique o usuário e o perfil de acesso em requisições posteriores, sem precisar reconsultar o banco a cada chamada. No entanto, o token deve guardar apenas dados mínimos e necessários para autenticação e autorização.

# Política de expiração: Informe o tempo de validade do token. ?

- 1 hora.

# Justifique sua escolha.

A expiração em 1 hora é uma política equilibrada para aplicações web e APIs em ambiente de desenvolvimento e estudo. Ela reduz o tempo de risco caso o token seja roubado ou vazado, sem prejudicar a experiência do usuário com sessões que expiram muito rapidamente.

Além disso, a expiração atua como uma camada de proteção adicional: mesmo que um token seja comprometido, ele perde validade após um intervalo curto. Esse comportamento é importante porque o JWT, por si só, não garante proteção contra roubo; ele precisa ser combinado com boas práticas de assinatura, uso de chave secreta, validação de claims e controle de autorização.

No projeto atual, essa escolha é funcional e simples, mas ainda pode ser melhorada com:

- uso de secret em variável de ambiente;
- armazenamento seguro da senha com hash;
- inclusão de mais claims apenas quando realmente necessários;
- política de refresh token em sistemas mais robustos.
