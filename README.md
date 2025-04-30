# Habit Calendar Checker API

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-Ready-blue?logo=docker)
![Postgres](https://img.shields.io/badge/Postgres-15+-336791?logo=postgresql)

## Overview
A production-ready RESTful API for tracking habits and daily routines. Built with Go, GORM, PostgreSQL, JWT authentication, role-based access, and full Docker support. Includes ready-to-use Postman Collection for easy API testing.

---

## Features
- User registration, login, JWT authentication
- Full CRUD for habits, habit lists, and daily checks
- Role-based access control (admin/user)
- Password hashing (bcrypt)
- GORM ORM with PostgreSQL
- Database migrations (auto & manual)
- Docker & Docker Compose support
- Ready Postman Collection for API testing
- Clean project structure & best practices

---

## Tech Stack
- **Go** (Gin, GORM)
- **PostgreSQL**
- **Docker / Docker Compose**
- **JWT** for authentication
- **bcrypt** for passwords

---

## Quick Start
1. **Clone the repository:**
   ```bash
   git clone https://github.com/bobur6/habit-calendar-checker.git
   cd habit-calendar-checker
   ```
2. **Configure environment:**
   - Copy `.env.example` to `.env` and edit if needed:
   ```env
   DB_HOST=postgres
   DB_PORT=5432
   DB_NAME=habits
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_SSLMODE=disable
   PORT=8080
   GIN_MODE=release
   JWT_SECRET=your-secret-key
   JWT_EXPIRATION=24h
   APP_NAME=Habits Tracker
   APP_ENV=development
   APP_DEBUG=false
   ```
3. **Run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```
4. **API available at:** `http://localhost:8080`

---

## Demo & Postman
- [HabitsTracker.postman_collection.json](./HabitsTracker.postman_collection.json) — Postman collection for API testing.
- Import the collection to Postman and test all endpoints easily.

---

## How to use API (Example)
**User registration:**
```http
POST /api/auth/register
Content-Type: application/json
{
  "username": "bobur",
  "email": "bobur@example.com",
  "password": "yourpassword"
}
```
**Response:**
```json
{
  "token": "<jwt_token>",
  "user": { ... }
}
```

**Add a habit:**
```http
POST /api/habits (with JWT in Authorization)
{
  "title": "Drink water",
  "emoji": "💧"
}
```

---

## Main Endpoints
- `POST /api/auth/register` — register
- `POST /api/auth/login` — login (get JWT)
- `GET /api/users/profile` — profile
- `POST /api/habits` — create habit
- `GET /api/habits` — get all habits
- `POST /api/habit-checks` — mark habit
- ...see Postman Collection for more

---

## Project Structure
```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── auth/
│   ├── db/
│   │   └── migrations/
│   ├── delivery/
│   │   └── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── services/
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum
```

## Security

- All endpoints except registration and login require JWT authentication
- Passwords are hashed using bcrypt
- Database connections are pooled and configured for optimal performance
- The application runs with a non-root user in Docker
- Environment variables are used for sensitive configuration

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.