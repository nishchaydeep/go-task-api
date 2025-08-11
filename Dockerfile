FROM golang:1.21-alpine
WORKDIR /go-task-api
COPY . .
RUN go build -o main .
CMD ["./main"]

