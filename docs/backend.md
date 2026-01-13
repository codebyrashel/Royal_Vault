# Backend Architecture

Framework: Go Fiber

## Structure
```
server/
 ├── config/     # DB configuration
 ├── controllers # Business logic
 ├── models      # Database models
 ├── routes      # API routes
 └── main.go     # App entry point

## API

Base URL:
http://localhost:5000/api

Health:
GET /

Auth:
POST /api/auth/register
POST /api/auth/login
```