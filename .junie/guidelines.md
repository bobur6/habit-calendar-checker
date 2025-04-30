# Project Guidelines for Junie

## Project Overview
This is a Habits Tracker API built with Go, designed to help users track their daily habits and routines. The application follows a clean architecture pattern and provides a RESTful API for client applications.

### Key Features
- User authentication (register, login, profile management)
- Habit list management
- Habit tracking with customizable emojis
- Daily habit checks
- Role-based access control
- PostgreSQL database with migrations
- Docker support

## Project Structure
The project follows a clean architecture pattern with the following structure:
```
.
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── auth/                   # Authentication logic
│   ├── db/                     # Database connection and migrations
│   │   └── migrations/         # SQL migration files
│   ├── delivery/               # HTTP handlers
│   │   └── http/
│   │       └── handlers/
│   ├── middleware/             # HTTP middleware
│   ├── models/                 # Domain models
│   ├── repository/             # Data access layer
│   ├── routes/                 # API routes
│   └── services/               # Business logic
├── docker-compose.yml          # Docker Compose configuration
├── Dockerfile                  # Docker build instructions
├── go.mod                      # Go module definition
└── go.sum                      # Go module checksums
```

## Testing Guidelines
When working with this project, Junie should:

1. Run tests to verify the correctness of proposed solutions:
   ```bash
   go test ./...
   ```

2. Ensure that any new code includes appropriate test coverage.

3. For API endpoint changes, verify functionality by testing the endpoint with appropriate requests.

## Building and Running the Project
Before submitting changes, Junie should:

1. Ensure the project builds successfully:
   ```bash
   go build -o app ./cmd/main.go
   ```

2. Verify that the application runs correctly with Docker:
   ```bash
   docker-compose up --build
   ```

3. Test the affected API endpoints to ensure they work as expected.

## Code Style Guidelines
When modifying code, Junie should follow these guidelines:

1. Follow Go's official style guide and conventions:
   - Use `gofmt` or `goimports` to format code
   - Follow the naming conventions (camelCase for private, PascalCase for public)
   - Keep functions small and focused on a single responsibility

2. Maintain the existing architecture pattern:
   - Models define the domain entities
   - Repositories handle data access
   - Services implement business logic
   - Handlers manage HTTP request/response

3. Error handling:
   - Always check errors and return them appropriately
   - Use meaningful error messages
   - Avoid panic in production code

4. Documentation:
   - Add comments to exported functions and types
   - Include examples where appropriate
   - Update README.md if adding new features

## Database Changes
When making changes that affect the database:

1. Create appropriate migration files in `internal/db/migrations/`
2. Test migrations both up and down
3. Ensure backward compatibility where possible

## Security Considerations
When implementing changes, Junie should:

1. Ensure all user inputs are properly validated
2. Use prepared statements for database queries to prevent SQL injection
3. Maintain proper authentication and authorization checks
4. Keep sensitive information in environment variables, not in code