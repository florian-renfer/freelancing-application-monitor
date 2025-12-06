# 📊 Freelancing Application Monitor

> [!IMPORTANT]
> This project is in a very early state and far from ready for production. So consider the project _work in progress_.

An interactive monitor that helps you keeping an eye on your applications for freelance projects.

## ✨ Technologies

### API

- Go
- Gin
- SQLite

### UI

- Next.JS
- Shadcn

## 🚀 Features

The application allows freelancers to keep track of their applications for freelance projects. It provides
functionalities such as creating and updating applications, and checking for possible duplicates using artificial intelligence (AI).

### Assumptions

The API assumes that data entering the `usecase` layer is validated before exectuing the actual usecase.

## 📍 The Process

The motivation behind this project is to create a tool that helps freelancers manage their applications more efficiently. This involves tracking the status of applications, deadlines, reminders, notes, and using AI to identify potential duplicate applications, etc.

By reducing the administrative burden, freelancers can focus more on their core work and increase their chances of securing projects.

## 🚦 Running the Project

1. Clone the repository: `git clone git@github.com:florian-renfer/freelancing-application-monitor.git`
2. Change directory: `cd freelancing-application-monitor`

### Running the API

1. Install dependencies: `go mod download`
2. Run development server: `go run cmd/api/main.go`
3. Execute requests against the server running on: `http://localhost:4000/v1/<api-endpoint>`

### Running the UI

1. Change directory: `cd web`
2. Install dependencies: `npm install`
3. Run development server: `npm run dev`

## 🏎️ Roadmap

- [ ] Create an Applcation
  - [x] Define API endpoint
  - [ ] Define Constraints
  - [ ] Add Validation
- [ ] AI-based Duplicate Detection
- [ ] Unit Tests
- [ ] Integration Tests
- [ ] User Interface
- [ ] User Authentication, OAuth2.0

## 📚 Lessons Learned

- Implementing enumerated types using `iota` while ensuring a proper string representation for better readability and maintainability
- Implementing interfaces such as `Scanner` from `database/sql` package to convert database rows from primitive types to custom types
- Implementing interfaces such as `Unmarshaler` and `Marshaler` from `encoding/json` package to convert between JSON and custom types

## Use Cases

> [!WARNING]
> The endpoints are listed for documentation purposes only. There is no
> guarantee of correct behavior, input validation, or anything else that would
> prevent the application or database from breaking.

> [!NOTE]
> I'm planning on releasing an OpenAPI specification once the API is in a stable
> state.

### Applcations

#### Create Application

```bash
curl -X POST http://localhost:4000/v1/applications\
  -H "Content-Type: application/json" \
  -d '{
    "title": "Java Entwickler | Hamburg gesucht",
    "description": "Wir sind auf der Suche nach einem senior Java support Entwickler mit folgenden technischen und funktionalen Fähigkeiten...",
    "url": "https://www.freelancermap.de/projekt/java-entwickler-hamburg-gesucht",
    "state": "APPLIED",
    "applied_at": "2025-11-27T20:59:00.895974+01:00"
  }' | jq
```

#### List all Applcations

```bash
curl -X GET http://localhost:4000/v1/applications | jq
```

### Users

#### Register User

```bash
curl -X POST http://localhost:4000/v1/auth/register\
  -H "Content-Type: application/json" \
  -d '{
    "email": "max.mustermann@mail.de",
    "password": "changeme",
    "firstname": "Max",
    "lastname": "Mustermann",
    "date_of_birth": "2000-11-27T20:59:00.895974+01:00"
  }' | jq
```

#### Delete User

```bash
curl -X DELETE http://localhost:4000/v1/users/{user_id} | jq
```
