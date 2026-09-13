# Análise de segurança

A API possui os seguintes riscos de segurança e respectivas medidas de mitigação:

| Risco                                                | Medida de mitigação                                                                                                                                                                    |
| :--------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Senhas armazenadas ou comparadas em texto puro       | Armazenar somente hashes seguros bcrypt e comparar a senha usando uma função própria para validação de hash. Nunca salvar ou registrar a senha original.                               |
| Roubo ou interceptação do token JWT                  | Utilizar HTTPS em todas as requisições, manter o token com validade curta, validar assinatura e expiração e adotar armazenamento seguro no frontend.                                   |
| Exposição ou comprometimento da chave secreta do JWT | Manter a chave em variável de ambiente ou serviço de secrets, usar uma chave longa e aleatória, restringir seu acesso e trocá-la periodicamente.                                       |
| Ataques de força bruta contra o endpoint de login    | Aplicar limitação de tentativas por IP e por conta, usar atrasos progressivos, registrar tentativas suspeitas e, quando necessário, bloquear temporariamente a conta.                  |
| Acesso indevido a recursos de outros usuários        | Aplicar RBAC em cada endpoint e verificar também a propriedade do recurso. Um cliente, por exemplo, deve conseguir consultar apenas o próprio `uid`, mesmo que possua um token válido. |
