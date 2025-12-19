# Age Calculator – DOB & Dynamic Age Calculation

A RESTful API built using **Go** and **Fiber** to manage users with their **name** and **date of birth (DOB)**.  
The API calculates the **age dynamically** using Go’s `time` package instead of storing it in the database.

This project follows **clean architecture** and uses **SQLC** for type-safe database access.

---

## 🚀 Tech Stack

- **Go (Golang)**
- **GoFiber** – Web framework
- **PostgreSQL**
- **SQLC** – Type-safe SQL queries
- **Uber Zap** – Structured logging
- **go-playground/validator** – Input validation
- **Docker & Docker Compose**

---

## 📁 Project Structure

```
cmd/server/main.go → Application entry point
config/ → Environment configuration
db/
├── migrations/ → Database migration files
└── sqlc/ → SQLC config & generated code
internal/
├── handler/ → HTTP handlers (controllers)
├── service/ → Business logic (age calculation)
├── repository/ → Database access layer
├── routes/ → API route definitions
├── middleware/ → RequestID, logging middleware
├── models/ → Request & response models
└── logger/ → Zap logger setup
```

---

## 🧱 Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

## 🧱⚙️ Setup & Run (Without Docker)
#### 1️⃣ Clone Repository
```
git clone https://github.com/PranayDhanke/AgeCalculatorUsingGo.git
cd go-user-apiAgeCalculatorUsingGo

#### 2️⃣ Set Environment Variable
```
set DATABASE_URL=postgres://postgres:password@localhost:5432/userdb
```

#### 3️⃣ Run Database Migration
```
psql -U postgres -d userdb -f db/migrations/001_create_users.sql
```

#### 4️⃣ Generate SQLC Code
```
cd db/sqlc
sqlc generate
cd ../../
```

#### 5️⃣ Run Application
```
go run cmd/server/main.go
```

### Server will start at:
```
http://localhost:8080
```
🧪 Running Unit Tests
```
go test ./...
```

#### Includes:

Unit tests for age calculation

Table-driven test cases


### 🐳 Run Using Docker (Recommended)
#### 1️⃣ Build & Start Services
```
docker-compose up --build
```

#### This will start:

API on 
``` 
http://localhost:8080 
```

PostgreSQL database on port 5432

#### 2️⃣ Stop Services
```
docker-compose down
```
