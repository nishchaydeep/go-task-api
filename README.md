# Go Task API

A simple and extensible **Task Management REST API** built with **Go**.


##  How to Run the Project

### Prerequisites

- Go 1.23+ installed
- `.env` file configured for SMTP

## 🛡️ Automatic Code Quality Enforcement

**No setup required!** This repository automatically enforces code quality standards via GitHub Actions.

### What's Enforced Automatically:
- ✅ No trailing whitespace in `.go`, `.yml`, `.yaml` files
- ✅ Runs on every push and pull request
- ✅ Cannot be bypassed by developers
- ✅ Blocks merges if checks fail

### Optional: Faster Local Feedback

Want to catch issues **before** pushing? Run this once:

```bash
git config core.hooksPath .githooks
```

This enables local pre-push checks for instant feedback (optional but recommended).


### Running Locally with Go

```bash
go run main.go
```

The server starts on:
**[http://localhost:8070/tasks](http://localhost:8070/tasks)**

### Running with Docker (via Rancher Desktop / nerdctl)

**1. Build the image:**

```bash
nerdctl build -f Dockerfile -t go-task-api .
```

**2. Run the container:**

```bash
nerdctl run -p 8070:8070 go-task-api
```


## Environment Variables 

| Key             | Description                         |
| --------------- | ----------------------------------- |
| `SMTP_EMAIL`    | Email address used to send the mail |
| `SMTP_PASSWORD` | SMTP auth password or app password  |
| `SMTP_TO`       | Recipient email address             |
| `SMTP_HOST`     | e.g., smtp.gmail.com                |
| `SMTP_PORT`     | Usually 587 (TLS)                   |

## Email Configuration (`.env`)

Create a `.env` file in your root directory with:

```env
SMTP_EMAIL=your@email.com
SMTP_PASSWORD=your_app_password
SMTP_TO=recipient@email.com
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

* App sends an email on startup with the filtered task summary based on the attribute user wants.


## API Endpoints

### Add Task

```http
POST /task
```

**Body (JSON):**

```json
{
        "name": "Finance Task",
        "completed": false,
        "created_at": "2025-08-12T14:16:45.294515+05:30",
        "updated_at": "2025-08-12T14:16:45.294515+05:30",
        "accessed_at": "2025-08-12T14:16:45.294515+05:30",
        "description": "handling finance tasks",
        "category": "office",
        "important" : true
}
```


### List All Tasks

```http
GET /tasks
```


### Get Task by Name

```http
GET /task?name=Required Task
```


### Get Task by any other Field

```http
GET /tasks?Field Name=Required category
```

Field Name can be description, category, name, important, completed


### Delete Task by Name

```http
DELETE /task?name=Required Task
```

## API Documentation (Swagger)

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
   `http://localhost:8070/swagger/index.html`





Made by [Nishchay Deep](https://github.com/nishchaydeep15) 

