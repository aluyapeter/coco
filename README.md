# 🚀 Task API - Go REST API with PostgreSQL

A production-ready RESTful API for task management built with Go, Gin framework, PostgreSQL, and GORM ORM.

## ✨ Features

- ✅ **Full CRUD Operations** - Create, Read, Update, Delete tasks
- 🔍 **Advanced Filtering** - Filter by completion status
- 📊 **Task Statistics** - Get insights on task completion
- 🗃️ **Soft Deletes** - Preserve data with soft deletion
- 🔒 **Environment Configuration** - Secure credential management
- 🏗️ **Clean Architecture** - Separated concerns (handlers, repository, models)
- ⚡ **High Performance** - Built with Gin framework
- 🐳 **Docker Support** - Containerized database setup
- 📝 **Auto-migration** - Database schema managed by GORM
- 🌐 **CORS Support** - Ready for frontend integration

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Web Framework**: Gin
- **Database**: PostgreSQL 15+
- **ORM**: GORM
- **Containerization**: Docker
- **Environment**: dotenv

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker & Docker Compose
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/aluyapeter/coco.git
cd coco
```

### 2. Set Up Environment

```bash
# Copy environment template
cp .env.example .env

# Edit with your configuration
nano .env
```

### 3. Start Database

```bash
# Start PostgreSQL with Docker
docker run --name taskapi-postgres \
  -e POSTGRES_USER=taskapi_user \
  -e POSTGRES_PASSWORD=your_password \
  -e POSTGRES_DB=coco_db \
  -p 5432:5432 \
  -d postgres:15
```

### 4. Install Dependencies & Run

```bash
# Install Go dependencies
go mod download

# Run the application
go run main.go
```

🎉 **API is now running at** `http://localhost:9090`

## 📖 API Documentation

### Base URL

```
http://localhost:9090/api/v1
```

### Endpoints

| Method   | Endpoint       | Description         |
| -------- | -------------- | ------------------- |
| `GET`    | `/health`      | Health check        |
| `POST`   | `/tasks`       | Create a new task   |
| `GET`    | `/tasks`       | Get all tasks       |
| `GET`    | `/tasks/:id`   | Get task by ID      |
| `PUT`    | `/tasks/:id`   | Update task         |
| `DELETE` | `/tasks/:id`   | Delete task         |
| `GET`    | `/tasks/stats` | Get task statistics |

### Example Requests

#### Create Task

```bash
curl -X POST http://localhost:9000/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Master Go programming language"
  }'
```

#### Get All Tasks

```bash
curl http://localhost:9090/api/v1/tasks
```

#### Filter Tasks

```bash
# Get completed tasks
curl "http://localhost:9090/api/v1/tasks?status=completed"

# Get pending tasks
curl "http://localhost:9090/api/v1/tasks?status=pending"
```

#### Update Task

```bash
curl -X PUT http://localhost:9090/api/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```

## 🔧 Configuration

### Environment Variables

| Variable      | Description       | Default   | Required |
| ------------- | ----------------- | --------- | -------- |
| `DB_HOST`     | Database host     | localhost | No       |
| `DB_PORT`     | Database port     | 5432      | No       |
| `DB_USER`     | Database username | -         | **Yes**  |
| `DB_PASSWORD` | Database password | -         | **Yes**  |
| `DB_NAME`     | Database name     | coco_db   | No       |
| `SERVER_PORT` | API server port   | 9090      | No       |

### Example .env

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=taskapi_user
DB_PASSWORD=secure_password_123
DB_NAME=coco_db
DB_SSLMODE=disable
SERVER_PORT=9090
APP_ENV=development
```

## 🧪 Testing

```bash
# Test health endpoint
curl http://localhost:9090/health

# Run all tests
go test ./...

# Test with coverage
go test -cover ./...
```

## 📊 Task Model

```go
type Task struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 🔒 Security Features

- Environment-based configuration
- Input validation
- SQL injection prevention (GORM)
- CORS configuration
- Soft deletes for data preservation

## 🚀 Deployment

### Docker Deployment

```bash
# Build image
docker build -t coco.

# Run container
docker run -p 9090:9090 coco
```

### Production Considerations

- Set `GIN_MODE=release`
- Use strong database passwords
- Configure proper CORS origins
- Set up SSL/TLS
- Use environment variables for secrets

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request
