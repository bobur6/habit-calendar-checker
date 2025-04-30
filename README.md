# Habits Tracker API

A RESTful API for tracking habits and daily routines.

## Features

- User authentication (register, login, profile management)
- Habit list management
- Habit tracking with customizable emojis
- Daily habit checks
- Role-based access control
- PostgreSQL database with migrations
- Docker support

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- PostgreSQL 15 or later (if running without Docker)

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-rest-project.git
   cd go-rest-project
   ```

2. Create a .env file:
   ```bash
   # Database configuration
   DB_HOST=postgres
   DB_PORT=5432
   DB_NAME=habits
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_SSLMODE=disable

   # Server configuration
   PORT=8080
   GIN_MODE=release

   # JWT configuration
   JWT_SECRET=your-secret-key-change-in-production
   JWT_EXPIRATION=24h

   # Application configuration
   APP_NAME=Habits Tracker
   APP_ENV=development
   APP_DEBUG=false
   ```

3. Start the application with Docker:
   ```bash
   docker-compose up --build
   ```

The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and get JWT token

### User Management

- `GET /api/users/profile` - Get user profile
- `PUT /api/users/profile` - Update user profile
- `DELETE /api/users/profile` - Delete user account

### Habit Lists

- `POST /api/habit-lists` - Create a new habit list
- `GET /api/habit-lists` - Get all user's habit lists
- `GET /api/habit-lists/{id}` - Get a specific habit list
- `PUT /api/habit-lists/{id}` - Update a habit list
- `DELETE /api/habit-lists/{id}` - Delete a habit list

### Habits

- `POST /api/habits` - Create a new habit
- `GET /api/habits` - Get habits by list
- `GET /api/habits/{id}` - Get a specific habit
- `PUT /api/habits/{id}` - Update a habit
- `DELETE /api/habits/{id}` - Delete a habit

### Habit Checks

- `POST /api/habit-checks` - Create a habit check
- `GET /api/habit-checks` - Get habit checks
- `GET /api/habit-checks/{id}` - Get a specific habit check
- `GET /api/habit-checks/date` - Get habit checks by date
- `PUT /api/habit-checks/{id}` - Update a habit check
- `DELETE /api/habit-checks/{id}` - Delete a habit check

## Development

### Running Tests
```bash
go test ./...
```

### Running Migrations
Migrations are automatically run when the application starts. To run them manually:
```bash
go run cmd/migrate/main.go
```

### Project Structure
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