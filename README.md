# Kredit Plus API

🚀 Backend API Service for Kredit Plus Test Project  
Built with Go (Echo Framework), MySQL, Docker, and Swagger Documentation.

---

## 📦 Features
- Clean Architecture (Handler → Service → Repository → Model)
- Mutex Locking Per User (Concurrency Safe Transaction)
- Auto Database Migration on Startup (Goose)
- Dockerized (App + MySQL)
- OpenAPI 3.0 Documentation (`docs/openapi.yaml`)
- Ready for Authentication Integration

---

## 🛠️ Tech Stack
- Go 1.24
- Echo Web Framework
- MySQL 8
- Docker & Docker Compose
- Goose DB Migration
- OpenAPI Spec (Swagger Format)

---

## 🚀 How to Run Locally

### 1. Clone This Project
```bash
git clone https://github.com/yourusername/kredit-plus-test.git
cd kredit-plus-test
```

### 2. Setup Environment Variables
Create a `.env` file in the project root:

```dotenv
# App Settings
PORT=8080

# Database Settings
DB_HOST=db
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=kredit_plus
```

### 3. Build and Run with Docker Compose
```bash
docker-compose up --build
```

- API Server will be accessible at:
  ➔ `http://localhost:8080`
- OpenAPI Specification:
  ➔ `docs/openapi.yaml`

---

## 📚 API Endpoints

| Method | Endpoint | Description | Response |
|:---|:---|:---|:---|
| POST | `/users` | Create new user | 201 Created (JSON) |
| GET | `/users/{id}` | Get user by ID | 200 OK (JSON) |
| POST | `/limits` | Create user limit | 201 Created (JSON) |
| GET | `/limits/{user_id}` | Get all limits by user | 200 OK (JSON Array) |
| POST | `/transactions` | Create new transaction | 201 Created (JSON) |
| GET | `/transactions/{user_id}` | Get all transactions by user | 200 OK (JSON Array) |
| GET | `/transactions/id/{id}` | Get transaction by ID | 200 OK (JSON) |

Detailed request/response structure is available in the OpenAPI Spec.

---



---

## 📄 Documentation
- OpenAPI Spec available at:
  ➔ `docs/openapi.yaml`
- ERD Diagram available at:
  ➔ `docs/erd.png`
- App Architecture available at:
  ➔ `docs/architechture.png`

---


---

## 🙌 Author
Made with ❤️ by alfa adriel monico.

