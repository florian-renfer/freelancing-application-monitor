# 📊 Freelancing Application Monitor

> [!IMPORTANT]
> This project is in a very early state and far from ready for production. So consider the project _work in progress_.

An interactive monitor that helps you keeping an eye on your applications for freelance projects.

## ✨ Technologies

- `Go`
- `SQLite`

## 🚀 Features

- Create users
- List created users

## 📍 The Process

The process and motivation need a clear description.

## 🚦 Running the Project

1. Clone the repository
2. Install dependencies: `go mod download`
3. Run development server: `go run cmd/api/main.go`
4. Execute requests against the server running on: `http://localhost:4000/v1/<api-endpoint>`

## 🎞️ Preview

For now, there's no preview - it's just API in the making.
I will add a preview as soon as the UI is on the roadmap.

## Uses Cases

### Register User

```bash
curl -X POST http://localhost:4000/v1/auth/register\
  -H "Content-Type: application/json" \
  -d '{
    "email": "max.mustermann@mail.de",
    "password": "changeme",
    "firstname": "Max",
    "lastname": "Mustermann",
    "date_of_birth": "2001-11-21T19:44:29.239Z",
    "created_at": "2025-11-21T20:47:36.418456625+01:00",
    "updated_at": "2025-11-21T20:47:36.418456675+01:00"
  }' | jq
```

### List all Users

```bash
curl http://localhost:4000/v1/users | jq
```

### Delete User

```bash
curl -X DELETE http://localhost:4000/v1/users/{user_id} | jq
```
