
# Go Task API

A simple task management API written in Go.

---


### Running Locally (with Go)

1. Make sure you have Go 1.23+ installed.

2. Clone the repo and navigate to the project folder:

   ```bash
   git clone https://github.com/nishchaydeep15/go-task-api.git
   cd go-task-api
````

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the app:

   ```bash
   go run main.go
   ```

5. The server will start on port `8070` .

---

### Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t go-task-api .
   ```

2. Run the container:

   ```bash
   docker run -p 8070:8070 go-task-api
   ```

3. The API will be accessible at `http://localhost:8070`.

---

## API Usage Examples

### Add a Task

**Request:**

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Finish report",
    "completed": false,
    "search": "report",
    "description": "Complete the quarterly report",
    "category": "work"
  }'
```

**Response:**

```json
{
  "Message": "Task added"
}
```

---

### List Tasks (filter by category)

**Request:**

```bash
curl "http://localhost:8080/tasks?category=work"
```

**Response:**

```json
[
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
]
```

---

### Get a Task by Name

**Request:**

```bash
curl "http://localhost:8080/task?name=Finish report"
```

---

### Delete a Task

**Request:**

```bash
curl -X DELETE "http://localhost:8080/task?name=Finish report"
```

**Response:**

```json
{
  "message": "Task Deleted"
}
```

---

## Environment Variables (for email sending)

* `SMTP_EMAIL` — your SMTP email address
* `SMTP_PASSWORD` — your SMTP app password which can be generated after checking app passwords in your google account
* `SMTP_HOST` — SMTP server host (e.g., smtp.gmail.com)
* `SMTP_PORT` — SMTP server port (e.g., 587)
* `SMTP_TO` — recipient email address for task summary emails

---

## Notes

* The API stores tasks persistently in a local file (`tasks.json`).
* Email sending is done as a background job filtering tasks by category.

```

```
