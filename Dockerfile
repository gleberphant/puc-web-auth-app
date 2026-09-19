# essa imagem terá um servidor openssh rodando e as dependencias para compilar e rodar aplicações em Go, Node.js e SQLite.
# para que possamos acessar o container via ssh e fausa-lo para desenvolvimento.

FROM golang:1.26-trixie

WORKDIR /app

# Instala dependências, Node.js e OpenSSH Server em camada otimizada
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    git \
    sqlite3 \
    ca-certificates \
    gnupg \
    openssh-server \
    && rm -rf /var/lib/apt/lists/* \
    && curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y --no-install-recommends nodejs \
    && rm -rf /var/lib/apt/lists/*

# Configuração do SSH: diretório de runtime, senha de root e permissões
RUN mkdir -p /var/run/sshd \
    && echo 'root:root' | chpasswd \
    && sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config \
    && sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config \
    && sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config

EXPOSE 4000 8080 22

# Script inline para iniciar o sshd e manter o container ativo com bash
CMD ["/bin/bash", "-c", "/usr/sbin/sshd && exec bash"]
    && rm -rf /var/lib/apt/lists/* \
    && curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y --no-install-recommends nodejs

EXPOSE 4000 8080 22

CMD ["bash"]


