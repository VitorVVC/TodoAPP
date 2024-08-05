# API PostgreSQL Project

This README is in English, but it is available in other languages: [PT-BR](docs/readmePTBR.md)

This project is an API built in Go that interacts with a PostgreSQL database. The frontend is a React application that consumes this API.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Project Setup

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository/api-postgresql
```

### 2. Set Up Environment Variables

Create a .env file in the root directory of the api-postgresql project with the following content:

```env
DATABASE_URL=postgresql://user:password@db-host:port/db-name?sslmode=require
API_PORT=4000
POSTGRES_HOST=your_host
POSTGRES_PORT=your_postgres_port
POSTGRES_USER=your_user
POSTGRES_PASS=your_pass
POSTGRES_NAME=your_name
```

### 3. Build and Start the Containers

Use Docker Compose to build and start the containers:

```bash
docker-compose up --build
```

### 4. Access the Application

- **Backend**: The API will be available at http://localhost:9000
- **Frontend**: The React application will be available at http://localhost:3000

## Useful Commands

### Stop the Containers

```bash
docker-compose down
```

### View Logs

```bash
docker-compose logs -f
```

### Remove Containers, Networks, and Volumes

```bash
docker-compose down -v
```

## Contribution

Contributions are welcome! Please submit a pull request or open an issue to discuss what you would like to change.