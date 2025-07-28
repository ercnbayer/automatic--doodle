# Automatic Doodle Backend

Automatic Doodle is a backend service designed to help people find daily simple jobs with guaranteed payment. It provides a robust API for job listings, applications, messaging, and user management.

## Features
- User registration and authentication (JWT-based)
- Job posting and application system
- Messaging between users
- Profile management
- File uploads (S3-compatible)
- Admin and customer roles

## Tech Stack
- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **Cache:** Redis
- **Object Storage:** S3 (via Localstack for development)
- **Containerization:** Docker Compose

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) 1.20+
- [Docker](https://www.docker.com/get-started)
- [Git](https://git-scm.com/)

### Installation
Clone the repository:
```bash
git clone git@github.com:Wodemy-Labs/automatic-doodle.git
cd automatic-doodle
```

### Environment Setup
Copy the sample environment file and edit as needed:
```bash
cp sample.env .env
```
Edit `.env` to match your local setup. See `sample.env` for all required variables.

### Start Dependencies (Postgres, Redis, Localstack)
```bash
docker compose -f ./dev/docker-compose.yaml up -d
```

### Run the Application
```bash
go run main.go
```

### Running Tests
```bash
go run test ./... -v
```

## Logging
This project uses [Crawl](https://github.com/Wodemy-Labs/crawl) for logging. Set up your `.env` accordingly.

## Contributing
Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License
[MIT](LICENSE)

