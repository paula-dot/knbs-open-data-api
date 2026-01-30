# GitHub Copilot Instructions for KNBS Open Data API

This repository contains a full-stack application for accessing, exploring, and visualizing Kenyan national statistics. Follow these guidelines when working with the codebase.

---

## Project Overview

This is a monorepo with:
- **Backend**: Go-based REST API using Chi router, PostgreSQL, and sqlc
- **Frontend**: React + Vite application with Tailwind CSS and shadcn/ui

---

## Repository Structure

```
knbs-open-data-api/
├── backend/          # Go backend service
│   ├── cmd/          # Application entrypoints
│   │   ├── server/   # Main API server
│   │   └── seeder/   # Database seeder
│   ├── internal/     # Internal packages
│   │   ├── database/ # Generated sqlc code
│   │   ├── handlers/ # HTTP request handlers
│   │   └── services/ # Business logic
│   ├── migrations/   # Database migration files
│   ├── sqlc/         # SQL queries for sqlc
│   └── sqlc.yaml     # sqlc configuration
├── frontend/         # React frontend application
└── docker-compose.yml # Docker orchestration
```

---

## Backend Guidelines (Go)

### Architecture
- Follow **Clean Architecture** principles
- Organize code into clear layers: handlers → services → database
- Use dependency injection for testability
- Keep business logic in the `services` package

### Database & Migrations
- **Database**: PostgreSQL 16+
- **Migrations**: Managed with Goose
- **Query Generation**: Use sqlc for type-safe SQL queries
- **Connection**: Use pgx/v5 driver

When modifying the database:
1. Create migration files in `backend/migrations/`
2. Write SQL queries in `backend/sqlc/queries/`
3. Run `sqlc generate` to regenerate Go code
4. Never modify generated code in `internal/database/`

### Running the Backend
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

### Generating sqlc Code
```bash
cd backend
sqlc generate
```

### Running Migrations
```bash
cd backend
goose -dir migrations postgres "connection-string" up
```

### Coding Standards
- Use Go 1.22+ features
- Follow standard Go formatting (gofmt, goimports)
- Use Chi router for HTTP routing
- Implement middleware for cross-cutting concerns (CORS, logging, recovery)
- Use pgx/v5 for database operations
- Return proper HTTP status codes
- Structure APIs under `/api/v1/` prefix

### Dependencies
- **Router**: Chi (github.com/go-chi/chi/v5)
- **Database**: pgx/v5 (github.com/jackc/pgx/v5)
- **Configuration**: Environment variables via godotenv (github.com/joho/godotenv)
  - Note: Viper may be added in the future for advanced configuration management

---

## Frontend Guidelines (React)

### Tech Stack
- **Build Tool**: Vite
- **Framework**: React
- **Styling**: Tailwind CSS
- **Components**: shadcn/ui
- **State Management**: Zustand
- **Routing**: React Router

### Running the Frontend
```bash
cd frontend
npm install
npm run dev
```

### Coding Standards
- Use functional components with hooks
- Prefer composition over inheritance
- Use TypeScript if present, otherwise modern JavaScript (ES6+)
- Follow Tailwind CSS utility-first approach
- Use shadcn/ui components for consistency
- Implement responsive design (mobile-first)
- Keep components small and focused

### Component Organization
- Place reusable components in `src/components/`
- Place page components in `src/pages/`
- Place state management in `src/store/`
- Place utilities in `src/lib/`

---

## Docker

The project includes Docker support for full-stack deployment:

```bash
docker-compose up --build
```

This launches:
- Go backend (port 8080)
- React frontend (port 5173)
- PostgreSQL instance

---

## Testing

### Backend Testing
- Write tests for handlers and services
- Use table-driven tests where appropriate
- Mock database dependencies
- Run tests with: `go test ./...`

### Frontend Testing
- Follow existing test patterns if present
- Test user interactions and edge cases

---

## API Design

### Versioning
- All endpoints under `/api/v1/`
- Use semantic versioning for major changes

### Response Format
- Return JSON responses
- Use consistent error response structure
- Include appropriate HTTP status codes

### Endpoints
- Dataset listing and retrieval
- Clean, RESTful resource naming
- Proper use of HTTP methods (GET, POST, PUT, DELETE)

---

## Code Quality

- Keep functions small and focused
- Use descriptive variable and function names
- Add comments for complex logic, but prefer self-documenting code
- Handle errors appropriately (don't ignore them)
- Validate input data
- Use proper logging (avoid excessive logging in production)

---

## Security

- Implement JWT authentication for protected endpoints
- Validate and sanitize all user inputs
- Use parameterized queries (sqlc handles this)
- Enable CORS properly
- Use environment variables for secrets
- Never commit sensitive data (API keys, passwords)

---

## Performance

- Use database indexes appropriately
- Implement pagination for large datasets
- Cache where beneficial
- Minimize database queries (N+1 problem)
- Optimize frontend bundle size

---

## Future Roadmap

When implementing new features, consider:
- Phase 2: Search, pagination, filters, OpenAPI docs, admin dashboard, charts, GIS layers
- Phase 3: User accounts, custom dashboards, dataset exports (CSV/XLSX/JSON)

---

## Contributing

- Make small, focused commits
- Write clear commit messages
- Test changes before committing
- Update documentation when changing functionality
- Follow existing code patterns and conventions
