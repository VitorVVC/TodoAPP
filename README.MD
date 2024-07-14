Claro! Aqui está o conteúdo formatado em Markdown para que você possa copiar diretamente:

# API PostgreSQL Project

Este projeto é uma API construída em Go que interage com um banco de dados PostgreSQL. O frontend é uma aplicação React que consome essa API.

## Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Configuração do Projeto

### 1. Clonar o Repositório

Clone o repositório para a sua máquina local:

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio/api-postgresql
```

### 2. Configurar Variáveis de Ambiente

Crie um arquivo `.env` no diretório raiz do projeto `api-postgresql` com o conteúdo abaixo:

```env
DATABASE_URL=postgresql://usuario:senha@endereco-do-banco:porta/nome-do-banco?sslmode=require
API_PORT=4000
POSTGRES_HOST=your_host
POSTGRES_PORT=your_postgres_port
POSTGRES_USER=your_user
POSTGRES_PASS=your_pass
POSTGRES_NAME=your_name
```

### 3. Construir e Iniciar os Contêineres

Use o Docker Compose para construir e iniciar os contêineres:

```bash
docker-compose up --build
```

### 4. Acessar a Aplicação

- **Backend:** A API estará disponível em `http://localhost:9000`
- **Frontend:** A aplicação React estará disponível em `http://localhost:3000`

## Comandos Úteis

### Parar os Contêineres

```bash
docker-compose down
```

### Ver Logs

```bash
docker-compose logs -f
```

### Remover Contêineres, Redes e Volumes

```bash
docker-compose down -v
```

## Contribuição

Contribuições são bem-vindas! Por favor, envie um pull request ou abra uma issue para discutir o que você gostaria de mudar.

