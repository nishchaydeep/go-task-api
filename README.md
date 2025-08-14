
```markdown
# 📌 Go Task API

A simple and extensible **Task Management REST API** built with **Go**

---

## 🚀 How to Run the Project

### Prerequisites

* Go 1.23+ installed
* `.env` file configured for SMTP 

---

### Running Locally with Go

```bash
go run main.go
```

The server starts on:
📍 **[http://localhost:8070](http://localhost:8070)**

---

### 🐳 Running with Docker (via Rancher Desktop / nerdctl)

**1. Build the image:**

```bash
nerdctl build -f Dockerfile -t go-task-api .
```

**2. Run the container:**

```bash
nerdctl run -p 8070:8070 --env-file .env go-task-api
```

---

## 📨 Email Configuration (`.env`)

Create a `.env` file in your root directory with:

```env
CATEGORY=work
SMTP_EMAIL=your@email.com
SMTP_PASSWORD=your_app_password
SMTP_TO=recipient@email.com
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

* `CATEGORY` is the task category you want to filter and email.
* App sends an email on startup with the filtered task summary.

---

## 📡 API Endpoints

### ➕ Add Task

```http
POST /tasks
```

**Body (JSON):**

```json
{
        "name": "Finance Task",
        "completed": false,
        "search": "fin",
        "created_at": "2025-08-12T14:16:45.294515+05:30",
        "updated_at": "2025-08-12T14:16:45.294515+05:30",
        "accessed_at": "2025-08-12T14:16:45.294515+05:30",
        "description": "handling finance tasks",
        "category": "office"
}
```

---

### 📋 List All Tasks

```http
GET /tasks
```

---

### 🔍 Get Task by Name

```http
GET /tasks?name=Required Task
```

---

### ❌ Delete Task by Name

```http
DELETE /tasks?name=Required Task
```

---

## 🔧 Environment Variables 

| Key             | Description                         |
| --------------- | ----------------------------------- |
| `CATEGORY`      | Task category to filter and email   |
| `SMTP_EMAIL`    | Email address used to send the mail |
| `SMTP_PASSWORD` | SMTP auth password or app password  |
| `SMTP_TO`       | Recipient email address             |
| `SMTP_HOST`     | e.g., smtp.gmail.com                |
| `SMTP_PORT`     | Usually 587 (TLS)                   |

---

## 📘 API Documentation (Swagger - Optional)

If using `swaggo`:

1. Install swag CLI:

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. Generate docs:

   ```bash
   swag init
   ```

3. Access Swagger UI:
   📍 `http://localhost:8070/swagger/index.html`





Made by [Nishchay Deep](https://github.com/nishchaydeep15) 

```

---

```
